package point
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

type PointRequest struct {
	point int64 `json:"point"`
}

func PointAdd(c *gin.Context) {
	context := action.GetContext(c)
	// request
	courseId := c.Param("course_id")
	request := new(PointRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	//core
	courseService := service.NewCourseService(context.Session, context.Log)
	addPointErr := courseService.AddPoint(request.point, courseId, context.UserId)
	if addPointErr != nil {
		context.Log.Println(addPointErr)
		return
	}
	go service.FlushPoint(courseId, context.Ds, context.Log)
}