package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

type CourseAddRequest struct {
	Name        string `json:"name"`
	Description string  `json:"Description"`
	Experiment  string `json:"experiment"`
	MajorId     string `json:"major_id"`
	CollegeId   string `json:"college_id"`
}

func CourseAddHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	request := new(CourseAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	courseId, addErr := courseService.Add(nil, context.UserId)
	if addErr != nil {
		return
	}
	// TODO
	context.Log.Println(courseId)
}