package user

import (
	image_service "just.com/service/image"
	"just.com/common"
	"errors"
	"just.com/service"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
)



/*return new userId*/
func (self *UserService) Add(mobile, name, iconId, roleName string) (userId string, err error) {
	// 1.check
	err = errors.New(service.SERVICE_USER_ADD_ERR)
	if common.IsEmpty(mobile, name, iconId, roleName) == true {
		return
	}
	// check the role
	// 2.get the icon
	imageService := image_service.ImageService{}
	imageService.Session = self.Session
	imageService.Log = self.Log
	icon, iconErr := imageService.FindById(iconId)
	if iconErr != nil {
		self.Log.Println(iconErr)
		return
	}
	// 3.insert
	user := new(table.UserTable)
	user.UUID = uuid.New()
	user.RoleName = roleName
	user.Name = name
	user.Number = ""
	user.Age = 0
	user.Sex = 0
	user.Mobile = mobile
	user.Email = ""
	user.IconId = icon.UUID
	user.IconUrl = icon.Url
	user.IconWidth = icon.Width
	user.IconHeight = icon.Height
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	user.FrozenStatus = "N"
	insertNum, insertErr := self.Session.InsertOne(user)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return
	}
	userId = user.UUID
	err = nil
	return
}
