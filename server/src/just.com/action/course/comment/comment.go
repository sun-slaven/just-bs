package comment
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
)

func Comment(c *gin.Context)  {
	switch c.Request.Method {
	case action.METHOD_POST:
		CommentAdd(c)
	case action.METHOD_DELETE:
		CommentDelete(c)
	}
}