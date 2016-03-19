package course
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"just.com/middleware"
	"just.com/action"
	"just.com/model/db/table"
	"just.com/query/vo/course"
)

func CourseListHandle(c *gin.Context)  {
	var response *middleware.Response
	context,contextFlag:=	action.GetContext(c)
	if contextFlag== false{
		response = middleware.NewResponse(http.StatusOK,nil,nil)
		return
	}
	session := context.Session
	log := context.Log
	courseVoList := make([]course.CourseVo,0)
	courseTableList := make([]table.CourseTable,0)
	sql := `SELECT * FROM "COURSE" WHERE "FROZEN_STATUS" = ?`
	findErr:= session.Sql(sql,"N").Find(&courseTableList)
	if findErr != nil{
		log.Println(findErr)
		response = middleware.NewResponse(http.StatusOK,nil,nil)
		return
	}
	for _,courseTable :=range courseTableList  {
		courseVo := course.NewCourseVo(&courseTable)
		courseVoList = append(courseVoList,*courseVo)
	}
	response = middleware.NewResponse(http.StatusOK,courseVoList,nil)
	log.Println(response)
	c.Set(middleware.RESPONSE,response)
}