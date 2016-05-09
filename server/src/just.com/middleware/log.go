package middleware
import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"just.com/service/log"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		context := c.MustGet(MLEARNING_CONTEXT).(*Context)
		method := c.Request.Method
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			context.Log.Println(err)
		}
		body := string(bodyByte)
		url := c.Request.URL.String()
		logService := log.NewLogHttpService(context.Ds, context.Log)
		var error string
		status := 200
		if context.Response.Error != nil {
			error = context.Response.Error.Error.Error()
			status = context.Response.Error.Status
		}
		data := context.Response.Data
		go logService.Log(method, url, body, context.UserId, error, status, data)
	}
}
