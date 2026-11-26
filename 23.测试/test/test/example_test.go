package test

import (
	"fmt"
	"learn/test/demo"
)

// 注意示例测试主要使用来展示如是使用的
// Output:这个注释不能少
// 示例测试不像其他三种测试一样是为了发现程序的问题所在
// 它更多的是展示某一个功能的使用方法，起到文档作用

// 关于Output有以下几种写法，第一种是只是一行输出
// 第二种是多行输出，即按顺序检测输出是否匹配

// 需要注意的是，对于测试函数而言，
// 仅当最后几行为Output注释才会被视为示例测试，
// 否则就只是一个普通的函数，不会被 Go 执行

// Hello函数示例测试
func ExampleHello() {
	demo.Hello()
	// Output:
	// hello 兹白
}

// GameStart函数示例测试
func ExampleGameStart() {
	demo.GameStart()
	// Output:
	// 原神启动!
}

// Sum[]函数示例测试
func ExampleSum() {
	sum := demo.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
	// Output:
	// 45
}
