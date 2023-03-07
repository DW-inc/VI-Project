package content

import "net"

type Player struct {
	Id    string
	Money int
	Body  int
	Conn  net.Conn
}
