# 解决Post请求乱码问题



## 配置过滤器——Spring框架

配置filter：

![image-20220212220557411](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220212220557411.png)

配置完filter之后，还需要配置映射

```xml
<filter-mapping>
    <!--过滤器名称与filter中配置的名称一致 -->
	<filter-name>CharacterEncodingFilter</filter-name>
    <!-- 拦截所有的请求-->
    <url-patter>/*</url-patter>
</filter-mapping>
```



## Spring Boot解决

去检查你的 filter 配置，**是否配置了一个最高优先级的 filter**, 这个最高优先级的 filter 会影响 springboot 自动配置的 CharacterEncodingFilter。原因如下：

　　在tomcat里:

1. request对象的parameter并不是一开始就解析的，它是等你第一次调用`getParameter*`等获得请求参数有关的方法的时候才解析的
2. paramter一旦被解析过一次，那就不会再次被解析

　　所以如果在`CharacterEncodingFilter`之前有另外一个filter，而这个filter调用了`getParameter*`方法，那么就有可能使用错误的encoding来解析，从而造成乱码问题。

方法：**修改你自己配置的 filter的优先级，至少比 CharacterEncodingFilter 的优先级低。**



SpringBoot中，在配置文件中加入该行即可（**SpringBoot为我们默认配置了该配置项为utf-8，实际不需要额外配置**）

```properties
server.tomcat.uri-encoding = utf-8
```



# 当请求切换为GET又变成乱码

* 最简单的方法：

  在tomcat找到`server.xml`的文件

  在connector中添加 `URIEncodeing=utf-8` 即可

  ![image-20220212221402610](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220212221402610.png)

