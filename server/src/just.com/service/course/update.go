package service
import (
	"just.com/dto"
	"just.com/common"
	"just.com/model/db/table"
	"time"
)


func (self *CourseService) Update(dto dto.CourseDto, courseId string, userId string) error {
	isEmpty := common.IsEmpty(dto.Name, dto.Introduction, dto.Syllabus, dto.Experiment, dto.Wish, dto.Major, dto.College)
	if isEmpty {
		return COURSE_UPDATE_ERR
	}
	courseTable := table.CourseTable{}
	courseTable.Name = dto.Name
	courseTable.Introduction = dto.Introduction
	courseTable.Syllabus = dto.Syllabus
	courseTable.Experiment = dto.Experiment
	courseTable.Wish = dto.Wish
	courseTable.MajorId = dto.Major
	courseTable.CollegeId = dto.College
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	updateNum, updateErr := self.Session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		self.Log.Println(updateErr)
		return COURSE_UPDATE_ERR
	}
	return nil
}