package user
import (
	"just.com/model/db/table"
	"time"
	"just.com/err"
	"just.com/value"
)

func (self *UserService)Frozen(userId string) (error *err.HttpError) {
	userTable := &table.UserTable{FrozenStatus:value.STATUS_DISABLED, FrozenTime:time.Now(), UpdateTime:time.Now()}
	updateNum, updateErr := self.Session.Id(userId).Update(userTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
			error = err.USER_FROZEN_ERR
			return
		}
		error = err.USER_ID_ERR
		return
	}
	return
}
