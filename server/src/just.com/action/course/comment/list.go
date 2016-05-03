package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
	"just.com/common"
	"just.com/err"
)

func CommentList(c *gin.Context) {
	context := action.GetContext(c)
	// request
	courseId := c.Param("course_id")
	if common.IsEmpty(courseId) {
		context.Response.Error = err.NO_COURSE_ID_FOUND
		return
	}
	commentListVo := course.LoadCommentVoList(courseId, context.Session, context.Log)
	context.Response.Data = commentListVo
}