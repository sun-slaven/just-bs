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
	group.GET("/", course.CourseListHandle)    // get list
	group.POST("/", course.CourseAddHandle)            //	add
	group.Any("/:course_id", course.CourseHandle)    // get put delete
	// comment
	group.GET("/:course_id/comments", comment.CommentList)        // get comment list
	group.POST("/:course_id/comments", comment.CommentAdd)    // add comment
	group.Any("/:course_id/comments/:comment_id", comment.Comment)    //comment(get put delete)
	// file
	group.GET("/:course_id/files", file.FileList)        //get list
	group.POST("/:course_id/files", file.FileAdd)        // file add
	group.Any("/:course_id/files/:file_id", file.File)    // file get update delete
	// mark
	group.POST("/:course_id/mark", mark.MarkAdd)
	group.DELETE("/:course_id/mark", mark.MarkCancel)
	// point
	group.POST("/:course_id/point", point.PointAdd)
}