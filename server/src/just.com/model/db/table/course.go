package table
import "time"

type CourseTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	Name         string `xorm:"'NAME'"`
	Description  string `xorm:"'DESCRIPTION'"`
	Introduction string `xorm:"'INTRODUCTION'"`
	Syllabus     string `xorm:"'SYLLABUS'"`
	Wish         string `xorm:"'WISH'"`
	Experiment   string `xorm:"'EXPERIMENT'"`
	IconId       string `xorm:"'ICON_ID'"`
	IconUrl      string `xorm:"'ICON_URL'"`
	IconWidth    int64 `xorm:"'ICON_WIDTH'"`
	IconHeight   int64 `xorm:"'ICON_HEIGHT'"`
	MarkSum      int64 `xorm:"'MARK_SUM'"`
	CommentSum   int64 `xorm:"'COMMENT_SUM'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CREATE_USER'"`
	UpdateTime   time.Time `xorm:"updated 'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"'FROZEN_TIME'"`
	MajorId      string    `xorm:"'MAJOR_ID'"`
	CollegeId    string    `xorm:"'COLLEGE_ID'"`
	TeacherId    string `xorm:"'TEACHER_ID'"`
	Points       int64    `xorm:"'POINTS'"`
	PointPerson  int64 `xorm:"'POINT_PERSON'"`
}

func (self *CourseTable)TableName() string {
	return "COURSE"
}

type CourseMarkTable struct {
	UUID       string    `xorm:"pk 'UUID'"`
	UserId     string    `xorm:"'USER_ID'"`
	CourseId   string    `xorm:"'COURSE_ID'"`
	CreateTime time.Time    `xorm:"created 'CREATE_TIME'"`
}


func (self *CourseMarkTable)TableName() string {
	return "COURSE_MARK"
}

type CourseCommentTable struct {
	UUID         string    `xorm:"pk 'UUID'"`
	CourseId     string    `xorm:"'COURSE_ID'"`
	Content      string `xorm:"'CONTENT'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CREATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"deleted 'FROZEN_TIME'"`
}

func (self *CourseCommentTable)TableName() string {
	return "COURSE_COMMENT"
}

type CoursePointTable struct {
	UUID       string `xorm:"pk 'UUID'"`
	Point      int64 `xorm:"'POINT'"`
	CourseId   string `xorm:"'COURSE_ID'"`
	UserID     string `xorm:"'USER_ID'"`
	CreateTime time.Time `xorm:"created 'CREATE_TIME'"`
	UpdateTime time.Time `xorm:"updated 'UPDATE_TIME'"`
}

func (self *CoursePointTable) TableName() string {
	return "COURSE_POINT"
}
