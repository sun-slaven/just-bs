package college
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/college"
)

func CollegeList(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	response := context.Response
	defer func() {
		context.Response = response
	}()
	collegeVoList := college.LoadCollegeVoList(context.Session, context.Log)
	response.Data = collegeVoList
}
