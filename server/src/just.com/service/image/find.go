package image
import (
	"just.com/model/db/table"
	"errors"
	"just.com/service"
	"strings"
)

func (self *ImageService) FindById(id string) (*table.ImageTable, error) {
	imageTable := new(table.ImageTable)
	getFlag, getErr := self.Session.Id(id).Get(imageTable)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return nil, errors.New(service.SERVICE_IMAGE_FIND_BY_ID_ERR)
	}
	return imageTable, nil
}

func (self *ImageService) FindByUrl(url string) (*table.ImageTable, error) {
	err := errors.New(service.SERVICE_IMAGE_FIND_BY_ID_ERR)
	if strings.TrimSpace(url) == "" {
		return nil, err
	}
	image := new(table.ImageTable)
	image.Url = url
	getFlag, getErr := self.Session.Get(image)
	if getFlag == false {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return nil, err
	}
	return image, nil
}

