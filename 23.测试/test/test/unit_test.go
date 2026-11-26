// 单元测试
package test

import (
	"learn/test/demo"
	"testing"
)

// 对于单元测试而言，每一个测试用例的命名风格为TestXXX
// 且函数的入参必须是t *testing.T，这个是testing包中提供的测试结构体
// 提供了很多可用的方法，例子中的t.Errof等于t.Logf，用于格式化输出测试失败的日志信息
// 其他中常用的还有t.Fail用于标记当前是测试失败用例
// 功能类似的还有t.FailNow同样会标记失败
// 但是现在我自己做项目基本不自己写测试了，直接让ai测试

// t.Helper可以将当前函数标记为帮助函数，帮助函数不会单独运行
// 而是在测试用例的后面补充，有点像gin中的中间件的感觉
// 嵌入在用例中使用

func CleanupHelper(t *testing.T) {
	t.Helper()
	t.Log("test finished")
}

// 数值累加求和函数单元测试
func TestSum(t *testing.T) {
	t.Cleanup(func() {
		CleanupHelper(t)
	})

	a, b := 10, 45
	e := 55

	r := demo.Sum(a, b)
	if r != e {
		t.Errorf("Sum(%d,%d) 期望是%d,实际结果是%d", a, b, e, r)
	}
}

// 两数相等单元测试
func TestEqusl(t *testing.T) {
	t.Cleanup(func() {
		CleanupHelper(t)
	})

	a, b := 10, 23
	e := false

	r := demo.Equal(a, b)
	if r != e {
		t.Errorf("Equal(%d,%d) 预期是%v,实际结果是%v", a, b, e, r)
	}
}

// 子测试
// 一些情况下会在一个测试用例中引入另一个测试用例
// 这种嵌套的关系就被称为子测试。通过t.Run()方法
func TestDemo(t *testing.T) {
	t.Run("demo.Sum(10,101)", TestSum)
	t.Run("demo.Equal(11,111)", TestEqusl)
}

// 执行结果如下
// === RUN   TestDemo
// === RUN   TestDemo/demo.Sum(10,101)
//     unit_test.go:28: test finished
// === RUN   TestDemo/demo.Equal(11,111)
//     unit_test.go:43: test finished
// --- PASS: TestDemo (0.00s)
//     --- PASS: TestDemo/demo.Sum(10,101) (0.00s)
//     --- PASS: TestDemo/demo.Equal(11,111) (0.00s)
// PASS
// ok      learn/test/test 0.476s
