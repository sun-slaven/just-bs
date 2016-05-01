package middleware
import (
	"github.com/gin-gonic/gin"
	"log"
	"just.com/model/db"
	"net/http"
	"github.com/go-xorm/xorm"
)

type Context struct {
	Ds       *db.DataSource
	Log      *log.Logger
	Session  *xorm.Session
	UserId   string
	Err      error
	Response *Response
}

func ContextMiddleWare(ds *db.DataSource, log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Log = log
		context.Ds = ds
		context.Session = ds.NewSession()
		response := NewResponse(http.StatusOK, nil, nil)
		context.Response = response
		defer context.Session.Close()
		beginErr := context.Session.Begin()
		if beginErr != nil {
			log.Println(beginErr)
			return
		}
		c.Set(MLEARNING_CONTENT, context)
		c.Next()
		// session rollback or commit
		response = context.Response
		c.Set(MLEARNING_RESPONSE, context.Response)
		if response.Status == http.StatusOK {
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
