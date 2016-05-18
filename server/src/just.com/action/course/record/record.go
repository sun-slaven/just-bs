package record

import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

type CourseRecordRequest struct {
	Process float64 `json:"process"`
}

func AddStudyRecord(c *gin.Context) {
	context := action.GetContext(c)
	request := new(CourseRecordRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	addError := courseService.AddStudyRecord(c.Param("course_id"), context.UserId, request.Process)
	if addError != nil {
		context.Response.Error = addError
	}
}
