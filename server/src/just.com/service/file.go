package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"errors"
)

type FileService struct {
	UUID    string
	Session *xorm.Session
	Log     *log.Logger
}

func (self *FileService) Add(courseId, userId, fileName, fileUrl string) (id string, err error) {
	fileTable := model.FileTable{}
	fileTable.UUID = uuid.New()
	fileTable.CourseID = courseId
	fileTable.Name = fileName
	fileTable.URL = fileUrl
	fileTable.CreateTime = time.Now()
	fileTable.CreateUser = userId
	fileTable.UpdateTime = time.Now()
	fileTable.UpdateUser = userId
	fileTable.FrozenStatus = "N"
	insertNum, insertErr := self.Session.InsertOne(&fileTable)
	if insertNum == 0 {
		self.Log.Println(insertErr)
		err = errors.New(SERVICE_FILE_ADD_ERR)
		return
	}
	return fileTable.UUID, nil
}

func (self *FileService) Update(userId, fileId, fileName, fileUrl string) (err error) {
	fileTable := model.FileTable{}
	fileTable.Name = fileName
	fileTable.URL = fileUrl
	fileTable.UpdateTime = time.Now()
	fileTable.UpdateUser = userId
	updateNum, updateErr := self.Session.Id(fileId).Update(&fileTable)
	if updateNum == 0 {
		self.Log.Println(updateErr)
		return errors.New(SERVICE_FILE_UPDATE_ERR)
	}
	return nil
}


func (self *FileService) Delete(fileId string) (err error) {
	fileTable := model.FileTable{}
	fileTable.FrozenStatus = "Y"
	fileTable.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Id(fileId).Update(&fileTable)
	if updateNum == 0 {
		self.Log.Println(updateErr)
		return errors.New(SERVICE_FILE_DELETE_ERR)
	}
	return nil
}