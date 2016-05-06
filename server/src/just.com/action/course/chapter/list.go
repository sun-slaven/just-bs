package chapter
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
)

func ChapterListHandle(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	chapterVoList, err := course.LoadChapterVoList(courseId, context.Session, context.Log)
	if err != nil {
		context.Log.Println(err)
		context.Response.Error = err
	}
	context.Response.Data = chapterVoList
}
