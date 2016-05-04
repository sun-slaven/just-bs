package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/user"
)

const INIT_PASSWORD = "123456"

func RestPassword(c *gin.Context) {
	context := action.GetContext(c)
	userService := user.NewUserService(context.Session, context.Log)
	userId := c.Param("user_id")
	resetErr := userService.RestPassword(userId, INIT_PASSWORD)
	if resetErr != nil {
		context.Log.Println(resetErr)
		context.Response.Error = resetErr
	}
}
