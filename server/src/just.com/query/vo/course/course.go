package course
import (
	"just.com/query/vo/file"
	"just.com/model/db/table"
	"github.com/go-xorm/xorm"
	"log"
)

type CourseVo struct {
	UUID         string `json:"id"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	Syllabus     string    `json:"syllabus"`
	Plan         string        `json:"plan"`
	Experiment   string        `json:"experiment"`
	Icon         file.ImageVo    `json:"icon"`
	MarkSum      int64    `json:"mark_sum"`
	CommentSum   int64 `json:"comment_sum"`
	Major        string    `json:"major"`
	College      string    `json:"college"`
	Point        int64 `json:"point"`
	PointPerson  int64 `json:"point_person"`
	PointStatus  string `json:"point_status"`
}

func NewCourseVo(t *table.CourseTable) *CourseVo {
	cv := new(CourseVo)
	cv.UUID = t.UUID
	cv.Name = t.Name
	cv.Introduction = t.Introduction
	cv.Syllabus = t.Syllabus
	cv.Plan = t.Plan
	cv.Experiment = t.Experiment

	icon := file.ImageVo{}
	icon.Url = t.IconUrl
	icon.Width = t.IconWidth
	icon.Height = t.IconHeight

	cv.Icon = icon
	cv.MarkSum = t.MarkSum
	cv.CommentSum = t.CommentSum
	cv.Major = t.Major
	cv.College = t.College
	if t.PointPerson > 0 {
		cv.Point = t.Points / t.PointPerson
	}else {
		cv.Point = 0
	}
	cv.PointPerson = t.PointPerson
	return cv
}

/*load point status*/
func (self *CourseVo) LoadPointStatus(userId string, session *xorm.Session, log *log.Logger) {
	if self.UUID == "" {
		return
	}
	pointTable := new(table.CoursePointTable)
	pointTable.CourseId = self.UUID
	pointTable.UserID = userId
	count, countErr := session.Count(pointTable)
	if count == 0 {
		if countErr != nil {
			log.Println(countErr)
		}
		self.PointStatus = "N"
		return
	}
	self.PointStatus = "Y"
}