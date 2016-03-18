package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"errors"
	"just.com/service"
)

var (
	FILE_ADD_ERR = errors.New(service.SERVICE_FILE_ADD_ERR)
	FILE_UPDATE_ERR = errors.New(service.SERVICE_FILE_UPDATE_ERR)
	FILE_DELETE_ERR = errors.New(service.SERVICE_FILE_DELETE_ERR)
)

type FileService struct {
	Session *xorm.Session
	Log     *log.Logger
}

