package service
import (
	"just.com/dto"
	"just.com/common"
	"just.com/model/db/table"
	"time"
)


func (self *CourseService) Update(dto dto.CourseAddRequest, courseId string, userId string) error {
	isEmpty := common.IsEmpty(dto.Name, dto.Introduction, dto.Experiment, dto.Wish, dto.MajorId, dto.CollegeId)
	if isEmpty {
		return COURSE_UPDATE_ERR
	}
	courseTable := table.CourseTable{}
	courseTable.Name = dto.Name
	courseTable.Introduction = dto.Introduction
	courseTable.Experiment = dto.Experiment
	courseTable.Wish = dto.Wish
	courseTable.MajorId = dto.MajorId
	courseTable.CollegeId = dto.CollegeId
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	updateNum, updateErr := self.Session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		self.Log.Println(updateErr)
		return COURSE_UPDATE_ERR
	}
	return nil
}