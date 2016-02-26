package model
import "time"

type User struct {
	UUID         string    `xorm:"pk 'UUID'"`
	Role         string        `xorm:"'ROLE'"`
	Name         string        `xorm:"'NAME'"`
	Number       string    `xorm:"'NUMBER'"`
	Age          int64    `xrom:"'AGE'"`
	Sex          int64    `xorm:"'SEX'"`
	Mobile       string    `xorm:"'MOBILE'"`
	Email        string    `xorm:"'EMAIL'"`
	Icon         ImageTable
	CreateTime   time.Time    `xorm:"created 'CREATE_TIME'"`
	UpdateTime   time.Time    `xorm:"updated 'UPDATE_TIME'"`
	FrozenStatus string    `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time    `xorm:"deleted 'FROZEN_TIME'"`
}
