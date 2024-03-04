# Redis 基础

### redis五大数据结构

- String
- Hash
- List
- Set
- SortedSet

### 一、String字符串类型

​	是redis中最基础的数据类型，一个key对应一个value

​	String类型是二进制安全的，意思是redis的string可以包含任何数据，如数字，字符串，jpg图片或者序列化后的存储对象

##### 实战场景：

1. 缓存：经典使用场景，把常用信息，字符串，图片或者视频信息放到redis中，redis作为缓存层，mysql做持久层，降低mysql的读写能力
2. 计数器：redis是单线程模型，一个命令执行完才会执行下一个，同时数据可以一步落到其他数据源
3. session：常见spring session+redis实现session共享



### 二、哈希

​	