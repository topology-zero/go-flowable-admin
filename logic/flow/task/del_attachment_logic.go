package task

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/task/task_attachment"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// DelAttachment 删除任务附件
func DelAttachment(req *task.DelAttachmentRequest, ctx *svc.ServiceContext) error {
	// todo 1.需要权限判断,只能删除自己的
	//      2.已完成的操作不能删除
	userInfo, _ := ctx.GinContext.Get("userInfo")
	claims := userInfo.(*jwt.Claims)

	err := task_attachment.Delete(req.TaskId, req.FileId, strconv.Itoa(claims.Id))
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("删除附件错误")
	}
	return err
}
