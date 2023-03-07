package main

import (
	"log"

	"github.com/DW-inc/Ludo_Server/content"
	"github.com/DW-inc/Ludo_Server/network"
	"github.com/DW-inc/Ludo_Server/utils"
)

func main() {
	utils.InitLogger()
	log.Print("LUDO SERVER START")

	content.ContentManagerInit()

	server := &network.Server{}
	server.RunTCP(":8000")
}
