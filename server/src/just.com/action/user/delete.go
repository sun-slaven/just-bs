package user
import (
	"just.com/action"
	"just.com/service/user"
	"github.com/gin-gonic/gin"
	"just.com/common"
	"just.com/err"
)

func UserDelete(c *gin.Context) {
	context := action.GetContext(c)
	userId := c.Param("user_id")
	if common.IsEmpty(userId) {
		context.Response.Error = err.NOT_USER_ID_FOUND
	}
	userService := user.NewUserService(context.Session, context.Log)
	frozenErr := userService.Frozen(userId)
	if frozenErr != nil {
		context.Log.Println(frozenErr)
		context.Response.Error = frozenErr
	}
}