package content

import (
	"encoding/binary"
	"encoding/json"
	"log"
)

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

func MakeSendBuffer[T any](pktname string, data T) []byte {
	sendData, err := json.Marshal(&data)
	if err != nil {
		log.Println("Marshal Error !", err)
	}
	packetname := []byte(pktname)

	namesize := len(packetname)
	datasize := len(sendData)

	sendBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint16(sendBuffer, uint16(namesize))
	binary.LittleEndian.PutUint16(sendBuffer[2:], uint16(datasize))

	sendBuffer = append(sendBuffer, packetname...)
	sendBuffer = append(sendBuffer, sendData...)

	return sendBuffer
}

type S_SignUp struct {
	EMail    string
	Id       string
	Password string
}

type R_SignUp struct {
	IsSignUp bool
}

type S_CheckUserName struct {
	Id string
}

type R_CheckUserName struct {
	IsCheckUserName bool
}

type S_CheckEMail struct {
	EMail string
}

type R_CheckEMail struct {
	IsCheckEMail bool
}

type S_Login struct { // 그냥 Login
	Id       string
	Password string
}

type R_SetAutoLogin struct { // 그냥 Login 했을때 Id, Token 전달해서 클라가 저장
	Id    string
	Token string
}

type S_AutoLogin struct { // AutoLogin 으로 로그인 요청했을때
	Id    string
	Token string
}

type R_Login struct { // Autologin + Login 결과 값
	IsLogin     bool
	HasNickName bool
}
