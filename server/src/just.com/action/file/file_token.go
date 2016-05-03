package file
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"just.com/model/qiniu"
)

type FileTokenRequest struct {
	Type string `json:"type"`
}

func FileTokenHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(FileTokenRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		action.BindErrHandle(context, bindErr)
		return
	}
	fileSystem := c.MustGet(middleware.MIDDLEWARE_FILE_SYSTEM).(*qiniu.QiniuFileSystem)
	context.Response.Data = fileSystem.MakeToken(request.Type)
}
