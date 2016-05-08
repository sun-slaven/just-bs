package service
import (
	"just.com/model/db/table"
	"time"
	"just.com/err"
	"just.com/value"
)

func (self *CourseService) Delete(courseId string) *err.HttpError {
	courseTable := &table.CourseTable{FrozenStatus:value.STATUS_DISABLED, FrozenTime:time.Now()}
	updateNum, updateErr := self.Session.Id(courseId).Update(courseTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return err.COURSE_DELETE_ERR
	}
	return nil
}