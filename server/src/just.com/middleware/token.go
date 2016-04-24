package middleware
import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"just.com/model/db/table"
	"just.com/service/token"
	"encoding/json"
)

func TokenMiddleWare(c *gin.Context) {
	contextInter, exists := c.Get(MLEARNING_CONTENT)
	context := contextInter.(*Context)
	if exists == false {
		c.Abort()
	}
	// is login
	// login
	userId := c.Param("user_id")
	if strings.TrimSpace(userId) != "" {
		c.Next()
	}else {
		// other
		xTokenStr := c.Request.Header.Get("X-Token")
		if strings.TrimSpace(xTokenStr) == "" {
			c.Abort()
			return
		}
		userToken := new(service.UserToken)
		unmarshalErr := json.Unmarshal([]byte(xTokenStr), userToken)
		if unmarshalErr != nil {
			c.Abort()
			return
		}
		tokenService := service.TokenService{}
		tokenService.Session = context.Ds.NewSession()
		defer tokenService.Session.Close()
		tokenService.Log = context.Log
		if tokenService.Check(userToken) == false {
			response := table.Response{}
			response.Ok = 0
			response.Err = ""
			c.JSON(http.StatusOK, response)
			c.Abort()
		}
	}
	c.Next()
}
