package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/dto"
)

func CourseAddHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.CourseAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	courseVo, addErr := courseService.Add(request, context.UserId)
	if addErr != nil {
		context.Log.Println(addErr)
		context.Response.Error = addErr
		return
	}
	context.Response.Data = courseVo
}