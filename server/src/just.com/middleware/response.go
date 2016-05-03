package middleware
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"just.com/err"
)


type Response  struct {
	Data  interface{}
	Error *err.HttpError
}

func NewResponse(data interface{}, err *err.HttpError) *Response {
	return &Response{
		Data:data,
		Error:err,
	}
}

//返回错误代码的response
func NewErrResponse(err *err.HttpError) *Response {
	return NewResponse(nil, err)
}

//返回data的response
func NewDataResponse(data interface{}) *Response {
	return NewResponse(data, nil)
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Header("Access-Control-Allow-Origin", "*")
		response := c.MustGet(MLEARNING_RESPONSE).(*Response)
		if response.Error == nil {
			c.JSON(http.StatusOK, response.Data)
		}else {
			errorMessage := make(map[string]string)
			errorMessage["message"] = response.Error.Error.Error()
			c.JSON(response.Error.Status, errorMessage)
		}
	}
}
