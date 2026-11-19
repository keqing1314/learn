package epic

import "log"

type EpicGamer struct {
	Nickname string
	LV       int
}

// EpicGamer
func (g *EpicGamer) Start() {
	log.Printf("epic玩家%s开始游戏", g.Nickname)
}

func (g *EpicGamer) Play() {
	log.Printf("epic玩家%s游玩游戏,当前游戏等级%d", g.Nickname, g.LV)
}

func (g *EpicGamer) Exit() {
	log.Println("epic退出游戏")
}
