package main

import (
	"log"

	"github.com/DW-inc/VI-Project/content"
	"github.com/DW-inc/VI-Project/network"
	"github.com/DW-inc/VI-Project/utils"
)

func main() {
	utils.InitLogger()
	log.Print("LUDO SERVER START")

	content.ContentManagerInit()

	server := &network.Server{}
	server.RunTCP(":8000")
}
