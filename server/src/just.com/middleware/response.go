package middleware
import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response  struct {
	Status int
	Data   interface{}
	Error  error
}

func NewResponse(status int, data interface{}, err error) *Response {
	response := new(Response)
	response.Status = status
	response.Data = data
	response.Error = err
	return response
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		response := c.MustGet(MLEARNING_RESPONSE).(*Response)
		if response.Status == http.StatusOK {
			c.JSON(http.StatusOK, response.Data)
		}else {
			c.JSON(response.Status, response.Error.Error())
		}
	}
}
