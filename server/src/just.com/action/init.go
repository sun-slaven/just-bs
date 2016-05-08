package action
import (
	"github.com/gin-gonic/gin"
	"just.com/middleware"
	"just.com/err"
	"just.com/common"
)

func GetContext(c *gin.Context) (*middleware.Context) {
	defer func() {
		if e := recover(); e != nil {
			c.Set(middleware.MLEARNING_CONTEXT, middleware.NewErrResponse(err.NO_CONTEXT))
		}
	}()
	context := c.MustGet(middleware.MLEARNING_CONTEXT).(*middleware.Context)
	if common.IsNil(context.Session, context.Log, context.Ds, context.Response) {
		panic(err.NO_CONTEXT)
	}
	return context
}

func BindErrHandle(context *middleware.Context, e error) {
	context.Log.Println(e)
	context.Response.Error = err.PARAM_ERR
}