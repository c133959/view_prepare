# Object类

> Object类是所有类的父类：也就是说，Java 允许把任何类型的对象赋给 Object 类型的变量
>
> 当一个类被定义后，如果没有指定继承的父类，那么默认父类就是 Object 类

由于 Java 所有的类都是 Object 类的子类，所以任何 Java 对象都可以调用 Object 类的方法。常见的方法如表 1 所示。

| 方法                       | 说明                                                         |
| -------------------------- | ------------------------------------------------------------ |
| Object clone()             | 创建与该对象的类相同的新对象                                 |
| boolean **equals**(Object) | 比较两**对象****是否相等**                                   |
| void finalize()            | 当垃圾回收器确定不存在对该对象的更多引用时，对象垃圾回收器调用该方法 |
| Class **getClass**()       | 返回一个对象**运行时的实例类**                               |
| int hashCode()             | 返回该对象的散列码值                                         |
| void notify()              | 激活等待在该对象的监视器上的一个线程                         |
| void notifyAll()           | 激活等待在该对象的监视器上的全部线程                         |
| String **toString**()      | **返回该对象的字符串表示**                                   |
| void wait()                | 在其他线程调用此对象的 notify() 方法或 notifyAll() 方法前，导致当前线程等待 |


其中，toString()、equals() 方法和 getClass() 方法在 Java 程序中比较常用。



## toString()

toString() 方法返回该对象的字符串，当程序输出一个对象或者把某个对象和字符串进行连接运算时，系统会自动调用该对象的 toString() 方法返回该对象的字符串表示。

Object 类的 toString() 方法返回“**运行时类名@十六进制哈希码**”格式的字符串，但很多类都重写了 Object 类的 toString() 方法，用于返回可以表述该对象信息的字符串。

> 因此对于一些System.out.println(对象实例) ，可以重写toString()，来输出定制的信息 

```java
public class Person {
    private String name;
    private int age;
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
    public String toString() {
        return "姓名：" + this.name + "：年龄" + this.age;
    }
    public static void main(String[] args) {
        Person per = new Person("C语言中文网", 30);// 实例化Person对象
        System.out.println("对象信息：" + per);// 打印对象调用toString()方法
    }
}
```



## equals()

`==`运算符是比较两个引用变量**是否指向同一个实例**，

`equals() `方法是比较两个**对象的内容**是否相等



## getClass()

getClass() 方法返回对象所属的类，是一个 Class 对象。

通过 Class 对象可以获取该类的各种信息，包括类名、父类以及它所实现接口的名字等。

```java
public class Test02 {
    public static void printClassInfo(Object obj) {
        // 获取类名
        System.out.println("类名：" + obj.getClass().getName());
        // 获取父类名
        System.out.println("父类：" + obj.getClass().getSuperclass().getName());
        System.out.println("实现的接口有：");
        // 获取实现的接口并输出
        for (int i = 0; i < obj.getClass().getInterfaces().length; i++) {
            System.out.println(obj.getClass().getInterfaces()[i]);
        }
    }
    // 获取String类的信息
    public static void main(String[] args) {
        String strObj = new String();
        printClassInfo(strObj);
    }
}


类名：java.lang.String
父类：java.lang.Object
实现的接口有：
interface java.io.Serializable
interface java.lang.Comparable
interface java.lang.CharSequence
```

## 接收任意引用类型的对象

既然 Object 类是所有对象的父类，则所有的对象都可以向 Object 进行转换，在这其中也包含了数组和接口类型，即一切的引用数据类型都可以使用 Object 进行接收。

```java
interface A {
    public String getInfo();
}
class B implements A {
    public String getInfo() {
        return "Hello World!!!";
    }
}
public class ObjectDemo04 {
    public static void main(String[] args) {
        // 为接口实例化
        A a = new B();
        // 对象向上转型
        Object obj = a;
        // 对象向下转型
        A x = (A) obj;
        System.out.println(x.getInfo());
    }
}
```

输出结果为：

> Hello World!!!

通过以上代码可以发现，虽然**接口不能继承一个类，但是依然是 Object 类的子类**，

因为**接口本身是引用数据类型，所以可以进行向上转型操作。**



同理，也可以使用 Object 接收一个数组，因为**数组本身也是引用数据类型。**

```java
public class ObjectDemo05 {
    public static void main(String[] args) {
        int temp[] = { 1, 3, 5, 7, 9 };
        // 使用object接收数组
        Object obj = temp;
        // 传递数组引用
        print(obj);
    }
    public static void print(Object o) {
        // 判断对象的类型
        if (o instanceof int[]) {
            // 向下转型
            int x[] = (int[]) o;
            // 循环输出
            for (int i = 0; i < x.length; i++) {
                System.out.print(x[i] + "\t");
            }
        }
    }
}
```

输出结果为：

>  1 3 5 7 9

以上程序使用 Object 接收一个整型数组，因为数组本身属于引用数据类型，所以可以使用 Object 接收数组内容，在输出时通过 instanceof 判断类型是否是一个整型数组，然后进行输出操作。

**提示：因为 Object 类可以接收任意的引用数据类型，所以在很多的类库设计上都采用 Object 作为方法的参数，这样操作起来会比较方便。**

