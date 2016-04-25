package user
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/common"
	"just.com/service/token"
	"strings"
)

type UserLoginVo struct {
	UserVo *UserVo `json:"user"`
	Token  *service.UserToken `json:"token"`
}

// CheckUser 验证用户
func CheckUser(email string, password string, session *xorm.Session, log *log.Logger) (userLoginVo *UserLoginVo, flag bool) {
	userTable := new(table.UserTable)
	userTable.Email = email
	getFlag, getErr := session.Get(userTable)
	if getFlag == false {
		log.Println(getErr)
		return userLoginVo, false
	}
	if common.Md5(password) != strings.ToLower(userTable.Password) {
		return userLoginVo, false
	}
	// token
	token, err := service.NewTokenService(session, log).Make(userTable.UUID)
	if err != nil {
		log.Println(err)
		return userLoginVo, false
	}
	// uv
	userVo := LoadUserVoByTable(userTable)
	return &UserLoginVo{UserVo:userVo, Token:token}, true
}