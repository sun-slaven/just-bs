package id
import "github.com/gin-gonic/gin"

func BuildRouter(g *gin.RouterGroup) {
	g.Any("/token")
}