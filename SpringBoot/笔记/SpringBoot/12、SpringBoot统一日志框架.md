# Spring Boot统一日志框架

> 在项目开发中，日志十分的重要，不管是记录运行情况还是定位线上问题，都离不开对日志的分析。在 Java 领域里存在着多种日志框架，如 JCL、SLF4J、Jboss-logging、jUL、log4j、log4j2、logback 等等。

## 日志框架的选择

分类：日志门面（日志抽象层）和日志实现

| 日志分类               | 描述                                                         | 举例                                                         |
| ---------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 日志门面（日志抽象层） | 为 Java 日志访问提供一套标准和规范的 API 框架，其主要意义在于**提供接口**。 | JCL（Jakarta Commons Logging）、SLF4j（Simple Logging Facade for Java）、jboss-logging |
| 日志实现               | 日志门面的**具体的实现**                                     | **Log4j**、JUL（java.util.logging）、Log4j2、Logback         |

Spring Boot 选用 SLF4J + Logback 的组合来搭建日志系统。

SLF4J ：简化代码、可读性强

## SLF4J 的使用

> 在项目开发中，记录日志时不应该直接调用日志实现层的方法，而应该调用日志门面（日志抽象层）的方法。

```java
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HelloWorld {
    public static void main(String[] args) {
        Logger logger = LoggerFactory.getLogger(HelloWorld.class);
       //调用 sl4j 的 info() 方法，而非调用 logback 的方法
        logger.info("Hello World");
    }
}
```

这里我们需要注意一点，每一个日志的实现框架都有自己的配置文件。使用 slf4j 记录日志时，配置文件应该使用**日志实现框架（例如 logback、log4j 和 JUL 等等）自己本身的配置文件。**

## 统一日志框架（通用）

> 通常一个完整的应用下会依赖于多种不同的框架，而且它们记录日志使用的日志框架也不尽相同，例如，Spring Boot（slf4j+logback），Spring（commons-logging）、Hibernate（jboss-logging）等等。那么如何统一日志框架的使用呢？

1. 排除应用中的原来的日志框架；
2. 引入替换包替换被排除的日志框架；
3. 导入 SLF4J 实现。

## 统一日志框架（Spring Boot）

> 我们在使用 Spring Boot 时，同样可能用到其他的框架，例如 Mybatis、Spring MVC、 Hibernate 等等，这些框架的底层都有自己的日志框架，此时我们也需要对日志框架进行统一。

我们知道，统一日志框架的使用一共分为 3 步，Soring Boot 作为一款优秀的开箱即用的框架，已经为用户完成了其中 2 步：**2.引入替换包**和**3.导入 SLF4J 实现。**


Spring Boot 的核心启动器 spring-boot-starter 引入了 spring-boot-starter-logging，使用 IDEA 查看其依赖关系，如下图。

![Spring Boot 依赖关系图](http://c.biancheng.net/uploads/allimg/210726/154Z34204-2.png)

SpringBoot 底层使用 slf4j+logback 的方式记录日志，**当我们引入了依赖了其他日志框架的第三方框架（例如 Hibernate）时，只需要把这个框架所依赖的日志框架排除**，即可实现日志框架的统一，示例代码如下。

```html
<dependency>
    <groupId>org.apache.activemq</groupId>
    <artifactId>activemq-console</artifactId>
    <version>${activemq.version}</version>
    <exclusions>
        <exclusion>
            <groupId>commons-logging</groupId>
            <artifactId>commons-logging</artifactId>
        </exclusion>
    </exclusions>
</dependency>
```

