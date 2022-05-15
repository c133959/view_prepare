# Spring AOP



## 代理机制

Spring 在运行期会为目标对象生成一个动态代理对象，并在代理对象中实现对目标对象的增强。

Spring AOP 的底层是通过以下 2 种动态代理机制，为目标对象（Target Bean）执行横向织入的。

| 代理技术       | 描述                                                         |
| -------------- | ------------------------------------------------------------ |
| JDK 动态代理   | Spring AOP 默认的动态代理方式，若目标对象实现了若干接口，Spring 使用 JDK 的 java.lang.reflect.Proxy 类进行代理。 |
| CGLIB 动态代理 | 若目标对象没有实现任何接口，Spring 则使用 CGLIB 库生成目标对象的子类，以实现对目标对象的代理。 |



## Spring AOP 连接点

> AOP 的核心概念，指的是程序执行期间明确定义的一个点，例如**方法的调用、类初始化、对象实例化等**。  在 Spring 中，连接点则指可以被动态代理拦截目标类的方法。

Spring AOP 并没有像其他 AOP 框架（例如 AspectJ）一样提供了完成的 AOP 功能，它是 Spring 提供的一种简化版的 AOP 组件。其中最明显的简化就是，Spring AOP 只支持一种连接点类型：**方法调用**。

您可能会认为这是一个严重的限制，但实际上 Spring AOP 这样设计的原因是为了让 Spring 更易于访问。

**如果需要使用其他类型的连接点（例如成员变量连接点），我们可以将 Spring AOP 与其他的 AOP 实现一起使用，最常见的组合就是 Spring AOP + ApectJ。** 



## Spring AOP 通知类型

Spring AOP 按照通知（Advice）织入到目标类方法的连接点位置，为 Advice 接口提供了 6 个子接口，如下表。

| 通知类型     | 接口                                            | 描述                                             |
| ------------ | ----------------------------------------------- | ------------------------------------------------ |
| 前置通知     | org.springframework.aop.MethodBeforeAdvice      | 在目标方法执行前实施增强。                       |
| 后置通知     | org.springframework.aop.AfterReturningAdvice    | 在目标方法执行后实施增强。                       |
| 后置返回通知 | org.springframework.aop.AfterReturningAdvice    | 在目标方法执行完成，并返回一个返回值后实施增强。 |
| 环绕通知     | org.aopalliance.intercept.MethodInterceptor     | 在目标方法执行前后实施增强。                     |
| 异常通知     | org.springframework.aop.ThrowsAdvice            | 在方法抛出异常后实施增强。                       |
| 引入通知     | org.springframework.aop.IntroductionInterceptor | 在目标类中添加一些新的方法和属性。               |



## Spring AOP 通知类型

Spring 使用 org.springframework.aop.Advisor 接口表示切面的概念，实现对通知（Adivce）和连接点（Joinpoint）的管理。

**在 Spring AOP 中，切面可以分为三类：一般切面、切点切面和引介切面。**

| 切面类型 | 接口                                                  | 描述                                                         |
| -------- | ----------------------------------------------------- | ------------------------------------------------------------ |
| 一般切面 | org.springframework.aop.<br />**Advisor**             | Spring AOP 默认的切面类型。  由于 Advisor 接口仅包含一个 Advice（通知）类型的属性，而没有定义 PointCut（切入点），因此它表示一个不带切点的简单切面。  这样的切面会对目标对象（Target）中的所有方法进行拦截并织入增强代码。由于这个切面太过宽泛，因此我们一般不会直接使用。 |
| 切点切面 | org.springframework.aop.<br />**PointcutAdvisor**     | Advisor 的子接口，用来表示带切点的切面，该接口在 Advisor 的基础上还维护了一个 PointCut（切点）类型的属性。  使用它，我们可以通过包名、类名、方法名等信息更加灵活的定义切面中的切入点，提供更具有适用性的切面。 |
| 引介切面 | org.springframework.aop.<br />**IntroductionAdvisor** | Advisor 的子接口，用来代表引介切面，引介切面是对应引介增强的特殊的切面，它应用于类层面上，所以引介切面适用 ClassFilter 进行定义。 |

