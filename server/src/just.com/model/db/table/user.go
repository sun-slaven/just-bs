package table
import "time"

type UserTable struct {
	UUID         string    `xorm:"pk 'UUID'"`
	RoleName         string        `xorm:"'ROLE_NAME'"`
	Name         string        `xorm:"'NAME'"`
	Number       string    `xorm:"'NUMBER'"`
	Age          int64    `xorm:"'AGE'"`
	Sex          int64    `xorm:"'SEX'"`
	Mobile       string    `xorm:"'MOBILE'"`
	Email        string    `xorm:"'EMAIL'"`
	IconId       string `xorm:"'ICON_ID'"`
	IconUrl      string `xorm:"'ICON_URL'"`
	IconWidth    int64 `xorm:"'ICON_WIDTH'"`
	IconHeight   int64 `xorm:"'ICON_HEIGHT'"`
	CreateTime   time.Time    `xorm:"created 'CREATE_TIME'"`
	UpdateTime   time.Time    `xorm:"updated 'UPDATE_TIME'"`
	FrozenStatus string    `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time    `xorm:"deleted 'FROZEN_TIME'"`
}
func (self *UserTable)TableName() string {
	return "USER"
}
