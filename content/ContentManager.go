package content

import (
	"github.com/DW-inc/VI-Project/utils"
)

var (
	ph  *PacketHandler
	sm  *SessionManager
	dbm *DBManager
	lm  *LoginManager
)

func ContentManagerInit() {
	utils.Print("INIT_ContentManager")
	dbm = &DBManager{}
	sm = &SessionManager{}
	lm = &LoginManager{}
	ph = &PacketHandler{}

	dbm.Init()
	sm.Init()
	lm.Init()
	ph.Init()
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
