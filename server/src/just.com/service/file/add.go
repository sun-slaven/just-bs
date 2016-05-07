package file
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/err"
)

func (self *FileService) Add(fileType, url string) (error *err.HttpError) {
	fileTable := new(table.FileTable)
	fileTable.UUID = uuid.New()
	fileTable.Url = url
	fileTable.Type = fileType
	fileTable.CreateTime = time.Now()
	inserNum, insertErr := self.Session.Insert(fileTable)
	if inserNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return err.FILE_ADD_ERR
	}
	return
}
