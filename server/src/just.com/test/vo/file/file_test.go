package file
import (
	"testing"
	"just.com/test"
	"just.com/query/vo/file"
	"encoding/json"
)

func TestFileList(t *testing.T)  {
	session := test.DataSource.NewSession()
	defer session.Close()
	fileVoList:= file.LoadFileVoList("35bfcb48-9f2a-4938-bb54-cab72e951970",session,test.Log)
	v,err:= json.Marshal(fileVoList)
	if err!=nil{
		test.Log.Println(err)
		t.Fail()
	}
	test.Log.Println(string(v))
}