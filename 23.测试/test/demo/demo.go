package demo

import "fmt"

func Hello() {
	fmt.Println("hello 兹白")
}

func GameStart() {
	fmt.Println("原神启动!")
}

// 数值求和，接受可变参数，支持泛型，但是只限制数字
func Sum[T Number](nums ...T) T {
	var total T
	for _, num := range nums {
		total = num + total
	}
	return total
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

// 两数相等函数
func Equal[T Number](a, b T) bool {
	if a == b {
		return true
	}
	return false
}
