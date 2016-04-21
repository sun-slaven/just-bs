package rbac
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"errors"
	"just.com/service"
)

var (
	ROLE_ADD_ERR = errors.New(service.SERVICE_RBAC_ROLE_ADD_ERR)
	ROLE_FIND_BY_NAME_ERR = errors.New(service.SERVICE_RBAC_ROLE_ADD_ERR)
	ROLE_ADD_PERMISSION__ERR = errors.New(service.SERVICE_RBAC_PERMISSION_ADD_ERR)
	PERMISSION_ADD_ERR = errors.New(service.SERVICE_RBAC_PERMISSION_ADD_ERR)
)

type PermissonService struct {
	Session *xorm.Session
	Log     *log.Logger
}

func(self *PermissonService)  Add(target,action string) (string,error) {
	p := new(table.PermissionTable)
	p.UUID = uuid.New()
	p.Target = target
	p.Action = action
	p.CreateTime = time.Now()
	p.UpdateTime = time.Now()
	insertNum,insertErr:= self.Session.InsertOne(p)
	if insertNum == 0 {
		if insertErr !=nil{
			self.Log.Println(insertErr)
		}
		return "",PERMISSION_ADD_ERR
	}
	return p.UUID,nil
}

type RoleService struct {
	Session *xorm.Session
	Log     *log.Logger
}
func (self *RoleService) Add(name, desc string) (string, error) {
	roleTable := table.RoleTable{}
	roleTable.UUID = uuid.New()
	roleTable.Name = name
	roleTable.Desc = desc
	roleTable.CreateTime = time.Now()
	roleTable.UpdateTime = time.Now()
	insertNum, insertErr := self.Session.InsertOne(&roleTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.Log.Println(insertErr)
		}
		self.Log.Println(ROLE_ADD_ERR)
		return "", ROLE_ADD_ERR
	}
	return roleTable.UUID, nil
}
func (self *RoleService) FindByName(name string) (*table.RoleTable,error) {
	roleTable := new(table.RoleTable)
	roleTable.Name = name
	getFlag,getErr:= self.Session.Get(roleTable)
	if getFlag == false{
		if getErr !=nil{
			self.Log.Println(getErr)
		}
		return nil,ROLE_FIND_BY_NAME_ERR
	}
	return roleTable,nil
}

func (self *RoleService) AddPermission(roleName,pTarget,pAction string) error {

	// get role
	roleTable,roleErr:= self.FindByName(roleName)
	if roleErr!=nil{
		self.Log.Println(roleErr)
		return ROLE_ADD_PERMISSION__ERR
	}
	// get permission
	p := new(table.PermissionTable)
	p.Target = pTarget
	p.Action = pAction
	getFlag,getErr:= self.Session.Get(p)
	if getFlag == false{
		if getErr != nil{
			self.Log.Println(getErr)
		}
		return ROLE_ADD_PERMISSION__ERR
	}
	// add
	rp := new(table.RolePermissionTable)
	rp.UUID = uuid.New()
	rp.RoleId = roleTable.UUID
	rp.PermissionId = p.UUID
	insertNum,insertErr:= self.Session.InsertOne(rp)
	if insertNum == 0{
		if insertErr != nil{
			self.Log.Println(insertErr)
		}
		return ROLE_ADD_PERMISSION__ERR
	}
	return nil
}