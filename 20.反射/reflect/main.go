// go 反射记录
package main

import (
	"fmt"
	"reflect"
)

// 前面在类型的记录中，说到过go是一个静态强类型的语言
// 尽管静态语言在类型安全上有着天然的优势，适用于工程化项目
// 但是go和大多数静态语言一样，提供反射机制，提供一定程度上的运行时动态操作
// 例如java，这个语言和go一样都是静态语言，同样提供发射机制
// 通过反射机制实现了很多的框架，例如知名web框架Spring
// go中也是一样的，凡是反射带来的运行时消耗可不小，主要是类型检查带来的
// 正常情况，编译时就将类型确定好了，后续的运行不会有类型检查的消耗了
// 但是发射就是提供在运行时动态特性，为了安全肯定在运行时有类型检查，运行时开销增大了

// 反射是一种在运行时检查语言自身结构的机制
// 特点就是给静态语言带来了动态语言的特性，缺点就是性能问题
// go中反射与空接口interface{}（也就是any）有关系
// 很大程度上有空接口出现的地方基本上就有反射

// 空接口和反射的关系
// 1、空接口可以转换为反射对象
// 2、反射对象可以还原成空接口
// 3、要修改反射对象，其值必须要能修改

// 这里是从reflect包中复制过来的
// 就是所有的顶层类型
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

func main() {
	// 在go中，理解发射不是简单的事，毕竟设计的的确复杂晦涩，但是使用就简单了
	// reflect.Type和reflect.Value
	// 基本上变量就是这两个类型就可以表示了，确定类型和值

	// 函数TypeOf和ValueOf就可以将空接口变量转换为对应的反射类型和值
	var num any = 100_000
	rt := reflect.TypeOf(num)
	fmt.Println("反射类型", rt) // 反射类型 int
	rv := reflect.ValueOf(num)
	fmt.Println("反射值", rv) // 反射值 100000

	// 注意TypeOf获取的类型不是基本类型
	type myint int
	var i myint = 10
	fmt.Println("反射类型", reflect.TypeOf(i)) // 反射类型 main.myint

	// reflect.Type.Kind 是获取底层类型的
	fmt.Println("Kind类型", reflect.TypeOf(i).Kind()) // Kind类型 int

	// reflect.Type.Elem 是获取像数组，切片，map，通道，指针这种元素类型的
	// 除了这些元素，是不能用的，用就会panic
	var mp = map[int]string{}
	fmt.Println("mp值类型", reflect.TypeOf(mp).Elem()) // Elem类型是 string

	// reflect.Type.Size 是获取对应类型所占的字节大小
	// reflect.TypeOf(10).Size() 这种方式可以使用下面的泛型方式
	fmt.Println("========获取字节大小========")
	fmt.Println(reflect.TypeFor[int]().Size())       // 8
	fmt.Println(reflect.TypeFor[string]().Size())    // 16
	fmt.Println(reflect.TypeFor[complex64]().Size()) // 8
	fmt.Println(reflect.TypeFor[float64]().Size())   // 8
	fmt.Println(reflect.TypeFor[[]string]().Size())  // 24

	// reflect.Type.Comparable 是用来判断类型是否是可比较的
	// reflect.TypeOf("字符串").Comparable() 使用下面的泛型方式
	fmt.Println("========判断比较========")
	fmt.Println(reflect.TypeFor[string]().Comparable())   // true
	fmt.Println(reflect.TypeFor[int]().Comparable())      // true
	fmt.Println(reflect.TypeFor[[]int]().Comparable())    // false
	fmt.Println(reflect.TypeFor[struct{}]().Comparable()) // true

	// reflect.Type.Implements 判断一类型是否实现接口
	fmt.Println("========判断实现========")
	rIface := reflect.TypeFor[*MyInterface]().Elem()
	rIface1 := reflect.TypeFor[*fmt.Stringer]().Elem()
	fmt.Println(reflect.TypeFor[*MyStruct]().Elem().Implements(rIface))   // true
	fmt.Println(reflect.TypeFor[*HisStruct]().Elem().Implements(rIface1)) // true
	fmt.Println(reflect.TypeFor[*HisStruct]().Elem().Implements(rIface))  // false

	// refelct.Type.ConvertibleTo 判断一个类型是否可以转换为另一个类型
	fmt.Println("=======判断类型转换=======")
	fmt.Println(reflect.TypeFor[*MyStruct]().Elem().ConvertibleTo(rIface))
	fmt.Println(reflect.TypeFor[*HisStruct]().Elem().ConvertibleTo(rIface))

	// reflect.Value 反射类型的值
	// 其代表着反射接口的值，使用reflect.ValueOf()函数可以将变量转换成reflect.Value
	fmt.Println("=======reflect.Value=======")
	str := "绯雪的小绯雪"
	fmt.Println(reflect.ValueOf(str))

	// reflect.Valut.Type获取反射值的类型
	// 这个获取的是reflect.Type类型，上面说的都可以使用
	fmt.Println(reflect.ValueOf(str).Type()) // string

	// 今天暂时到这里了

}

type MyInterface interface {
	My() string
}

type MyStruct struct {
}

func (m MyStruct) My() string {
	return "my"
}

type HisStruct struct {
}

func (h HisStruct) String() string {
	return "his"
}
