package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
)

func Comment(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	session := context.Session
	log := context.Log
	courseService := service.NewCourseService(session, log)
	// request
	courseId := c.Param("course_id")
	commentId := c.Param("comment_id")
	//	var data interface{}
	switch c.Request.Method {
	case action.METHOD_GET:
	case action.METHOD_DELETE:
		err := courseService.DeleteComment(courseId, commentId)
		if err != nil {
			log.Println(err)
		}
	}
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	c.Set(middleware.RESPONSE, response)
}