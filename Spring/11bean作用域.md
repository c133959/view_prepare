# Spring bean 作用域

> 默认情况下，所有的 Spring Bean 都是**单例**的，也就是说在整个 Spring 应用中， Bean 的实例只有一个。

Spring 5 共提供了 6 种 scope 作用域，如下表。



| 作用范围    | 描述                                                         |
| ----------- | ------------------------------------------------------------ |
| singleton   | 默认值，单例模式，表示在 Spring 容器中只有**一个** Bean 实例 |
| prototype   | 原型模式，表示每次通过 Spring 容器获取 Bean 时，容器都会创建一个**新的** Bean 实例。 |
| request     | **每次 HTTP 请求**，容器都会创建一个 Bean 实例。该作用域只在**当前 HTTP Request 内有效**。 |
| session     | 同一个 HTTP Session 共享一个 Bean 实例，不同的 Session 使用不同的 Bean 实例。该作用域仅在当前 HTTP Session 内有效。 |
| application | **同一个 Web 应用**共享一个 Bean 实例，该作用域在当前 ServletContext 内有效。  与 singleton 类似，但 singleton 表示每个 IoC 容器中仅有一个 Bean 实例，而一个 Web 应用中可能会存在多个 IoC 容器，但一个 Web 应用只会有一个 ServletContext，也可以说 application 才是 Web 应用中货真价实的单例模式。 |
| websocket   | websocket 的作用域是 WebSocket ，即在整个 WebSocket 中有效。 |

> 注意：在以上 6 种 Bean 作用域中，除了 singleton 和 prototype 可以直接在常规的 Spring IoC 容器（例如 ClassPathXmlApplicationContext）中使用外，剩下的都只能在基于 Web 的 ApplicationContext 实现（例如 XmlWebApplicationContext）中才能使用，否则就会抛出一个 **IllegalStateException** 的异常。



## singleton

```java
ApplicationContext context = new ClassPathXmlApplicationContext("Beans.xml");
SingletonBean bean1 = context.getBean("singletonBean", SingletonBean.class);
SingletonBean bean2 = context.getBean("singletonBean", SingletonBean.class);
System.out.println(bean1);// 输出地址，观察是否指向同一个对象
System.out.println(bean2);
```



```xml
<bean id="singletonBean" class="com.ssl.test3.SingletonBean" scope="singleton">
    <!-- 默认为单例模式  -->
    <property name="str" value="singleTest"></property>
</bean>
```

```shell
com.ssl.test3.SingletonBean@149494d8
com.ssl.test3.SingletonBean@149494d8
```



## prototype

```xml
<bean id="singletonBean" class="com.ssl.test3.SingletonBean" scope="prototype">
    <!--  默认为单例模式  -->
    <property name="str" value="singleTest"></property>
</bean>
```

```shell
com.ssl.test3.SingletonBean@78186a70
com.ssl.test3.SingletonBean@306279ee
```



