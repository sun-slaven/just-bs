package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/middleware"
	"net/http"
	"just.com/query/vo/course"
	"just.com/model/db/table"
)

func CourseHandle(c *gin.Context) {
	context, contextFlag := action.GetContext(c)
	if contextFlag == false{
		return
	}
	session := context.Session
	log := context.Log
	token,tokenFlag :=action.GetToken(c)
	if tokenFlag == false{
		return
	}
	courseId := c.Param("course_id")
	cs := service.NewCourseService(context.Session, context.Log)
	// response data
	var data interface{}
	switch c.Request.Method {
	// 根据id获取course
	case action.METHOD_GET:
		courseTable := new(table.CourseTable)
		getFlag,getErr:= session.Id(courseId).Get(courseTable)
		if getFlag == false{
			if getErr !=nil{
				log.Println(getErr)
			}
		}
		courseVo := course.NewCourseVo(courseTable)
		courseVo.LoadPointStatus(token.Id,session,log)
		data = courseVo
	case action.METHOD_PUT:
	// 根据id删除课程
	case action.METHOD_DELETE:
		cs.Delete(courseId)
	}
	context.Response = middleware.NewResponse(http.StatusOK,data , nil)
}
