# go面试资料整理

### go语言基础
熟悉语法，撸个百十道基础面试题就差不多了。

### go语言进阶
CSP并发模型？
```markdown
CSP并发模型它并不关注发送消息的实体，而关注的是发送消息时使用的channel，
go语言借用了process和channel这两个概念，process表现为go里面的goroutine，
是实际并发执行的实体，每个实体之间是通过channel来进行匿名传递消息使之解藕，
从而达到通讯来实现数据共享。

不要通过共享内存来通信，而要通过通信来实现内存共享。

1、sync.mutex 互斥锁（获取锁和解锁可以不在同一个协程，当获取到锁之后，
未解锁，此时再次获取锁将会阻塞）
2、通过channel通信
3、sync.WaitGroup
```
GPM模型指的是什么？goroutine的调度时机有哪些？如果syscall阻塞会发生什么？
```markdown
在go中是通过channel通信来共享内存的。

G：指的是Goroutine，也就是协程，go中的协程做了优化处理，内存占用仅几kb
且调度灵活，切换成本低。
P：指的是processor,也就是处理器，感觉也可理解为协程调度器。
M：指的是thread，内核线程。

调度器的设计策略：
1、线程复用：当本线程无可运行的G时，M-P-G0会处于自旋状态，尝试从全局队列
获取G，再从其他线程绑定的P队列中偷取G，而不是销毁线程；当本线程因为G进行
系统调用阻塞时，线程会释放绑定的P队列，如果有空闲的线程可用就复用空闲的
线程，不然就创建一个新的线程来接管释放出来的P队列。
2、利用并行：GOMAXPROCS设置P的数量，最多有这么多个线程分布在多个cpu上
同时运行。
3、抢占：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go
中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死。

go func的流程：
1、创建一个G，新建的G优先保存在P的本地队列中，如果满了则会保存到全局队列中。
2、G只能运行在M中，一个M必须持有一个P，M与P时1:1关系，M会从P的本地队列
弹出一个可执行状态的G来执行。
3、一个M调度G执行的过程是一个循环机制。
4、如果G阻塞，则M也会被阻塞，runtime会把这个线程M从P摘除，再创建或者
复用其他线程来接管P队列。
5、当G、M不在被阻塞，即系统调用结束，会先尝试找会之前的P队列，如果之前
的P队列已经被其他线程接管，那么这个G会尝试获取一个空闲的P队列执行，并放
入到这个P的本地队列。否则这个线程M会变成休眠状态，加入空闲线程队列，而G
则会被放入全局队列中。

M0：
M0是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，
不需要在heap上分配，M0负责执行初始化操作和启动第一个G，之后M0与其他的M一样。
G0：
G0是每次启动一个M都会第一个创建的goroutine，G0仅负责调度，不指向任何可执行
函数，每个M都会有一个自己的G0，在调度或者系统调用时会使用G0的栈空间，全局变量
的G0是M0的。

N:1-----出现阻塞的瓶颈，无法利用多个cpu
1:1-----跟多线程/多进程模型无异，切换协程代价昂贵
M:N-----能够利用多核，过于依赖协程调度器的优化和算法


同步协作式调度
异步抢占式调度

```
channel底层的数据结构是什么？发送和接收元素的本质是什么？
```go
type hchan struct {
    qcount   uint           // *chan里元素数量
    dataqsiz uint           // *底层循环数组的长度，就是chan的容量
    buf      unsafe.Pointer // *指向大小为dataqsiz的数组，有缓冲的channel
    elemsize uint16         // chan中的元素大小
    closed   uint32         // chan是否被关闭的标志
    elemtype *_type         // chan中元素类型
    recvx    uint           // *当前可以接收的元素在底层数组索引(<-chan)
    sendx    uint           // *当前可以发送的元素在底层数组索引(chan<-)
    recvq    waitq          // 等待接收的协程队列(<-chan)
    sendq    waitq          // 等待发送的协程队列(chan<-)
    lock     mutex          // 互斥锁,保证每个读chan或者写chan的操作都是原子的
}

// waitq是sudog的一个双向链表，sudog实际上是对goroutine的一个封装。
type waitq struct {
	first *sudog
	last  *sudog
}

// channel的发送和接收操作本质上都是"值的拷贝"(并不是将指针"发送"到了chan里面，
// 只是拷贝它的值而已)，无论是从sender goroutine的栈到chan buf，还是
// 从chan buf到receiver goroutine，或者是直接从sender goroutine到receiver goroutine。

```
channel在哪些情况下会死锁/阻塞？
```markdown
1、一个无缓冲channel在一个主go程里同时进行读和写；
2、无缓冲channel在go程开启之前使用通道；
3、通道1中调用了通道2，通道2中调用了通道1；
4、读取空的channel；
5、超过channel缓存继续写入数据；
6、向已经关闭的channel中写入数据不会导致死锁，但会Panic异常。
```
那些类型不能作map的为key？map的key为什么是无序的？
```markdown
map的key必须可以比较，func、map、slice这三种类型不可比较，
只有在都是nil的情况下，才可与nil (== or !=)。因此这三种类型
不能作为map的key。

数组或者结构体能够作为key？？？？有些能，有些不能，要看字段或者元素是否可比较

1、map在扩容后，会发生key的搬迁，原来落在同一个bucket中的key可能分散，key的位置发生了变化。
2、go中遍历map时，并不是固定从0号bucket开始遍历，每次都是从一个随机值序号的bucket开始遍历，
并且是从这个bucket的一个随机序号的cell开始遍历。
3、哈希查找表用一个哈希函数将key分配到不同的bucket(数组的下标index)。不同的哈希函数实现也
会导致map无序。

"迭代map的结果是无序的"这个特性是从go1.0开始加入的。
```
如何解决哈希查找表存在的"碰撞"问题（hash冲突）？
```markdown
hash碰撞指的是：两个不同的原始值被哈希之后的结果相同，也就是不同的key被哈希分配到了同一个bucket。

链表法：将一个bucket实现成一个链表，落在同一个bucket中的key都会插入这个链表。

开放地址法：碰撞发生后，从冲突的下标处开始往后探测，到达数组末尾时，从数组开始处探测，直到找到一个
空位置存储这个key，当找不到位置的情况下会触发扩容。
```
map是线程安全的么？
```markdown
map不是线程安全的，sync.map是线程安全的。

在查找、赋值、遍历、删除的过程中都会检测写标志，一旦发现写标志"置位"等于1，则直接panic,
因为这表示有其他协程同时在进行写操作。赋值和删除函数在检测完写标志是"复位"之后，先将
写标志位"置位"，才会进行之后的操作。

思考：为什么sync.map为啥是线程安全？？
```
map的底层实现原理是什么？
```markdown
type hmap struct {
    count      int   // len(map)元素个数
    flags      uint8 //写标志位
    B          uint8 // buckets数组的长度的对数，buckets数组的长度是2^B
    noverflow  uint16
    hash0      uint32
    buckets    unsafe.Pointer // 指向buckets数组
    oldbuckets unsafe.Pointer // 扩容的时候，buckets长度会是oldbuckets的两倍
    nevacuate  uintptr
    extra      *mapextra
}

// 编译期间动态创建的bmap
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}

在go中map是数组存储的，采用的是哈希查找表，通过哈希函数将key分配到不同的bucket，
每个数组下标处存储的是一个bucket，每个bucket中可以存储8个kv键值对，当每个bucket
存储的kv对到达8个之后，会通过overflow指针指向一个新的bucket，从而形成一个链表。
```
map的扩容过程是怎样的？
```markdown
相同容量扩容
2倍容量扩容

扩容时机:
1、当装载因子超过6.5时，表明很多桶都快满了，查找和插入效率都变低了，触发扩容。

扩容策略：元素太多，bucket数量少，则将B加1，buctet最大数量(2^B)直接变为
原来bucket数量的2倍，再渐进式的把key/value迁移到新的内存地址。

2、无法触发条件1，overflow bucket数量太多，查找、插入效率低，触发扩容。
(可以理解为：一座空城，房子很多，但是住户很少，都分散了，找起人来很困难)

扩容策略：开辟一个新的bucket空间，将老bucket中的元素移动到新bucket，使得
同一个bucket中的key排列更紧密，节省空间，提高bucket利用率。
```
map的key的定位过程是怎样的？
```markdown
对key计算hash值，计算它落到那个桶时，只会用到最后B个bit位，再用哈希值的高8位
找到key在bucket中的位置。桶内没有key会找第一个空位放入，冲突则从前往后找到第一个空位。
```
iface和eface的区别是什么？值接收者和指针接收者的区别？
```markdown
iface和eface都是Go中描述接口的底层结构体，区别在于iface比eface多了itab结构，包含方法。
而eface则是不包含任何方法的空接口：interface{}

如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用
者；如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。

如果类型具备"原始的本质"，如go中内置的原始类型，就定义值接收者就好。
如果类型具备"非原始的本质"，不能被安全的复制，这种类型总是应该被共享，则可定义为指针接收者。
```
context是什么？如何被取消？有什么作用？
```markdown
type Context interface {
    // 当context被取消或者到了deadline，返回一个被关闭的channel
    Done() <-chan struct{}
    // 在channel Done关闭后，返回context取消原因
    Err() error
    // 返回context是否会被取消以及自动取消时间(即deadline)
    Deadline() (deadline time.Time,ok boll)
    // 获取key对应的value
    Value(key interface{}) interface{}
}

type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{} 
}

context：goroutine的上下文，包含goroutine的运行状态、环境、现场等信息。

实现了canceler接口的Context，就表明是可取消的。

context用来解决goroutine之间退出通知、元数据传递的功能。比如并发控制和超时控制。

注意事项：
1、不要将Context塞到结构体里，直接将Context类型作为函数的第一参数，而且一般都
命名为ctx。
2、不要向函数传入一个nil的Context，如果你实在不知道传什么，标准库给你准备好了
一个Context：todo
3、不要把本应该作为函数参数的类型塞到Context中，Context存储的应该是一些共同
的数据。例如：登陆的session、cookie等
4、同一个Context可能会被传递到多个goroutine，Context是并发安全的。
```
slice的底层数据结构是怎样的？
```markdown
type slice struct {
    array unsafe.Pointer
    len int
    cap int
}

slice的底层数据是数组，slice是对数组的封装，它描述一个数组的片段。
slice可以向后扩展，不可以向前扩展。
s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)。
```
你了解GC么？常见的GC实现方式有哪些？
```markdown
GC即垃圾回收机制：引用计数、三色标记法+混合写屏障机制
```
go的GC有那三个阶段？流程是什么？如果内存分配速度超过了标记清除速度怎么办？
```markdown
goV1.3之前采用的是普通标记清除，流程如下：
1、开始STW，暂停程序业务逻辑，找出不可达的对象和可达对象；
2、给所有可达对象做上标记；
3、标记完成之后，开始清除未标记的对象；
4、停止STW，让程序继续运行，然后循环重复这个过程，直到程序生命周期结束。

goV1.5三色标记法，流程如下：
1、只要是新创建的对象，默认的颜色都标记为白色；
2、每次GC回收开始，从根节点开始遍历所有对象，把遍历到的对象从白色集合
放入灰色集合；
3、遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色
对象放入黑色集合；
4、重复3中内容，直到灰色集合中无任何对象；
5、回收白色集合中的所有对象。

犹如剥洋葱一样，一层一层的遍历着色，但同时满足以下条件会导致对象丢失：
条件1：一个白色对象被黑色对象引用；
条件2：灰色对象与白色对象之间的可达关系同时被解除。

强三色：强制性的不允许黑色对象引用白色对象。
弱三色：黑色对象可以引用白色对象，但白色对象存在其他灰色对象对它的引用，或者
可达它的链路上游存在灰色对象。

goV1.8三色+混合写屏障机制，栈不启动屏障，流程如下：
1、GC开始将栈上的对象全部扫描并标记为黑色(之后不再进行重复扫描，无需STW)；
2、GC期间，任何在栈上创建的新对象均标记为黑色；
3、被删除的对象和被添加的对象均标记为灰色；
4、回收白色集合中的所有对象。

总结：
v1.3普通标记清除法，整体过程需要STW，效率极低；
v1.5三色标记法+屏障，堆空间启动写屏障，栈空间不启动，全部扫描之后，需要重新扫描
一次栈(需要STW)，效率普通；
v1.8三色标记法+混合写屏障，堆空间启动，栈空间不启动屏障，整体过程几乎不需要STW,
效率较高。

如果申请内存的速度超过预期，运行时就会让申请内存的应用程序辅助完成垃圾收集的扫描阶段，
在标记和标记终止阶段结束之后就会进入异步的清理阶段，将不用的内存增量回收。并发标记会
设置一个标志，并在mallocgc调用时进行检查，当存在新的内存分配时，会暂停分配内存过快
的哪些goroutine，并将其转去执行一些辅助标记的工作，从而达到放缓内存分配和加速GC工作
的目的。
```
内存泄漏如何解决？
```markdown
1、通过pprof工具获取内存相差较大的两个时间点heap数据。htop可以查看内存增长情况。
2、通过go tool pprof比较内存情况，分析多出来的内存。
3、分析代码、修复代码。
```
内存逃逸分析？
```markdown
在函数中申请一个新的对象，如果分配在栈中，则函数执行结束可自动将内存回收；
如果分配在堆中，则函数执行结束可交给GC处理。

案例：
函数返回局部变量指针；
申请内存过大超过栈的存储能力。
```

### mysql
索引底层实现？为什么选择B+树作为索引结构？B+树的叶子节点都可以存哪些东西？
```
B+树/hash；减少查询磁盘io次数；主键/行数据
```

覆盖索引是什么？回表？
```
不回表；数据不全通过Id查找所有数据
```

什么情况下不会命中索引？
```
最左前缀原则；字段类型转换；字段函数处理；LiKe 前后%；非等值非范围查找；OR关键字 等
```

mvvc的实现原理是什么？
```
多版本一致性读
```

redolog、undolog、binlog作用是什么？有什么区别？
```
重做日志；持久性
回滚日志：原子性
同步日志：数据同步
```

乐观锁和悲观锁的实现方式？
```
乐观锁：版本号
悲观锁：锁
```

事务隔离级别有哪些？如何解决脏读和幻读？
```
RU/RC/RR/S
脏读：提交可读
幻读：MVVC+next-key lock
```

应该如何恰当的建立索引？
```
区分度
```

mysql的存储引擎有哪些？都有什么区别？
```
MySAM：
InnoDB：
CSV：
Memory：
```

mysql的主从复制？
```
??
```

### 分布式系统、微服务架构
什么是分布式事务？

分布式事务解决方案有哪些？

分布式锁实现原理？有哪些实现方式？

微服务架构设计？

### kafka
kafka为什么性能高？
```
分区，日志分段，顺序写入
```

kafka重复消费可能的原因以及处理方式？
```
生产端：at least once
消息队列：选举机制
消费端：不幂等
```

kafka消息丢失的原因以及解决方式？
```
生产端：ACK
消息队列：选举机制
消费端：ACK
```

kafka如何保证消息的顺序性？
```
同分片，hash一致
```
什么是kafka的Rebalance？
```
？
```
kafka集群消息积压问题如何处理？
```
并发+offset
```
### redis
redis单线程为什么效率也这么高？
```markdown
1. 纯内存
2. IO多路复用
3. 减少锁的竞争
```
redis有那五种常用的数据结构？应用场景以及实现原理是什么？
```markdown
1. string  记数 SDS(RAW，embstr，int)
2. hash  缓存结构数据 quicklist（hashtable，ziplist）
3. list 队列 （ziplist，linkedlist）
4. set 集合 （intset，hashtable）
5. zset 延迟队列 （ziplist，skiplist）+ 字典
```
redis如何实现延迟队列？
```markdown
zset
```
redis的过期策略？
```markdown
被动+主动
```
redis与mysql双写一致性解决方案？
```markdown
CAP；先数据库，然后缓存（高可用）；
```
发生缓存穿透、击穿、雪崩的原因以及解决方案？
```markdown
穿透：非法数据；数据检验，布隆过滤器
雪崩：服务器不可用/数据大规模失效；随机过期时间、均衡负载、限流
击穿：热点数据过期；不过期/异步更新，限流
```
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
题目：将6，2，10，32，9，5，18，14，30，29从小到大进行排列,使用冒泡排序
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
