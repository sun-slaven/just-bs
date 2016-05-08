package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/query/vo/course"
	"just.com/err"
	"just.com/value"
)
/*return commentVo and error*/
func (self *CourseService)AddComment(content, courseId, userId string) (commentVo *course.CourseCommentVo, error *err.HttpError) {
	commentVo = new(course.CourseCommentVo)
	commentTable := new(table.CourseCommentTable)
	commentTable.UUID = uuid.New()
	commentTable.Content = content
	commentTable.CourseId = courseId
	commentTable.CreateUser = userId
	commentTable.CreateTime = time.Now()
	commentTable.FrozenStatus = value.STATUS_ENABLED
	insertNum, insertErr := self.Session.InsertOne(commentTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		error = err.COURSE_COMMENT_INSERT_ERR
		return
	}
	commentVo = course.NewCommentVo(commentTable, self.Session, self.Log)
	error = nil
	return
}

func (self *CourseService) DeleteComment(courseId string, commentId string) (error *err.HttpError) {
	condiComment := new(table.CourseCommentTable)
	condiComment.UUID = commentId
	condiComment.CourseId = courseId
	newComment := new(table.CourseCommentTable)
	newComment.FrozenStatus = value.STATUS_DISABLED
	newComment.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Update(newComment, condiComment)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return err.COURSE_COMMENT_DELETE_ERR
	}
	return nil
}