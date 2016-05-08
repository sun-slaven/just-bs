package service
import (
	"just.com/model/db/table"
	"time"
	"just.com/service/image"
	"just.com/dto"
	"just.com/err"
	"just.com/query/vo/course"
)

/*return courseId*/
func (self *CourseService) Update(request *dto.CourseAddRequest, userId string) (courseVo *course.CourseVo, error *err.HttpError) {
	courseTable, courseErr := self.GetCourse(request.Id)
	if courseErr != nil {
		return nil, courseErr
	}
	if request.IconUrl != "" {
		imageService := image.NewImageService(self.Session, self.Log)
		imageTable, imageTableErr := imageService.FindByUrl(request.IconUrl)
		if imageTableErr != nil {
			self.Log.Println(imageTableErr)
			error = err.NO_IMAGE_FOUND_BY_URL
			return
		}
		// icon
		courseTable.IconId = imageTable.UUID
		courseTable.IconWidth = imageTable.Width
		courseTable.IconHeight = imageTable.Height
		courseTable.IconUrl = imageTable.Url
	}

	if request.CollegeId != "" {
		getFlag, getErr := self.Session.Get(&table.CollegeTable{UUID:request.CollegeId})
		if !getFlag {
			if getErr != nil {
				self.Log.Println(getErr)
			}
			error = err.NO_COLLEGE_OR_MAJOR_FOUND
			return
		}
		courseTable.CollegeId = request.CollegeId
	}
	if request.MajorId != "" {
		getFlag, getErr := self.Session.Get(&table.MajorTable{UUID:request.MajorId})
		if !getFlag {
			if getErr != nil {
				self.Log.Println(getErr)
			}
			error = err.NO_COLLEGE_OR_MAJOR_FOUND
			return
		}
		courseTable.MajorId = request.MajorId
	}
	// TODO role name
	if request.TeacherId != "" {
		getFlag, getErr := self.Session.Get(&table.UserTable{UUID:request.TeacherId, RoleName:"TEACHER"})
		if !getFlag {
			if getErr != nil {
				self.Log.Println(getErr)
			}
			error = err.NO_TEACHER_ID_FOUND
			return
		}
		courseTable.TeacherId = request.TeacherId
	}
	if request.Name != "" {
		courseTable.Name = request.Name
	}
	if request.Description != "" {
		courseTable.Description = request.Description
	}
	if request.Introduction != "" {
		courseTable.Introduction = request.Introduction
	}
	if request.Experiment != "" {
		courseTable.Experiment = request.Experiment
	}
	if request.Syllabus != "" {
		courseTable.Syllabus = request.Syllabus
	}
	if request.Wish != "" {
		courseTable.Wish = request.Wish
	}
	if request.VideoUrl != "" {
		courseTable.VideoUrl = request.VideoUrl
	}
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()

	insertNum, insertErr := self.Session.Id(request.Id).Update(courseTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		error = err.COURSE_UPDATE_ERR
		return
	}
	_, attachErr := self.AddAttachmentList(courseTable.UUID, userId, request.AttachmentList)
	if attachErr != nil {
		self.Log.Println(attachErr)
		error = attachErr
		return
	}
	return course.LoadCourseVoFromTable(courseTable, userId, self.Session, self.Log)
}