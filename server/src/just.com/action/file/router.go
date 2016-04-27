package file
import "github.com/gin-gonic/gin"

func BuildRouter(group *gin.RouterGroup) {
	// file
	group.GET("/", FileList)        //get list
	group.POST("/tokens", FileTokenHandle)        //get list
	group.POST("/", FileAdd)        // file add
	group.GET("/:id", File)    // file get update delete
}