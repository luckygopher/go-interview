# go面试资料整理

### go语言基础
熟悉语法，撸个百十道基础面试题就差不多了。

### go语言进阶
gmp，map，slice，channel，wg，gc，内存泄漏，context包，内存逃逸分析

### mysql
索引底层实现，btree+树，回表，事物隔离级别（脏读，不可重复读，幻读），mvcc的实现原理，binlog，undo log，redo log都是用来干什么的，
日志先行，刷脏页，悲观锁的实现方式

### Kafka
partition分区，一致性算法

### redis
5种常用的数据结构实现，如何防止雪崩，击穿，布隆过滤器，底层实现bitmap，setnx，redLock，事务（watch）

### 网络协议
重点是http和tcp，3次握手和4次挥手，二者timeout区别，tcp里的timewait多了怎么办，更狠点的会问为什么要设计这个timewait状态

### 排序算法
基本没碰到问这个的

### 分布式事务
特别是分布式锁，两阶段提交

### 微服务架构
acid，base，cap

### 负载均衡

### 区块链业务
也没遇到这个