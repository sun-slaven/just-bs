package action
import (
	"github.com/gin-gonic/gin"
	"just.com/middleware"
	"just.com/service/token"
)

const (
	METHOD_GET = "GET"
	METHOD_POST = "POST"
	METHOD_PUT = "PUT"
	METHOD_DELETE = "DELETE"
)

func GetContext(c *gin.Context) (*middleware.Context, bool) {
	contextTemp, contextTempFlag := c.Get(middleware.MLEARNING_CONTENT)
	if contextTempFlag == false {
		return nil,contextTempFlag
	}
	context, contextFlag := contextTemp.(*middleware.Context)
	if contextFlag == false {
		return nil, contextFlag
	}
	return context, true
}

func GetToken(c *gin.Context) (*service.UserToken, bool) {
	tokenTemp, tokenTempFlag := c.Get(middleware.MIDDLEWARE_TOKEN)
	if tokenTempFlag == false {
		return nil,tokenTempFlag
	}
	token, tokenFlag := tokenTemp.(*service.UserToken)
	if tokenFlag == false {
		return nil, tokenFlag
	}
	return token, true
}