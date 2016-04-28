package file
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/model/qiniu"
)

type FileTokenRequest struct {
	Suffix string `json:"suffix"`
	Type   string `json:"type"`
}

func FileTokenHandle(c *gin.Context) {
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	defer func() {
		context.Response = response
	}()
	request := new(FileTokenRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	fileSystem := c.MustGet(middleware.MIDDLEWARE_FILE_SYSTEM).(*qiniu.QiniuFileSystem)
	response = middleware.NewResponse(http.StatusOK, fileSystem.MakeToken(request.Suffix, request.Type), nil)
}
