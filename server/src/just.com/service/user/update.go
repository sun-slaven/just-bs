package user
import (
	"just.com/model/db/table"
	"just.com/service/image"
	"just.com/query/vo/user"
	"just.com/service"
)

func (self *UserService)Update(userId, name, email, iconUrl string) (userVo *user.UserVo, err error) {
	err = service.SERVICE_USER_UPDATE_ERR
	userVo = new(user.UserVo)
	userTable := new(table.UserTable)
	getFlag, getErr := self.Session.Id(userId).Get(userTable)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return
	}
	if name != "" {
		userTable.Name = name
	}
	if email != "" {
		// TODO 邮箱验证
		userTable.Email = name
	}
	if iconUrl != "" {
		imageService := image.NewImageService(self.Session, self.Log)
		imageTable, imageErr := imageService.FindByUrl(iconUrl)
		if imageErr != nil {
			self.Log.Panic(imageService)
			return
		}
		userTable.IconId = imageTable.UUID
		userTable.IconUrl = imageTable.Url
		userTable.IconWidth = imageTable.Width
		userTable.IconHeight = imageTable.Height
	}
	updateNum, updateErr := self.Session.Id(userId).Update(userTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Panic(updateErr)
		}
		return
	}
	userVo = user.LoadUserVoByTable(userTable)
	err = nil
	return
}
