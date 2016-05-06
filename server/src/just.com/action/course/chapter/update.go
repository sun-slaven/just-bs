package chapter
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/course"
)

func ChapterUpdateHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.CourseChapterRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
	}
	chapterId := c.Param("chapter_id")
	courseService := service.NewCourseService(context.Session, context.Log)
	chapterVo, updateErr := courseService.UpdateChapter(chapterId, context.UserId, request)
	if updateErr != nil {
		context.Log.Println(updateErr)
		context.Response.Error = updateErr
	}
	context.Response.Data = chapterVo
}
