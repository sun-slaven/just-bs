package table
import "time"

type TokenTable struct {
	UUID       string `xorm:"pk 'UUID'"`
	UserId     string `xorm:"'USER_ID'"`
	CreateTime time.Time    `xorm:"created 'CREATE_TIME'"`
	DeadTime   time.Time    `xorm:"'DEAD_TIME'"`
}

func (self *TokenTable )TableName() string {
	return "TOKEN"
}