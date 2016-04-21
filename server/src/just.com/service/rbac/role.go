package rbac

/*per 保存的是Permission对应的name*/
type Role struct {
	permissions map[string]*Permission
}
func NewRole() *Role {
	r := new(Role)
	r.permissions = make(map[string]*Permission)
	return r
}
/*是否被授权*/
func (self *Role) IsGranted(p *Permission) bool {
	if self.permissions[p.Name()] == nil {
		return true
	}
	return false
}
/*委托权限*/
func (self *Role) Assign(p *Permission) {
	self.permissions[p.Name()] = p
}
/*取消授权*/
func (self *Role)Revoke(p *Permission) {
	delete(self.permissions, p.Name())
}

type Permission struct {
	Target string
	Action string
}

func (self *Permission) Name() string {
	return self.Target + "_" + self.Action
}
func NewPermission(target, action string) *Permission {
	p := new(Permission)
	p.Target = target
	p.Action = action
	return p
}

