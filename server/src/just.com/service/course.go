package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model"
	"just.com/dto"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"strings"
	"errors"
)

type CourseService struct {
	Id      string
	Session *xorm.Session
	Log     *log.Logger
	Data    *model.CourseTable
}

func (self *CourseService) Add(dto dto.CourseDto, userId string) (id string, err error) {
	courseTable := model.CourseTable{}
	courseTable.UUID = uuid.New()
	courseTable.Name = dto.Name
	courseTable.Introduction = dto.Introduction
	courseTable.Syllabus = dto.Syllabus
	courseTable.Experiment = dto.Experiment
	courseTable.Plan = dto.Plan
	courseTable.Major = dto.Major
	courseTable.College = dto.College
	courseTable.CreateUser = userId
	courseTable.CreateTime = time.Now()
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	courseTable.MarkSum = 0
	courseTable.FrozenStatus = "N"
	insertNum, insertErr := self.Session.InsertOne(&courseTable)
	if insertNum == 0 {
		if insertErr != nil {
			err = insertErr
		}
		return
	}
	id = courseTable.UUID
	return
}

func (self *CourseService) Update(dto dto.CourseDto, courseId string, userId string) (err error) {
	isEmpty := IsEmpty(dto.Name, dto.Introduction, dto.Syllabus, dto.Experiment, dto.Plan, dto.Plan, dto.Major, dto.College)
	if isEmpty {
		return errors.New(SERVICE_COURSE_UPDATE_ERR)
	}
	courseTable := model.CourseTable{}
	courseTable.Name = dto.Name
	courseTable.Introduction = dto.Introduction
	courseTable.Syllabus = dto.Syllabus
	courseTable.Experiment = dto.Experiment
	courseTable.Plan = dto.Plan
	courseTable.Major = dto.Major
	courseTable.College = dto.College
	courseTable.UpdateUser = userId
	courseTable.UpdateTime = time.Now()
	updateNum, _ := self.Session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		err = errors.New(SERVICE_COURSE_UPDATE_ERR)
	}
	return nil
}

func (self *CourseService) Delete() (success bool) {
	if strings.TrimSpace(self.Id) == "" {
		return
	}
	courseTable := model.CourseTable{}
	courseTable.FrozenStatus = "Y"
	courseTable.FrozenTime = time.Now()
	updateNum, _ := self.Session.Id(self.Id).Update(&courseTable)
	if updateNum == 0 {
		success = false
		return
	}
	success = true
	return
}

/*mark the course*/
func (self *CourseService) Mark(courseId string, userId string) (id string, err error) {
	courseMarkTable := model.CourseMarkTable{}
	courseMarkTable.UUID = uuid.New()
	courseMarkTable.CourseId = courseId
	courseMarkTable.UserId = userId
	courseMarkTable.CreateTime = time.Now()
	courseMarkTable.FrozenStatus = "N"
	insertNum, _ := self.Session.InsertOne(&courseMarkTable)
	if insertNum == 0 {
		err = errors.New(SERVICE_COURSE_MARK_ERR)
		return
	}
	id = courseMarkTable.UUID
	go self.flushMarkSum(courseId)
	return
}

func (self *CourseService) MarkCancel(markId string, courseId string) (err error) {
	courseMarkTable := model.CourseMarkTable{}
	courseMarkTable.FrozenStatus = "Y"
	courseMarkTable.FrozenTime = time.Now()
	updateNum, _ := self.Session.Id(markId).Update(&courseMarkTable)
	if updateNum == 0 {
		err = errors.New(SERVICE_COURSE_MARK_CANCEL_ERR)
	}
	go self.flushMarkSum(courseId)
	return nil
}

func (self *CourseService) flushMarkSum(courseId string) {
	countSql := `SELECT COUNT("UUID") FROM "COURSE_MARK" WHERE "COURSE_ID" = ?
			AND "FROZEN_STATUS" = "N"`
	count, countErr := self.Session.Sql(countSql).Count(&model.CourseMarkTable{})
	if countErr != nil {
		self.Log.Println(countErr)
	}
	courseTable := model.CourseTable{}
	courseTable.MarkSum = count
	updateNum, _ := self.Session.Id(courseId).Update(&courseTable)
	if updateNum == 0 {
		self.Log.Println(SERVICE_COURSE_FLUSH_ERR)
	}
}

/*load according course id*/
func (self *CourseService) Load(courseId string) (err error) {
	getFlag, _ := self.Session.Id(courseId).Get(self.Data)
	if getFlag == false {
		err = errors.New(SERVICE_COURSE_LOAD_ERR)
		return
	}
	return nil
}
