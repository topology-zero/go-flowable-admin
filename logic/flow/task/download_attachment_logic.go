package task

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/task/task_attachment"
)

// DownloadAttachment 下载任务附件
func DownloadAttachment(req *task.DownloadAttachmentRequest, ctx *svc.ServiceContext) (data []byte, err error) {
	data, err = task_attachment.Detail(req.TaskId, req.AttachmentId)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取附件错误")
	}
	return
}
