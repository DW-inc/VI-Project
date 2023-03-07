package db

import "time"

type BattleScore struct {
	Phonenum string `korm:"varchar(50)" json:"phonenum"`
	Nickname string `korm:"varchar(50)"`
	Ranking  string `korm:"varchar(10)" json:"rank"`
	Type     int32  `korm:"integer"`
	Teamid   int32  `korm:"integer"`

	KillCount  int32 `korm:"integer" json:"kill"`
	Death      int32 `korm:"integer"`
	GiveDamage int32 `korm:"integer"`
	TakeDamage int32 `korm:"integer"`

	KeyActiveTime int32 `korm:"integer"`
	Judgement     int32 `korm:"integer"`
	IsWinner      bool  `korm:"boolean"`
	Score         int32 `korm:"integer"`
}

type Inventory struct {
	Playerid string    `korm:"varchar(50)"`
	Itemid   string    `korm:"varchar(50)"`
	Itemdate time.Time `korm:"datetime(3)"`
}
