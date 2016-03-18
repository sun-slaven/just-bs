package middleware
import (
	"github.com/gin-gonic/gin"
	"log"
	"just.com/model/db"
	"net/http"
	"github.com/go-xorm/xorm"
)

type Context struct {
	Ds      *db.DataSource
	Log     *log.Logger
	Session *xorm.Session
}

type Response  struct {
	status int
	data   interface{}
	err    error
}

func NewResponse(status int, data interface{}, err error) *Response {
	response := new(Response)
	response.status = status
	response.data = data
	response.err = err
	return response
}

func ContextMiddleWare(ds *db.DataSource, log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Log = log
		context.Ds = ds
		context.Session = ds.NewSession()
		defer context.Session.Close()
		beginErr := context.Session.Begin()
		if beginErr != nil {
			return
		}
		c.Set(MLEARNING_CONTENT, context)
		c.Next()    //next
		response, err := c.MustGet(RESPONSE).(*Response)
		if err != nil {
			return
		}
		// session rollback or commit
		if response.status != http.StatusOK {
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
		//response
		switch response.status {
		case http.StatusOK:
			c.JSON(http.StatusOK, response.data)
		case http.StatusNonAuthoritativeInfo:
			c.JSON(http.StatusNonAuthoritativeInfo, response.err)
		default:
			c.JSON(http.StatusBadRequest, response.err)
		}
	}
}
