package db

type Players struct {
	Id       string `gorm:"primaryKey"`
	EMail    string
	Password string
	NickName string
	Token    string
	IP       string
}
