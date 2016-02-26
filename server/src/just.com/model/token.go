package model
import "time"

type TokenTable struct {
	UUID       string `xorm:"pk 'UUID'"`
	UserId     string `xorm:"'USER_ID'"`
	CreateTime time.Time    `xorm:"created 'CREATE_TIME'"`
	FrozenTime time.Time    `xorm:"'FROZEN_TIME'"`
}