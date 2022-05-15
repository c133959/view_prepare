# bean的作用域

​	在Spring中，可以在<bean>元素的scope属性里设置bean的作用域，以决定这个bean是单实例还是多实例



​	默认情况下，Spring只为每个在IOC容器里声明的bean创建**唯一实例**，整个IOC容器范围内都可以共享该实例：所有的后续的getbean()和bean 引用都将返回这个唯一的bean实例。**该作用域称为singleton，他是所有bean的默认作用域**



​	![image-20220210214607555](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220210214607555.png)

## 总结：

singleton：默认值。单例

prototype：原型的。当ICO容器一创建，就不会再实例化该bean，调用getbean()方法时再实例化

request：每次请求实例化一个bean

session：在一次会话中共享一个bean