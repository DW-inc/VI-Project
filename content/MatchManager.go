package content

import (
	"math/rand"
	"net"

	"github.com/DW-inc/Ludo_Server/libs/cmap"
	"github.com/DW-inc/Ludo_Server/utils"
)

type MatchManager struct {
	Match cmap.ConcurrentMap[*Player]
	Count int
}

func (mm *MatchManager) Init() {

	utils.Print("INIT_LobbySession")
	// mm.Session = cmap.Cmap{}
	// ls.Session.Store("", "")
	mm.Match = cmap.New[*Player]()
	mm.Count = 0
}

func (mm *MatchManager) EnterGame(conn net.Conn, json string) {
	packet := R_EnterGame{Money: 999, Body: 1}
	buffer := utils.MakeSendBuffer("EnterGame", packet)
	conn.Write(buffer)
}

func (mm *MatchManager) MatchingStart(conn net.Conn, json string) {
	//recvpkt := utils.JsonStrToStruct[S_MatchingStart](json)
	Id := string(rand.Int())
	mm.Count++
	mm.Match.Set(mm.Count, &Player{Id: Id, Money: 999, Conn: conn})
	if mm.Count == 2 {
		utils.Print("START")
		packet := R_GameStart{GameTime: 100}
		buffer := utils.MakeSendBuffer("GameStart", packet)
		for _, v := range mm.Match.Items() {
			v.Conn.Write(buffer)
		}
	}
}
