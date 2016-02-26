package model
import "time"

type CourseTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	Name         string `xorm:"'NAME'"`
	Introduction string `xorm:"'INTRODUCTION'"`
	Syllabus     string `xorm:"'SYLLABUS'"`
	Plan         string `xorm:"'PLAN'"`
	Experiment   string `xorm:"'EXPERIMENT'"`
	MarkSum      int64 `xorm:"'MARK_SUM'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CRATE_USER'"`
	UpdateTime   time.Time `xorm:"updated 'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"deleted 'FROZEM_TIME'"`
	Version      int64 `xorm:"version 'VERSION'"`
	Major        string    `xorm:"'MAJOR'"`
	College      string    `xorm:"'COLLEGE'"`
	Points       int64    `xorm:"'POINTS'"`
}

type CourseMarkTable struct {
	UUID         string    `xorm:"pk 'UUID'"`
	UserId       string    `xorm:"'USER_ID'"`
	CourseId     string    `xorm:"'COURSE_ID'"`
	CreateTime   time.Time    `xorm:"created 'CREATE_TIME'"`
	FrozenStatus string    `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time    `xorm:"deleted 'FROZEN_TIME'"`
}

type Comment struct {
	UUID         string    `xorm:"pk 'UUID'"`
	Content      string `xorm:"'CONTENT'"`
	CourseId     string    `xorm:"'COURSE_ID'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CRATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"deleted 'FROZEM_TIME'"`
}