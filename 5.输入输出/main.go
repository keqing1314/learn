// go 输入输出记录
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 有过其他编程语言学习经历，肯定了解输入输出
// 一般分为：
// 标准输入
// 标准输出
// 标准错误
// 在go中也是如此
// 在os包下有三个暴露的文件描述符
// os.Stdin 标准输入
// os.Stdout 标准输出
// os.Stderr 标准错误

// main 函数
func main() {
	// go中存在多种输出方法
	// 标准输出本身就是一个文件
	os.Stdout.WriteString("我是玉衡星刻晴")

	// print和println
	// 这是go内置的两个函数，用于将参数输出到标准错误中（不推荐使用）
	print("你好\n")
	println("我是刻晴")

	// fmt
	// 其实输出常用的包就有fmt，就是格式化的缩写
	// fmt.Println() 该函数默认会将参数输出到标准输出中
	// 因为内部使用了反射机制，所以不推荐大量使用
	fmt.Println("我是甘雨")

	// bufio
	// bufio中提供了可缓冲的输出方法，也就是将内存写入到内存中
	// 积累到一定的阈值才会输出到指定的Writer中，默认缓冲大小事4KB
	// 通常在文件IO,网络IO场景中使用这个包
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.WriteString("你好呀，我是刻晴")

	// 格式化
	// 上面提到了fmt，其实重要的用法事格式化输出
	// 这就是一个格式化输出
	fmt.Printf("我的年龄是%d", 23)

	// 0	格式化	描述	接收类型
	// 1	%%	输出百分号%	任意
	// 2	%s	输出string/[] byte值	string,[] byte
	// 3	%q	格式化字符串，输出的字符串两端有双引号""	string,[] byte
	// 4	%d	输出十进制整型值	整型
	// 5	%f	输出浮点数	浮点
	// 6	%e	输出科学计数法形式 ,也可以用于复数	浮点
	// 7	%E	与%e相同	浮点
	// 8	%g	根据实际情况判断输出%f或者%e,会去掉多余的 0	浮点
	// 9	%b	输出整型的二进制表现形式	数字
	// 10	%#b	输出二进制完整的表现形式	数字
	// 11	%o	输出整型的八进制表示	整型
	// 12	%#o	输出整型的完整八进制表示	整型
	// 13	%x	输出整型的小写十六进制表示	数字
	// 14	%#x	输出整型的完整小写十六进制表示	数字
	// 15	%X	输出整型的大写十六进制表示	数字
	// 16	%#X	输出整型的完整大写十六进制表示	数字
	// 17	%v	输出值原本的形式，多用于数据结构的输出	任意
	// 18	%+v	输出结构体时将加上字段名	任意
	// 19	%#v	输出完整 Go 语法格式的值	任意
	// 20	%t	输出布尔值	布尔
	// 21	%T	输出值对应的 Go 语言类型值	任意
	// 22	%c	输出 Unicode 码对应的字符	int32
	// 23	%U	输出字符对应的 Unicode 码	rune,byte
	// 24	%p	输出指针所指向的地址	指针
	fmt.Printf("%%%s\n", "hello world")

	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")
	fmt.Printf("%d\n", 2<<7-1)

	fmt.Printf("%f\n", 1e2)
	fmt.Printf("%e\n", 1e2)
	fmt.Printf("%E\n", 1e2)
	fmt.Printf("%g\n", 1e2)

	fmt.Printf("%b\n", 2<<7-1)
	fmt.Printf("%#b\n", 2<<7-1)
	fmt.Printf("%o\n", 2<<7-1)
	fmt.Printf("%#o\n", 2<<7-1)
	fmt.Printf("%x\n", 2<<7-1)
	fmt.Printf("%#x\n", 2<<7-1)
	fmt.Printf("%X\n", 2<<7-1)
	fmt.Printf("%#X\n", 2<<7-1)

	type person struct {
		name    string
		age     int
		address string
	}
	fmt.Printf("%v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%#v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%t\n", true)
	fmt.Printf("%T\n", person{})
	fmt.Printf("%c%c\n", 20050, 20051)
	fmt.Printf("%U\n", '码')
	fmt.Printf("%p\n", &person{})

	// 输入，接下来介绍下，go中是如何输入的

	// read 向度文件一样，读取输入内容
	// 但是一般不这样使用，很麻烦，有更简单的方式
	var buf [1024]byte
	n, _ := os.Stdin.Read(buf[:])
	os.Stdout.Write(buf[:n])

	// fmt
	// 直接使用fmt包中提供的函数
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)

	n1 := 10
	s := make([]int, n1)
	for i := range n1 {
		fmt.Scan(&s[i])
	}
	fmt.Println(s)
	// 1 2 3 4 5 6 7 8 9 10
	// [1 2 3 4 5 6 7 8 9 10]

	// bufio
	// 这个提供的函数应该挺多的
	reader := bufio.NewReader(os.Stdin)
	var a1, b1 int
	fmt.Fscanln(reader, &a1, &b1)
	fmt.Printf("%d + %d = %d\n", a1, b1, a1+b1)

	// scanner
	// 与bufio.NewReader(os.Stdin)类似
	// 不同的是，这个是按行读取的，上面的用的是缓冲区
	// 默认单行上限 64KB
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		fmt.Println("scan", line)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("err:%v", err)
	}
}
