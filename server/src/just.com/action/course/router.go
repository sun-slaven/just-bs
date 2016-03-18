package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action/course/course"
	"just.com/action/course/comment"
	"just.com/action/course/file"
)

func BuildRouter(group *gin.RouterGroup) {
	// course
	group.GET("/index", course.CourseListHandle)    // get list
	group.POST("", course.CourseAddHandle)            //	add
	group.Any("/:course_id", course.CourseHandle)    // get put delete
	// comment
	group.GET("/:course_id/comment/index", comment.CommentList)        // get comment list
	group.POST("/:course_id/comment")    // add comment
	group.Any("/:course_id/comment/:comment_id", comment.Comment)    //comment(get put delete)
	// file
	group.GET("/:course_id/file/index", file.FileList)        //get list
	group.POST("/:course_id/file", file.FileAdd)        // file add
	group.Any("/:course_id/file/:file_id", file.File)    // file get update delete
}