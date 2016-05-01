package router

import (
	"github.com/gin-gonic/gin"
	"just.com/action/college"
	"just.com/action/course/course"
	"just.com/action/file"
	"just.com/action/course/comment"
	"just.com/action/course/mark"
	"just.com/action/course/point"
	"just.com/action/user"
	"just.com/action/token"
)

func BuildRouter(mainRouter *gin.RouterGroup) {
	collegeRouter := mainRouter.Group("/colleges")
	{
		collegeRouter.GET("/", college.CollegeList)
	}
	courseRouter := mainRouter.Group("/courses")
	{
		// course
		courseRouter.GET("/", course.CourseListHandle)
		courseRouter.POST("/", course.CourseAddHandle)
		courseRouter.GET("/:course_id", course.CourseGetHandle)
		courseRouter.DELETE("/:course_id", course.CourseDeleteHandle)
		// comment
		courseRouter.GET("/:course_id/comments", comment.CommentList)
		courseRouter.POST("/:course_id/comments", comment.CommentAdd)
		courseRouter.DELETE("/:course_id/comments/:comment_id", comment.CommentDeleteHandle)
		// mark
		courseRouter.POST("/:course_id/marks", mark.MarkAdd)
		courseRouter.DELETE("/:course_id/marks", mark.MarkCancel)
		// point
		courseRouter.POST("/:course_id/points", point.PointAdd)
	}
	userRouter := mainRouter.Group("/users")
	{
		userRouter.POST("/", user.RegisterHandle)
		userRouter.GET("/:user_id/courses", user.CourseListHandle)
	}
	tokenRouter := mainRouter.Group("/tokens")
	{
		tokenRouter.POST("/", token.LoginHandle)    //sign in
		tokenRouter.DELETE("/")    //sign out
	}
	fileRouter := mainRouter.Group("/files")
	{
		fileRouter.GET("/", file.FileList)
		fileRouter.POST("/tokens", file.FileTokenHandle)
		fileRouter.POST("/", file.FileAdd)
		fileRouter.GET("/:id", file.File)
	}
}
