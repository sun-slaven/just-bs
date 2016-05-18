package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/err"
	"just.com/value"
)

// 获取关注的课程列表
func LoadMarkedCourseVo(userId string, session *xorm.Session, log *log.Logger) (courseVoList []*CourseVo, error *err.HttpError) {
	courseVoList = make([]*CourseVo, 0)
	error = err.COURSE_MARKED_LIST_FIND_ERR
	markTableList := make([]table.CourseMarkTable, 0)
	findErr := session.Find(&markTableList, &table.CourseMarkTable{UserId:userId})
	if findErr != nil {
		log.Println(findErr)
		return
	}

	for _, markTable := range markTableList {
		courseTable := &table.CourseTable{UUID:markTable.CourseId, FrozenStatus:value.STATUS_ENABLED}
		courseVo, courseVoErr := LoadCourseVo(courseTable, userId, session, log)
		if courseVoErr != nil {
			log.Println(courseVo)
			return
		}
		// 之前没有删除关注数据,所有这边可能会有问题
		if courseVo != nil {
			courseVoList = append(courseVoList, courseVo)
		}
	}
	error = nil
	return
}

func LoadMarkStatus(courseId string, userId string, session *xorm.Session, log *log.Logger) bool {
	getFlag, getErr := session.Get(&table.CourseMarkTable{CourseId:courseId, UserId:userId})
	if getErr != nil {
		log.Println(getErr)
	}
	return getFlag
}
