package content

import (
	"github.com/DW-inc/VI-Project/libs/cmap"
	"github.com/DW-inc/VI-Project/utils"
)

type MainManager struct {
	Players cmap.ConcurrentMap[*Player]
	Conns   cmap.ConcurrentMap[string]
}

func (mm *MainManager) Init() {
	utils.Print("INIT_MainManager")
	mm.Players = cmap.New[*Player]()
	mm.Conns = cmap.New[string]()
}

func (mm *MainManager) NewPlayer(data *Player) {
	mm.Players.Set(data.Id, data)
	mm.Conns.Set(data.Conn, data.Id)
}
