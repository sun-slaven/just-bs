package user
import "just.com/model/db/table"

func (self *UserService)GetByEmail(email string) (userTable *table.UserTable, flag bool) {
	userTable = new(table.UserTable)
	userTable.Email = email
	getFlag, getErr := self.Session.Get(userTable)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return nil, false
	}
	flag = true
	return
}