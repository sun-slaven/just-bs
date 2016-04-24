package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/dto"
	"just.com/middleware"
	"net/http"
)

type CourseAddRequest struct {
	Name       string
	Desc       string
	Syllabus   string
	Plan       string
	Experiment string
	Major      string
	College    string
}

func CourseAddHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	token, tokenFlag := action.GetToken(c)
	if tokenFlag == false {
		return
	}
	request := new(CourseAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		return
	}
	session := context.Session
	log := context.Log
	courseService := service.NewCourseService(session, log)
	courseDto := dto.NewCouseDto(request.Name, request.Desc, request.Syllabus, request.Plan, request.Experiment, request.Major, request.College)
	courseId, addErr := courseService.Add(courseDto, token.UserId)
	if addErr != nil {
		return
	}
	context.Response = middleware.NewResponse(http.StatusOK, courseId, nil)
}