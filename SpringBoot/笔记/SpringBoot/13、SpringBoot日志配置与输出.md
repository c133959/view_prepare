# Spring Boot日志配置及输出

## 默认配置

> Spring Boot 默认使用 SLF4J+Logback 记录日志，并提供了默认配置，即使我们不进行任何额外配，也可以使用 SLF4J+Logback 进行日志输出。

常见的日志配置包括**日志级别**、**日志的输入出格式**等内容。

### 日志级别

日志的输出都是分级别的，当一条日志信息的级别**大于或等于**配置文件的级别时，就对这条日志进行记录。

常见的日志级别如下（优先级依次升高）。

| 序号 | 日志级别 | 说明                                                        |
| ---- | -------- | ----------------------------------------------------------- |
| 1    | trace    | 追踪，指明程序运行轨迹。                                    |
| 2    | debug    | 调试，实际应用中一般将其作为最低级别，而 trace 则很少使用。 |
| 3    | info     | 输出重要的信息，使用较多。                                  |
| 4    | warn     | 警告，使用较多。                                            |
| 5    | error    | 错误信息，使用较多。                                        |

### 输出格式

我们可以通过以下常用日志参数对日志的输出格式进行修改，如下表。

| 序号 | 输出格式                     | 说明                                                         |
| ---- | ---------------------------- | ------------------------------------------------------------ |
| 1    | %d{yyyy-MM-dd HH:mm:ss, SSS} | 日志生产时间,输出到毫秒的时间                                |
| 2    | %-5level                     | 输出日志级别，-5 表示左对齐并且固定输出 5 个字符，如果不足在右边补 0 |
| 3    | %logger 或 %c                | logger 的名称                                                |
| 4    | %thread 或 %t                | 输出当前线程名称                                             |
| 5    | %p                           | 日志输出格式                                                 |
| 6    | %message 或 %msg 或 %m       | 日志内容，即 logger.info("message")                          |
| 7    | %n                           | 换行符                                                       |
| 8    | %class 或 %C                 | 输出 Java 类名                                               |
| 9    | %file 或 %F                  | 输出文件名                                                   |
| 10   | %L                           | 输出错误行号                                                 |
| 11   | %method 或 %M                | 输出方法名                                                   |
| 12   | %l                           | 输出语句所在的行数, 包括类名、方法名、文件名、行数           |
| 13   | hostName                     | 本地机器名                                                   |
| 14   | hostAddress                  | 本地 ip 地址                                                 |

### 实例

**Spring Boot 日志默认级别为 info**，日志输出内容默认包含以下元素：

- 时间日期
- 日志级别
- 进程 ID
- 分隔符：---
- 线程名：方括号括起来（可能会截断控制台输出）
- Logger 名称
- 日志内容

### 修改默认日志配置

> 我们可以根据自身的需求，通过全局配置文件（application.properties/yml）修改 Spring Boot 日志级别和显示格式等默认配置。

在 **application.properties** 中，修改 Spring Boot 日志的默认配置，代码如下。

```python
#日志级别
logging.level.net.biancheng.www=trace
#使用相对路径的方式设置日志输出的位置（项目根目录目录\my-log\mylog\spring.log）
#logging.file.path=my-log/myLog
#绝对路径方式将日志文件输出到 【项目所在磁盘根目录\springboot\logging\my\spring.log】
logging.file.path=/spring-boot/logging
#控制台日志输出格式
logging.pattern.console=%d{yyyy-MM-dd hh:mm:ss} [%thread] %-5level %logger{50} - %msg%n
#日志文件输出格式
logging.pattern.file=%d{yyyy-MM-dd} === [%thread] === %-5level === %logger{50} === - %msg%n
```

