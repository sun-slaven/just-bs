package file
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/common"
)

type FileVo struct  {
	UUID string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	CreateTime string `json:"create_time"`
}

func LoadFileVoList(courseId string,session *xorm.Session,log *log.Logger) []FileVo {
	fileVoList := make([]FileVo,0)
	fileTableList := make([]table.FileTable,0)
 	findErr:= session.Find(&fileTableList,&table.FileTable{CourseID:courseId, FrozenStatus:"N"})
	if findErr != nil{
		log.Println(findErr)
		return nil
	}
	for _,fileTable := range fileTableList{
		fileVo := FileVo{}
		fileVo.UUID = fileTable.UUID
		fileVo.Name = fileTable.Name
		fileVo.Url = fileTable.Url
		fileVo.CreateTime = common.TimeFormat(fileTable.CreateTime)
		fileVoList = append(fileVoList,fileVo)
	}
	return fileVoList
}
