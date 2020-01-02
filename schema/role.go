package schema

// Role 角色对象
type Role struct {
	RecordID string `json:"record_id" swaggo:"false,记录ID"`
	Name     string `json:"name" binding:"required" swaggo:"true,角色名称"`
	Sequence int    `json:"sequence" swaggo:"false,排序值"`
	Memo     string `json:"memo" swaggo:"false,备注"`
	Creator  string `json:"creator" swaggo:"false,创建者"`
	// Menus     RoleMenus `json:"menus" binding:"required,gt=0" swaggo:"false,菜单权限"`
}
