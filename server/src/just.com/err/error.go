package err
import (
	"errors"
	"net/http"
)

type HttpError struct {
	Status int
	Error  error
}

func NewHttpErrorWithStatus(status int, error error) *HttpError {
	return &HttpError{
		Status:status,
		Error:error,
	}
}

func NewDefaultHttpError(error error) *HttpError {
	return &HttpError{
		Status:http.StatusBadRequest,
		Error:error,
	}
}

var (
	NO_COURSE_FOUND = NewDefaultHttpError(errors.New("课程id有误,请确认该课程id"))
	NO_REQUIERED_PARAM_FOUND = NewDefaultHttpError(errors.New("缺少必须的参数"))
	PARAM_ERR = NewDefaultHttpError(errors.New("请求参数有误"))
	NO_COURSE_LIST_FOUND = NewDefaultHttpError(errors.New("课程列表获取有误,请核对过滤信息"))
	STATUS_UNAUTHORIZED = NewHttpErrorWithStatus(http.StatusUnauthorized, errors.New("请登录后重试"))
	NO_CONTEXT = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("没有上下文执行环境"))

	NO_IMAGE_FOUND_BY_URL = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("没有找到该url对应的图片"))
	NO_COLLEGE_OR_MAJOR_FOUND = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("学院号或专业号有误"))
	NO_TEACHER_ID_FOUND = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("教师号有误"))
	COURSE_INSERT_ERR = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("创建课程失败"))
)
