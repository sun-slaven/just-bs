package file
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
)

func File(c *gin.Context)  {
	switch c.Request.Method {
	case action.METHOD_GET:
	case action.METHOD_POST:
	case action.METHOD_PUT:
		FileUpdate(c)
	case action.METHOD_DELETE:
		FileDelete(c)
	}
}