// go 并发记录
package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
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
	// 而且顺序还是乱序，简单的处理使用time.Sleep等一会而，实际一般使用waitGroup
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
	context.WithValue(context.Background(), "name", "刻晴")
	// 今天暂时到这里了
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
