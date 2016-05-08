package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"log"
	"just.com/common"
	"just.com/err"
	"just.com/model/db/table"
	"strings"
)

/**
根据角色判断,学生则显示关注课程,教师则显示创建课程
 */
func CourseListHandle(c *gin.Context) {
	context := action.GetContext(c)
	userId := c.Param("user_id")
	if common.IsEmpty(userId) {
		context.Response.Error = err.NOT_USER_ID_FOUND
		return
	}
	isCreated := c.Query("is_created")
	// 显示创建的课程
	if strings.TrimSpace(isCreated) == "Y" {
		courseVoList, courseVoListErr := course.LoadCourseVoList(&table.CourseTable{TeacherId:userId}, userId, context.Session, context.Log)
		if courseVoListErr != nil {
			log.Println(courseVoListErr)
			context.Response.Error = courseVoListErr
			return
		}
		context.Response.Data = courseVoList
	}else {
		courseVoList, courseVoListErr := course.LoadMarkedCourseVo(userId, context.Session, context.Log)
		if courseVoListErr != nil {
			log.Println(courseVoListErr)
			context.Response.Error = courseVoListErr
			return
		}
		context.Response.Data = courseVoList
	}
}