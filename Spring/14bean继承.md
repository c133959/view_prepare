# Spring bean继承

> 在 Spring 中，Bean 和 Bean 之间也存在继承关系。我们将被继承的 Bean 称为父 Bean，将继承父 Bean 配置信息的 Bean 称为子 Bean。

在 Spring XML 配置中，我们通过子 Bean 的 parent 属性来指定需要继承的父 Bean，配置格式如下。

```xml
<!--父Bean-->
<bean id="parentBean" class="xxx.xxxx.xxx.ParentBean" >
    <property name="xxx" value="xxx"></property>
    <property name="xxx" value="xxx"></property>
</bean> 
<!--子Bean--> 
<bean id="childBean" class="xxx.xxx.xxx.ChildBean" parent="parentBean"></bean>
```

Spring Bean 的定义中可以包含很多配置信息，例如构造方法参数、属性值。

**子 Bean 既可以继承父 Bean 的配置数据，也可以根据需要重写或添加属于自己的配置信息。**



## Bean定义模板

在父 Bean 的定义中，有一个十分重要的属性，那就是 **abstract** 属性。**如果一个父 Bean 的 abstract 属性值为 true，则表明这个 Bean 是抽象的。**（java中的抽象类）

抽象的父 Bean 只能作为模板被子 Bean 继承，**它不能实例化**，也不能被其他 Bean 引用，更不能在代码中根据 id 调用 getBean() 方法获取，否则就会返回错误。

在父 Bean 的定义中，既可以指定 class 属性，也可以不指定 class 属性。**如果父 Bean 定义没有明确地指定 class 属性，那么这个父 Bean 的 abstract 属性就必须为 true。**



```xml
<bean id="animal" abstract="true">
    <property name="name" value="动物"></property>
    <property name="age" value="10"></property>
</bean>


<bean id="dog" class="net.biancheng.c.Dog" parent="animal">
    <property name="name" value="小狗"></property>
    <property name="call" value="汪汪汪……"></property>
</bean>
```

```shell
Dog setName：小狗
Dog setAge：10
Dog setCall：汪汪汪……
Dog{name='小狗', age=10, call='汪汪汪……'}
```

**通过控制台输出可以看出，在 Spring 启动时，并没有实例化 animal。**

