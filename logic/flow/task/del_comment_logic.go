package task

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/task/task_comment"
)

// DelComment 删除备注
func DelComment(req *task.DelCommentRequest, ctx *svc.ServiceContext) error {
	// todo 1.需要权限判断,只能删除自己的
	//      2.已完成的操作不能删除

	err := task_comment.Delete(req.TaskId, req.CommentId)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("删除备注错误")
	}
	return err
}
