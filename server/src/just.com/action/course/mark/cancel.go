package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
)

func MarkCancel(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	token, tokenFlag := action.GetToken(c)
	if tokenFlag == false {
		return
	}
	session := context.Session
	log := context.Log
	// request
	courseId := c.Param("course_id")
	//core
	courseService := service.NewCourseService(session, log)
	addPointErr := courseService.AddPoint(100, courseId, token.UserId)
	if addPointErr != nil {
		log.Println(addPointErr)
		return
	}
	context.Response = middleware.NewResponse(http.StatusOK, nil, nil)
	go service.FlushPoint(courseId, context.Ds, log)
}