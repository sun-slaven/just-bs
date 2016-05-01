package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
)

func UploadCallbackHandle(c *gin.Context) {
	_, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}

}
