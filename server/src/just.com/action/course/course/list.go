package course
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CourseListHandle(c *gin.Context)  {
	c.JSON(http.StatusOK ,"LIST")
}