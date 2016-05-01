package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
)

func CommentDeleteHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if !contextFlag {
		return
	}
	courseId := c.Param("course_id")
	commentId := c.Param("comment_id")
	courseService := service.NewCourseService(context.Session, context.Log)
	err := courseService.DeleteComment(courseId, commentId)
	if err != nil {
		context.Log.Println(err)
		return
	}
	go service.FlushCommentSum(courseId, context.Ds, context.Log)
}
