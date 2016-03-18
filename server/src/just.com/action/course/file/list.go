package file
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileList(c *gin.Context)  {
	c.JSON(http.StatusOK,"FILE_LIST")
}