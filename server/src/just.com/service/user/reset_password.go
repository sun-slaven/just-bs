package user
import (
	"just.com/err"
	"just.com/common"
)


func (self *UserService)RestPassword(userId string, password string) (error *err.HttpError) {
	userTable, useErr := self.GetById(userId)
	if useErr != nil {
		self.Log.Println(useErr)
		return useErr
	}
	userTable.Password = common.Md5(password)
	updateNum, updateErr := self.Session.Id(userId).Update(userTable)
	if updateNum == 0 {
		if updateErr != nil {
			return err.USER_RESET_PASSWORD_ERR
		}
		return err.USER_ID_ERR
	}
	return nil
}
