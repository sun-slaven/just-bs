package middleware
import (
	"github.com/gin-gonic/gin"
	"strings"
	"just.com/service/token"
	"encoding/json"
	"just.com/etc"
	"just.com/err"
)

func TokenMiddleWare(whiteList []etc.White) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := c.MustGet(MLEARNING_CONTEXT).(*Context)
		whiteFlag := false
		for _, white := range whiteList {
			if white.Match(c.Request.Method, c.Request.URL.Path) {
				whiteFlag = true
				break
			}
		}
		flag := false        // 验证成功的标志
		// 验证白名单
		if whiteFlag {
			c.Next()
		} else {
			defer func() {
				if !flag {
					context.Response = NewErrResponse(err.STATUS_UNAUTHORIZED)
					c.Abort()
				}
			}()
			// 权限验证
			userTokenText := c.Request.Header.Get(MLEARNING_HEADER_AUTHORIZATION)
			if strings.TrimSpace(userTokenText) == "" {
				return
			}
			userToken := new(service.UserToken)
			unmarshalErr := json.Unmarshal([]byte(userTokenText), userToken)
			if unmarshalErr != nil {
				return
			}
			tokenService := service.NewTokenService(context.Session, context.Log)
			// 验证不成功
			if !tokenService.Check(userToken) {
				return
			}
			flag = true
			context.UserId = userToken.UserId
			c.Next()
		}
	}
}

