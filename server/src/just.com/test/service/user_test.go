package service_test
import (
	"testing"
	"just.com/service/user"
	"just.com/test"
	"log"
	"os"
)

func TestUserAdd(t *testing.T)  {
	userService := user.UserService{}
	userService.Session = test.DataSource.NewSession()
	userService.Log = log.New(os.Stdout, "mlearing", log.Llongfile)
	userId,addErr:= userService.Add("13918503047","www","1bf178e0-3402-475b-a4b2-3740c286bb3d","STUDENT")
	if addErr != nil{
		t.Fail()
	}
	userService.Session.Commit()
	log.Println(userId)
}
