package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/common"
	"just.com/err"
)

func CommentDeleteHandle(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	commentId := c.Param("comment_id")
	if common.IsEmpty(courseId, commentId) {
		context.Response.Error = err.NO_REQUIRED_PARAM_FOUND
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	deleteErr := courseService.DeleteComment(courseId, commentId)
	if deleteErr != nil {
		context.Log.Println(deleteErr)
		context.Response.Error = deleteErr
		return
	}
	go service.FlushCommentSum(courseId, context.Ds, context.Log)
}
