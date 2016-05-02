package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"log"
)

func CourseListHandle(c *gin.Context) {
	context := action.GetContext(c)
	userId := c.Param("user_id")
	courseVoList, err := course.LoadMarkedCourseVo(userId, context.Session, context.Log)
	if err != nil {
		log.Println(err)
		return
	}
	context.Response.Data = courseVoList
}