// go 数据类型学习记录
package main

import (
	"fmt"
	"log"
)

// main函数
func main() {
	// go 中的数据类型，整体类型不多就8种左右
	// 大致可以分为：布尔类型，整型，浮点型，负数类型，字符类型，派生类型
	// 接下来一次详细介绍每种分类：

	// 【布尔类型】
	// 关键字 bool
	var f = true
	log.Println(f == false)
	// bool类型只有两种值，true（真）和false（假）
	// 在 Go 中，整数 0 并不代表假值，非零整数也不能代表真值，
	// 即数字无法代替布尔值进行逻辑判断，两者是完全不同的类型。

	// 【整型】
	// 在go中区分有符号和无符号，另外指针也可以使用整型，所以类型如下
	// uint 无符号整型 至少32位和所用的机器有关
	// uint8 无符号8位整型 1字节
	// uint16 无符号16位整型 2字节
	// uint32 无符号32位整型 4字节
	// uint64 无符号64位整型 8字节
	// int 有符号整型 至少32位和所用的机器有关
	// int8 有符号8位整型 1字节
	// int16 有符号16位整型 2字节
	// int32 有符号32位整型 4字节
	// int64 有符号64位整型 8字节
	// uintptr 等价于无符号 64 位整型，但是专用于存放指针运算，用于存放死的指针地址
	var (
		uint1 uint   = 1
		uint2 uint8  = 128
		uint3 uint16 = 256
		uint4 uint32 = 456
		uint5 uint64 = 123456789
	)
	log.Println(uint1 + uint(uint2) + uint(uint3) + uint(uint4) + uint(uint5))
	var (
		int1 int   = 1
		int2 int8  = -128
		int3 int16 = 256
		int4 int32 = -456
		int5 int64 = 123456789
	)
	log.Println(int1 + int(int2) + int(int3) + int(int4) + int(int5))
	// 使用uintptr
	var ptr uintptr
	log.Println(ptr)

	// 【浮点型】
	// IEEE-754浮点数，主要区分单精度和双精度
	// java中就是float和double
	// 在go中就是float32和float64
	// 我个人的感觉这种方式更好记忆，而且见名知义，内存大小也清晰32位4字节和64位8字节
	var (
		f1 float32 = 0.6
		f2 float64 = 0.99
	)
	log.Println(float64(f1) + f2)

	// 【负数类型】
	// 这个很少见，使用场景也是，常见的场景就是科学计算了，我基本没有使用过这个类型
	// complex128 	64 位实数和虚数
	// complex64 32 位实数和虚数
	var (
		c1 complex64  = 2 + 2i
		c2 complex128 = 1 + 2i
	)
	log.Println(c1 + complex64(c2))

	// 【字符类型】
	// 在go中最常见的类型之一就是字符类型，公众字符串就是字节序列，类型如下
	// byte  等价于uint8 用来表达 ANSCII 字符
	// rune  等于int32 用来表达UTF-8字符，合法的UTF-8字符最大的也就是4字节（32位正好）
	// string 字节序列也就是字符串，可以转换位byte切片
	var (
		s1 byte   = 'a'
		s2 rune   = '中'
		s3 string = "我是字符串"
	)
	log.Printf("字节:%v 字符:%v 字符串:%v", s1, s2, s3)

	// 【派生类型】
	// 这个类型就很广泛了，像是数组，切片，通道，接口，还有struct结构体等等都是这一类的

	// 数组：声明[5]int 也就是[数组长度]类型
	var array = [5]int{1, 2, 3, 4, 5}
	// 数组的声明如上，注意的是数组长度不能是变量，也就是编译时期就得确定
	// 字面量或者常量和常量表达式
	// 数组的访问就是通过下标访问的
	log.Println(array[0])

	// 切片：声明[]int 也就是[]int
	// 可以是使用字面量初始化和make初始化，如果声明没有初始化，就是nil
	var slice = []int{1, 2, 3}
	// 使用make函数创建初始化make([]类型，len，cap)
	// 后面会详细介绍，这里只是数据类型介绍
	var slice2 = make([]int, 0, 5)
	log.Printf("%v %v", slice, slice2)

	// 映射：map
	// 同样也是可使用字面量初始化和make初始化，只声明值就是nil
	var mp = map[string]any{"keqing": "我是刻晴"}
	// 访问方式，注意如果map的key不存在，返回的就是默认类型的零值
	// 但是如果值就是零值的如何判断，其实map的访问提供两个返回值，value和一个bool值标记是否有值
	description := mp["keqing"]
	log.Println(description)

	// 结构体：struct
	type notes struct {
		data  string
		size  int
		title string
	}
	// notes就是一个标准的结构体
	// 声明方式就是type 结构体名 struct，支持匿名结构体
	p := struct {
		Name string
		Age  int
	}{
		Name: "Tom",
		Age:  18,
	}
	fmt.Println(p.Name, p.Age)

	// 指针：
	// 在go中保留了指针，但是限制了操作可以取地址&，和解引用*, 但是不支持指针运算
	num := 1
	ptr1 := &num
	// 简单说，取地址就是获取该变量的内存地址，解引用就是通过地址取到值，默认值是nil
	log.Printf("num=%d, 指针:%v, 解引用:%v", num, ptr1, *ptr1)

	// 函数：
	// 在go中函数是一种类型：type fn func(string)string
	// 这就是定义了一个函数入参是string，返回值是string的函数类型，函数的默认值也是nil
	type fn func(string) string
	fn1 := func(s string) string {
		return fmt.Sprintf("我是%v", s)
	}
	// 调用函数
	fn1("刻晴")

	// 接口的使用：
	var d duck = yollowDuck{}
	d.fly("我会飞")

	// 通道：chan
	// 这是go中特有的类型，用于协程建通信的一种方式
	// 创建可用的通道有且只有一种方式：make
	ch := make(chan int)
	// 创建一个无缓冲的通道,像其中写入一个值
	ch <- 1
	go func() {
		// 从另一个协程中读取
		// 这里只是一个简单的示例，channel，在协程中和重要，后续会详细介绍的
		_ = <-ch
	}()

	// 【注意点】前面我说到了零值，这个算是go的一大特点了
	// go中声明了类型，但是没有初始化，编译器会给一个该类型的默认零值
	// 例如bool的零值是false，int类型的零值是0，函数类型，指针类型，切片类型，通道类型，映射类型，接口类型的零值是nil
	// 所以不能给一个类型推断为nil, 如果要设用nil，就必须指明类型，不然编译器推断不出
}

// 接口：interface
// go中的接口声明式type duck interface
// 【注意】接口的默认值也是nil
type duck interface {
	fly(string)
}

// 这就是定义了一个duck接口，有一个fly方法，如果一个类型要实现这个接口
// 只需要实现这个结构的所有方法，不需要通过关键词来显示实现，go中实现接口是隐式实现
// 【注意】当一个类型实现了接口后就可以使用接口类型存储该类型
type yollowDuck struct{}

// 这个类型就实现了上面的duck接口
func (d yollowDuck) fly(s string) {
	log.Println(s)
}
