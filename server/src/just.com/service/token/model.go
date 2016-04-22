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

func NewTokenService(session *xorm.Session, log *log.Logger) *TokenService {
	ts := new(TokenService)
	ts.Session = session
	ts.Log = log
	return ts
}

/*token*/
type XToken  struct {
	Id     string
	UserId string
}

func NewXToken(id, userId string) *XToken {
	token := new(XToken)
	token.Id = id
	token.UserId = userId
	return token
}