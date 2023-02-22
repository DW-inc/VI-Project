package content

import (
	"log"
	"net"
	"sync"
)

type PacketHandler struct {
	HandlerFunc map[int]func(net.Conn, string)
}

var PH_Ins *PacketHandler
var PH_once sync.Once

func GetPacketHandler() *PacketHandler {
	PH_once.Do(func() {
		PH_Ins = &PacketHandler{}
	})
	return PH_Ins
}

const (
	Err = iota

	ESignIn = 1
	EMessage
)

func (ph *PacketHandler) Init() {
	log.Println("INIT_PacketHandler")

	ph.HandlerFunc = make(map[int]func(net.Conn, string))
	ph.HandlerFunc[ESignIn] = GetContentManager().SignIn
	ph.HandlerFunc[EMessage] = GetContentManager().Message
}
