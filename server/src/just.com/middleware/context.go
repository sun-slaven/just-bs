package middleware
import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"log"
)

type Context struct {
	Session *xorm.Session
	Log     *log.Logger
}

func ContextMiddleWare(c *gin.Context) {
	context := Context{}
	context.Session = nil
	context.Log = nil
	c.Set(JUST_CONTEXT, context)
	c.Next()
}
