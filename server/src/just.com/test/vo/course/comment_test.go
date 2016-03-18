package course_test
import (
	"testing"
	"just.com/query/vo/course"
	"just.com/test"
	"encoding/json"
)

func TestCommentLoad(t *testing.T)  {
	session := test.DataSource.NewSession()
	defer session.Close()
	commentVoList:= course.LoadCommentVoList("35bfcb48-9f2a-4938-bb54-cab72e951970",session,test.Log)
	v,err:= json.Marshal(commentVoList)
	if err!=nil{
		test.Log.Println(err)
		t.Fail()
	}
	test.Log.Println(string(v))
}