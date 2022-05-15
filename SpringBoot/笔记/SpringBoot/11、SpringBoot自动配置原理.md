# Spring Boot自动配置原理

> 我们知道，Spring Boot 项目创建完成后，即使不进行任何的配置，也能够顺利地运行，这都要归功于 **Spring Boot 的自动化配置。**

[SpringBoot官方文档：属性配置](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/#application-properties)

## 1.Spring Factories 机制

> Spring Boot 的自动配置是基于 Spring Factories 机制实现的。

Spring Factories 机制是 Spring Boot 中的一种服务发现机制，这种扩展机制与 Java SPI 机制十分相似。Spring Boot 会自动扫描所有 Jar 包类路径下 META-INF/spring.factories 文件，并读取其中的内容，进行实例化，这种机制也是 Spring Boot Starter 的基础。

### spring.factories 

spring.factories 文件**本质上与 properties 文件相似**，其中包含一组或多组键值对（key=vlaue），其中，key 的取值为接口的完全限定名；value 的取值为接口实现类的完全限定名，一个接口可以设置多个实现类，不同实现类之间使用“，”隔开，例如：

```java
org.springframework.boot.autoconfigure.AutoConfigurationImportFilter=\
org.springframework.boot.autoconfigure.condition.OnBeanCondition,\
org.springframework.boot.autoconfigure.condition.OnClassCondition,\
org.springframework.boot.autoconfigure.condition.OnWebApplicationCondition
```

### Spring Factories 实现原理

spring-core 包里定义了 `SpringFactoriesLoader` 类，这个类会扫描所有 Jar 包类路径下的` META-INF/spring.factories` 文件，并获取指定接口的配置。在 `SpringFactoriesLoader `类中定义了两个对外的方法，如下表。

| 返回值       | 方法                                                         | 描述                                                         |
| ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| <T> List<T>  | loadFactories(Class<T> factoryType, @Nullable ClassLoader classLoader) | 静态方法； 根据接口获取其实现类的实例； 该方法返回的是**实现类对象列表**。 |
| List<String> | loadFactoryNames(Class<?> factoryType, @Nullable ClassLoader classLoader) | 公共静态方法； 根据接口l获取其实现类的名称； 该方法返回的是**实现类的类名的列表** |

以上两个方法的关键都是从指定的 ClassLoader 中获取 spring.factories 文件，并解析得到类名列表

### 自动配置的加载

Spring Boot 自动化配置也是基于 Spring Factories 机制实现的，在 `spring-boot-autoconfigure-xxx.jar` 类路径下的 `META-INF/spring.factories` 中设置了 Spring Boot 自动配置的内容 。

以上配置中，value 取值是由多个 xxxAutoConfiguration （使用逗号分隔）组成，每个 xxxAutoConfiguration 都是一个自动配置类。**Spring Boot 启动时，会利用 Spring-Factories 机制，将这些 xxxAutoConfiguration 实例化并作为组件加入到容器中，以实现 Spring Boot 的自动配置。**

# @SpringBootApplication 注解

> 所有 Spring Boot 项目的主启动程序类上都使用了一个 @SpringBootApplication 注解，该注解是 Spring Boot 中最重要的注解之一 ，也是 Spring Boot 实现自动化配置的关键。 

@SpringBootApplication 是一个组合元注解，其主要包含两个注解：

**@SpringBootConfiguration** 和 **@EnableAutoConfiguration**

![@SpringBootApplication 注解](http://c.biancheng.net/uploads/allimg/210726/1546335957-0.png)

其中 **@EnableAutoConfiguration** 注解是 SpringBoot 自动化配置的核心所在。

## @EnableAutoConfiguration 注解

@EnableAutoConfiguration 注解用于**开启 Spring Boot 的自动配置功能**， 它使用 Spring 框架提供的 @Import 注解通过 **AutoConfigurationImportSelector类（选择器）**给容器中*导入自动配置组件*。

![AutoConfigurationImportSelector 类](http://c.biancheng.net/uploads/allimg/210726/1546332K5-1.png)

## AutoConfigurationImportSelector 类

AutoConfigurationImportSelector 类实现了 DeferredImportSelector 接口，AutoConfigurationImportSelector 中还包含一个静态内部类 AutoConfigurationGroup，它实现了 DeferredImportSelector 接口的内部接口 Group（Spring 5 新增）。

AutoConfigurationImportSelector 类中包含 3 个方法，如下表。

| 返回值                 | 方法声明                                                     | 描述                                                         | 内部类方法 | 内部类                 |
| ---------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ---------- | ---------------------- |
| Class<? extends Group> | getImportGroup()                                             | 该方法获取实现了 **Group 接口的类**，并实例化                | 否         |                        |
| void                   | process(AnnotationMetadata annotationMetadata, DeferredImportSelector deferredImportSelector) | 该方法用于**引入自动配置的集合**                             | 是         | AutoConfigurationGroup |
| Iterable<Entry>        | selectImports()                                              | **遍历**自动配置类集合（Entry 类型的集合），并逐个解析集合中的配置类 | 是         | AutoConfigurationGroup |


AutoConfigurationImportSelector 内各方法执行顺序如下。

1. getImportGroup() 方法
2. process() 方法
3. selectImports() 方法

#### 1. getImportGroup() 方法

AutoConfigurationImportSelector 类中 getImportGroup() 方法主要用于获取实现了 DeferredImportSelector.Group 接口的类，代码如下。

```java
    public Class<? extends Group> getImportGroup() {
        //获取实现了 DeferredImportSelector.Gorup 接口的 AutoConfigurationImportSelector.AutoConfigurationGroup 类
        return AutoConfigurationImportSelector.AutoConfigurationGroup.class;
    }
```

#### 2. process() 方法

静态内部类 AutoConfigurationGroup 中的核心方法是 process()，该方法通过调用 getAutoConfigurationEntry() 方法读取 spring.factories 文件中的内容，获得自动配置类的集合

#### 3. selectImports() 方法

以上所有方法执行完成后，AutoConfigurationImportSelector.AutoConfigurationGroup#selectImports() 会将 process() 方法处理后得到的自动配置类，进行过滤、排除，最后将所有自动配置类添加到容器中。