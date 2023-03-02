// Code generated by goctl. DO NOT EDIT.
package base

type BaseRoleResponse struct {
	Data []BaseRole `json:"data"`
}

type BaseRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BaseAuthResponse struct {
	Tree []BaseAuth `json:"tree"` // 节点树
}

type BaseAuth struct {
	Id    int        `json:"id"`
	Pid   int        `json:"pid"`
	Name  string     `json:"name"`  // 节点名
	Child []BaseAuth `json:"child"` // 子节点
}

type UserInfoResponse struct {
	Id       int    `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
	Rolename string `json:"rolename"` // 角色名
	Authkeys string `json:"authkeys"` // 角色权限KEY
}

type ChangeSelfPwdRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required,min=6,max=255" label:"老密码"`                     // 老密码
	NewPassword     string `json:"newPassword" binding:"required,nefield=OldPassword,min=6,max=255" label:"新密码"` // 新密码
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=NewPassword" label:"确认密码"`          // 确认密码
}

type AdminUsersResponse struct {
	List []AdminUsers `json:"list"` // 后台用户列表
}

type AdminUsers struct {
	Id   int    `json:"id"`   // 后台用户ID
	Name string `json:"name"` // 用户ID
}
