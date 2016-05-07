package chapter
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

func ChapterDeleteHandle(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	chapterId := c.Param("chapter_id")
	courserService := service.NewCourseService(context.Session, context.Log)
	deleteErr := courserService.DeleteChapter(courseId, chapterId, context.UserId)
	if deleteErr != nil {
		context.Response.Error = deleteErr
	}
}
