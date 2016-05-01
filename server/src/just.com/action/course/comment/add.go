package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/err"
)

type CommentAddRequest struct {
	Content string `json:"content"`
}

func CommentAdd(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	// request
	courseId := c.Param("course_id")
	request := new(CommentAddRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	commentVo, commentVoErr := courseService.AddComment(request.Content, courseId, context.UserId)
	if commentVoErr != nil {
		context.Log.Println(commentVoErr)
		context.Err = err.COURSE_NOT_FOUND
		return
	}
	context.Response.Data = commentVo
	go service.FlushCommentSum(courseId, context.Ds, context.Log)
}