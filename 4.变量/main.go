// go 变量相关记录
package main

import "log"

// 变量可以分为两种作用域
// 包级和函数内，遵循go的可见性，首字母大写共有，小写私有
// 前面提到过常量使用的const关键字声明
// 变量使用的是var关键字声明
var name string = "刻晴"

// 和常量一样支持编译器类型推断
var Age = 23

// 和常量不同的一点是，变量可以声明不初始化，在使用的时候才初始化
// 和常量一样，包级别的变量可以不使用，但是函数内的就必须使用
// 否则编译不通过
var addr string

// 相同类型的变量可以一行声明
var num1, num2, num3 int

// 同样也可以是使用()将变量括起来，增加可读性
var (
	dsn  string
	host string
	port int
)

// main函数
func main() {
	// 函数内和包级别的变量赋值基本没有区别
	// 前面也提到了重要的使用点就是声明了必须使用
	// 不然就会报错declared and not used: username
	var username = "刻晴"
	log.Println(username) // 通过日志打印使用这个变量
	// 不同的就是短变量操作符
	password := "qwerty_1024"
	log.Printf("密码是%v", password)

	// 支持多变量赋值
	num1, num2, num3, num4 := 3, 2, 1, 0
	log.Println(num1 + num2 + num3 + num4)

	// 匿名处理，使用下划线可以匿名处理
	_ = Sum(num1, num2, num3, num4)

	// 比较：对于变量的比较，必须类型相同，才能比较
	//【注意】这里说的比较不是之前说的大类分类，而是具体的类型
	var a int64 = 10
	var b int = 10
	// invalid operation: a == b (mismatched types int64 and int)
	// if a == b {
	// 	log.Println("相等")
	// }
	// 上述示例就会报错，因为int和int64是两种类型
	//【注意】go中不存在类型装箱和拆箱
	// 想要两个不类型的变量比较，就要类型装换
	// 前提是在同一个大类里面可以类型转换才行
	if int64(b) == a {
		log.Println("相等")
	}

	// 同时类型也要可以比较才行
	// 有一个 comparable 接口，就可以限制可比较类型
	// go中可比较类型有：
	// bool
	// 数字 整型 浮点型 负数
	// 字符串 指针 通道
	// 数组，结构体

	//【注意】变量和作用域有关系，不同的作用域变量是没有关联的，相互独立
	outer := 18

	{
		// 尽管这里再次声明outer变量
		// 编译一样通过，代码块中的作用域是相对独立的
		// 重复声明不会报错
		// 但是子作用域可以访问父作用域
		// 比如在这里不重新声明，而是重新赋值 outer = 16 那外面和里面打印都是16
		outer := 16
		log.Println(outer) // 打印16
	}

	log.Println(outer) // 打印18

}

func Sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}
