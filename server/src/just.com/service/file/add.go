package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
)

func (self *FileService) Add(courseId, fileName, fileUrl, userId  string) (id string, err error) {
	// 1.check the course id
	count, countErr := self.Session.Id(courseId).Count(&table.CourseTable{})
	if count == 0 {
		if countErr != nil {
			self.Log.Println(countErr)
		}
		return "", FILE_ADD_ERR
	}
	// 2.insert
	fileTable := table.FileTable{}
	fileTable.UUID = uuid.New()
	fileTable.CourseID = courseId
	fileTable.Name = fileName
	fileTable.Url = fileUrl
	fileTable.CreateTime = time.Now()
	fileTable.CreateUser = userId
	fileTable.UpdateTime = time.Now()
	fileTable.UpdateUser = userId
	fileTable.FrozenStatus = "N"
	insertNum, insertErr := self.Session.InsertOne(&fileTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return "", FILE_ADD_ERR
	}
	return fileTable.UUID, nil
}
