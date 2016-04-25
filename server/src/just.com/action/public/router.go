package public
import "github.com/gin-gonic/gin"

func BuildRouter(group *gin.RouterGroup) {
	group.GET("user_active", UserActiveHandle)
}
