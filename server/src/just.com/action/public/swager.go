package public
import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func SwaggerHandle(path, host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		swagger, err := template.ParseFiles(path + "/res/dist/swagger.yaml")
		if err != nil {
		}
		result := make(map[string]string)
		result["host"] = host
		swagger.Execute(c.Writer, result)
	}
}
