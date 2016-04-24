package token
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/query/vo/user"
)

type LoginRequest struct {
	Email    string
	Password string
}


func LoginHandle(c *gin.Context) {
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		context.Response = middleware.NewResponse(http.StatusBadRequest, nil, nil)
		return
	}
	defer func() {
		context.Response = response
	}()
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
	response = middleware.NewResponse(http.StatusOK, userVo, nil)
}
