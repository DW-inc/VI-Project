package content

import (
	"net"
)

type Player struct {
	EMail    string
	Id       string
	password string
	Conn     net.Conn
}

func (p *Player) Store() {

}
