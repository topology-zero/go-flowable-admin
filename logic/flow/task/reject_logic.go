package task

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	processTask "github.com/topology-zero/flowablesdk/task"
	"github.com/topology-zero/flowablesdk/variable"
	"go-flow-admin/logic/common"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// Reject 审批驳回
func Reject(req *task.TaskRejectRequest, ctx *svc.ServiceContext) error {
	user, _ := ctx.GinContext.Get("userInfo")
	claims := user.(*jwt.Claims)

	adminUser := common.GetAdminUser()
	result := fmt.Sprintf("[ %s ]审批驳回,驳回意见 : %s", adminUser[strconv.Itoa(claims.Id)].Realname, req.Message)

	err := processTask.Action(req.Id, processTask.ActionRequest{
		Action:   "complete",
		Assignee: strconv.Itoa(claims.Id),
		Outcome:  result,
		Variables: []variable.VariableRequest{
			{
				Name:  req.Id + "|operate_user",
				Type:  "string",
				Value: adminUser[strconv.Itoa(claims.Id)].Realname,
			},
			{
				Name:  req.Id + "|operate_type",
				Type:  "string",
				Value: "驳回",
			},
			{
				Name:  req.Id + "|operate_memo",
				Type:  "string",
				Value: req.Message,
			},
		},
		TransientVariables: []variable.VariableRequest{
			{
				Name:  "pass",
				Type:  "boolean",
				Value: false,
			},
		},
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("完成任务错误")
		return err
	}
	return nil
}
