package table
import "time"

// 学院
type CollegeTable struct {
	UUID         string        `xorm:"pk 'UUID'"`
	Name         string        `xorm:"'NAME'"`
	CreateTime   time.Time `xorm:"'CREATE_TIME'"`
	FrozenStatus string        `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time        `xorm:"'FROZEN_TIME'"`
}

func (self *CollegeTable) TableName() string {
	return "COLLEGE"
}

// 专业
type MajorTable struct {
	UUID         string        `xorm:"pk 'UUID'"`
	Name         string        `xorm:"'NAME'"`
	CollegeId    string `xorm:"'COLLEGE_ID'"`
	CreateTime   time.Time `xorm:"'CREATE_TIME'"`
	FrozenStatus string        `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time        `xorm:"'FROZEN_TIME'"`
}

func (self *MajorTable) TableName() string {
	return "MAJOR"
}
