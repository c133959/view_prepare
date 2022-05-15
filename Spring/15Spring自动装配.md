# Spring自动装配

> 我们把 Spring 在 **Bean 与 Bean 之间建立依赖关系的行为**称为“装配”。

Bean 与 Bean 之间建立依赖关系的行为

在前面的学习中，我们都是在 XML 配置中通过 <constructor-arg>和 <property> 中的 **ref 属性**，手动维护 Bean 与 Bean 之间的依赖关系的。

**缺点：**

* 不适用于过于复杂的XML配置

* 可读性差

* 编写易出错



## 自动装配

Spring 框架式默认不支持自动装配的，要想使用自动装配，则需要对 Spring XML 配置文件中 <bean> 元素的 autowire 属性进行设置。

```xml
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
   http://www.springframework.org/schema/beans/spring-beans-3.0.xsd">
    <!--部门 Dept 的 Bean 定义-->
    <bean id="dept" class="net.biancheng.c.Dept"></bean>
   
    <!--雇员 Employee 的 Bean 定义,通过 autowire 属性设置自动装配的规则-->
    <bean id="employee" class="net.biancheng.c.Employee" autowire="byName">
    </bean>
</beans>
```



## 自动装配规则

Spring 共提供了 5 中自动装配规则，它们分别与 autowire 属性的 5 个取值对应，具体说明如下表。



| 属性值      | 说明                                                         | 问题                             |
| ----------- | ------------------------------------------------------------ | -------------------------------- |
| byName      | 按名称自动装配。  Spring 会根据的 Java 类中对象属性的名称，在整个应用的上下文 ApplicationContext（IoC 容器）中查找。若某个 Bean 的 **id 或 name 属性值**与这个对象属性的名称相同，则获取这个 Bean，并与当前的 Java 类 Bean 建立关联关系。 | 当bean的id修改时（不一致）会出错 |
| byType      | 按类型自动装配。  Spring 会根据 Java 类中的**对象属性的类型**，在整个应用的上下文 ApplicationContext（IoC 容器）中查找。若某个 Bean 的 class 属性值与这个对象属性的类型相匹配，则获取这个 Bean，并与当前的 Java 类的 Bean 建立关联关系。 | 当定义了多个类型一致的bean会出错 |
| constructor | 与 byType 模式相似，不同之处在与它应用于构造器参数（依赖项），如果在容器中没有找到与构造器参数类型一致的 Bean，那么将抛出异常。  其实就是根据构造器参数的数据类型，进行 byType 模式的自动装配。 | 配合 <constructor-arg>使用       |
| default     | 表示默认采用**上一级元素** <beans> 设置的自动装配规则（default-autowire）进行装配。 |                                  |
| no          | 默认值，表示**不使用自动装配**，Bean 的依赖关系必须通过 <constructor-arg>和 <property> 元素的 ref 属性来定义。 |                                  |



## 示例

[autowire使用示例](http://c.biancheng.net/spring/autowire-xml.html)

