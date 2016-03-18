package course
import (
	"testing"
	"just.com/query/vo/course"
	"just.com/test"
	"just.com/model/db/table"
	"encoding/json"
)

func TestCourseLoad(t *testing.T) {
	session := test.DataSource.NewSession()
	courseTable := new(table.CourseTable)
	_, getErr := session.Id("35bfcb48-9f2a-4938-bb54-cab72e951970").Get(courseTable)
	if getErr != nil {
		test.Log.Println(getErr)
		t.Fail()
	}
	cv := course.NewCourseVo(courseTable)
	cv.LoadPointStatus("aa5eba0a-703c-4801-955b-1f44997738fe", session, test.Log)
	vo, voErr := json.Marshal(cv)
	if voErr != nil{
		test.Log.Println(voErr)
		t.Fail()
	}
	test.Log.Println(string(vo))
}