package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/service/course"
	"github.com/gin-gonic/gin/binding"
)

type CommentAddRequest struct {
	Content string `form:"content"`
}

func CommentAdd(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	token, tokenFlag := action.GetToken(c)
	if tokenFlag == false {
		return
	}
	session := context.Session
	log := context.Log
	// request
	courseId := c.Param("course_id")
	request := new(CommentAddRequest)
	bindErr := c.BindWith(request, binding.Form)
	if bindErr != nil {
		log.Println(bindErr)
	}
	//core
	courseService := service.NewCourseService(session, log)
	commentId, commentErr := courseService.AddComment(request.Content, courseId, token.UserId)
	if commentErr != nil {
		log.Println(commentErr)
		return
	}
	response := middleware.NewResponse(http.StatusOK, commentId, nil)
	c.Set(middleware.RESPONSE, response)
	go service.FlushCommentSum(courseId, context.Ds, log)
}