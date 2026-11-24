// go 反射记录
package main

import (
	"fmt"
	"reflect"
	"unsafe"
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
// type Kind uint

// const (
// 	Invalid Kind = iota
// 	Bool
// 	Int
// 	Int8
// 	Int16
// 	Int32
// 	Int64
// 	Uint
// 	Uint8
// 	Uint16
// 	Uint32
// 	Uint64
// 	Uintptr
// 	Float32
// 	Float64
// 	Complex64
// 	Complex128
// 	Array
// 	Chan
// 	Func
// 	Interface
// 	Map
// 	Pointer
// 	Slice
// 	String
// 	Struct
// 	UnsafePointer
// )

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

	// reflect.Value.Type获取反射值的类型
	// 这个获取的是reflect.Type类型，上面说的都可以使用
	fmt.Println(reflect.ValueOf(str).Type()) // string

	// reflect.Value.Elem 获取反射值的元素反射值
	// 不能使用值类型不然就会运行时panic
	fmt.Println(reflect.ValueOf(&str).Elem())

	// 指针
	// 获取反射值指针的方式有两种
	// 返回一个表示v地址的指针反射值
	// func (v Value) Addr() Value

	// 返回一个指向v的原始值的uinptr 等价于 uintptr(Value.Addr().UnsafePointer())
	// func (v Value) UnsafeAddr() uintptr

	// 返回一个指向v的原始值的uintptr
	// 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
	// func (v Value) Pointer() uintptr

	// 返回一个指向v的原始值的unsafe.Pointer
	// 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
	// func (v Value) UnsafePointer() unsafe.Pointer

	fmt.Println("=========反射指针=========")
	num2 := 1024
	ele := reflect.ValueOf(&num2).Elem()

	fmt.Println("&num", &num)
	fmt.Println("Addr", ele.Addr())
	fmt.Println("UnsafeAddr", unsafe.Pointer(ele.UnsafeAddr()))
	fmt.Println("Pointer", unsafe.Pointer(ele.Addr().Pointer()))
	fmt.Println("UnsafePointer", ele.Addr().UnsafePointer())
	// 输出结果
	// &num 0x3fa9a98e4070
	// Addr 0x3fa9a98c61b0
	// UnsafeAddr 0x3fa9a98c61b0
	// Pointer 0x3fa9a98c61b0
	// UnsafePointer 0x3fa9a98c61b0

	//【注意】fmt.Println()会反射获取参数类型，如果是reflect.Value类型的话
	// 会自动调用Value.Interface()来获取其原始值

	// 设置值
	// func (v Value) Set(x Value)
	// 想要通过反射来修改反射值，那么其值必须要时可取值的，通过指针来修改值
	// 不是直接修改值
	fmt.Println("=======设置值=======")
	nptr := new(10)
	ele1 := reflect.ValueOf(nptr).Elem()
	fmt.Println(ele1.Interface())
	ele1.SetInt(3345)
	fmt.Println(ele1.Interface())

	// 获取值
	fmt.Println("=========获取值=========")
	str2 := "刻晴啊啊啊"
	rv2 := reflect.ValueOf(str2)
	// reflect.TypeAssert[]() 相当于rv2.Interface().(string)
	if v, ok := reflect.TypeAssert[string](rv2); ok {
		fmt.Println("类型断言成功", v)
	}

	// 函数
	// 上面都是获取变量的类型和值，通过指针修改反射值之类的
	// 下面介绍如果通过反射机制获取函数
	fmt.Println("========函数=========")
	rFn := reflect.TypeOf(Max[int])
	// 输出函数名称,字面量函数的类型没有名称
	fmt.Println("函数名称", rFn.Name())
	// 输出参数，返回值的数量
	fmt.Println(rFn.NumIn(), rFn.NumOut())
	rpt := rFn.In(0)
	fmt.Println("第一个参数类型", rpt.Kind())
	rot := rFn.Out(0)
	fmt.Println("第一个返回值类型", rot.Kind())
	// 调用
	// 通过反射值来调用函数
	rrv := reflect.ValueOf(Max[int]).
		Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(13)})
	// 因为返回值可以有多个，所以返回的是一个切片
	for _, v := range rrv {
		fmt.Println("返回值", v)
	}

	// 结构体
	// 访问字段
	// 1、通过索引来访问
	fmt.Println("=========结构体========")
	rt3 := reflect.TypeFor[*Person]().Elem()
	fmt.Println("结构体字段数量", rt3.NumField())
	for sf := range rt3.Fields() {
		fmt.Println(sf.Index, sf.Name, sf.Type, sf.Offset, sf.IsExported())
	}
	// 2、通过名称来访问
	if sf, ok := rt3.FieldByName("money"); ok {
		fmt.Println(sf.Index, sf.Name, sf.Type, sf.Offset, sf.IsExported())
	}

	// 修改字段
	// 想要修改字段，必须传入一个结构体指针
	p := &Person{Name: "刻晴", Age: 23, Address: "璃月", money: 13659456}
	rv4 := reflect.ValueOf(p).Elem()
	// 获取字段
	age := rv4.FieldByName("Age")
	empty := reflect.Value{}
	// 修改字段
	if age != empty {
		// 如果等于字段就不存在
		age.SetInt(24)
	}

	fmt.Println("输出结构体", rv4.Interface())

	// 对于私有字段而言操作要多一点
	money := rv4.FieldByName("money")
	if money != empty {
		// 构造指向该结构体未导出字段的指针反射值
		p := reflect.NewAt(money.Type(), money.Addr().UnsafePointer())
		// 获取该指针所指向的元素，也就是要修改的字段
		field := p.Elem()
		field.SetInt(52013141)
	}
	fmt.Println("输出结构体", rv4.Interface())

	// 访问tag
	// 获取到结构体字段后，便可以直接访问其tag
	name, ok := rt3.FieldByName("Name")
	if ok {
		// LookUp返回值和bool，不存在bool是false
		fmt.Println(name.Tag.Lookup("json"))
		// 返回值，不存在为空
		fmt.Println(name.Tag.Get("json"))
	}

	// 访问方法
	// 访问方法与访问字段的过程很相似，只是函数签名略有区别
	// type Method struct {
	// 	// 方法名
	// 	Name string
	// 	// 包名
	// 	PkgPath string
	// 	// 方法类型
	// 	Type Type
	// 	// 方法对应的函数，第一个参数是接收者
	// 	Func Value
	// 	// 索引
	// 	Index int
	// }
	fmt.Println("方法个数", rt3.NumMethod())
	// 遍历方法信息
	for m := range rt3.Methods() {
		fmt.Println(m.Index, m.Name, m.Type, m.IsExported())
		fmt.Println("方法参数")
		for in := range m.Func.Type().Ins() {
			fmt.Println(in.String())
		}
		fmt.Println("方法返回值")
		for out := range m.Func.Type().Outs() {
			fmt.Println(out.String())
		}
	}

	// 调用方法
	// 调用方法和调用红函数过程类似，都是使用Call来调用，返回值都是一个切片
	talk := reflect.ValueOf(p).Elem().MethodByName("Talk")
	if talk != empty {
		res := talk.Call([]reflect.Value{reflect.ValueOf("刻晴天下第一可爱!")})
		for _, r := range res {
			fmt.Println(r.Interface())
		}
	}

	// 创建类型
	// 通过反射可以构建新的值，reflect包同时根据一些特殊的类型
	// 提供了不同的更为方便的函数
	fmt.Println("=======创建类型=======")
	rv5 := reflect.New(reflect.TypeFor[string]())
	rv5.Elem().SetString("刻晴的玉足")
	fmt.Println(rv5.Elem().Interface())

	rt5 := reflect.TypeFor[*Person]().Elem()
	p1 := reflect.New(rt5).Elem()
	fmt.Println(p1.Interface())
	// 可创建的类型除了基本类型还有介科狗提，切片，Map，管道，函数

	// 完全相等
	// 是反射包下提供的一个个用于判断两个变量是否完全相等函数
	// 数组：数组中的每一个元素都完全相等
	// 切片：都为nil时，判为完全相等，或者都不为空时，长度范围内的元素完全相等
	// 结构体：所有字段都完全相等
	// 映射表：都为nil时，为完全相等，都不为nil时，每一个键所映射的值都完全相等
	// 指针：指向同一个元素或指向的元素完全相等
	// 接口：接口的具体类型完全相等时
	// 函数：只有两者都为nil时才是完全相等，否则就不是完全相等
	fmt.Println("======完全相等======")
	a := make([]int, 0, 12)
	b := make([]int, 0, 12)
	fmt.Println(reflect.DeepEqual(a, b))

	keqing := Person{
		Name: "刻晴",
		Age:  39,
	}

	ganyu := Person{
		Name: "甘雨",
		Age:  18,
	}

	ningguang := Person{
		Name: "凝光",
		Age:  18,
	}
	fmt.Println(reflect.DeepEqual(keqing, ganyu))
	fmt.Println(reflect.DeepEqual(ningguang, ganyu))
	fmt.Println(reflect.DeepEqual(ganyu, ganyu))

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

type Number interface {
	Int | Uint | float32 | float64
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// 定义一个比较最大值的函数，使用泛型有一定的通用性
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	money   int
}

func (p Person) Talk(msg string) string {
	return msg
}
