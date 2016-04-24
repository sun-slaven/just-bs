package token
import (
	"github.com/gin-gonic/gin"
)

func BuildRouter(group *gin.RouterGroup)  {
	group.POST("/", LoginHandle)    //sign in
	group.DELETE("/")    //sign out
}
