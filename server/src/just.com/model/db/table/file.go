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

/**
用来存储资源的表
 */
type FileTable struct {
	UUID       string `xorm:"pk 'UUID'"`
	Name       string `xorm:"'NAME'"`
	Url        string `xorm:"'URL'"`
	Type       string `xorm:"'TYPE'"`
	CreateTime time.Time `xorm:"created 'CREATE_TIME'"`
}

func (self *FileTable) TableName() string {
	return "FILE"
}