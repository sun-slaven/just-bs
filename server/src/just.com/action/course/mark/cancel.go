package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/common"
	"just.com/err"
)

func MarkCancel(c *gin.Context) {
	context := action.GetContext(c)
	// request
	courseId := c.Param("course_id")
	if common.IsEmpty(courseId) {
		context.Response.Error = err.NO_COURSE_ID_FOUND
		return
	}
	//core
	courseService := service.NewCourseService(context.Session, context.Log)
	markCancelErr := courseService.MarkCancel(courseId, context.UserId)
	if markCancelErr != nil {
		context.Log.Println(markCancelErr)
		context.Response.Error = markCancelErr
		return
	}
	go service.FlushMarkSum(courseId, context.Ds, context.Log)
}