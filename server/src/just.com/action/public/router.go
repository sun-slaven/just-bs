package public
import "github.com/gin-gonic/gin"

func BuildRouter(group *gin.RouterGroup, path, port string) {
	group.GET("/user_active", UserActiveHandle)
	group.GET("/callback", )
	group.GET("/swagger", SwaggerHandle(path, port))
}
