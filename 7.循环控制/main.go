// go中循环控制记录
package main

import (
	"fmt"
	"log"
)

// 有过学习过其他编程语言的经历
// 对于循环控制肯定不陌生，for循环 while循环 do while循环之类的
// 但是在go中循环语句只存在for循环
// 其他语言中的while循环通过for关键字就行了

func main() {

	// for
	// 语句格式如下：
	// 	for init statement; expression; post statement {
	//   execute statement
	// }
	// 这是经典的三段式
	for i := 0; i < 10; i++ {
		log.Print(i)
	}
	// i := 0 就是初始化块，用来初始化变量
	// i < 10 就是表达式，用来作为判断循环继续的条件
	// i++ 就是后置块，一轮循环结束后的操作
	// 这单三个条件不是一定要有的，例如缺失后置块
	for i := 0; i < 10; {
		i++
		log.Print(i)
	}
	// 这就缺失了后置块，但是将后置块的逻辑放到了循环里面
	// 同样初始化块也可以不使用
	var i int
	for i < 10 {
		i++
		log.Println(i)
	}
	// 这就将初始化块和后置块都省略了，其实就等价于其他语言中的while i < 10
	// 同时条件表达式也可以不要
	for {
		i++
		if i >= 10 {
			break
		}
		log.Println(i)
	}
	// 这就是其他语言中的while true 语句，也就是死循环
	// 不过这里我还是在循环中添加了退出循环的条件分支

	// 这里的三个块的表达式可以很复杂
	// 例如在初始化块中初始多个变量，复杂的表达式，多个后置操作
	for i, j := 1, 2; i < 100 && j < 1000; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// for range
	// 这是一个很好的迭代方式
	// 比如上面我的第一个例子，使用for range 简化
	for i := range 10 {
		log.Println(i)
	}
	// 其中的格式如下：
	// for index, value := range iterable {
	// 	// body
	// }
	// index和vaule 不一定都有，例如上面的示例
	// 但是一般都是这样的，例如遍历一个切片
	s := []int{1, 2, 5, 3, 4}
	for i, v := range s {
		// 索引是从0开始的
		log.Printf("索引:%d,值:%d", i, v)
	}

	// 如果只是迭代次数不遍历值的话
	// 可以这样做，这样就是将循环中的操作重复10次
	for range 10 {
		log.Println("这是一些操作")
	}

	// 终止循环
	// break
	// 主流编程语言中使用break来终止循环，go也不列外
	for i := range 10 {
		for j := range 10 {
			if i == j {
				break
			}
			fmt.Println(i, j)
		}
	}
	// 这是一个简单的循环终止示例
	// 上述的break只会终止内层循环，也就是当条件满足的时候终止内层j循环，外层不会终止
	// 如果像同时终止，就需要接口标签使用
Outer:
	for i := range 10 {
		for j := range 10 {
			if i == j {
				break Outer
			}
			fmt.Println(i, j)
		}
	}
	// 这样满足条件就可以将外部循环一起终止

	// 跳过迭代
	// continue
	// go中跳过迭代也是使用的continue
	for i := range 10 {
		for j := range 10 {
			if i == j {
				continue
			}
			fmt.Println(i, j)
		}
	}
	// 结合标签
Out:
	for i := range 10 {
		for j := range 10 {
			if i == j {
				continue Out
			}
			fmt.Println(i, j)
		}
	}

	// 注意一点for range 在1.22前后有变动
	// 以前版本for range 里面有个坑，就是value使用的是同一个临时变量
	// 如果使用指针就会导致所有的值都是最后遍历的值，解决办法就是重复赋值
	// 所以你经常会看到，这样的写法
	for i, v := range []int{1, 2, 3, 5} {
		num := v
		log.Println(i, &num)
	}
	// 结果输出
	// 2026/11/17 11:13:43 0 0x22f1855560b0
	// 2026/11/17 11:13:43 1 0x22f185556608
	// 2026/11/17 11:13:43 2 0x22f185556630
	// 2026/11/17 11:13:43 3 0x22f185556638
	// 但是在1.22之后这个坑被官方给添上了
	// 所以你的版本如果是后续版本，你就可以放心的使用
	for i, v := range []int{1, 2, 3, 5} {
		log.Println(i, &v)
	}
	// 结果输出
	// 2026/11/17 11:14:04 0 0x1c423945c0b0
	// 2026/11/17 11:14:04 1 0x1c423945c608
	// 2026/11/17 11:14:04 2 0x1c423945c630
	// 2026/11/17 11:14:04 3 0x1c423945c638

	// 上述你可以看到两种结果都是指针每个都是不同的
	// 但是如果在之前的版本后者的写法，最后输出的指针地址一定是相同的
	// 你在网上找教程或者网课，绝对会有人写这个注意，不过基本都是1.22前的教程和课程
	// 很少有人说这个变化
}
