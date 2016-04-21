package notification
import "github.com/gin-gonic/gin"

func BuildRouter(g *gin.RouterGroup)  {
	// notification CRUD
	g.GET("")
	g.Any("/:notifi_id")
	g.Any("/:notifi_id")
}