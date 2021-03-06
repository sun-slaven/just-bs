package middleware
import (
	"github.com/gin-gonic/gin"
	"just.com/service/rbac"
)

// TODO ROLE
func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := c.MustGet(MLEARNING_CONTEXT).(*Context)
		rbacService := rbac.NewRbacService(context.Session, context.Log)
		rbacService.Load("TEACHER")
		context.Log.Println(rbacService.GetData())
	}
}
