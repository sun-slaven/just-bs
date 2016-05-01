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
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	request := new(FileTokenRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	fileSystem := c.MustGet(middleware.MIDDLEWARE_FILE_SYSTEM).(*qiniu.QiniuFileSystem)
	context.Response.Data = fileSystem.MakeToken(request.Type)
}
