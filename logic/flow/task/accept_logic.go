package task

import (
	"fmt"
	"strconv"

	"go-flow-admin/logic/common"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"

	"github.com/pkg/errors"
	processTask "github.com/topology-zero/flowablesdk/task"
	"github.com/topology-zero/flowablesdk/variable"
)

// Accept 审批同意
func Accept(req *task.TaskAcceptRequest, ctx *svc.ServiceContext) error {
	user, _ := ctx.GinContext.Get("userInfo")
	claims := user.(*jwt.Claims)

	adminUser := common.GetAdminUser()
	result := fmt.Sprintf("[ %s ]审批同意", adminUser[strconv.Itoa(claims.Id)].Realname)
	if len(req.Message) > 0 {
		result += fmt.Sprintf(",审批意见 : %s", req.Message)
	}

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
				Value: "同意",
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
				Value: true,
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
