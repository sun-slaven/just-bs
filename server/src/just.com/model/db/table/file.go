package table
import "time"

type ImageTable struct {
	UUID       string    `xorm:"pk 'UUID'"`
	Name       string    `xorm:"'NAME'"`
	Url        string    `xorm:"'URL'"`
	Width      int64    `xorm:"'WIDTH'"`
	Height     int64    `xorm:"'HEIGHT'"`
	CreateTime time.Time    `xorm:"created 'CREATE_TIME'"`
	CreateUser string `xorm:"'CREATE_USER'"`
}

func (self *ImageTable)TableName() string {
	return "IMAGE"
}

type CourseFileTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	Name         string `xorm:"'NAME'"`
	Url          string `xorm:"'URL'"`
	Type         string `xorm:"'TYPE'"`
	CourseID     string `xorm:"'COURSE_ID'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CREATE_USER'"`
	UpdateTime   time.Time `xorm:"updated 'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"'FROZEN_TIME'"`
}

func (self *CourseFileTable) TableName() string {
	return "COURSE_FILE"
}