package role

import (
	"strconv"

	"go-flow-admin/model"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/admin/role"

	"github.com/pkg/errors"
)

// Del 删除角色
func Del(req *role.RoleDeleteRequest, ctx *svc.ServiceContext) error {
	if req.Id == AdminRoleId {
		return errors.New("无法修改超级管理员角色")
	}

	userModel := query.AdminUserModel
	user, _ := userModel.Where(userModel.RoleID.Eq(req.Id)).First()
	if user != nil {
		return errors.New("当前角色正在使用,无法删除")
	}

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.Where(tx.AdminRoleModel.ID.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}

		_, err = tx.AdminCasbinRuleModel.Where(tx.AdminCasbinRuleModel.Ptype.Eq("p"), tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.Id))).Delete()
		return err
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
