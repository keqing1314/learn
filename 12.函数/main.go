// go 函数记录
package main

import (
	"fmt"
	"log"
)

// 函数在go中可谓是基石一般的存在
// 经常有一句话说函数是go的一等公民
// 在go中实现逻辑都是函数
// 方法某种意义上说也是函数
// 回调函数，闭包函数这些都是函数优秀的使用
// 下面来介绍下具体函数是个什么样的

// 声明格式
// func 函数名(...参数列表)(...函数返回值)
// 一个函数的声明就是如此，当让还有匿名函数
// 在go中函数可以作为参数传递给另一个函数，函数也是一种类型

// 定义一个递增函数
// 声明函数关键字是func
// 这里的函数名是increment
// 参数列表是count int
// 返回值是int类型
// {}里面就是函数的逻辑
// 这是一个典型的函数声明
func increment(count int) int {
	count++
	return count
}

//【注意】go中不支持函数重载，如果参数列表不同，那么就是不同的函数
// 所以函数名也不应该相同

// 函数中参数的说明
// 同种类型的参数，可以声明一次类型
func somethingDo(name, addr string) {
	// 这里的name和addr都是string类型，那么就只需要声明一次类型
}

//【注意】go中函数参数都是按值传递的，对于大参数建议使用指针类型
// 避免拷贝带来的性能损耗

// 返回值
// go中允许多个返回值，有返回值就写，没有就不写
// 不像其他编程语言写void，go中不用
func fn1() (bool, error) {
	return true, nil
}

// go中允许具名返回值
// 同样的可以有多个具名返回值
// 【注意】具名返回值，要么全部具名，要么都不具名，不然编译不通过
func fn2() (res int) {
	// res 默认值是零值所以返回0
	return
}

func main() {
	// 来一个匿名函数声明示例
	// 这里看到声明函数类型，参数可以不写变量名
	// 不过建议是带上，能增加可读性的话
	// go中支持变长参数，多个参数，变长参数要在最后面，源码中就可以看到这个
	// 	func Printf(format string, a ...any) (n int, err error) {
	// 	return Fprintf(os.Stdout, format, a...)
	// }
	var fn func(...int) int

	fn = func(nums ...int) int {
		total := 0
		for _, v := range nums {
			total = total + v
		}
		return total
	}

	fmt.Println(fn(10, 23, 66, 55))

	// 上述是匿名函数的一中，先声明类型，再赋值，看起来就像变量
	// 还有一种处理方式，直接声明
	func(num int) {
		log.Println("这里是匿名函数操作", num)
	}(10)
	// 注意匿名函数如果由返回值的话最好用第一种方式

	// defer 延迟调用
	// 这个通常就和匿名函数第二种方式结合
	// 比如panic之后的recover
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// 这里就是延迟函数的使用
	// 延迟函数调用，就是延迟函数的调用
	// 在函数return之前才开始调用
	// 多个延迟调用按照注册顺序的【逆序调用】
	// 延迟调用存在"快照机制"这个名是我取的
	// 具体说延迟的是调用，但是参数会在注册时就确定了
	str := "我是刻晴"
	defer fmt.Println(str)
	str = "我是甘雨"
	// 最后defer中打印的一定是【我是刻晴】
	// 虽然在返回前才会调用，但是str在defer注册时就确定了，后面就算修改了也不会变化
	// 不过如果使用的指针就会变化，因为指向的地址是同一个
	// 还有就是匿名函数持有外部变量
	defer func() {
		fmt.Println(str) // 打印的就是【我是甘雨】
	}()
	//【注意】defer看起来简单，但是结合panic就会有很多坑
	// 后面错误处理会详细介绍的
}

// 闭包函数
// 这是一个经典的使用
// 简单说就是一个函数返回另一个函数，通过这种方式可以实现函数工厂
func counter(num int) func() int {
	n := num
	return func() int {
		n++
		return n
	}
	// 这就是闭包，同时由变量逃逸
	// n变量从栈逃逸到堆上
}

// 回调函数
func callIncrement(fn func(int) int) int {
	// 这里就是回调
	// 这里的例子场景不一定恰当，主要示例
	r := fn(42)
	return r
}
