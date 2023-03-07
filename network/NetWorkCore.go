package network

import (
	"errors"
	"io"
	"log"
	"net"
	"syscall"

	"github.com/DW-inc/VI-Project/content"
	"github.com/DW-inc/VI-Project/utils"
)

type Server struct {
}

func (s *Server) RunTCP(addr string) {

	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		utils.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		utils.Fatal(err)
	}

	utils.Print("Run TCP:" + tcpaddr.String())

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			utils.Fatal(err)
			return
		}

		utils.Print("New Conn:" + conn.RemoteAddr().String())

		go s.EventLoop(conn)
	}
}

func (s *Server) EventLoop(conn net.Conn) {

	header := make([]byte, 4)

	for {
		n, err := conn.Read(header)
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, syscall.Errno(10054)) {
				// content.GetGlobalSession().Disconnect(conn)
				log.Println("Client DIscoonect", conn.RemoteAddr().String())
				conn.Close()
			}
			return
		}

		if n == 4 {
			namesize, datasize := utils.ParseHeader(header)
			recvdata := make([]byte, namesize+datasize)
			recvn, err := conn.Read(recvdata)
			if err != nil {
				if errors.Is(err, io.EOF) || errors.Is(err, syscall.Errno(10054)) {
					utils.Print("Connection is closed from client : " + conn.RemoteAddr().String())
				}
				return
			}

			if recvn == namesize+datasize {
				pktname, jsondata := utils.ExtractData(namesize, datasize, recvdata)

				if _, ok := content.GetPacketHandler().HandlerFunc[pktname]; ok {
					utils.PrintRecvPkt(pktname, jsondata)
					content.GetPacketHandler().HandlerFunc[pktname](conn, jsondata)
				} else {
					utils.Print("Unknown Packet : " + pktname)
				}

			} else {
				utils.Print("Recv Wrong Datasize !")
			}
		} else {
			utils.Print("Recv Wrong Header !")
		}
	}
}

func (s *Server) Test(conn net.Conn) {
}
