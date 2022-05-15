# redis持久化



## Redis持久化有几种类型？区别？

类型：

2种：

* RDB (Redis DataBase)

  > 在指定的时间间隔内将内存中的数据集快照写入磁盘，即Snapshot快照
  >
  > 它恢复时是将快照文件直接读到**内存**中
  >
  > ![image-20220213194144930](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220213194144930.png)
  >
  > ![image-20220213194308225](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220213194308225.png)
  >
  > 说明：每次保存的都是**全量**的数据

* AOF (Append Of File)

  > 以日志的形式来记录每个读写操作
  >
  > 将Redis执行过的所有**写指令**记录下来，只许**追加文件** 不许改写文件
  >
  > Redis启动之初，会读取该文件，并按照里面的写操作来重新构建数据
  >
  > ![image-20220213194811447](C:\Users\26442\AppData\Roaming\Typora\typora-user-images\image-20220213194811447.png)
  >
  > 说明：AOF记录每个写操作，RDB是隔一段时间保存一次，明显AOF的粒度更小，所以记录的内容会更多，从而更占用磁盘

