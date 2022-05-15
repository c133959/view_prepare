# Spring自动装配（基于注解）

> Spring 通过注解实现自动装配的步骤如下：
>
> 1. 引入依赖
> 2. 开启组件扫描
> 3. 使用注解定义 Bean
> 4. 依赖注入



## 1. 引入依赖

使用注解的第一步，就是要在项目中引入以下 Jar 包。

- org.springframework.**core**-5.3.13.jar
- org.springframework.**beans**-5.3.13.jar
- spring-**context**-5.3.13.jar
- spring-**expression**-5.3.13.jar
- commons.**logging**-1.2.jar
- spring-**aop**-5.3.13.jar

> 注意，除了 spring 的四个基础 jar 包和 commons-logging-xxx.jar 外，想要使用注解实现 Spring 自动装配，还需要引入Spring 提供的 spring-aop 的 Jar 包。



## 2. 开启组件扫描

Spring **默认不使用注解装配 Bean**，因此我们需要在 Spring 的 XML 配置中，通过<context : componet-scan>元素**开启** Spring Beans的自动扫描功能。

开启此功能后，Spring 会自动从**扫描指定的包（base-package 属性设置）及其子包下的所有类**

**如果类上使用了 @Component 注解，就将该类装配到容器中。**

```xml
    <!--开启组件扫描功能-->
    <context:component-scan base-package="net.biancheng.c"></context:component-scan>
```

> 注意：在使用<context : componet-scan>元素开启自动扫描功能前，首先需要在 XML 配置的一级标签 <beans> 中添加 context 相关的约束。

```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:context="http://www.springframework.org/schema/context"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
    http://www.springframework.org/schema/beans/spring-beans-3.0.xsd
    http://www.springframework.org/schema/context
            http://www.springframework.org/schema/context/spring-context.xsd">
    <!--开启组件扫描功能-->
    <context:component-scan base-package="net.biancheng.c"></context:component-scan>
</beans>
```



## 3. 使用注解定义bean

Spring 提供了以下多个注解，这些注解可以直接标注在 Java 类上，将它们定义成 Spring Bean。

| 注解            | 说明                                                         |
| --------------- | ------------------------------------------------------------ |
| @**Component**  | 该注解用于描述 Spring 中的 **Bean**，它是一个泛化的概念，仅仅表示容器中的一个组件（Bean），并且**可以作用在应用的任何层次**，例如 Service 层、Dao 层等。  使用时只需将该注解标注在相应类上即可。 |
| @**Repository** | 该注解用于将数据访问层（**Dao 层**）的类标识为 Spring 中的 Bean，其功能与 @Component 相同。 |
| @**Service**    | 该注解通常作用在业务层（**Service 层**），用于将业务层的类标识为 Spring 中的 Bean，其功能与 @Component 相同。 |
| @**Controller** | 该注解通常作用在控制层（如 Struts2 的 Action、SpringMVC 的 **Controller**），用于将控制层的类标识为 Spring 中的 Bean，其功能与 @Component 相同。 |



## 4. 基于注解方式实现依赖注入

我们可以通过以下注解将定义好 Bean 装配到其它的 Bean 中。



| 注解           | 说明                                                         |
| -------------- | ------------------------------------------------------------ |
| @**Autowired** | 可以应用到 Bean 的属性变量、setter 方法、非 setter 方法及构造函数等，默认按照 Bean 的类型进行装配。  <br />@Autowired 注解默认按照 **Bean 的类型**进行装配，默认情况下它要求依赖对象必须存在，如果允许 null 值，可以设置它的 required 属性为 false。如果我们想使用按照名称（byName）来装配，可以结合 @Qualifier 注解一起使用 |
| @**Resource**  | 作用与 Autowired 相同，区别在于 @Autowired 默认按照 Bean 类型装配，而 @Resource 默认按照 **Bean 的名称**进行装配。  <br />@Resource 中有两个重要属性：**name 和 type**。<br /> **Spring 将 name 属性解析为 Bean 的实例名称，type 属性解析为 Bean 的实例类型**。<br />如果指定 name 属性，则按实例名称进行装配；如果指定 type 属性，则按 Bean 类型进行装配；如果都不指定，则先按 Bean 实例名称装配，如果不能匹配，则再按照 Bean 类型进行装配；如果都无法匹配，则抛出 NoSuchBeanDefinitionException 异常。 |
| @**Qualifier** | 与 @Autowired 注解配合使用，会将默认的按 Bean **类型装配**修改为按 Bean 的**实例名称装配**，<br />**Bean 的实例名称由 @Qualifier 注解的参数指定。** |

