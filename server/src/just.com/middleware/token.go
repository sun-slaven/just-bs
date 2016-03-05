package middleware
import (
	"github.com/gin-gonic/gin"
	"strings"
	"just.com/service"
	"net/http"
	"just.com/model/db/table"
)

type XToken struct {
	Id     string
	UserId string
}

func TokenMiddleWare(c *gin.Context) {
	contextInter, exists := c.Get(JUST_CONTEXT)
	context := contextInter.(Context)
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
		xToken := c.Request.Header.Get("X-Token")
		if strings.TrimSpace(xToken) == "" {
			c.Abort()
		}
		tokenService := service.TokenService{}
		tokenService.Session = context.Session
		tokenService.Log = context.Log
		if tokenService.Check() == false {
			response := table.Response{}
			response.Ok = 0
			response.Err = ""
			c.JSON(http.StatusOK, response)
			c.Abort()
		}
	}
	c.Next()
}
