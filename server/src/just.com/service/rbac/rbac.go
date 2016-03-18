package rbac
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"errors"
	"just.com/service"
)

var RBAC_LOAD_ERR = errors.New(service.SERVICE_RBAC_LOAD_ERR)

type RbacService struct {
	Session *xorm.Session
	Log     *log.Logger
	data    *Role
}

func (self *RbacService)GetData() *Role {
	return self.data
}

func (self *RbacService) Load(roleName string) error {
	// role
	role := NewRole()
	// 1. find the role
	rs := RoleService{}
	rs.Session = self.Session
	rs.Log = self.Log
	roleTable, roleErr := rs.FindByName(roleName)
	if roleErr != nil {
		self.Log.Println(roleErr)
		return RBAC_LOAD_ERR
	}
	// 2. find the permission
	sql := `SELECT * FROM "RBAC_ROLE_PERMISSION" WHERE "ROLE_ID" = ?`
	rolePerList := make([]table.RolePermissionTable, 0, 0)
	findErr := self.Session.Sql(sql, roleTable.UUID).Find(&rolePerList)
	if findErr != nil {
		self.Log.Println(findErr)
	}
	for _, rp := range rolePerList {
		sql := `SELECT * FROM "RBAC_PERMISSION" WHERE "UUID" = ?`
		p := new(table.PermissionTable)
		getFlag, getErr := self.Session.Sql(sql, rp.PermissionId).Get(p)
		if getFlag == false {
			if getErr != nil {
				self.Log.Println(getErr)
				return RBAC_LOAD_ERR
			}
		}
		permission := NewPermission(p.Target, p.Action)
		role.Assign(permission)
	}
	self.data = role
	return nil
}