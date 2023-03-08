package db

type Players struct {
	Id       string `gorm:"primaryKey"`
	Password string
	NickName string `gorm:"unique"`
	EMail    string
	Token    string
	IP       string
}
