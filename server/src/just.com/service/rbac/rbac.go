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

func NewRbacService(session *xorm.Session, log *log.Logger) *RbacService {
	return &RbacService{
		Session:session,
		Log:log,
	}
}

func (self *RbacService)GetData() *Role {
	return self.data
}

func (self *RbacService) Load(roleName string) error {
	// role
	role := NewRole()
	// 1. find the role
	rs := NewRoleService(self.Session, self.Log)
	roleTable, roleErr := rs.FindByName(roleName)
	if roleErr != nil {
		self.Log.Println(roleErr)
		return RBAC_LOAD_ERR
	}
	// 2. find the permission
	rolePerList := make([]table.RolePermissionTable, 0, 0)
	findErr := self.Session.Find(&rolePerList, &table.RolePermissionTable{RoleId:roleTable.UUID})
	if findErr != nil {
		self.Log.Println(findErr)
		return findErr
	}
	for _, mapping := range rolePerList {
		permissionTable := new(table.PermissionTable)
		getFlag, getErr := self.Session.Id(mapping.PermissionId).Get(permissionTable)
		if getFlag == false {
			if getErr != nil {
				self.Log.Println(getErr)
				return RBAC_LOAD_ERR
			}
		}
		permission := NewPermission(permissionTable.Target, permissionTable.Action)
		role.Assign(permission)
	}
	self.data = role
	return nil
}