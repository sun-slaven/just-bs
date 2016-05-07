package token
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/token"
)

func LogoutHandle(c *gin.Context) {
	context := action.GetContext(c)
	tokenService := service.NewTokenService(context.Session, context.Log)
	deleteErr := tokenService.Delete(context.UserId)
	if deleteErr != nil {
		context.Log.Println(deleteErr)
		context.Response.Error = deleteErr
	}
}