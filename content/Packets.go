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

func MakeSendBuffer[T any](pktid uint16, data T) []byte {
	sendData, err := json.Marshal(&data)
	if err != nil {
		log.Println("MakeSendBuffer : Marshal Error", err)
	}
	sendBuffer := make([]byte, 4)

	pktsize := len(sendData) + 4

	binary.LittleEndian.PutUint16(sendBuffer, uint16(pktsize))
	binary.LittleEndian.PutUint16(sendBuffer[2:], pktid)

	sendBuffer = append(sendBuffer, sendData...)

	return sendBuffer
}

type S_SignIn struct {
	Id string
}

type S_Message struct {
	Message string
}

type R_Message struct {
	Id      string
	Message string
}
