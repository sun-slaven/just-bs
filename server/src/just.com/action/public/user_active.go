package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/user"
	"just.com/service/token"
)

func UserActiveHandle(c *gin.Context) {
	context := action.GetContext(c)
	userId := c.Query("user_id")
	id := c.Query("id")
	// 1. check token
	tokenService := service.NewTokenService(context.Session, context.Log)
	checkFlag := tokenService.Check(service.NewUserToken(id, userId))
	if checkFlag == false {
		return
	}
	// 2. active
	userService := user.NewUserService(context.Session, context.Log)
	activeFlag := userService.Active(userId)
	if activeFlag == false {
		return
	}
}
