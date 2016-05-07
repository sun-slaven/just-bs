package file

import (
	"github.com/go-xorm/xorm"
	"log"
)

type FileService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func NewFileService(session *xorm.Session, log *log.Logger) *FileService {
	return &FileService{
		Session:session, Log:log,
	}
}

