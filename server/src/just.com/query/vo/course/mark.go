package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/query"
)

func LoadMarkedCourseVo(userId string, session *xorm.Session, log *log.Logger) (courseVoList []*CourseVo, err error) {
	courseVoList = make([]*CourseVo, 0)
	err = query.QUERY_COURSE_LOAD_MARK_ERR
	markTableList := make([]table.CourseMarkTable, 0)
	findErr := session.Find(&markTableList, &table.CourseMarkTable{UserId:userId})
	if findErr != nil {
		log.Println(findErr)
		return
	}

	for _, markTable := range markTableList {
		courseTable := &table.CourseTable{UUID:markTable.CourseId}
		getFlag, getErr := session.Get(courseTable)
		if !getFlag {
			if getErr != nil {
				log.Println(getErr)
			}
			return
		}
		courseVo, loadErr := LoadCourseVoFromTable(courseTable, session, log)
		if loadErr != nil {
			log.Println(loadErr)
			return
		}
		courseVoList = append(courseVoList, courseVo)
	}
	err = nil
	return
}
