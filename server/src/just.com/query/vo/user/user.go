package user
import (
	"just.com/query/vo/file"
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
)

type UserVo struct {
	UUID string `json:"id"`
	Name string `json:"name"`
	Icon *file.ImageVo `json:"icon"`
}

func LoadUserVo(userId string, session *xorm.Session, log *log.Logger) *UserVo {
	userTable := new(table.UserTable)
	getFlag, getErr := session.Id(userId).Get(userTable)
	if getFlag == false {
		if getErr != nil {
			log.Println(getErr)
		}
		return nil
	}
	userVo := new(UserVo)
	userVo.UUID = userTable.UUID
	userVo.Name = userTable.Name
	userVo.Icon = file.NewImageVo(userTable.IconUrl,userTable.IconWidth,userTable.IconHeight)
	return userVo
}