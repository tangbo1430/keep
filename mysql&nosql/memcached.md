# Memcached基础

### 一、memcached是什么，有什么作用？

- 它是一个开源的，高性能的内存缓存软件，从名称上看mem就是内存的意思，cache是缓存的意思。
- 作用是在事先规划好的内存空间中，临时缓存数据库中的各类数据，以达到减少业务对数据库的直接高并发访问

### 二、memcached和redis有什么区别

1. memcache只支持简单的数据类型，需要自己序列化/反序列化
2. 数据依赖于服务，服务炸了数据也没了，没有容灾功能
3. memcache的key限制250个字符，redis的key可以512k
4. redis是单线程模型，保住了数据顺序提交，memcache需要用cas保证数据一致性（乐观锁操作）
5. redis单核，memcache多核，存小数据redis性能高，存100k以上的数据，memcache性能高
6. mamcache预先申请了一块内存，避免了内存碎片化，但是不能变长，没用完的，空着的资源一直空着

### 三、memcached能够更有效的使用内存吗？

1. memcache客户端仅根据哈希算法来决定某个key存储在哪个节点上，而不考虑节点的内存大小，因此你可以在不同的节点上使用大小不等的缓存。但是一般做法是：拥有较多内存的节点上可以允许多个memcached实例，每个实例使用内存跟其他节点上的实例相同

### 四、memcached与msql的query cache相比，有什么优缺点？

1. mysql改表后，query cache会刷新，存储一个memcache item只需要很少的时间，但是当写操作很频繁时，mysql的query cache会让所有缓存都失效。
2. query cache只会缓存select的结果，并不全。memcached可以缓存更全面的数据
3. memcached资源消耗代价更低，有空闲的内存就使用，mysql的query cache容易受到mysql服务器空闲内存限制。