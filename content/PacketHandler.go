package content

import (
	"net"

	"github.com/DW-inc/VI-Project/utils"
)

type PacketHandler struct {
	HandlerFunc map[string]func(net.Conn, string)
}

func (ph *PacketHandler) Init() {
	utils.Print("INIT_PacketHandler")

	ph.HandlerFunc = make(map[string]func(net.Conn, string))
	ph.HandlerFunc["CheckUserName"] = lm.CheckUserName
	ph.HandlerFunc["CheckEMail"] = lm.CheckEMail
	ph.HandlerFunc["SignUp"] = lm.SignUp
	ph.HandlerFunc["Login"] = lm.Login
	ph.HandlerFunc["AutoLogin"] = lm.AutoLogin

	ph.Test()
}

func (ph *PacketHandler) Test() {
	// p := db.Players{Id: "123", Password: "123", EMail: "3"}
	// if err := dbm.db.Where("id = ? && e_mail = ?", "123", "3").First(&p).Error; err != nil {
	// 	log.Println("FAIL:", p)
	// } else {
	// 	log.Println("SUCCESS:", p)
	// }
}
