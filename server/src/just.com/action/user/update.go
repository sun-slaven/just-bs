package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/user"
	"just.com/dto"
)

func UserUpdateHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if !contextFlag {
		return
	}
	userId := c.Param("user_id")
	request := new(dto.UserUpdateRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	userService := user.NewUserService(context.Session, context.Log)
	userVo, updateErr := userService.Update(userId, request.Name, request.Email, request.IconUrl)
	if updateErr != nil {
		context.Log.Println(updateErr)
		return
	}
	context.Response.Data = userVo
}
