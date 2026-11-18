// go 接口记录
package main

import "log"

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

// 实现接口
func (g *SteamGamer) Start() {
	log.Println("开始游戏")
}

func main() {}
