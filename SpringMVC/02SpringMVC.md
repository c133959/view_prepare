# SpringMVC

Spring MVC 是 Spring 提供的一个基于 MVC 设计模式的轻量级 Web 开发框架，本质上相当于 Servlet。

在 Spring MVC 框架中**，Controller 替换 Servlet 来担负控制器的职责，用于接收请求，调用相应的 Model 进行处理，处理器完成业务处理后返回处理结果。Controller 调用相应的 View 并对处理结果进行视图渲染，最终客户端得到响应信息。**



# SpringMVC优点

- 清晰地角色划分，Spring MVC 在 Model、View 和 Controller 方面提供了一个非常清晰的角色划分，这 3 个方面真正是各司其职，各负其责。
- 灵活的配置功能，可以把类当作 Bean 通过 XML 进行配置。
- 提供了大量的控制器接口和实现类，开发者可以使用 Spring 提供的控制器实现类，也可以自己实现控制器接口。
- 真正做到与 View 层的实现无关。它不会强制开发者使用 JSP，可以根据项目需求使用 Velocity、FreeMarker 等技术。
- 国际化支持
- 面向接口编程
- 与 Spring 框架无缝集成

