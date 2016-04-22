package token
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/token"
)

func Login(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return nil
	}
	log := context.Log
	session := context.Session

	tokenService := service.NewTokenService(session, log)
	tokenService.Make("123")
}
