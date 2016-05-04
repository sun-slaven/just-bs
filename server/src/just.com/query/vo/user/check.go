package user
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/common"
	"just.com/service/token"
	"strings"
	"just.com/err"
)

type UserLoginVo struct {
	UserVo *UserVo `json:"user"`
	Token  *service.UserToken `json:"token"`
}

// CheckUser 验证用户
func CheckUser(email string, password string, session *xorm.Session, log *log.Logger) (userLoginVo *UserLoginVo, error *err.HttpError) {
	error = err.USER_PASSWORD_OR_EMAIL_ERR
	userTable := new(table.UserTable)
	if common.IsEmpty(email) {
		return
	}
	userTable.Email = email
	userTable.ActiveStatus = "Y"
	userTable.FrozenStatus = "N"
	getFlag, getErr := session.Get(userTable)
	if !getFlag {
		if getErr != nil {
			log.Println(getErr)
		}
		return
	}
	if common.Md5(password) != strings.ToLower(userTable.Password) {
		return
	}
	// token
	token, tokenErr := service.NewTokenService(session, log).Make(userTable.UUID)
	if tokenErr != nil {
		log.Println(tokenErr)
		error = err.TOKEN_CREATE_ERR
		return
	}
	// uv
	userVo := LoadUserVoByTable(userTable)
	return &UserLoginVo{UserVo:userVo, Token:token}, nil
}