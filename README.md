# go面试资料整理

### go语言基础
熟悉语法，撸个百十道基础面试题就差不多了。

### go语言进阶
什么是CSP？

channel底层的数据结构是什么？发送和接收元素的本质是什么？

channel在哪些情况下会死锁/阻塞？

那些类型不能作map的为key？map的key为什么是无序的？

map是线程安全的么？

map的底层实现原理是什么？

map的扩容过程是怎样的？

map的赋值过程和查找过程是怎样的？

iface和eface的区别是什么？值接收者和指针接收者的区别？

接口的构造过程是怎样的？

context是什么？有什么作用？如何被取消？

go中的指针有什么限制？

GPM是什么？goroutine的调度时机有哪些？如果syscall阻塞会发生什么？

slice的底层数据结构是怎样的？

你了解go的GC么？常见的GC实现方式有哪些？

什么是三色标记法和混合写屏障？

go的GC流程是什么？如果内存分配速度超过了标记清除速度怎么办？

内存泄漏是如何发生的，如何解决？

内存逃逸分析是怎么进行的？

### mysql
索引底层实现？为什么选择B+树作为索引结构？B+树的叶子节点都可以存哪些东西？

覆盖索引是什么？回表？

什么情况下不会命中索引？

mvcc的实现原理是什么？

redolog、undolog、binlog作用是什么？有什么区别？

乐观锁和悲观锁的实现方式？

事务隔离级别有哪些？如何解决脏读和幻读？

应该如何恰当的建立索引？

mysql的存储引擎有哪些？都有什么区别？

mysql的主从复制？

### 分布式系统、微服务架构
什么是分布式事务？

分布式事务解决方案有哪些？

分布式锁实现原理？有哪些实现方式？

微服务架构设计？

### kafka
kafka为什么性能高？

kafka重复消费可能的原因以及处理方式？

kafka消息丢失的原因以及解决方式？

kafka如何保证消息的顺序性？

什么是kafka的Rebalance？

kafka集群消息积压问题如何处理？

### redis
redis与memcached的区别？

redis单线程为什么效率也这么高？

redis有那五种常用的数据结构？应用场景以及实现原理是什么？

redis如何实现延迟队列？

redis的过期策略？

redis集群有那三种模式？

redis的事务？

redis与mysql双写一致性解决方案？

发生缓存穿透、击穿、雪崩的原因以及解决方案？

布隆过滤器的了解？

### 网络协议
浏览器访问一个网站，都经历了怎样一个流程？

什么是HTTP协议？

什么是HTTP报文？HTTP报文由那三部分组成？

HTTP常见的状态码有哪些？

HTTPS协议的底层原理是什么？

TCP协议和UDP协议有什么区别？

TCP协议的三次握手和四次挥手？为什么是三次和四次？

### 负载均衡
暂无

### 数据结构与算法
// 题目：将6，2，10，32，9，5，18，14，30，29从小到大进行排列,使用冒泡排序
```go
package main

import "fmt"

func main() {
    // 定义数组
    arr := [10]int{6, 2, 10, 32, 9, 5, 18, 14, 30, 29}
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println(arr)
}
```

快排

选择排序

堆排

### 区块链业务
暂无