package mark
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
)

func MarkAdd(c *gin.Context) {
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
	markErr := courseService.Mark(courseId, token.UserId)
	if markErr != nil {
		log.Println(markErr)
		return
	}
	context.Response = middleware.NewResponse(http.StatusOK, nil, nil)
	go service.FlushMarkSum(courseId, context.Ds, log)
}