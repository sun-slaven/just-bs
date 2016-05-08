package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/dto"
	"just.com/common"
	"just.com/err"
)

func CourseUpdateHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.CourseAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
		return
	}
	request.Id = c.Param("course_id")
	if common.IsEmpty(request.Id) {
		context.Response.Error = err.NO_COURSE_ID_FOUND
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	courseVo, updateErr := courseService.Update(request, context.UserId)
	if updateErr != nil {
		context.Log.Println(updateErr)
		context.Response.Error = updateErr
		return
	}
	context.Response.Data = courseVo
}