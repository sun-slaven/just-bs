package point
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
)

func PointAdd(c *gin.Context) {
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
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	c.Set(middleware.RESPONSE, response)
	log.Println(response)
	go service.FlushPoint(courseId, context.Ds, log)
}