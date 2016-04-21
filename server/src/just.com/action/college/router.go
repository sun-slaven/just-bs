package college
import "github.com/gin-gonic/gin"

func BuildRouter(group *gin.RouterGroup) {
	// colleges
	group.GET("/", CollegeList)
	group.GET("/:college_id/majors", MajorList)
}
