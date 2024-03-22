# RocketMQ基础

### 一、描述一下RocketMQ

###### 	Java 开发，⾯向互联⽹集群化功能丰富，对在线业务的响应时延做了很多的优化，⼤多数情况下可以做到毫秒级的响应，每秒钟⼤概能处理⼏⼗万条消息。

### 二、why RocketMQ

1. ##### 解耦：降低系统耦合度，减少依赖关系，通过消息传递业务

2. ##### 异步：提高系统响应效率

3. ##### 削峰：请求达到瓶颈后，后端服务还可以保持固定消费速率消费，不会压垮

### 三、RocketMQ组成

| 角色       | 作用                                                         |
| ---------- | ------------------------------------------------------------ |
| Nameserver | ⽆状态，动态列表；这是和 ZooKeeper 的重要区别之⼀。ZooKeeper 是有状态的 |
| Producer   | 消息⽣产者，负责发消息到 Broker。                            |
| Broker     | 就是 MQ 本身，负责收发消息、持久化消息等。                   |
| Consumer   | 消息消费者，负责从 Broker 上拉取消息进⾏消费，消费完进⾏ ack。 |

### 四、RocketMQ中的topic和jms的queue有什么区别？

​	jms（java message service） queue就是来源于数据结构的FIFO队列，而topic是个抽象的概念，每个topic底层对应n个queue，而数据也真实存在queue上。

### 五、RocketMQ消费模式有哪几种

​	消费模型由consumer决定，消费维度为topic

- 集群模式（多节点服务，只有一个节点会消费，不会重复消费）
  1. 一条消息只会被同group中一个consumer消费
  2. 多个group同时消费一个topic时，每个group都会有一个consumer消费到数据
- 广播模式（多节点服务，每个节点都会消费，重复消费）
  - 消息将对一个consumer group下的各个consumer实例都消费一遍，即使这些consumer属于同一个consumer group。消息也会被consumer group中每个consumer都消费一次

### 六、RocketMQ怎么保证消息不丢失

​	首先在如下三个部分，可能会出现丢失

1. producer端
   - 采取send()同步发送消息，发送结果是感知的
   - 发送失败后可以重试，设置重试次数，默认三次。
2. broker端
   - 修改刷盘策略为同步刷盘，默认情况下是异步刷盘的，FLUSHDISKTYPE = SYNC_FLUSH
   - 集群部署，主从模式，高可用
3. consumer端
   - 完全消费正常后再进行手动ack确认

### 七、Broker如何处理拉取请求？

​	consumer首次请求broker

- broker中是否有符合条件的消息
  - 有
    - 响应consumer，等待下次consumer的请求
  - 没有
    - DefaultMessageStore#ReputMessageServoce#Run方法-PullRequestHoldService来hold连接，每隔五秒执行一次检测pullRequestTable有没有消息，有的话立即推送
    - 每隔1ms检查commitLog中是否有消息，有的话写入到pullRequestTable
    - 当有新消息的时候返回请求
    - 挂起consumer的请求，即不断开连接，也不返回数据
    - 使用consumer的offset

### 八、重复消费问题

​	影响消息正常发送和消费的重要原因是网络的不确定性

​	重复消费原因：

- ack
  - 正常情况下consumer真正消费完消息后应该发ack，通知broker，该消息已经正常消费，从queue中移除。但是当ack因为网络原因无法发送到broker，broker会认为词条消息没消费掉，此后会开启消息重投机制，把消息再次投递到consumer
- 消费模式
  - 在clustering模式下，消息在broker中会保证相同group的consumer消费一次，但是针对不同的group的consumer会多推送几次

​	解决方案

- 数据库表
  - 处理消息前，使用消息主键在表中带有约束的字段中insert
- map
  - 单机时可以使用map concurrentHashMap->putifAbsent guava cache-redis分布式锁

### 九、高吞吐量下如何优化生产者和消费者的性能

- 同一group下，多机部署，并行消费（job服务做成多节点）

- 单个consumer提高消费线程数，比如设置10个协程一起消费

  ```
  // GOMAXPROCS set max goroutine to work.
  func (g *Group) GOMAXPROCS(n int) {
  	if n <= 0 {
  		panic("errgroup: GOMAXPROCS must great than 0")
  	}
  	g.workerOnce.Do(func() {
  		g.ch = make(chan func(context.Context) error, n)
  		for i := 0; i < n; i++ {
  			go func() {
  				for f := range g.ch {
  					g.do(f)
  				}
  			}()
  		}
  	})
  }
  ```

- 消息生产时，批量处理，比如应用市场上报，攒够10条再发送消息

### 十、如何防止消息丢了？

1. 一般是使用三节点部署
2. 生产者发送消息时，如果失败了，重试三次，三次后还失败，就要打log，log多了就得报警了
