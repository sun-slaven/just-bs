package user
import (
	"github.com/go-xorm/xorm"
	"log"
)

type UserService struct {
	Session *xorm.Session
	Log     *log.Logger
}