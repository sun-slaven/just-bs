package college
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/college"
)

func CollegeList(c *gin.Context) {
	context := action.GetContext(c)
	collegeVoList := college.LoadCollegeVoList(context.Session, context.Log)
	context.Response.Data = collegeVoList
}
