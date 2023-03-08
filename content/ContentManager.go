package content

import (
	"github.com/DW-inc/VI-Project/utils"
)

var (
	ph  *PacketHandler
	sm  *SessionManager
	dbm *DBManager
	lm  *LoginManager
	mm  *MainManager
)

func ContentManagerInit() {
	utils.Print("INIT_ContentManager")
	dbm = &DBManager{}
	sm = &SessionManager{}
	lm = &LoginManager{}
	ph = &PacketHandler{}
	ph = &PacketHandler{}
	mm = &MainManager{}

	dbm.Init()
	sm.Init()
	lm.Init()
	ph.Init()
	mm.Init()
}

func GetPacketHandler() *PacketHandler {
	return ph
}

func GetGlobalSession() *SessionManager {
	return sm
}

func GetDBManager() *DBManager {
	return dbm
}

func GetMainManager() *MainManager {
	return mm
}
