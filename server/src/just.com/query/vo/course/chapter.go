package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/err"
	"just.com/model/db/table"
	"just.com/common"
)

type CourseChapterVo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Order      int64 `json:"order"`
	CreateTime string `json:"create_time"`
}

func LoadChapterVoList(courseId string, session *xorm.Session, log  *log.Logger) (courseVoList []*CourseChapterVo, error *err.HttpError) {
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
	chapterTableList := make([]*table.CourseChapterTable, 0)
	findErr := session.Asc("ORDER").Find(&chapterTableList, &table.CourseChapterTable{CourseId:courseId, FrozenStatus:"N"})
	if findErr != nil {
		log.Println(findErr)
		return nil, err.NO_CHAPTER_FOUND
	}
	return NewChapterVoList(chapterTableList), nil
}

func NewChapterVo(table *table.CourseChapterTable) *CourseChapterVo {
	chapterVo := new(CourseChapterVo)
	chapterVo.Id = table.UUID
	chapterVo.Order = table.Order
	chapterVo.Name = table.Name
	chapterVo.Content = table.Content
	chapterVo.CreateTime = common.TimeFormat(table.CreateTime)
	return chapterVo
}

func NewChapterVoList(tableList []*table.CourseChapterTable) []*CourseChapterVo {
	chapterVoList := make([]*CourseChapterVo, 0)
	for _, table := range tableList {
		chapterVoList = append(chapterVoList, NewChapterVo(table))
	}
	return chapterVoList
}
