# spring-boot-starter-web（web启动器）

> Spring Boot 是在 Spring 的基础上创建一款开源框架，它提供了 spring-boot-starter-web（Web 场景启动器） 来为 Web 开发予以支持。spring-boot-starter-web 为我们提供了嵌入的 Servlet 容器以及 SpringMVC 的依赖，并为 Spring MVC 提供了大量自动配置，可以适用于大多数 Web 开发场景。

## Spring Boot Web快速开发

Spring Boot 为 Spring MVC 提供了自动配置，并在 Spring MVC 默认功能的基础上添加了以下特性：

- 引入了 ContentNegotiatingViewResolver 和 BeanNameViewResolver（视图解析器）
- 对包括 WebJars 在内的静态资源的支持
- 自动注册 Converter、GenericConverter 和 Formatter （转换器和格式化器）
- 对 HttpMessageConverters 的支持（Spring MVC 中用于转换 HTTP 请求和响应的消息转换器）
- 自动注册 MessageCodesResolver（用于定义错误代码生成规则）
- 支持对静态首页（index.html）的访问
- 自动使用 ConfigurableWebBindingInitializer


只要我们在 Spring Boot 项目中的 **pom.xml** 中引入了 **spring-boot-starter-we**b ，即使不进行任何配置，也可以直接使用 Spring MVC 进行 Web 开发。

## 实例

1. 创建一个名为 spring-boot-springmvc-demo1 的 Spring Boot 工程，并在其 pom.xml 的dependencies 节点中添加 spring-boot-starter-web 的依赖，代码如下。

   ```java
   <dependency>
       <groupId>org.springframework.boot</groupId>
       <artifactId>spring-boot-starter-web</artifactId>
   </dependency>
   ```

2. 创建一个HelloController

   ```java
   import org.springframework.stereotype.Controller;
   import org.springframework.web.bind.annotation.RequestMapping;
   import org.springframework.web.bind.annotation.ResponseBody;
   @Controller
   public class HelloController {
       @ResponseBody
       @RequestMapping("/hello")
       public String hello() {
           return "hello";
       }
   }
   ```

   