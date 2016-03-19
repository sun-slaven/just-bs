package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/dto"
	"just.com/middleware"
	"net/http"
)

func CourseAddHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	token,tokenFlag := action.GetToken(c)
	if tokenFlag == false{
		return
	}
	session := context.Session
	log := context.Log
	courseService:= service.NewCourseService(session,log)
	courseDto := dto.NewCouseDto("数据库","简单db","syllbus","plan","experiment","major","college")
	courseId,addErr:= courseService.Add(*courseDto,token.UserId)
	if addErr != nil{
		return
	}
	response := middleware.NewResponse(http.StatusOK,courseId,nil)
	c.Set(middleware.RESPONSE,response)
}