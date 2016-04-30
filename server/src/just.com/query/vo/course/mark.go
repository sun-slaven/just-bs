package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/query"
)

// 获取关注的课程列表
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
		courseVo, courseVoErr := LoadCourseVo(courseTable, session, log)
		if courseVoErr != nil {
			log.Println(courseVo)
			return
		}
		courseVoList = append(courseVoList, courseVo)
	}
	err = nil
	return
}
