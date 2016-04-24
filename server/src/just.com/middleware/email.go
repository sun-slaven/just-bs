package middleware
import (
	"github.com/gin-gonic/gin"
	"just.com/etc"
	"just.com/service/email"
)

func EmailMiddleware(config etc.SendCloudConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailService := new(service.EmailService)
		emailService.Config = config
		c.Set(MLEARNING__EMAIL, emailService)
		c.Next()
	}
}
