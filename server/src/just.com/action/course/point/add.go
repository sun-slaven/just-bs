package point
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
)

type PointRequest struct {
	point int64 `json:"point"`
}

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
	request := new(PointRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		log.Println(bindErr)
		return
	}
	//core
	courseService := service.NewCourseService(session, log)
	addPointErr := courseService.AddPoint(request.point, courseId, token.UserId)
	if addPointErr != nil {
		log.Println(addPointErr)
		return
	}
	context.Response = middleware.NewResponse(http.StatusOK, nil, nil)
	go service.FlushPoint(courseId, context.Ds, log)
}