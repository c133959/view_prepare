# Spring Boot定制Spring MVC

> Spring Boot 抛弃了传统 xml 配置文件，通过配置类（标注 **@Configuration** 的类，相当于一个 xml 配置文件）以 JavaBean 形式进行相关配置。

Spring Boot 对 Spring MVC 的自动配置可以满足我们的大部分需求，但是我们也可以通过

1. 自定义配置类（标注 @Configuration 的类）并

2. 实现 **WebMvcConfigurer** 接口来定制 Spring MVC 配置，例如拦截器、格式化程序、视图控制器等等。

| efault void addInterceptors(InterceptorRegistry registry) {} | 添加 Spring MVC 生命周期拦截器，对请求进行拦截处理。         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| default void addResourceHandlers(ResourceHandlerRegistry registry) {} | 添加或修改静态资源（例如图片，js，css 等）映射； Spring Boot 默认设置的静态资源文件夹就是通过重写该方法设置的。 |
| default void addViewControllers(ViewControllerRegistry registry) {} | 主要用于实现无业务逻辑跳转，例如主页跳转，简单的请求重定向，错误页跳转等 |

在 Spring Boot 项目中，我们可以通过以下 2 中形式定制 Spring MVC:

- 扩展 Spring MVC
- 全面接管 Spring MVC

