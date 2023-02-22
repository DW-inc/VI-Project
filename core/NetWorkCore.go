package core

import (
	"encoding/binary"
	"log"
	"net"
	"sync"

	"github.com/lkjfrf/content"
)

type NetworkCore struct {
}

var instance *NetworkCore
var once sync.Once

func GetNetworkCore() *NetworkCore {
	once.Do(func() {
		instance = &NetworkCore{}
	})
	return instance
}

func (nc *NetworkCore) Init() {
	log.Println("INIT_NetworkCore")
	nc.Connect()
}

func (nc *NetworkCore) Connect() {
	if ln, err := net.Listen("tcp", ":8001"); err != nil {
		log.Println(err)
	} else {
		if conn, err := ln.Accept(); err != nil {
			log.Println(err)
		} else {
			nc.RecvPacket(conn)
		}
	}
}

func (nc *NetworkCore) RecvPacket(conn net.Conn) {
	go func() {
		for {
			header := make([]byte, 4)
			conn.Read(header)
			id, size := nc.ParseHeader(header)

			if size > 0 && 0 < id && id > 100 {
				data := make([]byte, size-4)
				n, _ := conn.Read(data)
				if n > 0 && content.GetPacketHandler().HandlerFunc[id] != nil {
					content.GetPacketHandler().HandlerFunc[id](conn, string(data))
					log.Println(conn, "Send - ", string(data))
				}
			} else {
				log.Println("Header Err")
			}
		}
	}()
}

func (nc *NetworkCore) ParseHeader(json []byte) (int, int) {
	pktSize := binary.LittleEndian.Uint16(json[:2])
	pktId := binary.LittleEndian.Uint16(json[2:4])

	return int(pktId), int(pktSize)
}
