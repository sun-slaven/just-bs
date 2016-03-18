package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
)

/*mark the course
 return markId*/
func (self *CourseService) Mark(courseId string, userId string) error {

	courseMarkTable := new(table.CourseMarkTable)
	courseMarkTable.CourseId = courseId
	courseMarkTable.UserId = userId
	// 删除这个课程的关注
	_, deleteErr := self.Session.Delete(courseMarkTable)
	if deleteErr != nil {
		self.Log.Println(deleteErr)
		return COURSE_MARK_ERR
	}
	courseMarkTable.UUID = uuid.New()
	courseMarkTable.CreateTime = time.Now()
	insertNum, updateErr := self.Session.InsertOne(courseMarkTable)
	if insertNum == 0 {
		self.Log.Println("insert num == 0")
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return COURSE_MARK_ERR
	}
	return nil
}
/*cancel mark*/
func (self *CourseService) MarkCancel(courseId string, userId string) error {
	// 1.get the mark
	cm := new(table.CourseMarkTable)
	cm.CourseId = courseId
	cm.UserId = userId
	deleteNum, deleteErr := self.Session.Delete(cm)
	if deleteNum == 0 {
		if deleteErr != nil {
			self.Log.Println(deleteErr)
		}
		return COURSE_MARK_CANCEL_ERR
	}
	return nil
}