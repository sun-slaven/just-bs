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
	NO_COURSE_ID_FOUND = NewDefaultHttpError(errors.New("缺少课程id"))
	NO_REQUIRED_PARAM_FOUND = NewDefaultHttpError(errors.New("缺少必须的参数"))
	PARAM_ERR = NewDefaultHttpError(errors.New("请求参数有误"))
	NO_COURSE_LIST_FOUND = NewDefaultHttpError(errors.New("课程列表获取有误,请核对过滤信息"))
	STATUS_UNAUTHORIZED = NewHttpErrorWithStatus(http.StatusUnauthorized, errors.New("请登录后重试"))
	NO_CONTEXT = NewHttpErrorWithStatus(http.StatusInternalServerError, errors.New("没有上下文执行环境"))

	NO_IMAGE_FOUND_BY_URL = NewDefaultHttpError(errors.New("没有找到该url对应的图片"))
	NO_COLLEGE_OR_MAJOR_FOUND = NewDefaultHttpError(errors.New("学院号或专业号有误"))
	NO_TEACHER_ID_FOUND = NewDefaultHttpError(errors.New("教师号有误"))
	COURSE_INSERT_ERR = NewDefaultHttpError(errors.New("创建课程失败"))
	COURSE_LIST_FIND_ERR = NewDefaultHttpError(errors.New("课程列表获取失败"))
	COURSE_MARKED_LIST_FIND_ERR = NewDefaultHttpError(errors.New("课程列表获取失败"))
	COURSE_DELETE_ERR = NewDefaultHttpError(errors.New("删除课程失败,id无效"))
	COURSE_UPDATE_ERR = NewDefaultHttpError(errors.New("更新课程失败,id无效"))
	COURSE_COMMENT_INSERT_ERR = NewDefaultHttpError(errors.New("创建课程评论失败"))
	COURSE_COMMENT_DELETE_ERR = NewDefaultHttpError(errors.New("删除评论失败"))
	COURSE_MARK_ERR = NewDefaultHttpError(errors.New("关注课程失败,用户id或课程id无效"))
	COURSE_MARK_CANCEL_ERR = NewDefaultHttpError(errors.New("取消课程关注,用户id或课程id无效"))

	USER_PASSWORD_OR_EMAIL_ERR = NewDefaultHttpError(errors.New("邮箱或密码错误"))
	USER_REGISTER_ERR = NewDefaultHttpError(errors.New("用户注册失败"))
	USER_REGISTER_EMIAL_ERR = NewDefaultHttpError(errors.New("该邮箱已被注册"))
	NOT_USER_ID_FOUND = NewDefaultHttpError(errors.New("缺少用户id"))
	USER_ID_ERR = NewDefaultHttpError(errors.New("用户id有误,找不到该用户"))
	USER_RESET_PASSWORD_ERR = NewDefaultHttpError(errors.New("用户重置密码错误"))
	USER_FROZEN_ERR = NewDefaultHttpError(errors.New("用户禁用失败"))
	TOKEN_CREATE_ERR = NewDefaultHttpError(errors.New("登录token创建失败"))

	NO_CHAPTER_FOUND = NewDefaultHttpError(errors.New("章节获取失败"))
	NO_ATTACHMENT_FOUND = NewDefaultHttpError(errors.New("附件获取失败"))
	CHAPTER_ADD_ERR = NewDefaultHttpError(errors.New("章节新增失败"))
	CHAPTER_FORMAT_ERR = NewDefaultHttpError(errors.New("章节请求格式错误"))
	CHAPTER_UPDATE_ERR = NewDefaultHttpError(errors.New("章节更新失败,可能是章节id无效"))

	USER_LIST_ERR = NewDefaultHttpError(errors.New("用户列表获取失败"))
// file
	FILE_ADD_ERR = NewDefaultHttpError(errors.New("文件添加失败"))
	TOKEN_DELETE_ERR = NewDefaultHttpError(errors.New("登出失败"))
)
