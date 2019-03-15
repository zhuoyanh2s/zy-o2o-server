package models

import (
	"time"
)

type MemberPet struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn     time.Time `xorm:"not null DATETIME"`
	ModifiedOn    time.Time `xorm:"not null DATETIME"`
	PetName       string    `xorm:"not null VARCHAR(50)"`
	PetSpecies    string    `xorm:"not null VARCHAR(50)"`
	PetAge        int64     `xorm:"not null BIGINT"`
	PetWeight     string    `xorm:"not null NUMERIC"`
	PetBirthday   time.Time `xorm:"not null DATE"`
	PetGender     string    `xorm:"default 'NULL::character varying' index VARCHAR(6)"`
	PetCharacter  string    `xorm:"default 'NULL::character varying' VARCHAR(50)"`
	PetDetailInfo string    `xorm:"TEXT"`
	UserId        int       `xorm:"not null index INTEGER"`
}
