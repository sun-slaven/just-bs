package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

func MarkAdd(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	// request
	courseId := c.Param("course_id")
	courseService := service.NewCourseService(context.Session, context.Log)
	markErr := courseService.Mark(courseId, context.UserId)
	if markErr != nil {
		context.Log.Println(markErr)
		return
	}
	go service.FlushMarkSum(courseId, context.Ds, context.Log)
}