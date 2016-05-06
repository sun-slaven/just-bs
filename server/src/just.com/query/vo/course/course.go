package course
import (
	"just.com/query/vo/file"
	"just.com/model/db/table"
	"github.com/go-xorm/xorm"
	"log"
	"just.com/query/vo/college"
	"just.com/query/vo/user"
	"just.com/err"
	"just.com/common"
)

type CourseVo struct {
	UUID         string `json:"id"`
	Name         string    `json:"name"`
	Description  string `json:"description"`
	Introduction string    `json:"introduction"`
	Syllabus     string    `json:"syllabus"`
	Wish         string        `json:"wish"`
	Experiment   string        `json:"experiment"`
	Icon         *file.ImageVo    `json:"icon"`
	MarkSum      int64    `json:"mark_sum"`
	CommentSum   int64 `json:"comment_sum"`
	Major        *college.MajorVo    `json:"major"`
	College      *college.CollegeVo    `json:"college"`
	//	Point        int64 `json:"point"`
	//	PointPerson  int64 `json:"point_person"`
	//	PointStatus  string `json:"point_status"`
	Teacher      *user.UserVo `json:"teacher"`
	VideoUrl     string `json:"video_url"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

// 什么都没有用这个就行了
func LoadCourseVo(courseTable *table.CourseTable, session *xorm.Session, log *log.Logger) (cv *CourseVo, error *err.HttpError) {
	getFlag, getErr := session.Get(courseTable)
	if getFlag == false {
		if getErr != nil {
			log.Println(getErr)
		}
		return
	}
	return LoadCourseVoFromTable(courseTable, session, log)
}

// 已经有 courseTable用这个
func LoadCourseVoFromTable(courseTable *table.CourseTable, session *xorm.Session, log *log.Logger) (cv *CourseVo, error *err.HttpError) {
	cv = new(CourseVo)
	cv.UUID = courseTable.UUID
	cv.Name = courseTable.Name
	cv.Description = courseTable.Description
	cv.Introduction = courseTable.Introduction
	cv.Syllabus = courseTable.Syllabus
	cv.Wish = courseTable.Wish
	cv.Experiment = courseTable.Experiment

	// icon
	icon := file.NewImageVo(courseTable.IconUrl, courseTable.IconWidth, courseTable.IconHeight)
	cv.Icon = icon
	cv.MarkSum = courseTable.MarkSum
	cv.CommentSum = courseTable.CommentSum

	// college
	collegeTable := &table.CollegeTable{UUID:courseTable.CollegeId}
	getFlag, getErr := session.Get(collegeTable)
	if getFlag == false {
		if getErr != nil {
			log.Println(getErr)
		}
		error = err.NO_COLLEGE_OR_MAJOR_FOUND
	}
	cv.College = college.NewCollegeVo(collegeTable)

	// major
	majorTable := &table.MajorTable{UUID:courseTable.MajorId}
	getFlag, getErr = session.Get(majorTable)
	if getFlag == false {
		if getErr != nil {
			log.Println(getErr)
		}
		error = err.NO_COLLEGE_OR_MAJOR_FOUND
	}
	cv.Major = college.NewMajorVo(majorTable)

	// teacher
	cv.Teacher = user.LoadUserVo(courseTable.TeacherId, session, log)

	cv.VideoUrl = courseTable.VideoUrl
	cv.CreateTime = common.TimeFormat(courseTable.CreateTime)
	cv.UpdateTime = common.TimeFormat(courseTable.UpdateTime)

	//	if courseTable.PointPerson > 0 {
	//		cv.Point = courseTable.Points / courseTable.PointPerson
	//	}else {
	//		cv.Point = 0
	//	}
	//	cv.PointPerson = courseTable.PointPerson
	//	cv.PointStatus = "N"
	return cv, nil
}

/*load point status*/
func (self *CourseVo) LoadPointStatus(userId string, session *xorm.Session, log *log.Logger) {
	if self.UUID == "" {
		return
	}
	pointTable := new(table.CoursePointTable)
	pointTable.CourseId = self.UUID
	pointTable.UserID = userId
	//	count, countErr := session.Count(pointTable)
	//	if count == 0 {
	//		if countErr != nil {
	//			log.Println(countErr)
	//		}
	//		self.PointStatus = "N"
	//		return
	//	}
	//	self.PointStatus = "Y"
}