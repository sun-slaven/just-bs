package service
import (
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
)
/*return commentId and error*/
func (self *CourseService)AddComment(content, courseId, userId string) (string, error) {
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
		return "", COURSE_COMMENT_ADD_ERR
	}
//	go self.flushCommentSum(courseId)
	return commentTable.UUID, nil
}

func (self *CourseService) DeleteComment(commentId string) error {
	comment := new(table.CourseCommentTable)
	comment.FrozenStatus = "Y"
	comment.FrozenTime = time.Now()
	updateNum, updateErr := self.Session.Id(commentId).Update(comment)
	if updateNum == 0 {
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
		return COURSE_COMMENT_DELETE_ERR
	}
//	go self.flushCommentSum(comment.CourseId)
	return nil
}