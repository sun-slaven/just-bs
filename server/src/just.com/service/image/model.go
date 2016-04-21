package image
import (
	"github.com/go-xorm/xorm"
	"log"
)

type ImageService struct {
	Session *xorm.Session
	Log     *log.Logger
}