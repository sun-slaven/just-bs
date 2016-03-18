package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"net/http"
	"just.com/service/course"
	"just.com/dto"
	"just.com/middleware"
)

func CourseHandle(c *gin.Context) {
	context, contextErr := action.GetContext(c)
	if contextErr != nil {
		return
	}
	courseId := c.Params("course_id")
	cs := service.NewCourseService(context.Session, context.Log)
	switch c.Request.Method {
	case action.METHOD_GET:
	case action.METHOD_PUT:
	case action.METHOD_DELETE:
		cs.Delete(courseId)
	}
	response := middleware.NewResponse(1, nil, nil)
	c.Set(middleware.RESPONSE, response)
}
