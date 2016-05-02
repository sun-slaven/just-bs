package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"just.com/model/db/table"
)

func CourseGetHandle(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	courseVo, courseVoErr := course.LoadCourseVo(&table.CourseTable{UUID:courseId}, context.Session, context.Log)
	if courseVoErr != nil {
		context.Log.Println(courseVoErr)
		return
	}
	context.Response.Data = courseVo
}
