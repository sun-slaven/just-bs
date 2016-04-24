package user
import (
	"github.com/go-xorm/xorm"
	"log"
)


type UserService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func NewUserService(session *xorm.Session, log *log.Logger) *UserService {
	return &UserService{Session:session, Log:log}
}