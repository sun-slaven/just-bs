package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/course"
)

func CommentList(c *gin.Context) {
	context := action.GetContext(c)
	session := context.Session
	log := context.Log
	// request
	courseId := c.Param("course_id")
	//core
	commentListVo := course.LoadCommentVoList(courseId, session, log)
	context.Response.Data = commentListVo
}