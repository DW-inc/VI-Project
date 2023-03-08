package content

import (
	"net"

	"github.com/DW-inc/VI-Project/utils"
)

type SessionManager struct {
}

func (sm *SessionManager) Init() {
	utils.Print("INIT_GlobalSession")

}

func (sm *SessionManager) SendByte(c net.Conn, data []byte) {
	if c != nil {
		sent, err := c.Write(data)
		if err != nil {
			utils.Print("SendPacket ERROR :", err)
		} else {
			if sent != len(data) {
				utils.Print("[Sent diffrent size] : SENT =", sent, "BufferSize =", len(data))
			}
			utils.Print("SendPacket : ", data)
		}
	}
}

func (sm *SessionManager) SendPacketById(Id string, data []byte) {
	if Id != "" {
		if p, ok := mm.Players.Get(Id); ok {
			sent, err := p.Conn.Write(data)
			if err != nil {
				utils.Print("SendPacket ERROR :", err)
			} else {
				if sent != len(data) {
					utils.Print("[Sent diffrent size] : SENT =", sent, "BufferSize =", len(data))
				}
				utils.Print("SendPacket : ", data)
			}
		} else {
			utils.Print("SendPacket Faile : ID not found in players")
		}
	}
}

func (sm *SessionManager) BroadCast(buff []byte) {

}
