package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/value"
	"just.com/err"
)

func (self *CourseService) AddStudyRecord(courseId, userId string, process float64) *err.HttpError {
	if process < 0 || process > 1 {
		return err.COURSE_SECONDS_ERR
	}
	course, courseErr := self.GetCourse(courseId)
	if courseErr != nil {
		return courseErr
	}
	recordTable := new(table.CourseStudyRecordTable)
	recordTable.CourseId = courseId
	recordTable.UserId = userId
	getFlag, getErr := self.Session.Get(recordTable)
	if getErr != nil {
		self.Log.Println(getErr)
		return err.COURSE_SECONDS_ERR
	}
	// 有这条记录则先更新
	if getFlag {
		updateNum, updateErr := self.Session.Id(recordTable.UUID).Update(&table.CourseStudyRecordTable{Process:process, UpdateTime:time.Now()})
		if updateNum == 0 {
			if updateErr != nil {
				self.Log.Println(updateErr)
			}
			return err.COURSE_SECONDS_ERR
		}
	} else {
		self.Session.InsertOne(&table.CourseStudyRecordTable{
			UUID:uuid.New(),
			CourseId:courseId,
			UserId:userId,
			VideoUrl:course.VideoUrl,
			Process:process,
			CreateTime:time.Now(),
			CreateUser:userId,
			UpdateTime:time.Now(),
			UpdateUser: userId,
			FrozenStatus:value.STATUS_ENABLED})
	}
	return nil
}

func (self *CourseService)DeleteStudyRecord(courseId, userId string) *err.HttpError {
	updateNum, updateErr := self.Session.Update(&table.CourseStudyRecordTable{
		FrozenStatus:value.STATUS_DISABLED,
		FrozenTime:time.Now(),
	},
		&table.CourseStudyRecordTable{
			CourseId:courseId,
			UserId:userId},
	)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
	}
	return nil
}
