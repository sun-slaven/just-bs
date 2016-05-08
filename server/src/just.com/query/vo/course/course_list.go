package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/err"
)

func LoadCourseVoList(condition *table.CourseTable, userId string, session *xorm.Session, log *log.Logger) (courseVoList []*CourseVo, error *err.HttpError) {
	condition.FrozenStatus = "N"
	courseVoList = make([]*CourseVo, 0)
	courseTableList := make([]*table.CourseTable, 0)
	findErr := session.Find(&courseTableList, condition)
	if findErr != nil {
		log.Println(findErr)
		error = err.COURSE_LIST_FIND_ERR
		return
	}
	for _, courseTable := range courseTableList {
		courseVo, courseVoErr := LoadCourseVoFromTable(courseTable, userId, session, log)
		if courseVoErr != nil {
			log.Println(courseVoErr)
			error = courseVoErr
			return
		}
		courseVoList = append(courseVoList, courseVo)
	}
	return courseVoList, nil
}
