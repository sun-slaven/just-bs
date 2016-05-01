package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/query/vo/course"
	"just.com/dto"
)


func CourseAddHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	request := new(dto.CourseAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	courseTable, addErr := courseService.Add(request, context.UserId)
	if addErr != nil {
		context.Log.Println(addErr)
		return
	}
	courseVo, couseVoErr := course.LoadCourseVoFromTable(courseTable, context.Session, context.Log)
	if couseVoErr != nil {
		context.Log.Println(couseVoErr)
		return
	}
	context.Response.Data = courseVo
}