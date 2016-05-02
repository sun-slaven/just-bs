package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/service/image"
	"just.com/common"
	"just.com/dto"
	"just.com/err"
	"strings"
)

/*return courseId*/
func (self *CourseService) Save(request *dto.CourseAddRequest, userId string) (courseTable *table.CourseTable, error *err.HttpError) {
	courseTable = new(table.CourseTable)
	// 1. check
	if common.IsEmpty(request.Name, request.CollegeId, request.MajorId, request.TeacherId, request.IconUrl) {
		error = err.NO_REQUIERED_PARAM_FOUND
		return
	}
	// 2. icon college major teacher
	imageService := image.NewImageService(self.Session, self.Log)
	imageTable, imageTableErr := imageService.FindByUrl(request.IconUrl)
	if imageTableErr != nil {
		self.Log.Println(imageTableErr)
		error = err.NO_IMAGE_FOUND_BY_URL
		return
	}
	getFlag, getErr := self.Session.Get(&table.CollegeTable{UUID:request.CollegeId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.NO_COLLEGE_OR_MAJOR_FOUND
		return
	}
	getFlag, getErr = self.Session.Get(&table.MajorTable{UUID:request.MajorId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.NO_COLLEGE_OR_MAJOR_FOUND
		return
	}

	// TODO role name
	getFlag, getErr = self.Session.Get(&table.UserTable{UUID:request.TeacherId})
	if !getFlag {
		if getErr != nil {
			self.Log.Println(getErr)
		}
		error = err.NO_TEACHER_ID_FOUND
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

	// insert
	courseId := strings.TrimSpace(request.Id)
	if courseId == "" {
		insertNum, insertErr := self.Session.InsertOne(courseTable)
		if insertNum == 0 {
			if insertErr != nil {
				self.Log.Println(insertErr)
			}
			error = err.COURSE_INSERT_ERR
			return
		}
	}
	if courseId != "" {
		updateNum, updateErr := self.Session.Id(courseId).Update(courseTable)
		if updateNum == 0 {
			if updateErr != nil {
				self.Log.Println(updateErr)
			}
		}
		error = err.NO_COURSE_FOUND
	}
	error = nil
	return
}