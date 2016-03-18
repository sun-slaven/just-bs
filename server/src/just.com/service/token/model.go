package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"errors"
	"just.com/service"
)

var (
	TOKEN_MAKE_ERR = errors.New(service.SERVICE_TOKEN_MAKE_ERR)
	TOKEN_CHECK_ERR = errors.New(service.SERVICE_TOKEN_CHECK_ERR)
)

type TokenService  struct {
	Session *xorm.Session
	Log     *log.Logger
}
/*token*/
type XToken  struct {
	Id     string
	UserId string
}