package content

import (
	"github.com/DW-inc/Ludo_Server/utils"
	//"github.com/lrita/cmap"
)

type LobbySession struct {
	//Session cmap.Cmap
}

func (ls *LobbySession) Init() {

	utils.Print("INIT_LobbySession")
	// ls.Session = cmap.Cmap{}
	// ls.Session.Store("", "")
}
