package user
import (
	"just.com/model/db/table"
	"just.com/service"
	"time"
)

//Active 激活
func (self *UserService)Active(userId string) bool {
	userTable := new(table.UserTable)
	userTable.ActiveStatus = "Y"
	userTable.ActiveTime = time.Now()
	updateNum, updateErr := self.Session.Update(userTable, &table.UserTable{UUID:userTable.UUID})
	if updateNum == 0 {
		self.Log.Println(service.SERVICE_USER_ACTIVE_ERR)
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return false
	}
	return true
}
