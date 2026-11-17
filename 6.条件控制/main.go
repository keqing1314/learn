// go 条件控制记录
package main

import (
	"errors"
	"fmt"
	"log"
)

// 条件控制就是分支语句
// 在java中就是if else swirch case 之类的
// 在go中也类似有if switch，但是多了select用于通道多路控制
// 下面简单介绍下

func main() {
	// if else
	// 在go中不存在三元运算符，全部直接使用if else
	// if else是简单的分支，一个条件和条件取反
	// 最多有两条分支
	a := 10
	if a == 10 {
		log.Println(10)
	}

	// 上面的示例是单条件分支
	// 这是双分支
	if a == 10 {
		log.Println(10)
	} else {
		log.Println("不是10")
	}

	// 格式就是
	// if expression {
	// } else {
	// }
	//【注意】expression是表达式，结果必须是bool
	// 另外就是在expression可以进行赋值操作，使用;分隔就行
	if a := 10; a == 10 {
		log.Println("赋值操作")
	}
	// 这样的变量作用域只限制与这个赋值语句中
	// 这种写法很常见
	// 给一个简单的示例
	if err := somethingDo(); err != nil {
		log.Println(err)
	}

	// 多分支语句
	// if else if else
	// 结构和上面的if else 基本一致
	// if expression1 {
	// } else if expression2 {
	// } else if expression3 {
	// } else {
	// }
	// 就是变成了else if 分支了
	// 给一个简单的示例
	score := 90
	var ans string
	if score == 100 {
		ans = "S"
	} else if score >= 90 && score < 100 {
		ans = "A"
	} else if score >= 80 && score < 90 {
		ans = "B"
	} else if score >= 70 && score < 80 {
		ans = "C"
	} else if score >= 60 && score < 70 {
		ans = "E"
	} else if score >= 0 && score < 60 {
		ans = "F"
	} else {
		ans = "nil"
	}
	fmt.Println(ans)

	// switch
	// 也是一种多分支语句，很多时候多分支使用switch
	// 格式如下
	// 	switch expr {
	//   case case1:
	//     statement1
	//   case case2:
	//     statement2
	//   default:
	//     default statement
	// }
	//【注意】go中没有break，想要执行下一个分支要使用fallthrough
	str := "a"
	switch str {
	case "a":
		str += "a"
		str += "c"
	case "b":
		str += "bb"
		str += "aaaa"
	default: // 当所有case都不匹配后，就会执行default分支
		str += "CCCC"
	}
	fmt.Println(str)

	var f func() int
	f = func() int {
		return 1
	}

	// 还可以在表达式之前编写一些简单语句，例如声明新变量
	switch num := f(); { // 等价于 switch num := f(); true {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		fallthrough
	case num < 0:
		num += num
	}

	// switch语句也可以没有入口处的表达式
	num := 2
	switch { // 等价于 switch true {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
	case num < 0:
		num *= num
	}
	fmt.Println(num)

	// 通过fallthrough关键字来继续执行相邻的下一个分支
	num2 := 2
	switch {
	case num2 >= 0 && num2 <= 1:
		num2++
	case num2 > 1:
		num2--
		fallthrough // 执行完该分支后，会继续执行下一个分支
	case num2 < 0:
		num2 += num2
	}
	fmt.Println(num2)

	// label
	// 标签，给代码打上标签，可以是goto，break，continue
	// 示例如下
	// A:
	// 	a1 := 1
	// B:
	// 	b1 := 2
	// 但是单纯的标签没有实际意义，要结合具体的关键字使用

	// goto
	// goto跳转，将控制权传递给同一函数中的标签语句
	a1 := 1
	if a1 == 1 {
		goto A
	} else {
		log.Println("B")
	}
A:
	log.Println("A")
	// 但是实际上使用跳转关键词goto的实际场景较少(反正我在开发中基本就没有使用过)
	// 不过源码中有看见过使用场景，但是也很少。性能是一个问题
	// 同时语句执行跳来跳去，用多了可读性也会下降

	// select
	// 这个是用于通道中，在并发编程中很常见
	// 例如要监听多个通道，这里简单说下语法就行，后续并发中会详细介绍
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	// 上面创建了三个通道，假设在其他协程写入数据，在主协程中监听并跳转对应的处理分支
	// 通常会结合for循环使用
	// 可以在case中读取数据，也可以写入数据，还可以有默认操作
	// 这就是监听多个通道，如果，同时有多个通道收到数据，会伪随机选一个来执行
	for {
		select {
		case <-ch1:
			log.Println("ch1中收到数据")
		case v := <-ch2:
			log.Printf("ch2中收到数据:%v", v)
		case ch3 <- 1:
			log.Println("向ch3中写入数据")
		default:
			log.Println("默认操作,不阻塞")
		}
	}
}

// 模拟返回错误的函数
func somethingDo() error {
	return errors.New("这是返回的错误")
}
