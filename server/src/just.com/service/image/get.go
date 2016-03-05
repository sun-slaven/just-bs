package image
import (
	"just.com/model/db/table"
	"errors"
	"just.com/service"
)

func (self *ImageService) GetById(id string) (*table.ImageTable, error) {
	imageTable := new(table.ImageTable)
	getFlag, getErr := self.Session.Id(id).Get(imageTable)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return nil, errors.New(service.SERVICE_IMAGE_GET_BY_ID_ERR)
	}
	return imageTable, nil
}

