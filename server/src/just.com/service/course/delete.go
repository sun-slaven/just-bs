package service
import (
	"strings"
	"just.com/model/db/table"
	"time"
)

func (self *CourseService) Delete(courseId string) error {
	if strings.TrimSpace(courseId) == "" {
		return COURSE_DELETE_ERR
	}
	courseTable := table.CourseTable{}
	courseTable.FrozenStatus = "Y"
	courseTable.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return COURSE_DELETE_ERR
	}
	return nil
}