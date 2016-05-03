package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"log"
	"just.com/common"
	"just.com/err"
)

func CourseListHandle(c *gin.Context) {
	context := action.GetContext(c)
	userId := c.Param("user_id")
	if common.IsEmpty(userId) {
		context.Response.Error = err.NOT_USER_ID_FOUND
	}
	courseVoList, courseVoListErr := course.LoadMarkedCourseVo(userId, context.Session, context.Log)
	if courseVoListErr != nil {
		log.Println(courseVoListErr)
		context.Response.Error = courseVoListErr
		return
	}
	context.Response.Data = courseVoList
}