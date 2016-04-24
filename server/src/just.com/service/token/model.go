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
type UserToken  struct {
	Id     string `json:"id"`
	UserId string        `json:"user_id"`
}

func NewUserToken(id, userId string) *UserToken {
	token := new(UserToken)
	token.Id = id
	token.UserId = userId
	return token
}

func NewTokenService(session *xorm.Session, log *log.Logger) *TokenService {
	return &TokenService{
		Session:session,
		Log:log,
	}
}