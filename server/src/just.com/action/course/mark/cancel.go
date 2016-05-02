package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

func MarkCancel(c *gin.Context) {
	context := action.GetContext(c)
	session := context.Session
	log := context.Log
	// request
	courseId := c.Param("course_id")
	//core
	courseService := service.NewCourseService(session, log)
	addPointErr := courseService.AddPoint(100, courseId, context.UserId)
	if addPointErr != nil {
		log.Println(addPointErr)
		return
	}
	go service.FlushPoint(courseId, context.Ds, log)
}