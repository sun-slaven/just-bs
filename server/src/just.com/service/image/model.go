package image
import (
	"github.com/go-xorm/xorm"
	"log"
)

type ImageService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func NewImageService(session *xorm.Session, log *log.Logger) *ImageService {
	return &ImageService{Session: session, Log:log}
}