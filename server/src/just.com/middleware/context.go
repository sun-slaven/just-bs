package middleware
import (
	"github.com/gin-gonic/gin"
	"log"
	"just.com/model/db"
	"github.com/go-xorm/xorm"
	"just.com/err"
)

type Context struct {
	Ds       *db.DataSource
	Log      *log.Logger
	Session  *xorm.Session
	UserId   string
	Response *Response
}

func ContextMiddleWare(ds *db.DataSource, log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Log = log
		context.Ds = ds
		context.Session = ds.NewSession()
		context.Response = NewResponse(nil, nil)
		defer context.Session.Close()
		beginErr := context.Session.Begin()
		if beginErr != nil {
			log.Println(beginErr)
			c.Set(MLEARNING_RESPONSE, NewErrResponse(err.NO_CONTEXT))
			c.Abort()
			return
		}
		c.Set(MLEARNING_CONTEXT, context)
		c.Next()
		c.Set(MLEARNING_RESPONSE, context.Response)
		// session rollback or commit
		if context.Response.Error != nil {
			// commit
			commitErr := context.Session.Commit()
			if commitErr != nil {
				context.Log.Println(commitErr)
				rollbackErr := context.Session.Rollback()
				if rollbackErr != nil {
					context.Log.Println(rollbackErr)
				}
			}
		} else {
			// rollback
			rollbackErr := context.Session.Rollback()
			if rollbackErr != nil {
				context.Log.Println(rollbackErr)
				return
			}
		}
	}
}
