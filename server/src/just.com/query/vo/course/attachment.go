package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/err"
	"just.com/model/db/table"
	"just.com/common"
	"just.com/value"
)

type CourseAttachmentVo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	CreateTime string `json:"create_time"`
}

func LoadAttachmentVoList(courseId string, session *xorm.Session, log  *log.Logger) (attachmentListVo []*CourseAttachmentVo, error *err.HttpError) {
	if common.IsEmpty(courseId) {
		return nil, err.NO_COURSE_ID_FOUND
	}
	getFlag, getErr := session.Id(courseId).Get(&table.CourseTable{})
	if !getFlag {
		if getErr != nil {
			log.Println(getErr)
		}
		return nil, err.NO_COURSE_FOUND
	}
	attachmentTableList := make([]*table.CourseAttachmentTable, 0)
	findErr := session.Find(&attachmentTableList, &table.CourseAttachmentTable{CourseId:courseId, FrozenStatus:"N"})
	if findErr != nil {
		log.Println(findErr)
		return nil, err.NO_ATTACHMENT_FOUND
	}
	return NewAttachmentVoList(attachmentTableList), nil
}

func NewAttachmentVo(table *table.CourseAttachmentTable) *CourseAttachmentVo {
	attachmentVo := new(CourseAttachmentVo)
	attachmentVo.Id = table.UUID
	attachmentVo.Name = table.Name
	url := table.Url
	if url == "" {
		url = value.DEFAULT_FILE
	}
	attachmentVo.Url = value.BASE_URL + url
	attachmentVo.CreateTime = common.TimeFormat(table.CreateTime)
	return attachmentVo
}

func NewAttachmentVoList(tableList []*table.CourseAttachmentTable) []*CourseAttachmentVo {
	attachmentVoList := make([]*CourseAttachmentVo, 0)
	for _, table := range tableList {
		attachmentVoList = append(attachmentVoList, NewAttachmentVo(table))
	}
	return attachmentVoList
}
