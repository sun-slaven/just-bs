package err
import "errors"

var (
	COURSE_NOT_FOUND = errors.New("课程id有误,请确认该课程id")
)
