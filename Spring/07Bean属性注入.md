# Bean属性注入

> 所谓 Bean 属性注入，简单点说就是将属性注入到 Bean 中的过程，
>
> 而这属性既可以**普通属性**，也可以是一个**对象（Bean）**。

Spring 主要通过以下 2 种方式实现属性注入：

- 构造函数注入
- setter 注入（又称设值注入）



## 构造函数注入



我们可以通过 Bean 的带参构造函数，以实现 Bean 的属性注入。

使用构造函数实现属性注入大致步骤如下：

1. 在 Bean 中添加一个有参构造函数，构造函数内的每一个参数代表一个需要注入的属性；
2. 在 Spring 的 XML 配置文件中，通过 <beans> 及其子元素 <bean> 对 Bean 进行定义；
3. 在 <bean> 元素内使用 <constructor-arg> 元素，对构造函数内的属性进行赋值，
4. **Bean 的构造函数内有多少参数，就需要使用多少个 <constructor-arg> 元素。**



### 代码示例

```java
package com.ssl.test1;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Student {
    private static final Log LOGGER = LogFactory.getLog(Student.class);
    private int id;
    private String name;
    private Grade grade;

    public Student(int id, String name, Grade grade) {
        this.id = id;
        this.name = name;
        this.grade = grade;
    }

    @Override
    public String toString() {
        return "Student{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", grade=" + grade +
                '}';
    }
}
```





```java
package com.ssl.test1;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Grade {

    private static final Log LOGGER = LogFactory.getLog(Grade.class);
    private Integer gradeId;
    private String gradeName;


    public Grade(Integer gradeId, String gradeName) {
        LOGGER.info("正在执行 Course 的有参构造方法，参数分别为：gradeId=" + gradeId + ",gradeName=" + gradeName);
        this.gradeId = gradeId;
        this.gradeName = gradeName;
    }

    @Override
    public String toString() {
        return "Grade{" +
                "gradeId=" + gradeId +
                ", gradeName='" + gradeName + '\'' +
                '}';
    }
}

```



```xml
<bean id="student" class="com.ssl.test1.Student">
        <constructor-arg index="0" value="2"/>
        <constructor-arg index="1" value="李四"/>
        <constructor-arg index="2" ref="grade"/>
    </bean>

    <bean id="grade" class="com.ssl.test1.Grade">
        <constructor-arg index="0" value="4"/>
        <constructor-arg index="1" value="四年级"/>
    </bean>
```



```shell
10:01:37.110 [main] DEBUG org.springframework.context.support.ClassPathXmlApplicationContext - Refreshing org.springframework.context.support.ClassPathXmlApplicationContext@1c655221
10:01:38.046 [main] DEBUG org.springframework.beans.factory.xml.XmlBeanDefinitionReader - Loaded 4 bean definitions from class path resource [Beans.xml]
10:01:38.293 [main] DEBUG org.springframework.beans.factory.support.DefaultListableBeanFactory - Creating shared instance of singleton bean 'message'
入参：constructor-arg
10:01:38.520 [main] DEBUG org.springframework.beans.factory.support.DefaultListableBeanFactory - Creating shared instance of singleton bean 'myRef'
10:01:38.524 [main] DEBUG org.springframework.beans.factory.support.DefaultListableBeanFactory - Creating shared instance of singleton bean 'student'
10:01:38.525 [main] DEBUG org.springframework.beans.factory.support.DefaultListableBeanFactory - Creating shared instance of singleton bean 'grade'
10:01:38.548 [main] INFO com.ssl.test1.Grade - 正在执行 Course 的有参构造方法，参数分别为：gradeId=4,gradeName=四年级
10:01:38.595 [main] INFO com.ssl.MainAPP - Student{id=2, name='李四', grade=Grade{gradeId=4, gradeName='四年级'}}
```





## setter注入

我们可以通过 Bean 的 setter 方法，将**属性值**注入到 Bean 的属性中。

> 即 <property>与setter() 方法搭配使用

使用 setter 注入的方式进行属性注入，大致步骤如下：

1. 在 Bean 中提供一个默认的无参构造函数（在没有其他带参构造函数的情况下，可省略），并**为所有需要注入的属性提供一个 setXxx() 方法；**
2. 在 Spring 的 XML 配置文件中，使用 <beans> 及其子元素 <bean> 对 Bean 进行定义；
3. 在 <bean> 元素内使用 <property> 元素对各个属性进行赋值。





## 短命名空间注入

上述书写繁琐，此方法可简化Spring的XML配置

如下表。

| 短命名空间 | 简化的 XML 配置                        | 说明                                     |
| ---------- | -------------------------------------- | ---------------------------------------- |
| p 命名空间 | <bean> 元素中嵌套的 <property> 元素    | 是 setter 方式属性注入的一种快捷实现方式 |
| c 命名空间 | <bean> 元素中嵌套的 <constructor> 元素 | 是构造函数属性注入的一种快捷实现方式     |

> 通过它，我们能够以 bean 属性的形式实现 setter 方式的属性注入，而不再使用嵌套的 <property> 元素

### p 命名空间注入

导入约束：

```xml
xmlns:p="http://www.springframework.org/schema/p"
```

实现属性注入：

```xml
<bean id="Bean 唯一标志符" class="包名+类名" p:普通属性="普通属性值" p:对象属性-ref="对象的引用">
```

使用 p 命名空间注入依赖时，必须注意以下 3 点：

- Java 类中**必须有 setter 方法**；
- Java 类中**必须有无参构造器**（类中不包含任何带参构造函数的情况，无参构造函数默认存在）；
- 在使用 p 命名空间实现属性注入前，XML 配置的 <beans> 元素内必须**先导入 p 命名空间的 XML 约束。**



示例：

```xml
<bean id="employee" class="net.biancheng.c.Employee" p:empName="小李" p:dept-ref="dept" p:empNo="22222"></bean>
    <bean id="dept" class="net.biancheng.c.Dept" p:deptNo="1111" p:deptName="技术部"></bean>
```



### c 命名空间注入

> 通过它，我们能够以 <bean> 属性的形式实现构造函数方式的属性注入，而不再使用嵌套的 <constructor-arg> 元素

1. 导入约束

   ```xml
   xmlns:c="http://www.springframework.org/schema/c"
   ```

2. 属性注入：

   ```xml
   <bean id="Bean 唯一标志符" class="包名+类名" c:普通属性="普通属性值" c:对象属性-ref="对象的引用">
   ```

   

使用 p 命名空间注入依赖时，必须注意以下 2 点：

- Java 类中必须包含对应的**带参构造器**；
- 在使用 c 命名空间实现属性注入前，XML 配置的 <beans> 元素内必须**先导入 c 命名空间的 XML 约束。**

