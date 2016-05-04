package user
import (
	"just.com/model/db/table"
	"just.com/service/image"
	"just.com/query/vo/user"
	"just.com/err"
)

func (self *UserService)Update(userId, name, email, iconUrl string) (userVo *user.UserVo, error *err.HttpError) {
	error = err.NO_IMAGE_FOUND_BY_URL
	userVo = new(user.UserVo)
	userTable := new(table.UserTable)
	getFlag, getErr := self.Session.Id(userId).Get(userTable)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.USER_ID_ERR
		return
	}
	if name != "" {
		userTable.Name = name
	}
	if email != "" {
		// TODO 邮箱验证
		userTable.Email = email
	}
	if iconUrl != "" {
		imageService := image.NewImageService(self.Session, self.Log)
		imageTable, imageErr := imageService.FindByUrl(iconUrl)
		if imageErr != nil {
			self.Log.Println(imageErr)
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
		error = err.USER_ID_ERR
		return
	}
	userVo = user.LoadUserVoByTable(userTable)
	error = nil
	return
}
