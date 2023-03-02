package base

import (
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/admin/base"
)

// AdminUsers 获取所有用户
func AdminUsers(ctx *svc.ServiceContext) (resp base.AdminUsersResponse, err error) {
	list, err := query.AdminUserModel.Find()
	for i := range list {
		resp.List = append(resp.List, base.AdminUsers{
			Id:   list[i].ID,
			Name: list[i].Username,
		})
	}
	return
}
