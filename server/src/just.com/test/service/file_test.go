package service
import (
	"testing"
	"just.com/service/file"
	"just.com/test"
)


func TestFileAdd(t *testing.T) {
	fs := service.FileService{}
	fs.Log = Logger
	session := test.DataSource.NewSession()
	fs.Session = session
	defer fs.Session.Close()
	_, addErr := fs.Add("35bfcb48-9f2a-4938-bb54-cab72e951970", "123", "123124","aa5eba0a-703c-4801-955b-1f44997738fe")
	if addErr != nil {
		t.Fail()
	}
}

func TestFileUpdate(t *testing.T) {
	fs := service.FileService{}
	fs.Log = Logger
	session := test.DataSource.NewSession()
	fs.Session = session
	updateErr := fs.Update("aa5eba0a-703c-4801-955b-1f44997738fe", "d73df59c-13b6-4ed0-86f4-f3eddf2155bb", "name", "url")
	if updateErr != nil {
		Logger.Println(updateErr)
		t.Fail()
	}
}

func TestFileDelete(t *testing.T) {
	fs := service.FileService{}
	fs.Log = Logger
	fs.Session = test.DataSource.NewSession()
	defer fs.Session.Close()
	deleteErr := fs.Delete("d73df59c-13b6-4ed0-86f4-f3eddf2155bb", "aa5eba0a-703c-4801-955b-1f44997738fe")
	if deleteErr != nil {
		Logger.Println(deleteErr)
		t.Fail()
	}
}
