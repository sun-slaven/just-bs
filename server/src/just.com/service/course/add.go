package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/service/image"
	"just.com/common"
	"just.com/service"
	"just.com/dto"
)

/*return courseId*/
func (self *CourseService) Add(request *dto.CourseAddRequest, userId string) (courseTable *table.CourseTable, err error) {
	err = service.SERVICE_COURSE_ADD_ERR
	courseTable = new(table.CourseTable)
	// 1. check
	if common.IsEmpty(request.Name, request.CollegeId, request.MajorId, request.TeacherId, request.ImageUrl) {
		return
	}
	// 2. icon college major teacher
	imageService := image.NewImageService(self.Session, self.Log)
	imageTable, imageTableErr := imageService.FindByUrl(request.ImageUrl)
	if imageTableErr != nil {
		self.Log.Println(imageTableErr)
		return
	}
	getFlag, getErr := self.Session.Get(&table.CollegeTable{UUID:request.CollegeId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return
	}
	getFlag, getErr = self.Session.Get(&table.MajorTable{UUID:request.MajorId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return
	}

	// TODO role name
	getFlag, getErr = self.Session.Get(&table.UserTable{UUID:request.TeacherId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		return
	}

	courseTable.UUID = uuid.New()
	courseTable.Name = request.Name
	courseTable.Introduction = request.Introduction
	courseTable.Syllabus = ""
	courseTable.Experiment = request.Experiment
	courseTable.Wish = request.Wish
	courseTable.MajorId = request.MajorId
	courseTable.CollegeId = request.CollegeId
	courseTable.CreateUser = userId
	courseTable.CreateTime = time.Now()
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	courseTable.MarkSum = 0
	courseTable.FrozenStatus = "N"
	courseTable.Points = 0
	courseTable.PointPerson = 0

	// icon
	courseTable.IconId = imageTable.UUID
	courseTable.IconWidth = imageTable.Width
	courseTable.IconHeight = imageTable.Height
	courseTable.IconUrl = imageTable.Url
	courseTable.TeacherId = request.TeacherId

	insertNum, insertErr := self.Session.InsertOne(courseTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return
	}
	err = nil
	return
}