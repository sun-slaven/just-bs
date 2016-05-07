package user
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/err"
	"just.com/model/db/table"
	"just.com/common"
)

type UserDetailVo struct {
	*UserVo
	FrozenStatus string `json:"frozen_status"`
	CreateTime   string `json:"create_time"`
}

/**
获取用户除了ADMIN以外的列表
 */
func LoadUserVoList(session *xorm.Session, log *log.Logger) (userVoList []*UserDetailVo, error *err.HttpError) {
	tableList := make([]table.UserTable, 0)
	userSql := `SELECT * FROM "USER" WHERE "ROLE_NAME" != 'ADMIN'`
	findErr := session.Sql(userSql).Find(&tableList)
	if findErr != nil {
		log.Println(findErr)
		return nil, err.USER_LIST_ERR
	}
	userVoList = make([]*UserDetailVo, 0)
	for _, userTable := range tableList {
		userVo := LoadUserVoByTable(&userTable)
		log.Println(userTable.CreateTime)
		userDetailVo := &UserDetailVo{UserVo:userVo, FrozenStatus:userTable.FrozenStatus, CreateTime: common.TimeFormat(userTable.CreateTime)}
		userVoList = append(userVoList, userDetailVo)
	}
	return
}
