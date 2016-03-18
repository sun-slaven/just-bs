package action
import (
	"github.com/gin-gonic/gin"
	"just.com/middleware"
	"errors"
)

const (
	METHOD_GET = "GET"
	METHOD_POST = "POST"
	METHOD_PUT = "PUT"
	METHOD_DELETE = "DELETE"
)

func GetContext(c *gin.Context) (*middleware.Context, error) {
	contextTemp, contextTempFlag := c.Get(middleware.MLEARNING_CONTENT)
	if contextTempFlag == false {
		return
	}
	context, contextFlag := contextTemp.(*middleware.Context)
	if contextFlag == false {
		return nil, errors.New("err")
	}
	return context, nil
}