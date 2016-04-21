package course
import (
	"just.com/query/vo/user"
	"just.com/model/db/table"
	"just.com/common"
	"github.com/go-xorm/xorm"
	"log"
)

type CourseCommentVo struct {
	UUID       string `json:"id"`
	Content    string `json:"content"`
	CreateUser *user.UserVo `json:"create_user"`
	CreateTime string `json:"create_time"`
}

func NewCommentVo(t *table.CourseCommentTable, session *xorm.Session, log *log.Logger) *CourseCommentVo {
	comment := new(CourseCommentVo)
	comment.UUID = t.UUID
	comment.Content = t.Content
	comment.CreateTime = common.TimeFormat(t.CreateTime)
	comment.CreateUser = user.LoadUserVo(t.CreateUser, session, log)
	return comment
}

/*根据courseId 加载评论*/
func LoadCommentVoList(courseId string, session *xorm.Session, log *log.Logger) []CourseCommentVo {
	commentVoList := make([]CourseCommentVo, 0)
	commentTableList := make([]table.CourseCommentTable, 0)
	sql := `SELECT * FROM "COURSE_COMMENT" WHERE "COURSE_ID" = ?
		AND "FROZEN_STATUS" = ?`
	findErr := session.Sql(sql,courseId,"N").Find(&commentTableList)
	if findErr != nil {
		log.Println(findErr)
		return nil
	}
	for _, commentTable := range commentTableList {
		comment := NewCommentVo(&commentTable, session, log)
		commentVoList = append(commentVoList, *comment)
	}
	return commentVoList
}