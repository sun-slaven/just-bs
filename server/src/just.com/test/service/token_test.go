package service
import (
	"testing"
	"just.com/service/token"
	"just.com/test"
)

func TestTokenMake(t *testing.T) {
	ts := service.TokenService{}
	ts.Log = Logger
	ts.Session = test.DataSource.NewSession()
	defer ts.Session.Close()

	_, makeErr := ts.Make("aa5eba0a-703c-4801-955b-1f44997738fe")
	if makeErr != nil {
		Logger.Println(makeErr)
		ts.Session.Rollback()
		t.Fail()
	}
	ts.Session.Commit()
}

func TestTokenCheck(t *testing.T) {
	ts := service.TokenService{}
	ts.Log = Logger
	ts.Session = test.DataSource.NewSession()
	defer ts.Session.Close()

	x := new(service.XToken)
	x.Id = "6688789c-1cb3-4303-9558-bcfd4c3b5d9e"
	x.UserId = "aa5eba0a-703c-4801-955b-1f44997738fe"
	b := ts.Check(x)
	Logger.Println(b)
}