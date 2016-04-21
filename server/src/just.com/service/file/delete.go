package service
import (
	"time"
	"just.com/model/db/table"
)

func (self *FileService) Delete(fileId string, userId string) error {
	fileTable := table.FileTable{}
	fileTable.UpdateTime = time.Now()
	fileTable.UpdateUser = userId
	fileTable.FrozenStatus = "Y"
	fileTable.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Id(fileId).Update(&fileTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return FILE_DELETE_ERR
	}
	return nil
}