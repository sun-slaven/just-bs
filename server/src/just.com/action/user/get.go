package user
import (
	"just.com/action"
	"just.com/query/vo/user"
	"github.com/gin-gonic/gin"
)

func UserGetHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if !contextFlag {
		return
	}
	userId := c.Param("user_id")
	context.Response.Data = user.LoadUserVo(userId, context.Session, context.Log)
	return
}