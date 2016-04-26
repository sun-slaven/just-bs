package course
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"just.com/middleware"
	"just.com/action"
	"log"
	"just.com/query/vo/course"
	"just.com/model/db/table"
)

type CourseListRequest struct {
	MajorId   string `form:"major_id"`
	CollegeId string `form:"college_id"`
	Mark      bool        `form:"mark"`
	Page      int64 `form:"page"`
	PageSize  int64 `form:"page_size"`
}

//func NewCourseLitRequest() *CourseListRequest {
//	return &CourseListRequest{Mark:true, Page:1, PageSize:20}
//}

/*需要过滤条件 major_id */
func CourseListHandle(c *gin.Context) {
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	defer func() {
		context.Response = response
	}()
	request := new(CourseListRequest)
	log.Println(request)
	bindErr := c.Bind(request)
	if bindErr != nil {
		log.Println(bindErr)
		return
	}
	courseVoList, err := course.LoadCourseVoList(&table.CourseTable{MajorId:request.MajorId, CollegeId:request.CollegeId}, context.Session, context.Log)
	if err != nil {
		context.Log.Println(err)
	}
	response = middleware.NewResponse(http.StatusOK, courseVoList, nil)
	return
}