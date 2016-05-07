package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/user"
)

func UserListHandle(c *gin.Context) {
	context := action.GetContext(c)
	userVoList, err := user.LoadUserVoList(context.Session, context.Log)
	if err != nil {
		context.Log.Println(context)
		context.Response.Error = err
	}
	context.Response.Data = userVoList
}
