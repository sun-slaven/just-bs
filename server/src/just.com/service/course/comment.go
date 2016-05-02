package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"just.com/query/vo/course"
)
/*return commentVo and error*/
func (self *CourseService)AddComment(content, courseId, userId string) (commentVo *course.CourseCommentVo, err error) {
	commentVo = new(course.CourseCommentVo)
	err = COURSE_COMMENT_ADD_ERR
	commentTable := new(table.CourseCommentTable)
	commentTable.UUID = uuid.New()
	commentTable.Content = content
	commentTable.CourseId = courseId
	commentTable.CreateUser = userId
	commentTable.CreateTime = time.Now()
	commentTable.FrozenStatus = "N"
	insertNum, insertErr := self.Session.InsertOne(commentTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		return
	}
	commentVo = course.NewCommentVo(commentTable, self.Session, self.Log)
	err = nil
	return
}

func (self *CourseService) DeleteComment(courseId string, commentId string) error {
	condiComment := new(table.CourseCommentTable)
	condiComment.UUID = commentId
	condiComment.CourseId = courseId
	newComment := new(table.CourseCommentTable)
	newComment.FrozenStatus = "Y"
	newComment.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Update(newComment, condiComment)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return COURSE_COMMENT_DELETE_ERR
	}
	//	go self.flushCommentSum(comment.CourseId)
	return nil
}