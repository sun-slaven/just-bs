package service
import (
	"just.com/model/db/table"
	"time"
)

func (self *FileService) Update(userId, fileId, fileName, fileUrl string) (err error) {
	fileTable := table.FileTable{}
	fileTable.Name = fileName
	fileTable.Url = fileUrl
	fileTable.UpdateTime = time.Now()
	fileTable.UpdateUser = userId
	updateNum, updateErr := self.Session.Id(fileId).Update(&fileTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return FILE_UPDATE_ERR
	}
	return nil
}