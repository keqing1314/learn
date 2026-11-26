// nil的实现原理
package main

import (
	"fmt"
	"reflect"
)

// 前面提到过nil是引用类型的零值
// 例如map 切片 管道 函数 指针 接口
// 前面5种都可以理解为指针指向底层数据地址
// 但是接口不太一样，它是指向地址和类型，要两个都是nil才是nil
// 其中一个是nil，另一个不是就不是nil

func main() {
	// 举个例子
	var p *int
	fmt.Println(p)        // nil
	fmt.Println(p == nil) // true
	var pa any
	pa = p
	fmt.Println(pa)        // nil
	fmt.Println(pa == nil) // false
	// 结果很奇怪呀，为什么打印的是nil但是不等于nil
	// 这是因为打印的是值，你看类型就不是了
	fmt.Println("p type", reflect.TypeOf(p))   // *int
	fmt.Println("p value", reflect.ValueOf(p)) // nil

	fmt.Println("pa type", reflect.TypeOf(pa))   // *int
	fmt.Println("pa value", reflect.ValueOf(pa)) // nil
	// 可以看到空接口的值是nil而类型是*int
	// 前面说到，对于空接口必须类型和值都是nil才是nil
	// 比较空接口类型是否相等，相比较类型，然后比较值
	// 但是空接口的零值是类型为nil，值为nil，但是上述类型是*int自然是false

	// 如果只想判断其值是不是就可意识用reflect.Value.IsNil的方法
	fmt.Println(reflect.ValueOf(pa).IsNil()) // true

	// 总结下
	// 对于接口类型判断是否为nil，要求是类型和值都为nil
	// 如果只判断接口值是否为nil可以使用reflect.Value.IsNil方法
	// 注意前面提到过的nil零值的注意
	// 当切片为nil的时候，可以访问它的长度和容量，也可以对其添加元素

	// 当 map 为nil的时候，还可以对其进行访问，
	// 但nil的 map 是只读的，一旦尝试写入就会引发panic

	// nil不能用来类型推断，例如
	// 错误: use of untyped nil in variable declaration
	// var mp = nil

	var mp map[string]string = nil // 这样倒是可以，不过没必要，声明不赋值就是零值nil，这个赋值为nil多此一举
	fmt.Println(reflect.TypeOf(mp), reflect.ValueOf(mp))
	fmt.Println(mp == nil)
}
