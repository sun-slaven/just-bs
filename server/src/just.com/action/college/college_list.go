package college
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/query/vo/college"
)

func CollegeList(c *gin.Context) {
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	defer func() {
		context.Response = response
	}()
	collegeVoList := college.LoadCollegeVoList(context.Session, context.Log)
	response = middleware.NewResponse(http.StatusOK, collegeVoList, nil)
}
