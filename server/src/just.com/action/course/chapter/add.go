package chapter
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/course"
)

func ChapterAddHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.CourseChapterRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
	}
	courseId := c.Param("course_id")
	courseService := service.NewCourseService(context.Session, context.Log)
	chapterVo, addErr := courseService.AddChapter(courseId, context.UserId, request)
	if addErr != nil {
		context.Log.Println(addErr)
		context.Response.Error = addErr
	}
	context.Response.Data = chapterVo
}
