package content

import (
	"net"
)

type Player struct {
	Id       string
	password string
	EMail    string
	Conn     net.Conn
}

func (p *Player) Store(data Player) {

}
