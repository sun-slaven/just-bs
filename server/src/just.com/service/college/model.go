package college
import (
	"github.com/go-xorm/xorm"
	"log"
)

type CollegeService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func NewCollegeService(session *xorm.Session, log *log.Logger) {
	return &CollegeService{session, log}
}
