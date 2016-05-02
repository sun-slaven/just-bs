package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/image"
	"net/http"
	"just.com/common"
)

func UploadCallbackHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.UploadCallbackRequest)
	bindErr := c.Bind(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	context.Log.Println(request)
	if common.IsEmpty(request.Key) {
		c.JSON(http.StatusOK, nil)
		return
	}
	imageService := image.NewImageService(context.Session, context.Log)
	addErr := imageService.Add("", request.Key, "", request.W, request.H)
	if addErr != nil {
		context.Log.Println(addErr)
	}
	context.Response.Data = &UploadReturn{Key:request.Key, Hash:request.Hash}
}

type UploadReturn struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}
