package models

type CoreConfiguration struct {
	Id    int    `xorm:"not null pk autoincr INTEGER"`
	Name  string `xorm:"not null VARCHAR(20)"`
	Key   string `xorm:"not null unique VARCHAR(30)"`
	Value string `xorm:"not null TEXT"`
}
