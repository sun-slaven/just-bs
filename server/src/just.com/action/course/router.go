package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action/course/course"
	"just.com/action/course/comment"
	"just.com/action/course/file"
	"just.com/action/course/point"
	"just.com/action/course/mark"
)

func BuildRouter(group *gin.RouterGroup) {
	// course
	group.GET("", course.CourseListHandle)    // get list
	group.POST("", course.CourseAddHandle)            //	add
	group.Any("/:course_id", course.CourseHandle)    // get put delete
	// comment
	group.GET("/:course_id/comment", comment.CommentList)        // get comment list
	group.POST("/:course_id/comment", comment.CommentAdd)    // add comment
	group.Any("/:course_id/comment/:comment_id", comment.Comment)    //comment(get put delete)
	// file
	group.GET("/:course_id/file", file.FileList)        //get list
	group.POST("/:course_id/file", file.FileAdd)        // file add
	group.Any("/:course_id/file/:file_id", file.File)    // file get update delete
	// mark
	group.POST("/:course_id/mark", mark.MarkAdd)
//	group.DELETE("/:course_id/mark", mark.MarkAdd)
	// point
	group.POST("/:course_id/point", point.PointAdd)
}