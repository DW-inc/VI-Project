package content

import (
	"encoding/binary"
	"encoding/json"

	"github.com/DW-inc/Ludo_Server/utils"
)

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

func MakeSendBuffer[T any](pktid uint16, data T) []byte {
	sendData, err := json.Marshal(&data)
	if err != nil {
		utils.Print("MakeSendBuffer : Marshal Error", err)
	}
	sendBuffer := make([]byte, 4)

	pktsize := len(sendData) + 4

	binary.LittleEndian.PutUint16(sendBuffer, uint16(pktsize))
	binary.LittleEndian.PutUint16(sendBuffer[2:], pktid)

	sendBuffer = append(sendBuffer, sendData...)

	return sendBuffer
}

type S_EnterGame struct {
	Id string
}

type R_EnterGame struct {
	Money int
	Body  int
}

type S_MatchingStart struct {
	Id string
}

type R_GameStart struct {
	GameTime int
}

type R_OtherPlayer struct {
	Body       []int
	SpawnPoint []int
}
