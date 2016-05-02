package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"log"
	"just.com/query/vo/course"
	"just.com/model/db/table"
	"just.com/err"
)

type CourseListRequest struct {
	MajorId   string `form:"major_id"`
	CollegeId string `form:"college_id"`
	Page      int64 `form:"page"`
	PageSize  int64 `form:"page_size"`
}

/*需要过滤条件 major_id */
func CourseListHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := &CourseListRequest{Page:1, PageSize:10}
	bindErr := c.Bind(request)
	if bindErr != nil {
		log.Println(bindErr)
		return
	}
	if request.Page <= 0 || request.PageSize <= 0 {
		return
	}
	table := new(table.CourseTable)
	if request.MajorId != "" {
		table.MajorId = request.MajorId
	}
	if request.CollegeId != "" {
		table.CollegeId = request.CollegeId
	}
	courseVoList, e := course.LoadCourseVoList(table, context.Session, context.Log)
	if e != nil {
		context.Log.Println(e)
		context.Response.Error = err.NO_COURSE_LIST_FOUND
	}
	context.Response.Data = courseVoList
	return
}