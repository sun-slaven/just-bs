package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/common"
	"just.com/err"
)

func MarkAdd(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	if common.IsEmpty(courseId) {
		context.Response.Error = err.NO_COURSE_ID_FOUND
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	markErr := courseService.Mark(courseId, context.UserId)
	if markErr != nil {
		context.Log.Println(markErr)
		return
	}
	go service.FlushMarkSum(courseId, context.Ds, context.Log)
}