# 创建Spring程序



 ## 1.结构

![image-20220214141153746](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220214141153746.png)

## 2.代码

**Message.java**

```java
public class Message {

    public void getMessage() {
        System.out.println("message" + message);
    }

    public void setMessage(String message) {
        this.message = message;
    }

    private String message;

}
```



**MainAPP.java**

```java
package com.ssl;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

/**
 * 1. 创建 ApplicationContext 对象时使用了 ClassPathXmlApplicationContext 类，
 * 这个类用于加载 Spring 配置文件、创建和初始化所有对象（Bean）。
 * 2. ApplicationContext.getBean() 方法用来获取 Bean，该方法返回值类型为 Object，
 * 通过强制类型转换为 HelloWorld 的实例对象，调用其中的 getMessage() 方法。
 */
public class MainAPP {
    public static void main(String[] args) {
        ApplicationContext context = new ClassPathXmlApplicationContext("Beans.xml");
        Message message = context.getBean("message", Message.class);
        message.getMessage();
    }
}

```

在Resources文件夹中创建Spring配置文件

```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd">
    
    <bean id="message" class="com.ssl.Message">
        <property name="message" value="hello spring"></property>
    </bean>
    
</beans>
```

我们也可以将该配置文件命名为其它有效的名称，但需要注意的是，该文件名必须与 MainApp.java 中读取的配置文件名称一致。

Beans.xml 用于给不同的 Bean 分配唯一的 ID，并给相应的 Bean 属性赋值。例如，在以上代码中，我们可以在不影响其它类的情况下，给 message 变量赋值。



## 运行结果

![image-20220214141545138](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220214141545138.png)