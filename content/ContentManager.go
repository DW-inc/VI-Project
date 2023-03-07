package content

import (
	"github.com/DW-inc/Ludo_Server/utils"
)

var (
	ph  *PacketHandler
	sm  *SessionManager
	dbm *DBManager
	mm  *MatchManager
)

func ContentManagerInit() {
	utils.Print("INIT_ContentManager")
	ph = &PacketHandler{}
	sm = &SessionManager{}
	dbm = &DBManager{}
	mm = &MatchManager{}

	ph.Init()
	sm.Init()
	dbm.Init()
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
