// go 接口记录
package main

import (
	"fmt"
	"log"
)

// 接口在go中是很重要的一个概念了
// go对于接口的定义和其他传统的面向对象语言有着较大区别
// 直观的一点就是，go中接口的实现是隐式的
// 只要一个类型实现了一个接口定义的方法，那么这个类型就是实现了这个接口
// 像是Java C++ 这种都与一个显式的关键字来表明实现一个接口
// 另外就是go对于接口的定义
// 接口可以理解成定义了抽象事物的行为
// 例如定义个Player接口
type PlayGame interface {
	Start() // 打开游戏
	Play()  // 玩游戏
	Exit()  // 退出游戏
}

// 上面简单的抽象了一个游戏玩家的接口
// 定义了三个行为方法
// 在go推荐的是定义小接口

type SteamGamer struct {
	Username string
	LV       int
}

// SteamGamer实现接口就是将接口定义的方法全部实现
func (g *SteamGamer) Start() {
	log.Println("开始游戏")
}

func (g *SteamGamer) Play() {
	log.Println("游玩游戏")
}

func (g *SteamGamer) Exit() {
	log.Println("退出游戏")
}

//【定义】
// 上面给了一个接口和实现示例，现在来说说接口在go中定义
// 接口定义在1.18出现了变化
// 之前是接口是一个方法集，只要一个类型的方法集是接口方法的超集，那么这个类型就实现了该接口
// 但是1.18引入泛型，接口就成了类型集
// 只要一个类型是接口的子集，那么这个类型就是实现了该接口
// 看起来差异很大，其实在普通接口中这两个定义是一个东西
// 先抛开类型集的通用接口，只看普通接口
// 就像是上面示例中定义的接口，包含方法
// 一个接口有多种类型可以实现，那么实现了该接口的类型不就是该接口类型的子集吗
// 实现接口就是类型的方法集是接口的超集，也就是一个类型要实现接口定义的所有方法
// 不过是从类型角度和方法角度，两种视角的描述，本质是一件事

// 【创建】
// 创建一个接口
// type 接口名 interface
// go中接口是一种类型，默认零值是nil，同时接口也遵循go的暴露规范
type GenerteView interface {
	Render(r, g, b, a int)
}

// 空接口
// 在go中这是一种特殊的接口，空接口内部没有定义方法，所有类型都可以说实现了空接口
// 所以空接口常被作为代指任何类型
// any是空接口的别名
func Demo1() {
	var one any
	one = 1
	println(one)
	fmt.Println(one)
}

func main() {
	// 上面定义了SteamGamer结构体，其指针类型实现了接口PlayGame
	// 注意接口方法和普通方法不一样，什么类型实现了接口，就是用什么类型
	var player PlayGame = &SteamGamer{}
	player.Start() // 开始游戏
	player.Play()  // 游玩游戏
	player.Exit()  // 退出游戏
	// 上述使用接口来调用SteamGamer的实现
	// 但是SteamGamer自己本身也可以使用
	var steamGamer = &SteamGamer{}
	steamGamer.Start()
	steamGamer.Play()
	steamGamer.Exit()
	// 可以发现这里&SteamGamer可以使用PlayGame接口变量存储，也可使用SteamGamer变量

	// 多态
	// gamers := []PlayGame{&SteamGamer{}, &EpicGamer{}}
	// for _, g := range gamers {
	// 	g.Start()
	// 	g.Play()
	// 	g.Exit()
	// }
	// 由于这里没有创建项目，只是写的示例，找不到我的epic包所以注释了
	// 我在实际开发项目中，实现绘制svg图形就是封装图元类型，抽象绘制接口
	// 然后构造不同的图元数据，采用这种多态机制将svg绘制出来的
}

// 通用接口
// 这个其实和泛型相关，这里直接描述下使用
// 这就是定义的通用接口，用在泛型中使用
// 和传统的接口有很大的差异
type Number interface {
	Integer | Float
}

type Integer interface {
	uint | int
}

type Float interface {
	float32 | float64
}

// 泛型接口
// 当然有了泛型后就有了泛型接口
type DBMethod[T any] interface {
	Get() (T, error)
}

func (s SteamGamer) Get() (string, error) {
	return "刻晴", nil
}

func Demo2() {
	// 泛型接口使用方式
	var d DBMethod[string] = SteamGamer{}
	d.Get()
}

//【注意】尽管在社区中推荐使用小接口，但是不是说大接口不好，不能用
// 在go的源码中，你会发现这两种风格的接口都有大量使用的情况，合适才是最重要的
