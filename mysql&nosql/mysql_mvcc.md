# MySql的MVCC原理

### 一、MVCC 的实现原理

##### 	MVCC 使用了“三个隐藏字段”来实现版本并发控制

| RowID        | DB_TRX_ID | DB_ROLL_PTR | id   | name | password |
| ------------ | --------- | ----------- | ---- | ---- | -------- |
| 自动创建的id | 事务id    | 回滚指针    | id   | name | password |

- RowID：隐藏的自增ID，当建表没有指定主键，InnoDB会使用该RowID创建一个聚簇索引。
- DB_TRX_ID：最近修改（更新/删除/插入）该记录的事务ID。
- DB_ROLL_PTR：回滚指针，指向这条记录的上一个版本。

> 还有一个删除的flag字段，用来判断该行记录是否已经被删除。