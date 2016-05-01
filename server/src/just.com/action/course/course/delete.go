package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

func CourseDeleteHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	courseId := c.Param("course_id")
	courseService := service.NewCourseService(context.Session, context.Log)
	deleteErr := courseService.Delete(courseId)
	if deleteErr != nil {
		context.Log.Println(deleteErr)
		return
	}
}
