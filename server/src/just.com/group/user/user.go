package user
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BuildRouter(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	userGroup.Any("/course", func(c *gin.Context) {
		log.Println("user")
		c.JSON(http.StatusOK,"12")
	})
}