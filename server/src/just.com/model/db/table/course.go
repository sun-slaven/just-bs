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
	VideoUrl     string `xorm:"'VIDEO_URL'"`
}

func (self *CourseTable)TableName() string {
	return "COURSE"
}

type CourseChapterTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	CourseId     string `xorm:"'COURSE_ID'"`
	Name         string `xorm:"'NAME'"`
	Content      string `xorm:"'CONTENT'"`
	Order        int64       `xorm:"'ORDER'"`
	VideoName    string `xorm:"'VIDEO_NAME'"`
	VideoUrl     string `xorm:"'VIDEO_URL'"`
	CreateUser   string `xorm:"'CREATE_USER'"`
	CreateTime   time.Time `xorm:"'CREATE_TIME'"`
	UpdateTime   time.Time `xorm:"'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenTime   time.Time `xorm:"'FROZEN_TIME'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
}

func (self *CourseChapterTable)TableName() string {
	return "COURSE_CHAPTER"
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

type CourseAttachmentTable struct {
	UUID         string        `xorm:"pk 'UUID'"`
	Name         string        `xorm:"'NAME'"`
	Url          string        `xorm:"'URL'"`
	CourseId     string        `xorm:"'COURSE_ID'"`
	CreateTime   time.Time        `xorm:"'CREATE_TIME'"`
	CreateUser   string        `xorm:"'CREATE_USER'"`
	UpdateTime   time.Time `xorm:"'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenTime   time.Time        `xorm:"'FROZEN_TIME'"`
	FrozenStatus string        `xorm:"FROZEN_STATUS"`
}

func (self *CourseAttachmentTable) TableName() string {
	return "COURSE_ATTACHMENT"
}


type CourseStudyRecordTable struct {
	UUID         string        `xorm:"pk 'UUID'"`
	VideoUrl     string        `xorm:"'VIDEO_URL'"`
	CourseId     string        `xorm:"'COURSE_ID'"`
	UserId       string `xorm:"'USER_ID'"`
	Process      float64        `xorm:"'PROCESS'"`
	CreateTime   time.Time        `xorm:"'CREATE_TIME'"`
	CreateUser   string        `xorm:"'CREATE_USER'"`
	UpdateTime   time.Time `xorm:"'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenTime   time.Time        `xorm:"'FROZEN_TIME'"`
	FrozenStatus string        `xorm:"FROZEN_STATUS"`
}

func (self *CourseStudyRecordTable) TableName() string {
	return "COURSE_STUDY_RECORD"
}

