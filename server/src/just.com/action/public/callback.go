package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/image"
	"net/http"
)

func UploadCallbackHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.UploadCallbackRequest)
	bindErr := c.Bind(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
	}
	imageService := new(image.ImageService)
	imageService.Add("", request.Key, "", request.W, request.H)
	c.JSON(http.StatusOK, nil)
}
