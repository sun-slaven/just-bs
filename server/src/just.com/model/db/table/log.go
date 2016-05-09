package table
import "time"

type LogHttp struct {
	UUID         string        `xorm:"'UUID'"`
	Method       string        `xorm:"'METHOD'"`
	Url          string        `xorm:"'URL'"`
	Body         string        `xorm:"'BODY'"`
	UserId       string        `xorm:"USER_ID"`
	Status       int `xorm:"'STATUS'"`
	ErrorMessage string `xorm:"'ERROR_MESSAGE'"`
	ResponseBody string        `xorm:"'RESPONSE_BODY'"`
	CreateTime   time.Time        `xorm:"'CREATE_TIME'"`
}

func (self *LogHttp) TableName() string {
	return "LOG_HTTP"
}