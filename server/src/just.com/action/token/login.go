package token
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/user"
)

type LoginRequest struct {
	Email    string
	Password string
}


func LoginHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(LoginRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
		return
	}
	userVo, userVoErr := user.CheckUser(request.Email, request.Password, context.Session, context.Log)
	if userVoErr != nil {
		context.Log.Println(userVoErr)
		context.Response.Error = userVoErr
		return
	}
	context.Response.Data = userVo
}
