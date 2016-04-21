package middleware
import (
	"github.com/gin-gonic/gin"
	"just.com/service/token"
)


func TokenTest(c *gin.Context)  {
	token := service.NewXToken("6688789c-1cb3-4303-9558-bcfd4c3b5d9e","aa5eba0a-703c-4801-955b-1f44997738fe")
	c.Set(MIDDLEWARE_TOKEN,token)
	c.Next()
}