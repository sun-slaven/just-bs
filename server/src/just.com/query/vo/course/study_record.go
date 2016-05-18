package course
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"just.com/value"
)

// 返回学习相关课程的秒数
func LoadStudyRecord(courseId, userId string, session *xorm.Session, log *log.Logger) float64 {
	record := &table.CourseStudyRecordTable{CourseId:courseId, UserId:userId, FrozenStatus:value.STATUS_ENABLED}
	getFlag, getErr := session.Get(record)
	if !getFlag {
		if getErr != nil {
			log.Println(getErr)
		}
		return 0
	}
	return record.Process
}
