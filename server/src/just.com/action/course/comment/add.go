package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/err"
	"just.com/common"
)

type CommentAddRequest struct {
	Content string `json:"content"`
}

func CommentAdd(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	request := new(CommentAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		context.Response.Error = err.PARAM_ERR
		return
	}
	if common.IsEmpty(courseId, request.Content) {
		context.Response.Error = err.NO_REQUIRED_PARAM_FOUND
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	commentVo, commentVoErr := courseService.AddComment(request.Content, courseId, context.UserId)
	if commentVoErr != nil {
		context.Log.Println(commentVoErr)
		context.Response.Error = commentVoErr
		return
	}
	context.Response.Data = commentVo
	go service.FlushCommentSum(courseId, context.Ds, context.Log)
}