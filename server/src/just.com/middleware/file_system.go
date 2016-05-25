package middleware
import (
	"github.com/gin-gonic/gin"
	"just.com/model/qiniu"
)

func FileSystemMiddleware(fs *qiniu.QiniuFileSystem) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set(MIDDLEWARE_FILE_SYSTEM, fs)
		c.Next()
	}
}
