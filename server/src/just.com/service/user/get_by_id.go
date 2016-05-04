package user
import (
	"just.com/model/db/table"
	"just.com/err"
	"just.com/common"
)

func (self *UserService) GetById(userId string) (userTable *table.UserTable, error *err.HttpError) {
	userTable = new(table.UserTable)
	if common.IsEmpty(userId) {
		return nil, err.NOT_USER_ID_FOUND
	}
	getFlag, getErr := self.Session.Id(userId).Get(userTable)
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
			return nil, err.USER_ID_ERR
		}
		return nil, err.USER_ID_ERR
	}
	return userTable, nil
}
