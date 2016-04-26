package service
import (
	"just.com/dto"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/service/image"
)

/*return courseId*/
func (self *CourseService) Add(dto *dto.CourseDto, userId string) (string, error) {
	// 1. check icon
	is := image.ImageService{}
	is.Log = self.Log
	is.Session = self.Session

	courseTable := table.CourseTable{}
	courseTable.UUID = uuid.New()
	courseTable.Name = dto.Name
	courseTable.Introduction = dto.Introduction
	courseTable.Syllabus = dto.Syllabus
	courseTable.Experiment = dto.Experiment
	courseTable.Wish = dto.Wish
	courseTable.MajorId = dto.Major
	courseTable.MajorId = dto.College
	courseTable.CreateUser = userId
	courseTable.CreateTime = time.Now()
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	courseTable.MarkSum = 0
	courseTable.FrozenStatus = "N"
	courseTable.Points = 0
	courseTable.PointPerson = 0
	insertNum, insertErr := self.Session.InsertOne(&courseTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return "", COURSE_ADD_ERR
	}
	return courseTable.UUID, nil
}