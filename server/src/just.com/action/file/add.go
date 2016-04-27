package file
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/file"
	"just.com/middleware"
	"net/http"
)

func FileAdd(c *gin.Context) {
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
	// core
	fileService := service.NewFileService(session, log)
	fileId, addErr := fileService.Add(courseId, "file name", "httpbaidu", token.UserId)
	if addErr != nil {
		log.Println(addErr)
		return
	}
	context.Response = middleware.NewResponse(http.StatusOK, fileId, nil)
}