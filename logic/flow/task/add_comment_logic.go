package task

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/comment"
	"github.com/topology-zero/flowablesdk/task/task_comment"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// AddComment 添加备注
func AddComment(req *task.AddCommentRequest, ctx *svc.ServiceContext) error {
	userInfo, _ := ctx.GinContext.Get("userInfo")
	claims := userInfo.(*jwt.Claims)

	_, err := task_comment.Add(req.TaskId, comment.AddComment{
		UserId:  strconv.Itoa(claims.Id),
		Message: req.Message,
	})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("添加备注失败")
	}
	return err
}
