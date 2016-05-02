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
		return
	}
	userVo, flag := user.CheckUser(request.Email, request.Password, context.Session, context.Log)
	if flag == false {
		context.Log.Println("用户名或密码错误")
		return
	}
	context.Response.Data = userVo
}
