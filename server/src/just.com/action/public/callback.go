package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/image"
	"net/http"
	"just.com/common"
	"just.com/model/qiniu"
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
	switch request.Type {
	case qiniu.UPLOAD_TYPE_ICON:
		imageService := image.NewImageService(context.Session, context.Log)
		addErr := imageService.Add("", request.Key, "", request.W, request.H)
		if addErr != nil {
			context.Log.Println(addErr)
		}
	case qiniu.UPLOAD_TYPE_ATTACHMENT:

		
	case qiniu.UPLOAD_TYPE_VIDEO:

	}
	context.Response.Data = &UploadReturn{Key:request.Key}
}

type UploadReturn struct {
	Key string `json:"key"`
}
