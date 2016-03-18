package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/service"
	"errors"
)

var (
	COURSE_ADD_ERR = errors.New(service.SERVICE_COURSE_ADD_ERR)
	COURSE_UPDATE_ERR = errors.New(service.SERVICE_COURSE_UPDATE_ERR)
	COURSE_DELETE_ERR = errors.New(service.SERVICE_COURSE_DELETE_ERR)
	COURSE_MARK_ERR = errors.New(service.SERVICE_COURSE_MARK_ERR)
	COURSE_MARK_CANCEL_ERR = errors.New(service.SERVICE_COURSE_MARK_CANCEL_ERR)
	COURSE_FLUSH_MARK_NUM_ERR = errors.New(service.SERVICE_COURSE_FLUSH_MARK_NUM_ERR)
	COURSE_FLUSH_COMMENT_NUM_ERR = errors.New(service.SERVICE_COURSE_FLUSH_COMMENT_NUM_ERR)
	COURSE_FLUSH_POINT_ERR = errors.New(service.SERVICE_COURSE_FLUSH_POINT_ERR)
	COURSE_COMMENT_ADD_ERR = errors.New(service.SERVICE_COURSE_COMMENT_ADD_ERR)
	COURSE_COMMENT_DELETE_ERR = errors.New(service.SERVICE_COURSE_COMMENT_DELETE_ERR)
	COURSE_POINT_ADD_ERR = errors.New(service.SERVICE_COURSE_POINT_ADD_ERR)
	COURSE_POINT_UPDATE_ERR = errors.New(service.SERVICE_COURSE_POINT_UPDATE_ERR)
)

type CourseService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func NewCourseService(session *xorm.Session, log *log.Logger) *CourseService {
	cs := new(CourseService)
	cs.Session = session
	cs.Log = log
	return cs
}