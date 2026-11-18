// go 结构体记录
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

// 结构体可以用将一组不同类型整合为一个类型
// 可以适应复杂数据场景，是一种复合类型
// go中弱化面向对象的概念，没有传统面向对象的继承和类的概念
// go中不是不能做面向对象，结构体就类似于类的作用，同时也有嵌入组合类似继承

// 创建一个结构体
// type 结构体名 struct
// 结构体同样遵循go的暴露方式
// 大写公有，小写私有
// 结构体中可以增加tag，序列化使用
type Character struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}

//【提示】声明结构体字段时，字段名不能与方法名重复

// 实例化
// go中实例化结构体很简单
func Demo1() {
	c := Character{
		Name:   "刻晴",
		Age:    23,
		Gender: "女",
		Height: 164,
		Weight: 48,
	}
	fmt.Println(c)
	// 除了这种方式实例化
	// 还可以省略字段名直接给值，不过这样就要全不给值，不建议使用
}

//【注意】前面函数中提到过go中不支持函数重载，所以不能想java中使用重载构造函数
// 一般是直接使用结构体名{}实例化，不过使用选项模式可以很优雅的实现
// 创建实例化不同的字段的结构体

// 创建选项类型，一个持有该类型指针的函数
type option func(*Character)

// 设置Name
func WithName(name string) option {
	return func(c *Character) {
		c.Name = name
	}
}

// 设置Age
func WithAge(age int) option {
	return func(c *Character) {
		c.Age = age
	}
}

// 设置Gender
func WithGender(gender string) option {
	return func(c *Character) {
		c.Gender = gender
	}
}

// 设置Height
func WithHeight(height float32) option {
	return func(c *Character) {
		c.Height = height
	}
}

// 设置Weight
func WithWeight(weight float32) option {
	return func(c *Character) {
		c.Weight = weight
	}
}

// 构造函数 option
// 返回 *Character
func NewCharacter(options ...option) *Character {
	res := &Character{}
	for _, opt := range options {
		opt(res)
	}
	return res
}

func main() {
	// 构建年龄和姓名示例 刻晴 23
	c1 := NewCharacter(WithName("刻晴"), WithAge(23))
	// 构建年龄，姓名，性别
	c2 := NewCharacter(WithName("甘雨"), WithAge(2000), WithGender("女"))
	// 构建身高和体重示例
	c3 := NewCharacter(WithHeight(160), WithWeight(50))
	fmt.Println(c1, c2, c3)
}

// 组合
// 这是go中扩展结构体的一种方式
// 和传统面向对象的继承不同，go中使用的组合也就是嵌入
type Person struct {
	name string
	age  int
}

type Student struct {
	stuID string
	Person
}

type Worker struct {
	workerID string
	Person
}

// 这就是组合嵌入
// 将一个结构图嵌入另一个结构体中
// 例如上面的学生就嵌入的人类结构体
// 那么学生结构体就拥有了人类的特性，如果人类实现了方法，学生也可以调用
// 最终实现的效果看起来和类大差不差
// 很多认为这就是go的类，但是不然，区别也很明显，这里不做过多的赘述
// 可以看到go中通过组合嵌入也能扩展结构体
// 【注意】这里版本变化也有区别，例如在1.27之前声明创建这种复合结构
var stu = Student{
	stuID: "stu001",
	Person: Person{
		name: "刻晴",
		age:  23,
	},
}

// 1.27就不用这么麻烦了，直接使用
var stu1 = Student{
	stuID: "stu001",
	name:  "刻晴",
	age:   23,
}

// 结构体指针
func Demo2() {
	stu := &Student{
		stuID: "stu001",
		name:  "刻晴",
		age:   23,
	}
	// 可以看到这里没有解引用也能访问结构体的字段值
	// 这是编译器底层做的优化，算是go的语法糖
	fmt.Println(stu.name, stu.age, stu.stuID)

}

// 标签
// 前面个也提到了标签，主要是在序列化中使用
// 【注意】go序列化结构体，该结构体必须公开，序列化的字段也是
type StructDemo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Demo3() {
	//这里就用json序列化为例子
	demo := StructDemo{
		Name: "刻晴",
		Age:  23,
	}
	b, err := json.Marshal(demo)
	if err != nil {
		log.Println(err.Error())
	}
	// 这里打印字段就是tag中写的标签
	fmt.Println(string(b))
}

// 内存对齐
// 在go中结构体字段的内存分布遵循内存对齐的规则
// 主要是为了减少cpu访问次数
// 这是一种空间换时间的机制
type Num struct {
	A int64 // 8字节
	B int32 // 4字节
	C int16 // 2字节
	D int8  // 1字节
	E int32 // 4字节
}

// 看起来是8+4+2+1+4 = 19
// 但是由于内存对齐，其实不是，会选择最大字段的整数倍，这里最大字段是8，所以是8的倍数才对
// A 不用补齐，B C D 加起来是7 补充一字节，E补充4字节 所以一共是24字节

// 下面结构体
type Num1 struct {
	A int8
	B int64
	C int8
}

// 你会发现调整字段的顺序，补齐的内存就会变化
// A B C 排列 A 补齐7字节，B 不补，C补齐7字节，补充了14字节，但是ABC一共才10字节
// A C B 排列 A C 补齐6字节
// 合理的排序可以节省内存，现在是AI时代，自己算起来麻烦，但是交给AI做就方便多了
// 这是优化系统占用的一种手段

// 空结构体
// 之前提到过空结构体，不占用内存空间
func Demo4() {
	type Empty struct{}
	fmt.Println(unsafe.Sizeof(Empty{})) // 0
}

// 泛型
// 这里只是提及不会深入介绍，后面会详细说明泛型的
// 1.18之后引入了泛型，结构体也就可是用了，创建通用性更强的结构体
// 不过匿名结构体不能使用泛型
// 这就是创建了一个通用的Set集合结构体
type Set[T comparable] struct {
	Data map[T]struct{}
}

// List通用
type List[T any] struct {
	Data []T
}
