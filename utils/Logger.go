package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var GLogger *Logger = nil

type Logger struct {
	LogQ chan string
}

// func GetLogger() *Logger {
// 	if GLogger == nil {
// 		GLogger = &Logger{}
// 		GLogger.Init()
// 	}

// 	return GLogger
// }

func InitLogger() {
	GLogger = &Logger{}
	GLogger.LogQ = make(chan string, 10000)
	go StartLogger()
	Print("INIT_Logger")
}

func (l *Logger) MakeLog(msg string) {
	GLogger.LogQ <- msg
}

func StartLogger() {
	for {
		log.Print(<-GLogger.LogQ)
	}
}

// func Print(msg string) {

// 	pc, _, _, ok := runtime.Caller(1)
// 	funcname := runtime.FuncForPC(pc).Name()
// 	if !ok || funcname == "" {
// 		GLogger.MakeLog("[LOG ERROR] LogPrintWithError")
// 	}

//		//GLogger.MakeLog("[LOG] " + "[" + funcname + "] " + msg)
//		GLogger.MakeLog("[LOG] " + msg)
//	}

func Print(v ...any) {

	//pc, _, _, ok := runtime.Caller(1)
	//funcname := runtime.FuncForPC(pc).Name()
	//if !ok || funcname == "" {
	//	GLogger.MakeLog("[LOG ERROR] LogPrintWithError")
	//}
	//GLogger.MakeLog("[LOG] " + "[" + funcname + "] " + msg)

	GLogger.MakeLog("[LOG] " + fmt.Sprint(v...))
}

func Fatal(v ...any) {

	//pc, _, _, ok := runtime.Caller(1)
	//funcname := runtime.FuncForPC(pc).Name()
	//if !ok || funcname == "" {
	//	GLogger.MakeLog("[LOG ERROR] LogPrintWithError")
	//}
	//
	//GLogger.MakeLog("[LOG] " + "[" + funcname + "] " + msg)
	GLogger.MakeLog("[LOG] " + fmt.Sprint(v...))
	os.Exit(1)
}

func PrintErr(msg string, err error) {

	pc, _, _, ok := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()
	if !ok || funcname == "" {
		GLogger.MakeLog("[LOG ERROR] LogPrintWithError")
	}

	GLogger.MakeLog("[LOG] " + "[" + funcname + "] " + msg + " " + err.Error())
}

func PrintSendPkt(n int, data []byte) {
	namesize, datasize := ParseHeader(data[:4])
	pktname, jsondata := ExtractData(namesize, datasize, data[4:n])

	if pktname != "OtherPlayerMove" && pktname != "UDPHeartBeat" {
		GLogger.MakeLog("[SEND] " + "[" + pktname + "]\n" + jsondata + "\n")
	}
}

func PrintRecvPkt(pktname string, jsonstr string) {

	if pktname != "PlayerMove" && pktname != "UDPHeartBeat" {
		GLogger.MakeLog("[RECV] " + "[" + pktname + "]\n" + jsonstr + "\n")
	}
}
