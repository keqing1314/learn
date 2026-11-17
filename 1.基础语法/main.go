// 本内容是go的基本语法的示例
package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// 可见性，在go中使用大小写区分可见性
// java中使用访问修饰符 private public区分私有还是公有
// go中大写字母开头就是公有的 变量/函数/方法
// 小写字母或者下划线开头的就是私有的，变量/函数/方法
var one int // 这里声明了一个int类型的变量one
// 该变量是私有的，仅在本包中可见

var Two int // 这里声明了一个int类型的变量Two
// 该变量是公有的，在外部包中也可访问

// 【注意】这里的可见性不只是变量，函数，方法这些都使用
// 这是go中的规则

// 在go中函数是构建项目的基石
// 声明一个函数
// func 函数名(...函数入参)(...函数返回值)
// 这是main函数，是go程序的入口函数
func main() {
	one = 10
	// 这里是将变量one打印输出
	// 在这使用了fmt和log两个包
	// 上面可以看到
	// 	import (
	// 	"fmt"
	// 	"log"
	// )
	// 这就是包导入，多个包就是用括号括起来
	// 【注意】在go中对于循环依赖零容忍
	// 不存在任何情况的循环依赖，
	// packageA -> packageB
	// packageB -> packageC
	// packageC -> packageA
	// 或者
	// packageA -> packageB
	// packageB -> packageA
	// 上述情况编译都不会通过
	// 【约定】如果一个包名是internal，那么外部包无法访问内部包
	// 当让内部包的祖包是可以访问的
	f := fmt.Sprintf("变量one=%d", one)
	log.Println(f)

	// 标识符
	// 标识符就是一个名称，用于包名，变量名，函数名
	// 命名规则如下
	// 1、只能由字母(不只是英文)，数字，下划线组成
	// 2、只能是字母或下划线开头
	// 3、严格区分大小写，可见性相关
	// 4、不能与任何已经存在的标识符重复，保证包内唯一
	// 5、不能和go中内置的关键字重复
	// 下面的就是合法的
	// var (
	// 	变量    = "变量"
	// 	num   = 100
	// 	Age   = 10
	// 	_name = "刻晴"
	// 	nil   = "nil"
	// )
	// 【注意】go中内置的关键字
	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var
	// 【注意】nil这个不是关键字，很容易误认为这个是关键字

	// 运算符
	// Go 语言中支持的运算符号的优先级排列
	// Precedence    Operator
	//     5             *  /  %  <<  >>  &  &^
	//     4             +  -  |  ^
	//     3             ==  !=  <  <=  >  >=
	//     2             &&
	//     1             ||
	//【注意】go 语言中没有选择将~作为取反运算符，而是复用了^符号
	a := 1
	b := 2
	log.Printf("这是异或操作：%v", a^b)
	log.Printf("这是取反操作: %v", ^b)

	// 【注意】go中不存在 ++i --i 这种写法只有 i++ i--
	a++
	b--

	// 字面量
	// 顾名思义就是字面上的值
	// 常用于变量赋值
	var username = "keqing1314"
	var age = 23
	const gender = "女"
	log.Printf("用户名:%s 年龄:%d 性别:%v", username, age, gender)
	// 上面就是典型的字面量使用
	// 变量和常量都可以使用字面量，可以不指明类型，编译器在编译时会推断类型
}

// 函数
// 上面提到了main函数
// 函数的声明是通过func关键字来声明的
// 这里声明一个累加函数，函数名式Sum，可见性是公开，返回值类型是int和error
// 调用在需要使用该函数的地方调用
// total,err := Sum(10,23,45,67)
// 【注意】函数的参数列表nums ...int是可变参数，可以传递多个同类型的变量值
// 同样的，返回值也可以接受多个返回值，go中的特点就是允许一个函数有多个返回值
// 具体是通过返回值入栈实现的，细节自行了解
// 这里接受int类型和error类型的两个返回值
func Sum(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("参数不能为空")
	}

	var total int
	for _, v := range nums {
		num := v
		total = total + num
	}
	return total, nil
}

// 编码风格，这个在go中自带格式化工具gofmt，可以自己手动运行
// 使用现代编辑器会自己集成该格式化插件 vscode goland nvim之类的
// 这一点很好，程序员可以更加专注于代码实现，减少对于编码风格的心智负担
// 同时不同gopher的代码编码风格是一样的，可读性更好
// func Foo()
// {
// 	log.Println("我是错误示例")
// }
// 上述给了一个错误示例
// 编译器会报错误: unexpected semicolon or newline before {

// 【注意】在go中，语句条件不需要(),但是{}不可省略
func Demo(a string) {
	if strings.TrimSpace(a) == "" {
		log.Println("参数为空")
	}
	// 上面的示例中if的条件不要使用(),但是{}不能省略，哪怕只有一行代码
}

// 【注意】在go中，另一点就是不存在三元运算表达式这种写法
// var username = name == "" ? "keqing" : name
// 这种写法是错误的，直接使用if分支语句就行

// 以上是go语法的一个基础介绍，后续会详细介绍
