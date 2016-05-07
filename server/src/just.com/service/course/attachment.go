package service
import (
	"just.com/query/vo/course"
	"just.com/err"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"just.com/dto"
	"just.com/common"
	"time"
)

func (self *CourseService) AddAttachment(courseId, userId string, request *dto.CourseAttachmentRequest) (attachmentVo *course.CourseAttachmentVo, error *err.HttpError) {
	if common.IsEmpty(request.Name, request.Url) {
		error = err.NO_REQUIRED_PARAM_FOUND
		return
	}
	courseTable, getErr := self.GetCourse(courseId)
	if getErr != nil {
		error = getErr
		return
	}
	table := &table.CourseAttachmentTable{
		UUID:uuid.New(),
		Name:request.Name,
		Url:request.Url,
		CourseId:courseTable.UUID,
		CreateUser:userId,
		CreateTime:time.Now(),
		FrozenStatus:"N",
	}
	insertNum, insertErr := self.Session.Insert(table)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		error = err.CHAPTER_ADD_ERR
		return
	}
	return course.NewAttachmentVo(table), nil
}

func (self *CourseService)AddAttachmentList(courseId, userId string, requestList []*dto.CourseAttachmentRequest) ([]*course.CourseAttachmentVo, *err.HttpError) {
	attachListVo := make([]*course.CourseAttachmentVo, 0)
	for _, request := range requestList {
		attachVo, addErr := self.AddAttachment(courseId, userId, request)
		if addErr != nil {
			self.Log.Println(addErr)
			return nil, addErr
		}
		attachListVo = append(attachListVo, attachVo)
	}
	return attachListVo, nil
}