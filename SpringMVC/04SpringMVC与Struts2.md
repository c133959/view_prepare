#### 1. Spring MVC 基于方法开发，Struts2 基于类开发。

在使用 Spring MVC 框架进行开发时，会将 URL 请求路径与 Controller 类的某个方法进行绑定，请求参数作为该方法的形参。当用户请求该 URL 路径时， **Spring MVC 会将 URL 信息与 Controller 类的某个方法进行映射，生成 Handler 对象**，该对象中只包含了一个 method 方法。方法执行结束之后，形参数据也会被销毁。

而在使用 Struts2 框架进行开发时，Action 类中所有方法使用的请求参数都是 Action 类中的成员变量，随着方法变得越来越多，就很难分清楚 Action 中那么多的成员变量到底是给哪一个方法使用的，整个 Action 类会变得十分混乱。

相比较而言，**Spring MVC 优点是其所有请求参数都会被定义为相应方法的形参，用户在网页上的请求路径会被映射到 Controller 类对应的方法上，此时请求参数会注入到对应方法的形参上。Spring MVC 的这种开发方式类似于 Service 开发。**

#### 2. Spring MVC 可以进行单例开发，Struts2 无法使用单例

Spring MVC 支持单例开发模式，而 Struts2 由于只能通过类的成员变量接受参数，所以无法使用单例模式，只能使用多例。

#### 3. 经过专业人员的大量测试，Struts2 的处理速度要比 SpringMVC 慢，原因是 Struts2 使用了 Struts 标签，Struts 标签由于设计原因，会出现加载数据慢的情况

这里仅仅比较了 Spring MVC 在某些方面相比 Struts2 的优势，但这并不能说明 Spring MVC 比 Struts2 优秀，仅仅因为早期 Struts2 使用广泛，所以出现的漏洞也比较多，但是在新版本的 Struts2 中也修复了许多漏洞。Spring MVC 自诞生以来，几乎没有什么致命的漏洞 。且 Spring MVC 基于方法开发，这一点较接近 Service 开发，这也是 Spring MVC 近年来备受关注的原因之一。