package image
import (
	"code.google.com/p/go-uuid/uuid"
	"just.com/model/db/table"
	"time"
	"just.com/service"
)

func (self *ImageService) Add(name, url, userId string, width, height int64) error {
	imageTable := new(table.ImageTable)
	imageTable.UUID = uuid.New()
	imageTable.Name = name
	imageTable.Url = url
	imageTable.Width = width
	imageTable.Height = height
	imageTable.CreateTime = time.Now()
	imageTable.CreateUser = userId
	insertNum, insertErr := self.Session.InsertOne(imageTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return service.SERVICE_IMAGE_ADD_ERR
	}
	return nil
}