package service
import (
	"log"
	"os"
	"testing"
	"just.com/service/rbac"
	"just.com/test"
)

var Logger = log.New(os.Stdout, "mlearing", log.Llongfile)

func TestRoleAdd(t *testing.T) {
	rs := rbac.RoleService{}
	rs.Session = test.DataSource.NewSession()
	rs.Log = Logger
	defer rs.Session.Close()
	roleId, addErr := rs.Add("ADMIN", "管理员")
	if addErr != nil {
		rs.Session.Rollback()
		t.Fail()
	}
	rs.Session.Commit()
	rs.Log.Println(roleId)
}

func TestPermissionAdd(t *testing.T)  {
	ps := rbac.PermissonService{}
	ps.Log = Logger
	ps.Session = test.DataSource.NewSession()
	defer ps.Session.Close()
	pId,pErr:= ps.Add("FILE","C")
	ps.Add("FILE","R")
	ps.Add("FILE","U")
	ps.Add("FILE","D")
	if pErr != nil{
		Logger.Println(pErr)
		ps.Session.Rollback()
		t.Fail()
	}
	ps.Session.Commit()
	Logger.Println(pId)
}

func TestRoleAddPermission(t *testing.T)  {
	rs := rbac.RoleService{}
	rs.Log = Logger
	rs.Session = test.DataSource.NewSession()
	defer rs.Session.Close()

	addErr:= rs.AddPermission("STUDENT","FILE","C")
	rs.AddPermission("STUDENT","FILE","R")
	rs.AddPermission("STUDENT","FILE","U")
	rs.AddPermission("STUDENT","FILE","D")
	if addErr != nil{
		Logger.Println(addErr)
		rs.Session.Rollback()
		t.Fail()
	}
	rs.Session.Commit()
}

func TestRoleService(t *testing.T)  {
	rs := rbac.RbacService{}
	rs.Log = Logger
	rs.Session = test.DataSource.NewSession()
	defer rs.Session.Close()
	loadErr:= rs.Load("STUDENT")
	if loadErr != nil{
		Logger.Println(loadErr)
		t.Fail()
	}
}