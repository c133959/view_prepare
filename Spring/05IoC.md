# IoC

> IoC 是 Inversion of Control 的简写，译为“**控制反转**”，它不是一门技术，而是**一种设计思想**，是一个重要的面向对象编程法则，能够指导我们**如何设计出松耦合、更优良的程序**。

* IoC容器作用：
  * 管理所有 Java 对象的**实例化和初始化**
  * 控制对象与对象之间的依赖关系

***我们将由 IoC 容器管理的 Java 对象称为 Spring Bean，它与使用关键字 new 创建的 Java 对象没有任何区别。***



## 控制反转 ( IoC )

**传统方法**，A想要调用B的属性或者方法：

```java
// new 一个 B的对象
public class A {
    B b = new B();
    // 创建出B的实例对象后，调用B的属性或方法
    sout(B.info);
    sout(b.getInfo());
   //通常把A称为调用者，B称为被调用者
}
```

**Spring方法**：

在 Spring 应用中，Java 对象创建的控制权是掌握在 IoC 容器手里的，其大致步骤如下。

1. 开发人员通过 XML 配置文件、注解、Java 配置类等方式，对 Java 对象进行定义，例如
   1. 在 XML 配置文件中**使用 <bean> 标签**
   2. 在 Java 类上使用 **@Component 注解**等。
2. Spring 启动时，IoC 容器会自动根据对象定义，**将这些对象创建并管理起来**。这些被 IoC 容器创建并管理的对象被称为 **Spring Bean**。
3. 当我们想要使用某个 Bean 时，可以直接从 IoC 容器中获取（例如通过 ApplicationContext 的 getBean() 方法），而不需要手动通过代码（例如 new Obejct() 的方式）创建。


IoC 带来的最大改变不是代码层面的，而是从思想层面上发生了“主从换位”的改变。原本调用者是主动的一方，它想要使用什么资源就会主动出击，自己创建；但在 Spring 应用中，IoC 容器掌握着主动权，调用者则变成了被动的一方，被动的等待 IoC 容器创建它所需要的对象（Bean）。

> 原：主动出击，直接new出来
>
> 现：IoC掌握主动权，等 IoC 容器创建它所需要的对象（Bean）
>
> 这个过程在职责层面发生了控制权的反转，把原本调用者通过代码实现的对象的创建，反转给 IoC 容器来帮忙实现，因此我们将这个过程称为 Spring 的“控制反转”。



## 依赖注入（DI）

依赖注入（Denpendency Injection，简写为 DI）

简单来说，依赖关系就是在一个对象中需要用到另外一个对象，即对象中存在一个属性，该属性是另外一个类的对象。

例如，有一个名为 B 的 Java 类，它的代码如下。

```java
public class B {    
    String bid;    
    A a;
}
```


从代码可以看出，B 中存在一个 A 类型的对象属性 a，此时我们就可以说 **B 的对象依赖于对象 A**。而依赖注入就是就是基于这种“依赖关系”而产生的。

我们知道，控制反转核心思想就是由 Spring 负责对象的创建。在对象创建过程中，S**pring 会自动根据依赖关系，将它依赖的对象(A)  注入  到当前对象(B)中，这就是所谓的“依赖注入”。**



## IoC工作原理

IoC 底层通过工厂模式、Java 的反射机制、XML 解析等技术，将代码的耦合度降低到最低限度，其主要步骤如下。

1. 在配置文件（例如 Bean.xml）中，对各个对象以及它们之间的**依赖关系**进行配置；
2. 我们可以把 IoC 容器当做一个**工厂**，这个工厂的产品就是 **Spring Bean**；
3. 容器启动时会加载并解析这些配置文件，得到**对象的基本信息**以及它们之间的**依赖关系**；
4. IoC 利用 Java 的反射机制，**根据类名生成相应的对象（即 Spring Bean）**，并根据依赖关系将这个对象**注入**到依赖它的对象中。



### 解耦原理

由于对象的基本信息、对象之间的依赖关系都是在配置文件中定义的，并没有在代码中紧密耦合

当对象发生改变，我们也只需要在配置文件中进行修改即可，而无须对 Java 代码进行修改



## IoC容器的两种实现 

> IoC 容器底层其实就是一个 Bean 工厂

Spring 框架为我们提供了两种不同类型 IoC 容器，它们分别是 **BeanFactory 和 ApplicationContext。**

### BeanFactory

> BeanFactory 是 IoC 容器的**基本实现**，也是 Spring 提供的最简单的 IoC 容器，它提供了 IoC 容器最基本的功能，由 org.springframework.beans.factory.BeanFactory 接口定义。

懒加载

```java
// 用法
BeanFactory context = new ClassPathXmlApplicationContext("Beans.xml");
```

> 注意：BeanFactory 是 Spring 内部使用接口，通常情况下不提供给开发人员使用。 

### ApplicationContext

> ApplicationContext 是 BeanFactory 接口的子接口，是对 BeanFactory 的扩展。ApplicationContext 在 BeanFactory 的基础上增加了许多企业级的功能，例如 AOP（面向切面编程）、国际化、事务支持等。

```java
public interface ApplicationContext extends 
 EnvironmentCapable, ListableBeanFactory, HierarchicalBeanFactory, MessageSource, ApplicationEventPublisher, ResourcePatternResolver {
    
}
```



ApplicationContext 接口有两个常用的实现类，具体如下表。

| 实现类                          | 描述                                                         | 示例代码                                                     |
| ------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| ClassPathXmlApplicationContext  | 加载**类路径 ClassPath 下**指定的 XML 配置文件，并完成 ApplicationContext 的实例化工作 | ApplicationContext applicationContext = new ClassPathXmlApplicationContext(String configLocation); |
| FileSystemXmlApplicationContext | 加载指定的**文件系统路径中**指定的 XML 配置文件，并完成 ApplicationContext 的实例化工作 | ApplicationContext applicationContext = new FileSystemXmlApplicationContext(String configLocation); |

```java
public static void main(String[] args) {
    //使用 FileSystemXmlApplicationContext 加载指定路径下的配置文件 Bean.xml
    BeanFactory context = new FileSystemXmlApplicationContext(
        "D:\\eclipe workspace\\spring workspace\\HelloSpring\\src\\Beans.xml");
    
    HelloWorld obj = context.getBean("helloWorld", HelloWorld.class);
    obj.getMessage();
}
```

