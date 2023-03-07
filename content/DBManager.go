package content

import (
	"log"

	"github.com/DW-inc/VI-Project/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBManager struct {
	db *gorm.DB
}

func (dbm *DBManager) Init() {
	if db, err := gorm.Open(mysql.Open("newuser:Tomatosoup1!@tcp(192.168.0.99:3306)/vi?charset=utf8mb4&parseTime=True&loc=Local")); err != nil {
		log.Fatal("GORM OPEN ERROR")
	} else {
		dbm.db = db
	}
	dbm.db.AutoMigrate(&db.Players{})
}

func (dbm *DBManager) SignUp() {

}

func (dbm *DBManager) Login() {

}

func (dbm *DBManager) BringInventory() {

}
