package table
import "time"

/*权限*/
type PermissionTable struct {
	UUID       string `xorm:"pk 'UUID'"`
	Target string   `xorm:"'TARGET'"`
	Action     string	 `xorm:"'ACTION'"`
	CreateTime time.Time	`xorm:"created 'CREATE_TIME'"`
	UpdateTime time.Time	`xorm:"updated 'UPDATE_TIME'"`
}
func (self *PermissionTable)TableName() string {
	return "RBAC_PERMISSION"
}

/*角色，权限的集合*/
type RoleTable struct {
	UUID       string	`xorm:"pk 'UUID'"`
	Name       string `xorm:"'NAME'"`
	Desc       string `xorm:"'DESC'"`
	CreateTime time.Time  `xorm:"created 'CREATE_TIME'"`
	UpdateTime time.Time  `xorm:"updated 'UPDATE_TIME'"`
}
func (self *RoleTable)TableName() string {
	return "RBAC_ROLE"
}

/*角色和权限的映射表*/
type RolePermissionTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	RoleId       string	`xorm:"'ROLE_ID'"`
	PermissionId string	`xorm:"'PERMISSION_ID'"`
}
func (self *RolePermissionTable) TableName()string{
	return "RBAC_ROLE_PERMISSION"
}