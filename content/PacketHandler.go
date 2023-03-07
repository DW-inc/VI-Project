package content

import (
	"net"

	"github.com/DW-inc/Ludo_Server/utils"
)

type PacketHandler struct {
	HandlerFunc map[string]func(net.Conn, string)
}

func (ph *PacketHandler) Init() {
	utils.Print("INIT_PacketHandler")

	ph.HandlerFunc = make(map[string]func(net.Conn, string))
	ph.HandlerFunc["EnterGame"] = mm.EnterGame
	ph.HandlerFunc["MatchingStart"] = mm.MatchingStart
}
