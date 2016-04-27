package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/query"
)

func LoadCourseVoList(condition *table.CourseTable, session *xorm.Session, log *log.Logger) (courseVoList []*CourseVo, err error) {
	err = query.QUERY_COURSE_LOAD_LIST_ERR
	courseVoList = make([]*CourseVo, 0)
	courseTableList := make([]*table.CourseTable, 0)
	findErr := session.Find(&courseTableList, condition)
	if findErr != nil {
		log.Println(findErr)
		return nil, err
	}
	for _, courseTable := range courseTableList {
		courseVo, err := LoadCourseVoFromTable(courseTable, session, log)
		if err != nil {
			log.Println(err)
		}
		log.Println(courseVo.Teacher)
		courseVoList = append(courseVoList, courseVo)
	}
	return courseVoList, nil
}
