package content

import (
	"log"
	"net"
	"sync"
)

type ContentManager struct {
	Players sync.Map
}

var CM_Ins *ContentManager
var CM_once sync.Once

func GetContentManager() *ContentManager {
	CM_once.Do(func() {
		CM_Ins = &ContentManager{}
	})
	return CM_Ins
}

func (cm *ContentManager) Init() {
	log.Println("INIT_ContentManager")
	// cm.Players = sync.Map{}
}

func (cm *ContentManager) SignIn(conn net.Conn, data string) {
	recvpkt := JsonStrToStruct[S_SignIn](data)

	cm.Players.Store(conn, recvpkt.Id)
}

func (cm *ContentManager) Message(conn net.Conn, data string) {
	if c, ok := cm.Players.Load(conn); ok {
		recvpkt := JsonStrToStruct[S_Message](data)

		packet := R_Message{}
		packet.Id = c.(string)
		packet.Message = recvpkt.Message
		sendBuffer := MakeSendBuffer(EMessage, recvpkt)
		GetGlobalSession().BroadCast(sendBuffer)
	}
}
