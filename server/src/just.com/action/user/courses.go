package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"log"
)

func CourseListHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if !contextFlag {
		return
	}
	courseVoList, err := course.LoadMarkedCourseVo(context.UserId, context.Session, context.Log)
	if err != nil {
		log.Println(err)
		return
	}
	context.Response.Data = courseVoList
}