package service
import (
	"just.com/model/db/table"
	"just.com/err"
	"just.com/common"
)

func (self *CourseService)GetCourse(courseId string) (courseTable *table.CourseTable, error *err.HttpError) {
	if common.IsEmpty(courseId) {
		error = err.NO_COURSE_ID_FOUND
		return
	}
	courseTable = new(table.CourseTable)
	getFlag, getErr := self.Session.Id(courseId).Get(courseTable)
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.NO_COURSE_FOUND
		return
	}
	return
}
