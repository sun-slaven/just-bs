package middleware
import "github.com/gin-gonic/gin"

func LogMiddleware(c *gin.Context) {
	context := c.MustGet(MLEARNING_CONTEXT).(Context)
	context.Log.Println("ONE REQUEST BEGIN")
	c.Next()
	context.Log.Println("ONE REQUEST ENDS")
}
