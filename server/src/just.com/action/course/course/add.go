package course
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"just.com/action"
)

func CourseAddHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	session := context.Ds.NewSession()
	defer session.Close()

	beginErr := session.Begin()
	if beginErr != nil {

	}
}