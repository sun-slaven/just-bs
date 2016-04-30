package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"strings"
	"net/http"
	"just.com/middleware"
	"just.com/dto"
)

func User(c *gin.Context) {
	contextTemp, contextTempFlag := c.Get(middleware.MLEARNING_CONTENT)
	if contextTempFlag == false {
		return
	}
	context, contextFlag := contextTemp.(*middleware.Context)
	if contextFlag == false {
		return
	}
	userId := c.Param("user_id")
	if strings.TrimSpace(userId) == "" {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	var result interface{}    // result
	var err error        // err code
	switch c.Request.Method {
	case action.METHOD_GET:
		result, err = UserGet(context, userId)
	case action.METHOD_PUT:
		userDto := new(dto.UserDto)
		bindErr := c.BindJSON(userDto)
		if bindErr != nil {

		}
		result, err = UserUpdate(context, userId, userDto)
	case action.METHOD_DELETE:
		result, err = UserDelete(context, userId)
	}
	switch err {
	case nil:
		c.JSON(http.StatusOK, result)
	case middleware.NO_AUTHORITATION_ERR:
		c.JSON(http.StatusNonAuthoritativeInfo, result)
	default:
		c.JSON(http.StatusNotModified, result)
	}
}