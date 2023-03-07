package content

import (
	"math/rand"
	"net"
	"strings"

	"github.com/DW-inc/VI-Project/db"
	"github.com/DW-inc/VI-Project/utils"
)

type LoginManager struct {
	Charset string
}

func (lm *LoginManager) Init() {
	utils.Print("INIT_LobbySession")
	lm.Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
}

func GetUnicRandomKey(length int) string {
	s := make([]byte, length)
	for i := range s {
		s[i] = lm.Charset[rand.Intn(len(lm.Charset))]
	}
	return string(s)
}

func (lm *LoginManager) SignUp(c net.Conn, json string) {
	recvpkt := utils.JsonStrToStruct[S_SignUp](json)
	packet := R_SignUp{IsSignUp: true}

	p := db.Players{EMail: recvpkt.EMail, Password: recvpkt.Password, Id: recvpkt.Id,
		IP: strings.Split(c.RemoteAddr().String(), ":")[0]}
	if err := dbm.db.Create(&p).Error; err != nil {
		packet.IsSignUp = false
	}
	c.Write(MakeSendBuffer("SignUp", packet))
}

func (ph *LoginManager) CheckUserName(c net.Conn, json string) {
	recvpkt := utils.JsonStrToStruct[S_CheckUserName](json)
	packet := R_CheckUserName{IsCheckUserName: true}

	p := db.Players{}
	if err := dbm.db.Where("id = ?", recvpkt.Id).First(&p).Error; err == nil {
		packet.IsCheckUserName = false
	}
	c.Write(MakeSendBuffer("CheckUserName", packet))
}

func (ph *LoginManager) CheckEMail(c net.Conn, json string) {
	recvpkt := utils.JsonStrToStruct[S_CheckEMail](json)
	packet := R_CheckEMail{IsCheckEMail: true}

	p := db.Players{}
	if err := dbm.db.Where("e_mail = ?", recvpkt.EMail).First(&p).Error; err == nil {
		packet.IsCheckEMail = false
	}
	c.Write(MakeSendBuffer("CheckEMail", packet))
}

func (ph *LoginManager) Login(c net.Conn, json string) {
	recvpkt := utils.JsonStrToStruct[S_Login](json)
	packet := R_Login{IsLogin: true, HasNickName: true}

	p := db.Players{}
	if err := dbm.db.Where("id = ? and password = ?", recvpkt.Id, recvpkt.Password).First(&p).Error; err != nil {
		packet.IsLogin = false
	} else {
		if p.NickName == "" {
			packet.HasNickName = false
		}
	}
	c.Write(MakeSendBuffer("Login", packet))

	if packet.IsLogin {
		token := GetUnicRandomKey(10)
		if err := dbm.db.Model(p).Update("token", token).Error; err == nil {
			packet2 := R_SetAutoLogin{Id: recvpkt.Id, Token: token}
			c.Write(MakeSendBuffer("SetAutoLogin", packet2))
		}
	}
}

func (ph *LoginManager) AutoLogin(c net.Conn, json string) {
	recvpkt := utils.JsonStrToStruct[S_AutoLogin](json)
	packet := R_Login{IsLogin: true, HasNickName: true}

	p := db.Players{}
	if err := dbm.db.Where("id = ? and token = ? and ip = ?", recvpkt.Id, recvpkt.Token,
		strings.Split(c.RemoteAddr().String(), ":")[0]).First(&p).Error; err != nil {
		packet.IsLogin = false
	} else {
		if p.NickName == "" {
			packet.HasNickName = false
		}
	}
	c.Write(MakeSendBuffer("Login", packet))
}
