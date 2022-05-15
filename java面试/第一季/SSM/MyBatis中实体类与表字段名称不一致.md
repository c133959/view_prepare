# MyBatis中实体类与表字段名称不一致



##  方法1 ： 写SQL语句时给字段加上别名

```sql
select id, user_name userName, # 此处给数据库中的user_name加上别名，与实体类中的名称一致
gender
from tbl_user
```



## 方法2 ： 在Mybatis全局配置文件中开启驼峰命名规则

```xml
<settings>
	<!--开启驼峰命名规则，可以将数据库中的下划线命名的字段映射为驼峰命名-->
    <setting name="mapUnderscoreToCamelCase" value="true"/>
</settings>
```



## 方法3 ： 在Mapper映射文件中使用resultMap来定义映射规则

```xml
<!--自定义高级映射-->
<resultMap id="RelationshipMap" type="com.sitech.irms.transfer.po.ProjectQuantitiesRrPO">
    <!--映射主键-->    
    <id column="id" jdbcType="DECIMAL" property="id" />
    <result column="project_id" jdbcType="DECIMAL" property="projectId" />
</resultMap>
```



## 方法4 ： SpringBoot @Column()注解

```java
@Table(name = "rr_tsf_checkpoint_project")
public class RrTsfCheckpointProject {
    /**
     * ID
     */
    @Id
    private BigDecimal id;

    /**
     * 检查点id
     */
    @Column(name = "checkpoint_id")
    private BigDecimal checkpointId;
}
```

