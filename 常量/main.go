// go常量的学习记录
package main

import "log"

// 常量的值是不可变的，声明初始化后就不会改变
// 值的来源只有以下几种：
// 1、字面量
// 2、其他常量标识符如iota
// 3、常量表达式或者其他常量

// 在go中常量只能是基本类型，也就是前面数据类型排除派生类型剩下的类型

// 初始化，我这里将初始化写在包级别，声明后可以不使用
const (
	C1      = "我是字符串吧"
	C2      = 1
	C3      = 2.3
	C4      = '中'
	C5      = 1 + 2i
	C6 int8 = 127
	C7      = iota
	C8      = C1
)

// 常量声明了就必须初始化，不然编译不通过
// const s1 string 这是错误示例 错误: missing init expr for s1

//【注意】iota这个的使用，这个是按行增加的
const (
	Num  = iota // 0
	Num1        // 1
	Num2        // 2
	Num3        // 3
	Num4        // 4
	Num5 = 6    // 6
	Num6        // 6
	_           // 7 占位
	_           // 8 占位
	Num7        //9
)

// 【注意】iota和常量在go中可以模拟枚举，go中没有专门的枚举类型
type Season uint8

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

var seasonValue = []string{"Spring", "Summer", "Autumn", "Winter"}

func (s Season) String() string {
	if s > 3 {
		return "unkown"
	}
	return seasonValue[s]
}

// 上述就是定义的一个枚举，其实表现力已经够用了

// main函数
func main() {
	log.Println("常量")
}
