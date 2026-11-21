// go 错误记录
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// 错误处理一直都是各个语言的一种处理重点
// 我想在这个上面的确要花点时间才行
// 部分编程人员，喜欢将代码会出错的地方尽量完善，但是很少在错误处理上花时间
// 都是极可能希望自己的代码少出错，这没错，但是完全忽略错误处理这是不行的
// 现如今绝大多数主流编程语言采用try catch这种处理方式
// 捕获错误，但是处理范围不限，捕获错误不限，全是自己设置
// 有的直接将整段代码捕获，但是这样颗粒度过于粗，对于排查错误也不友好

func main() {
	// go中对于错误的处理是不同的
	// go将错误作为返回值返回给开发者
	// 比如在读取一个文件的时候，文件可能不存在，文件异常之类的
	_, err := os.ReadFile("README.md")
	if err != nil {
		log.Printf("读取文件失败:%v", err)
	}
	// 这样可读性好多了，那里有错误一眼就看见了
	// 个人喜欢这种写法，比起try catch
	// 虽然处理错误多了，这种代码会很多，但是这很清晰的
	// 而且就算使用try catch 很多推荐的写法是减小颗粒度
	// 代码中也会大量充斥着try catch ,没不如这个直观

	//【注意】go的错误是一个返回值，那么就具有变量的特性
	// 可以将响应封装成带有错误的结构体，这样只需要在函数返回的时候查看
	// 有没有错误内容就行，尤其是并发编程中，这种处理真的很好

	// 创建
	// 1、使用errors包
	_ = errors.New("这是创建的错误")
	// 2、fmt包的Errorf函数
	_ = fmt.Errorf("这是错误%w", err)
	// go项目建议将需要使用的错误提前创建好
	// 在go的源代码中很常见

}

// 自定义错误
// error本质是一个接口，不属于任何包
// 只要实现error接口，就是自定义错误
func New(text string) error {
	return &errorString{s: text}
}

type errorString struct {
	s string
}

// 实现error接口
func (e *errorString) Error() string {
	fmt.Println("这是自定义错误")
	return e.s
}

// 传递
// 在一些情况中，调用者调用的函数返回了一个错误
// 但是调用者本身不负责处理错误，于是也将错误作为返回值返回
// 抛给上一层的调用者，这个过程就是传递
// 错误在传递中可能层层包装，当上层调用者想要判断错误的类型
// 来做出不同的处理时，可能会无法判断错误的类别或者误判
// 而链式错误就是为了解决这个问题的
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

// wappError同样实现了error接口，也多了一个Unwrap方法
// 用于返回其内部对于原error的引用，层层包装之下形成一条错误链表
// 顺着链表上寻找，很容就能找到原始错误，由于该结构体并不对外暴露
// 所以只能使用fmt.Errorf函数来进行创建

// 处理
// 错误处理中的最后一步就是检查和处理错误
// errors.Unwrap函数用于解包一个错误连，其内部实现大致如下
func Unwrap(err error) error {
	u, ok := err.(interface { // 类型断言，是否实现该方法
		Unwrap() error
	})
	if !ok { // 没有事先说明是一个基础的error
		return nil
	}
	return u.Unwrap() // 否则带哦用Unwrap
}

// 解包后会返回当前错误连所包裹的错误，被包裹的错误可能依旧是一个错误链
// errors.Is函数就喊判断错误链中是否存在指定错误
var originalErr = errors.New("这是原始错误")

func wrap1() error {
	return fmt.Errorf("wapp1 error%w", wrap2())
}

func wrap2() error {
	return originalErr
}

func Demo1() {
	err := wrap1()
	if errors.Is(err, originalErr) {
		fmt.Println("当前错误链存在原始错误")
	}
}

// errors.As函数的作用是在错误链中找到第一个匹配的错误

// panic
// panic中文就是恐慌
// 表示十分严重的程序问题，程序需要立即停止
func Demo2() {
	var mp map[string]int
	mp["a"] = '2'
	// 向nil map写入内容会panic
}

//【提示】程序存在多个协程的时候，只要任意协程触发panic，
// 不recove处理的话，整个程序就会崩溃

// 创建panic
// 使用panic()就能创建
func Demo3() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	panic("panic")
	// 这里注释了，后面永远不会执行
	//defer fmt.Println("D")
}

//【注意】panic之后，会在推出前做善后工作
// 上面的示例中存在defer延迟调用，会执行panic之前注册的调用
// C
// B
// A
// panic:panic
// 这个不是很好理解，尤其是defer多层嵌套的情况

// recover
// 触发panic后可以使用recover恢复的
//【注意】recover必须在defer中使用，多层嵌套是不行的

// fatal
// 这个是go最严重的错误，表示程序不应该运行，而是立即停止的情况
// 创建 fatal
func Demo4() {
	log.Fatal("创建fatal")
}
