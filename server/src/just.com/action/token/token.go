package id
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
)

func Token(c *gin.Context) {
	switch c.Request.Method {
	case action.METHOD_GET:
	case action.METHOD_POST:
	case action.METHOD_PUT:
	case action.METHOD_DELETE:
	}
}