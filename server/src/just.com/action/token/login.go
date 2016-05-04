package token
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/query/vo/user"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string        `form:"password" json:"password"`
}


func LoginHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(LoginRequest)
	bindErr := c.Bind(request)
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
