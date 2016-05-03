package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/query/vo/course"
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
	courseTable, addErr := courseService.Add(request, context.UserId)
	if addErr != nil {
		context.Log.Println(addErr)
		context.Response.Error = addErr
		return
	}
	courseVo, courseVoErr := course.LoadCourseVoFromTable(courseTable, context.Session, context.Log)
	if courseVoErr != nil {
		context.Log.Println(courseVoErr)
		context.Response.Error = courseVoErr
		return
	}
	context.Response.Data = courseVo
}