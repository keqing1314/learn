// go 并发记录
package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 这里是前言，不算是学习记录，而是我的个人感想
// 学习和使用go已经3个多月了，并发这个对于现在我来说理解起来很轻松
// 我的学习的第一门静态语言是java，现在想起来当时在大学学习这个的时候基本对于
// 并发这个概念理解的不够深刻，更多的是停留在理论的记忆上，而且java的并发其实用起来
// 不是很方便，这种通过线程而且是系统态切换起来开销大，所以还有了线程池技术来复用线程
// 尽管框架可以使用一个Async注解来将同步方法转换为并发的
// 在学习go之前我都认为这个是很好的，直到一个go关键字就可以开启一个用户态线程(协程)
// 如此的方便，而且是用户态线程，轻松实现数十万的并发，而且还是在配置一般的
// 服务器上，重点是创建协程的开销相比传统小的多，对于IO密集型任务十分的方便
// 其实并发的使用场景最多的就是IO密集型任务，说白了就是IO等待的耗时较长
// 充分利用等待时间做其他事，让CPU利用率不忙等
// 其实生活中充满了并发，就像我们网购一样，我们下单买东西，算好了点击付款就行
// 最多就几分钟，就像CPU执行指令耗时一样就事纳秒级别，但是等待快递回来的多半要两天
// 这两天不可能一直等是吧，这段时间做其他事晴，直到两天后收到快递，网购任务才算结束
// 同样的，IO等待时间长，可能要个微秒或者毫秒级别的时间，尽管对于人类来说这些时间
// 都差不多，但是对于计算机就像是人类等了几天一样，这段时间完全可以利用来做其他任务
// 你看这样说并发的使用场景不久很清晰了嘛
// 在go也是如此，使用并发是用来解决问题的，如果没有IO等待并发还有有串行执行来的块
// 毕竟串行不用切换上下文，没有这个开销

// Go 语言对于并发的支持是纯天然的，这是这门语言的核心所在，其上手难度相对较小
// 开发人员不太需要关注底层实现就能做出一个相当不错的并发应用，提高了开发人员的下限。

// 协程
// 协程事一种轻量级的线程，或者说是用户态的线程
// 不受操作系统直接调度，由Go语言自身的调度器进行运行时调度
// 所以上文开销小，这也是Go的并发性能不错的原因之一
// 而是创建协程也十分简单

func main() {
	go fmt.Println("我是刻晴，璃月七星的玉衡星")
	go hello()
	go func() {
		fmt.Println("我是甘雨")
	}()
	// 这是三种创建使用协程的方式
	//【注意】对于有返回值的内置函数不能直接在go后面调用
	// 可以使用匿名函数包裹后使用，具体看使用场景
	// 还有就是主协程退出后，子协程也会退出，一般是主协程等待子协程
	fmt.Println("start")
	for i := range 10 {
		i++
		go fmt.Println("协程", i)
	}
	fmt.Println("end")
	// 上述代码中主协程没有等待子协程执行完毕后再退出
	// 所以执行结果是不确定的，有时候子协程能执行完毕，有时候还没开始执行就退出了
	// 而且顺序还是乱序，简单的处理使用time.Sleep等一会而，实际一般使用wgGroup
	//在go中常用的并发控制手段由三种：
	// Channel：通道或者叫管道
	// WaitGroup
	// Context
	// Mutex 这种传统的锁

	// 管道 Channel
	// 在go中说到Channel就经常会提到，这个的核心思想通信
	//【不要通过共享内存来通信，相反，通过通信来共享内存】
	var ch chan int
	fmt.Println(ch)
	// 上面是定义了一个nil管道，没有什么实际价值，就是可以这样写
	// 如果要定义可以使用的管道，只能通过make函数
	ch1 := make(chan int)     // 没有缓冲区
	ch2 := make(chan int, 10) // 有缓冲区
	// 上面两者的区别主要在于，是否可以同步读写
	fmt.Println(ch1, ch2)

	// 约定，一般来说命名为类型Stream
	intStream := make(chan int, 10)
	fmt.Println(intStream)
	// 不过这个也不是强制规定，而是社区的主流约定

	// 管道是资源创建的就要关闭，不然就是在浪费资源
	// 使用close内置函数，常常结合defer延迟调用一起使用
	// 注意这里不关闭ch是因为，nil管道读写关闭都会panic
	defer close(ch2)

	// 读写
	// 管道读写数据很是形象
	// ch <- 1 数据指向管道就是写数据
	// v,ok := <-ch 数据箭头指出去，就是读数据
	intStream <- 10
	v, ok := <-intStream
	if ok {
		fmt.Println("intStream", v)
	}
	intStream <- 20
	// 之所以使用intStream是因为这个是有缓冲区的管道
	// 可以同步读写，尽管实际开发中不建议这样做，但这是符合语法的
	// Demo1()
	// Demo3()
	// Demo4()
	Demo5()
	// len 和 cap
	// 对于channel管道也可以使用这两个内置函数
	// len看到的是当前管道中有几个元素
	// cap看到的是当前管道的缓冲区有多大

	//【管道总结】
	// 下面操作会造成阻塞
	// 1、读写无缓冲管道
	// 2、读取空缓冲管道
	// 3、写入满缓冲区管道
	// 4、管道为nil，无论怎样读写都会导致当前协程阻塞
	// 下面的操作会造成panic
	// 1、关闭一个nil管道
	// 2、写入已关闭的管道
	// 3、关闭已关闭的管道

	// 单向管道
	// 就是只读或者只写管道
	// 常作为函数参数来限制操作，避免并发问题
	func(readStrem <-chan int) {
		fmt.Println("本函数只能进行读取操作")
		if v, ok := <-readStrem; ok {
			fmt.Println("读取到值", v)
		}
		// 如果进行写操作就会编译失败
		// 错误: invalid operation:
		// cannot send to receive-only channel <-chan int
		// readStrem (variable of type <-chan int)
		// readStrem <- 10
	}(intStream)

	func(writeStream chan<- int) {
		fmt.Println("本函数只能进行写操作")
		writeStream <- 10
		// 读取操作就会报错
		// 错误: invalid operation: cannot receive from
		// send-only channel chan<- int writeStream (variable of type chan<- int)
		// <- writeStream
	}(intStream)

	//【管道注意】管道是引用类型

	// for range
	// 处理上述说的v,ok := <-ch的取值方式，还可以使用for range
	strStream := make(chan string, 10)
	for i := range 10 {
		strStream <- strconv.Itoa(i)
	}
	// 注意管道使用完毕一定要关闭
	// 发送方关闭
	// 如果不关闭，for range不知道数据发送完了就会继续读取
	// 就会因为读取空缓冲区管道最后死锁
	close(strStream)
	for s := range strStream {
		fmt.Println(s)
	}

	// =========== WaitGroup =============
	// sync.WaitGroup是一个结构体，可以实现等待一组协程的效果
	var wg sync.WaitGroup
	// 对外暴露 四个方法
	// Add 方法用于指明要等待的协程的数量
	wg.Add(1)
	// Done  方法表示当前协程已经执行完毕
	wg.Done()
	// Go 方法开启一个新的协程执行任务，包含WaitGroup Add和Done
	// 这个就是一个回调
	wg.Go(func() {
		fmt.Println("刻晴哈哈哈")
	})
	// Wait 方法等待子协程结束，否则就阻塞
	wg.Wait()
	WGDemo2()
	// WaitGroup使用十分简单上面的示例就把使用情况全部展现了
	// 主协程调用Wait就会阻塞直到所有Add的任务全部都Done了
	//【WaitGroup注意】因为是结构体，所以函数参数传递的时候是按照值传递的
	// 会拷贝副本，所以不能最好使用指针传递

	// =========== Context ===========
	// 上下文，这个在go中很常见是并发控制的一种手段
	// 相比于管道和WaitGroup，它可以更好的控制子孙协程的这种层级关系
	// 比如在web后台服务中处理请求
	// 当接受到一个请求，服务会开启一个协程来处理这个请求，这个协程可能启动子协程
	// 来处理问题，但是如果请求关闭了呢，按理说这个请求的所有操作都该结束是吧
	// 如果没有Context，答案是不会结束，子协程会继续执行，这显然不合理，所以有了Context

	// Context就是一个接口
	// type Context interface {
	// 	Deadline() (deadline time.Time, ok bool)
	// 该方法具有两个返回值，deadline是截止时间，即上下文应该取消的时间
	// 第二个值是是否设置deadline,如果没有设置则一直为false

	// 	Done() <-chan struct{}
	// 其返回值是一个空结构体类型的只读管道，该管道仅仅起到通知作用，不传递任何数据
	// 当上下文所做的工作应该取消时，该通道就会被关闭，对于一些不出取消的上下文，可能会返回nil

	// 	Err() error
	// 该方法返回一个error，用于便是上下关闭的原因。当Done管道没有关闭时，返回nil
	// 如果关闭过后，会返回一个err来解释为什么会关闭

	// 	Value(key any) any
	// 改方法返回对应的键值，如果key不存在，或者不支持该方法，就会返回nil

	// }

	// emptyCtx
	// 空上下文，创建上下文的函数
	context.Background() // 创建空上下文
	context.TODO()       // 创建空上下文
	// 这两者本质上是回事，不过就是语义上的差距
	// emptyCtx 通常是用来当作最顶层的上下文，在创建其他三种上下文时作为父上下文传入

	// valueCtx
	// 这里是展示如何通过WithValue函数创建
	//【注意】valueCtx多用于在多级协程中传递一些数据，无法被取消
	// 因此ctx.Done永远返回nil
	valCtx := context.WithValue(context.Background(), "name", "刻晴")
	fmt.Println(valCtx.Value("name"))

	// cancelCtx
	// 这个实现了canceler接口
	// 一样的cancelCtx也不直接对外暴露而是通过WithCancel函数来创建
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx, cancel)
	// ctx 是创建的上下文，cancel是一个函数用来取消该上下文
	// 当调用cancel函数时，Done管道会关闭，该上下文的任何子级都会随之取消
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				// 收到Done管道数据直接打印关闭原因并退出循环
				fmt.Println(ctx.Err()) // context canceled
				return
			default:
				fmt.Println("等待取消中...")
			}
			time.Sleep(time.Millisecond * 200)
		}
	}(ctx)
	time.Sleep(time.Second)
	cancel()
	wg.Wait()

	// timerCtx
	// timerCtx其实就是在cancelCtx上增加超时机制，超时会自动取消
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*2)
	fmt.Println(ctx1, cancel1)
	// 同样的这个会创建context和取消函数，不过就算不手动取消，超时也会自动取消
	// 当然保险起见也可以手动调用cancel1函数，常见是配合defer延迟调用使用
	wg.Go(func() {
		for {
			select {
			case <-ctx1.Done():
				fmt.Println(ctx1.Err()) // context deadline exceeded
				return
			default:
				fmt.Println("等待超时...")
			}
			time.Sleep(time.Millisecond * 200)
		}
	})
	wg.Wait()
	// 除了使用WithTimeout，也可以使用WithDeadline

	// =========Select==========
	// 前面提到过Select使用来监听管道的，语法上和Switch有点像
	// Select可以监听多个管道，做不同的处理
	chA := make(chan int, 10)
	chB := make(chan int, 10)
	chC := make(chan int, 10)
	defer func() {
		close(chA)
		close(chB)
		close(chC)
	}()

	select {
	case <-chA:
		fmt.Println("处理管道A")
	case <-chB:
		fmt.Println("处理管道B")
	case <-chC:
		fmt.Println("处理管道C")
	default:
		fmt.Println("默认操作")
	}

	select {
	//case里面也可以向管道里面写入数据
	case chA <- 10:
		fmt.Println(<-chA)
	default: // 添加这个就不会阻塞
	}

	// 还可以实现超时机制
	select {
	case n := <-chA:
		fmt.Println(n)
	case <-time.After(time.Second):
		fmt.Println("超时")
	}

	// 如果你经常看源码，你会发现select经常会结合for{}循环使用

	// 永久阻塞，上述所说如果没有default，且监听的管道都没有数据来，就会阻塞
	// 想要不阻塞就结合使用超时机制

	// =========互斥锁==========
	// 前面说到管道中的通信，本质是CSP（通信顺序进程）的理论实践
	// 但是不是说所有的操作都是依靠管道实现，对于协程之间的通信，使用管道很适用
	// 但是对于共享资源的保护，使用互斥锁更好
	count := 0

	go func() {
		for range 10 {
			count++
		}
	}()

	go func() {
		for range 7 {
			count--
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("count", count)
	// 当你执行上述代码的时候，按理说你想的结果是3
	// 但是实际上4，6都有可能，因为资源竞争了，上述自增和自减操作都不是原子性的
	// 当一个协程准备加操作的时候，另一个协程进行了减操作，对方都不知道对方已经操作了
	// 导致结果是不可预期的，这种情况如果使用管道，效果不是很好，但是使用互斥锁
	// 将资源保护起来，每次加的时候就不能减，减的时候就不能加，让操作变得原子性

	count1 := 0
	var lock sync.Mutex

	go func() {
		for range 10 {
			lock.Lock()
			count1++
			lock.Unlock()
		}
	}()

	go func() {
		for range 7 {
			lock.Lock()
			count1--
			lock.Unlock()
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("count1", count1)

	// 其实锁这个可说的没有多少，和其他编程语言的用法基本一致
	// 不过在go中的互斥锁是不可重入锁，提供了普通互斥锁和读写互斥锁
	// 具体的自行了解
	//【注意】对于锁而言，不应该将其作为值传递和存储，应该使用指针

	// =========条件变量=========
	// 条件边聊，与互斥锁一同出现和使用，所以有些人可能会误称为条件所
	// 但是它不是锁，是一种通讯机制，go中的sync.Cond对此提供了实现
	// 创建的条件的函数签名如下
	// func NewCond(l Locker) *Cond
	// 可以看到创建条件变量的前提是要有一把锁
	// 提供Wait Signal Broadcast 几个方法
	// 例子我不写在main函数里面了，在函数CondDemo1()中
	CondDemo1()
	//【注意】对于条件变量，应该使用for而不是if，应该使用循环来判断条件是否满足
	// 因为协程被唤醒时并不保证当前条件就已经满足了
	// for !condition {
	// 	cond.Wait()
	// }

	// ==========sync==========
	// go中很大部分的并发相关的工具都是sync标准库提供的，上面已经说过了
	// sync.WaitGroup,sync.Locker,除此之外还有些常用的，下面一起介绍

	// Once
	// 当使用一些二数据结构时，如果这些数据结构体太过于庞大，就可以使用Once
	// 这是一个中懒加载的模式，只会在使用的时候才会初始化该数据结构体
	// 这种结构很常见，例如
	// func() {
	// 	if *m == nil { // 使用的时候才会初始化
	// 		*m = make([]int, 0, 10)
	// 	}
	// 	*m = append(*m, i)
	// }()
	// 但是有问题是吧，多协程的时候就有问题，
	// 不然的话协程A初始化了，但是还没完成，这时候协程B来了，它的速度开立马就初始化了
	// 最后会怎样，B的把A给覆盖了，在java中一般使用锁关键字双重检定
	// 但是在go中呢，其实提供了一个sync.Once来解决这个问题
	// Once顾名思义就是一次，只初始化一次，对外只暴露一个方法Do
	// 例子在OnceDemo1()函数
	OnceDemo1()

	// Pool
	// sync.Pool的设计目的时用于存储临时对象以便后续的复用
	// 是一个临时的并发安全对象池，将暂时用不到的对象临时存在对象池中
	// 在后续使用就不需要再额外的创建对象，直接在池中取对象复用就行
	// 减少内存的分配与释放频率，降低GC压力
	// 这个场景在游戏开发中很常见，对象复用，避免频繁创建
	// sync.Pool中总共只有两个方法
	// 申请一个对象
	// func (p *Pool) Get() any
	// 放入一个对象
	// func (p *Pool) Put(x any)
	// 并且sync.Pool有一个对外暴露的New字段
	// 示例在PoolDemo1()函数中
	PoolDemo1()

	//=========原子===========
	// 在计算机科学中，原子或者原语操作，通常用于表述一些不可再分割的操作
	// 由于这些操作无法再细化为更小的步骤，再执行完毕前，不会被其他的任何协程打断
	// 所以执行结果要么成功，要么失败，没有第三种情况可言，如果出现了其他情况
	// 那么他就是不是原子操作，如下代码
	a := 0
	if a == 0 {
		a = 1
	}
	fmt.Println(a)
	// 虽然上述代码很简单，但是不是原子操作
	// 原子操作是底层硬件指令层面支持的

	// 类型
	// go中提供了一些原子类型，至于底层的指令的封装，不需要程序员来处理
	// 编程语言中就会提前封装好
	_ = atomic.Bool{}
	_ = atomic.Pointer[string]{} // 现在支持泛型了
	_ = atomic.Int32{}
	_ = atomic.Int64{}
	_ = atomic.Uint32{}
	_ = atomic.Uint64{}
	_ = atomic.Uintptr{}
	_ = atomic.Value{}
	// Value可以支持存储任何类型，除此之外，还提供了许多的函数来方便操作
	// 因为原子操作的粒度过细，在大多情况下，更适合处理这些基础的数据类型
	// atomic包下的原子操作都是函数签名，具体实现是汇编实现的，不同的架构实现不一样的

	// 使用
	// 每一个原子类型都会提供以下三个方法：
	// Load():原子的获取值
	// Swap(newVal type) (old type):原子的交换之，并且返回旧值
	// Store(val type):原子的存储值

	var aint64 atomic.Int64
	aint64.Store(64) //存储值
	aint64.Swap(128) //交换值
	aint64.Add(112)  //增加
	// 加载值
	fmt.Println(aint64.Load())

	// CAS
	// 就是比较并交换，一种无锁数据结构的核心，也称之为乐观锁
	// 但是这不是锁，而是一种并发条件下无锁化并发控制方式
	// 线程/协程在修改数据前，不会先加锁，而是先读取数据，进行计算
	// 然后再提交修改时使用CAS来判断此期间是否有其他线程修改过数据
	// 如果没有（值还等于之前读取的值）则修改成功，否则失败重试
	// 之所以被称为乐观锁，是因为他总是乐观的假设共享数据不会被修改
	// 仅当发现数据未被修改时才会去执行对应操作
	// 但是这个CAS存在ABA问题，可以通过版本号机制来解决
	// CAS有三个值，内存值，期望值，新值
	// 逻辑关系大致是这样的
	// if 内存值 == 期望值 {
	// 	内存值 = 新值
	// 	return true // 成功
	// } else {
	// 	return false // 失败，内存值不变
	// }

	// Value
	// 前面说到atomic.Value这个可以存储任何类型的值
	// 但是nil不行啊
	var aval atomic.Value
	aval.Store("刻晴")
	aval.Store("兹白")
	// 还有就是存储类型要一致才行！！！，不然会panic
	// 只能存一个值
	// "sync/atomic: store of inconsistently typed value into Value"
	fmt.Println(aval.Load())

}

func hello() {
	fmt.Println("你好呀，我是刻晴")
}

// 无缓冲示例
func Demo1() {
	ch := make(chan int)
	defer close(ch)
	// 写入数据
	go func() {
		ch <- 3
	}()
	// 读取数据
	v := <-ch
	fmt.Println(v)
	// 尽管看起来没啥问题，但是会死锁
	// 无缓冲就是这样的
	// 一个协程写数据，就必须有另一个协程来读取数据，不然就会阻塞, 触发死锁
	// fatal error: all goroutines are asleep - deadlock!
	// 解决也很简单，给写数据单独开一个协程
}

// 有缓冲示例
func Demo2() {
	ch := make(chan int, 10) // 缓冲区为10的管道
	defer close(ch)
	ch <- 3
	v := <-ch
	fmt.Println(v)
	// 因为有缓冲区，所以同步操作不会阻塞
	// 但是向已满的缓冲区写入数据会阻塞
	// 向空的缓冲区中读数据也会阻塞
	// 所以之前才会说不建议同步操作，也会有风险的
}

func Demo3() {
	// 创建有缓冲guandao
	ch := make(chan int, 5)
	// 创建两个无缓冲管道
	chW := make(chan struct{})
	chR := make(chan struct{})
	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()
	// 负责写
	go func() {
		for i := range 10 {
			ch <- i
			fmt.Println("写入", i)
		}
		chW <- struct{}{}
	}()
	// 负责读
	go func() {
		for range 10 {
			// 模拟读取耗时
			time.Sleep(time.Millisecond)
			fmt.Println("读取", <-ch)
		}
		chR <- struct{}{}
	}()
	fmt.Println("写入完毕", <-chW)
	fmt.Println("读取完毕", <-chR)
	// 写入 0
	// 写入 1
	// 写入 2
	// 写入 3
	// 写入 4
	// 读取 0
	// 写入 5
	// 读取 1
	// 写入 6
	// 读取 2
	// 写入 7
	// 读取 3
	// 写入 8
	// 读取 4
	// 写入 9
	// 写入完毕 {}
	// 读取 5
	// 读取 6
	// 读取 7
	// 读取 8
	// 读取 9
	// 读取完毕 {}
	// 上述是执行效果
	// 可以看到读取耗时比写入耗时长，就会出现一次性将缓冲区写满的情况
	// 如果不是读写都是一个协程，这个代码就有问题了
}

func Demo4() {
	// 上面也看到了管道阻塞的情况，可以利用这点实现等待操作
	quit := make(chan int)
	go func() {
		fmt.Println("刻晴天下第一个可爱")
		time.Sleep(time.Second)
		quit <- 1
	}()
	<-quit
	fmt.Println("子协程执行完毕，主协程退出")
	// 这个函数实现了函数主协程，等待其子协程执行完毕后再退出
}

// 通过缓冲管道实现互斥锁
func Demo5() {
	count := 0
	// 创建一个缓冲区为1的管道
	// 由于缓冲区是1，只要有协程在干活，就会独占
	// 不过实际上遇到这种情况直接使用互斥锁就行
	lock := make(chan struct{}, 1)

	add := func() {
		lock <- struct{}{}
		count++
		<-lock
	}

	sub := func() {
		lock <- struct{}{}
		count--
		<-lock
	}

	go func() {
		for range 10 {
			add()
		}
	}()

	go func() {
		for range 8 {
			sub()
		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(count)
}

// 管道其实非常适合一种操作
// 就是流水线模式，处理数据
func Demo6() {
	nums := []int{2, 4, 5, 6, 1}
	// 第一步：切片转换为channel
	dataChannel := slicesToChannel(nums)
	// 第二步：元素平方
	finalChannel := sq(dataChannel)
	for v := range finalChannel {
		fmt.Println(v)
	}
}

func slicesToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func WGDemo1() {
	var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	fmt.Println("子协程执行任务")
	// 	wg.Done()
	// }()
	// 这个Go方法就是上面注释代码的等效
	wg.Go(func() {
		// 模拟子协程执行任务耗时1秒
		fmt.Println("子协程执行任务")
		time.Sleep(time.Second)
	})

	// 这里会阻塞主协程，直到子协程执行完毕
	wg.Wait()
	fmt.Println("子协程执行完毕，主协程退出")
}

func WGDemo2() {
	// 一个简单的示例优化，结合WaitGroup
	var mwg, wg sync.WaitGroup
	mwg.Add(10)
	fmt.Println("start")
	for i := range 10 {
		wg.Go(func() {
			fmt.Println(i)
			mwg.Done()
		})
		// 阻塞子协程，保证顺序
		wg.Wait()
	}
	// 阻塞主协程
	mwg.Wait()
	fmt.Println("end")
}

var (
	wg sync.WaitGroup
	rw sync.RWMutex
)

var count = 0

// 创建条件变量
var cond = sync.NewCond(rw.RLocker())

// 条件变量的示例函数
func CondDemo1() {
	wg.Add(12)
	// 读多写少
	go func() {
		for range 3 {
			go Write(&count)
		}
		wg.Done()
	}()

	go func() {
		for range 7 {
			go Read(&count)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("最终结果", count)

}

func Read(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	rw.RLock()
	fmt.Println("拿到读锁")
	// 条件不满足就一直阻塞
	for *i < 3 {
		cond.Wait()
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("释放读锁", *i)
	rw.RUnlock()
	wg.Done()
}

func Write(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	rw.Lock()
	fmt.Println("拿到写锁")
	temp := *i
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	*i = temp + 1
	fmt.Println("释放写锁", *i)
	rw.Unlock()
	// 唤醒所有因条件变量阻塞的协程
	cond.Broadcast()
	wg.Done()
}

var wait sync.WaitGroup

// Once示例函数
func OnceDemo1() {
	var slice MySlice
	wait.Add(4)
	for range 4 {
		go func() {
			slice.Add(1)
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(slice.Len())
}

type MySlice struct {
	s []int
	o sync.Once
}

func (m *MySlice) Get(i int) (int, bool) {
	if m.s == nil {
		return 0, false
	} else {
		return m.s[i], true
	}
}

func (m *MySlice) Add(i int) {
	// 当真正用到切片的时候，才会考虑去初始化
	m.o.Do(func() {
		fmt.Println("初始化")
		if m.s == nil {
			m.s = make([]int, 0, 10)
		}
	})
	m.s = append(m.s, i)
}

func (m *MySlice) Len() int {
	return len(m.s)
}

// 创建临时对象池
var pool sync.Pool

// 用于计数创建了多少个对象
var numOfObject atomic.Int64

// 大结构体
type BigMemData struct {
	M string
	N string
}

// Pool示例函数
func PoolDemo1() {
	pool.New = func() any {
		numOfObject.Add(1)
		return BigMemData{"刻晴", "兹白"}
	}

	// 开启100_000个协程
	for range 100_000 {
		wait.Go(func() {
			val := pool.Get()
			//【注意】中间的操作如果IO时间长，创建的对象也会很多
			// 难以实现复用效果
			if _, ok := val.(BigMemData); ok {
				// 这个打印就是耗时较长的IO
				// 如果耗时过于长，这个就会退化成一个协程创建一个对象了
				// fmt.Println(v.M, v.N)
			}
			// 释放对象
			pool.Put(val)
		})
	}
	wait.Wait()
	fmt.Println("对象数量", numOfObject.Load())
	// 这里如果不是使用pool复用，100_000个协程每个协程都要创建一个对象
	// 那么就是100_000个GC的压力直接拉爆，但是现在你看打印反正比100_000少
	// 临时对象：sync.Pool 只适合存放临时对象，池中的对象可能会在没有任何通知的情况下被 GC 移除，所以并不建议将网络链接，数据库连接这类存入 sync.Pool 中。
	// 不可预知：sync.Pool 在申请对象时，无法预知这个对象是新创建的还是复用的，也无法知晓池中有几个对象
	// 并发安全：官方保证 sync.Pool 一定是并发安全，但并不保证用于创建对象的 New 函数就一定是并发安全的，
	// New 函数是由使用者传入的，所以 New 函数的并发安全性要由使用者自己来维护，这也是为什么上例中对象计数要用到原子值的原因。
}
