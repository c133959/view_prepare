# AOP面向切面编程

> AOP 的全称是“Aspect Oriented Programming”，译为“面向切面编程”，和 OOP（面向对象编程）类似，它也是一种编程思想。
>
> ![应用程序的组成](http://c.biancheng.net/uploads/allimg/220119/1FS54P5-0.png)

## 概述

与 OOP 中纵向的父子继承关系不同，AOP 是通过**横向的抽取机制**实现的。它将应用中的一些**非业务的通用功能**抽取出来单独维护，并通过**声明的方式（例如配置文件、注解等）**定义这些功能要以何种方式作用在那个应用中，而不是在业务模块的代码中直接调用。



目前最流行的 AOP 实现（框架）主要有两个，分别为 [Spring AOP](http://c.biancheng.net/spring/aop-module.html) 和 [AspectJ](http://c.biancheng.net/spring/aspectj.html)。

| AOP 框架   | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| Spring AOP | 是一款基于 AOP 编程的框架，它能够有效的减少系统间的重复代码，达到松耦合的目的。  Spring AOP 使用纯 Java 实现，不需要专门的编译过程和类加载器，在运行期间通过**代理**方式向目标类植入增强的代码。Spring AOP 支持 2 种代理方式，分别是**基于接口的 JDK 动态代理**和**基于继承的 CGLIB 动态代理**。 |
| AspectJ    | 是一个基于 Java 语言的 AOP 框架，从 Spring 2.0 开始，Spring AOP 引入了对 AspectJ 的支持。  AspectJ 扩展了 Java 语言，提供了一个专门的编译器，在编译时提供横向代码的植入。 |



## AOP术语

| 名称                | 说明                                                         |
| ------------------- | ------------------------------------------------------------ |
| Joinpoint（连接点） | AOP 的核心概念，指的是程序执行期间明确定义的一个点，例如**方法的调用、类初始化、对象实例化等**。  在 Spring 中，连接点则指可以被动态代理拦截目标类的方法。 |
| Pointcut（切入点）  | 又称切点，指**要对哪些 Joinpoint 进行拦截**，即被拦截的连接点。 |
| Advice（通知）      | 指**拦截到 Joinpoint 之后要执行的代码**，即对切入点增强的内容。 |
| Target（目标）      | 指**代理的目标对象**，通常也被称为被通知（advised）对象。    |
| Weaving（织入）     | 指**把增强代码应用到目标对象上，生成代理对象的过程**。       |
| Proxy（代理）       | 指**生成的代理对象**。                                       |
| Aspect（切面）      | 切面是**切入点（Pointcut）和通知（Advice）的结合**。         |



**Advice** 直译为通知，也有人将其翻译为“增强处理”，共有 5 种类型，如下表所示。

| 通知                           | 说明                               |
| ------------------------------ | ---------------------------------- |
| before（前置通知）             | 通知方法在目标方法调用之前执行     |
| after（后置通知）              | 通知方法在目标方法返回或异常后调用 |
| after-returning（返回后通知）  | 通知方法会在目标方法返回后调用     |
| after-throwing（抛出异常通知） | 通知方法会在目标方法抛出异常后调用 |
| around（环绕通知）             | 通知方法会将目标方法封装起来       |



## AOP类型

### 动态 AOP

动态 AOP 的织入过程是在**运行时动态执行的**。其中最具代表性的动态 AOP 实现就是 **Spring AOP**，它会为所有被通知的对象创建代理对象，并通过代理对象对被原对象进行增强。

相较于静态 AOP 而言，动态 AOP 的**性能通常较差**，但随着技术的不断发展，它的性能也在不断的稳步提升。

动态 AOP 的优点是它可以轻松地对应用程序的所有切面进行修改，而**无须对主程序代码进行重新编译**。

### 静态 AOP

静态 AOP 是通过修改应用程序的实际 Java 字节码，根据需要修改和扩展程序代码来实现织入过程的。最具代表性的静态 AOP 实现是 AspectJ。

相较于动态 AOP 来说，性能较好。但它也有一个明显的缺点，那就是**对切面的任何修改都需要重新编译整个应用程序。**





## why AOP

在 Spring 框架中使用 AOP 主要有以下优势。

- 提供**声明式**企业服务，特别是作为 EJB 声明式服务的替代品，最重要的是，这种服务是声明式事务管理。
- **允许用户实现自定义切面**。在某些不适合用 OOP 编程的场景中，采用 AOP 来补充。
- 可以对业务逻辑的各个部分进行**隔离**，从而使业务逻辑各部分之间的**耦合度降低**，提高程序的**可重用性**，同时也提高了**开发效率**。

