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

##### 而 MVCC 使用的是其中的事务字段，回滚指针字段，是否删除字段。

| isDelete   | DB_TRX_ID | DB_ROLL_PTR | id   | name | password |
| ---------- | --------- | ----------- | ---- | ---- | -------- |
| true/false | 事务id    | 回滚指针    | id   | name | password |

那么如何通过这三个字段来实现 MVCC 的 可见性算法呢？ 还差点东西！undoLog(回滚日志) read-view(读视图)。

- undoLog: 事务的回滚日志，是可见性算法 的非常重要的部分，分为两类。
  - insert undo log：事务在插入新记录时产生的undo log，当事务提交之后可以直接丢弃
  - update undo log：事务在进行 update 或者 delete 的时候产生的 undo log，在快照读的时候还是需要的，所以不能直接删除，只有当系统没有比这个log更早的read-view了的时候才能删除。ps：所以长事务会产生很多老的视图导致undo log无法删除 大量占用存储空间。
- read-view: 读视图，是MySQL秒级创建视图的必要条件，比如一个事务在进行 select 操作(快照读)的时候会创建一个 read-view ，这个read-view 其实只是三个字段。
  - alive_trx_list(我自己取的)：read-view生成时刻系统中正在活跃的事务id。
  - up_limit_id：记录上面的 alive_trx_list 中的最小事务id。
  - low_limit_id：read-view生成时刻，目前已出现的事务ID的最大值 + 1。

可见性算法。 其实主要思路就是：当生成read-view的时候如何去拿获取的 DB_TRX_ID 去和 read-view 中的三个属性(上面讲了)去作比较。我来说一下三个步骤，如果不是很理解可以参考着我后面的实践结合着去理解。

- 首先比较这条记录的 DB_TRX_ID 是否是 小于 up_limit_id 或者 等于当前事务id。如果满足，那么说明当前事务能看到这条记录。如果大于则进入下一轮判断
- 然后判断这条记录的 DB_TRX_ID 是否 大于等于 low-limit-id。如果大于等于则说明此事务无法看见该条记录，不然就进入下一轮判断。
- 判断该条记录的 DB_TRX_ID 是否在活跃事务的数组中，如果在则说明这条记录还未提交对于当前操作的事务是不可见的，如果不在则说明已经提交，那么就是可见的。