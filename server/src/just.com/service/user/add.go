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
func (self *UserService) Add(mobile, name, iconId, role string) (userId string, err error) {
	// 1.check
	err = errors.New(service.SERVICE_USER_ADD_ERR)
	if common.IsEmpty(mobile, name, iconId, role) == true {
		return
	}
	if role == "ADMIN" || role == "STUDENT" || role == "TEACHER" {
		return
	}
	// 2.get the icon
	imageService := image_service.ImageService{}
	icon, iconErr := imageService.GetById(iconId)
	if iconErr != nil {
		self.Log.Println(iconErr)
		return
	}
	// 3.insert
	user := new(table.UserTable)
	user.UUID = uuid.New()
	user.Role = role
	user.Name = name
	user.Number = ""
	user.Age = 0
	user.Sex = "F"
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
