// go 泛型记录
package main

import "errors"

// 泛型在go中是1.18正式引入的
// 之前一直没有泛型，主要还是其团队对于这个技术存在分歧
// go一直是以编译速度快来作为自己的承诺
// 但是泛型的主流方案
// 生成模板代码降低编译速度，避免运行开销或者存储数据字段，用一套代码，降低编译耗时
// 但是两者各自都有一定的问题，go采用的是两者结合
// 其实这是很有问题的，这样编译会倍拖慢，又有运行时开销
// 如果把握不好，就是两种缺点都有
// 果不其然，1.18刚更新，编译速度相比以前大幅增加，运行时开销也不小
// 不过迭代了几个版本，现在优化了不少，一个折中的方案

// 来一个简单的引子
// 我需要使用切片，问题是切片基本没有封装的方法，这里实现几个方法
type myslice []int

func (s myslice) Add(num int) {
	s = append(s, num)
}

func (s myslice) Pop() int {
	return s[0]
}

// 但是如果你换一个类型就又要重新实现这些方法，重复逻辑，就只是存储的类型不同
// 但是有了泛型就方便很多
type myslice2[T any] []T

func (s myslice2[T]) Add(v T) {
	s = append(s, v)
}

func (s myslice2[T]) Pop() T {
	return s[0]
}

// 这样就得到了一个通用的mysilce类型
//【注意】这个类型是自定义的切片类型，和go自带的有区别的
// 如果要转换需要强制类型转换才行

// 这里使用泛型实现一个stack通用类型
// Stack是一个先进后出的数据结构
type Statck[T any] struct {
	data []T // 存储具体元素的切片
	size int // 栈的容量
}

// 这里是在main包中，才这样命名的
// 如果是在Stack保重，我就会命名成New
func NewStack[T any](size int) *Statck[T] {
	return &Statck[T]{size: size}
}

// Push 方法添加元素
func (s *Statck[T]) Push(v T) error {
	if s.size >= len(s.data) {
		return errors.New("栈容量已满")
	}
	s.data = append(s.data, v)
	s.size++
	return nil
}

// Pop 方法将栈顶元素弹出并返回
func (s *Statck[T]) Pop() (T, error) {
	var res T
	if s.size == 0 {
		return res, errors.New("栈为空")
	}
	res = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.size--
	return res, nil
}

// Peek 返回栈顶元素
func (s *Statck[T]) Peek() (T, error) {
	var res T
	if s.size == 0 {
		return res, errors.New("栈为空")
	}
	res = s.data[len(s.data)-1]
	return res, nil
}

// Size 返回大小
func (s *Statck[T]) Size() int {
	return s.size
}

// 这两个示例，体现除了泛型对于编写通用代码的优势，这也是使用泛型的好处
// go中泛型的另一个就是类型集
// 类型集 就是接口 通用接口
type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// 这就是一个类型集，可以用在泛型中使用
// 这就是使用无符号类型集泛型的求和函数
// 对于有些场景，也是要限制泛型的类型的，不一定是所有的类型都可以使用
// 这时候定义个类型集就是很好的方式
// 比如这里我就想整一个无符号的求和函数，使用any肯定不行吧
func Sum[T Uint](nums ...T) T {
	var total T = 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

// 类型集存在并集和交集两种格式
// 上卖弄的Uint类型集就是一个典型的并集
// 下面展示一个交集的示例
// Int类型集和Uint类型集没有交集，所以这个Zero就是一个空的类型集
// 实际项目开发中，我目前还没有使用交集来限定类型集
type Zero interface {
	Int
	Uint
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

// 使用 ~ 符号，来表示底层类型，
// 如果一个类型的底层类型属于该类型集，那么该类型就属于该类型集，如下所示
type Int2 interface {
	~int8 | ~int16 | ~int | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

//【注意】
// 1、泛型不能作为一个类型的基本类型
// 错误: cannot use a type parameter as RHS in type declaration
// type GenericType[T int | int32 | int64] T
// 错误: cannot use type Int outside a type constraint: interface contains type constraints
// var i Int = 10

// 2、泛型类型不能使用类型断言
func Demo1[T int | uint](v T) {
	// v1, ok := v.(int)
	// 错误: invalid operation: cannot use type assertion on type parameter value v (variable of type T constrained by int | uint)
}

// 3、匿名结构体不支持结构体
// 错误: expected declaration, found testStruct
// testStruct := struct[T int | string] {
//    Name string
//    Id T
// }[int]{
//    Name: "jack",
//    Id: 1
// }

// 4、匿名函数不支持自定义泛型
// var fn func[T int|Uint](v T)

// 5、不支持泛型方法（1.27版本之前）

// 6、类型集无法作为类型实参
type SignedInt interface {
	int8 | int16 | int | int32 | int64
}

func Do[T SignedInt](n T) T {
	return n
}

func Demo2() {
	// 错误: cannot use type SignedInt outside a type constraint: interface contains type constraints
	//Do[SignedInt](1) // 无法通过编译
	Do[int](10) // 具体类型才能通过编译
	Do(10)      // 这样也行，编译器会自行推断
}

// 7、类型集中的交集问题
// 对于非接口类型，类型并集中不能有交集

// 8、类型集无法直接或间接的并入自身

// 9、comparable 接口无法并入类型集

// 10、方法集无法并入类型集（但是可以有交集，不过这种写法没有意义）

func main() {
	// 创建Stack实例
	stack := NewStack[string](10)
	stack.Push("刻晴")
	stack.Push("甘雨")
	stack.Peek() // 甘雨
	stack.Size() // 2
	stack.Pop()  // 甘雨
	stack.Size() // 1
}
