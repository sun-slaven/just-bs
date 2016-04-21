package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/query/vo/course"
)

func CommentList(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	session := context.Session
	log := context.Log
	// request
	courseId := c.Param("course_id")
	//core
	commentListVo := course.LoadCommentVoList(courseId, session, log)
	response := middleware.NewResponse(http.StatusOK, commentListVo, nil)
	c.Set(middleware.RESPONSE, response)
}