package file
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
)

type FileTokenRequest struct {
	Name     string `json:"name"`
	FileType string `json:"file_type"`
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
}
