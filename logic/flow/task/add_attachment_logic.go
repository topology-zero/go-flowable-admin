package task

import (
	"io"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk/attachment"
	"github.com/MasterJoyHunan/flowablesdk/task/task_attachment"
	"github.com/pkg/errors"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// AddAttachment 添加任务附件
func AddAttachment(fs io.Reader, req *task.AddAttachmentRequest, ctx *svc.ServiceContext) (err error) {
	userInfo, _ := ctx.GinContext.Get("userInfo")
	claims := userInfo.(*jwt.Claims)

	if len(req.ExternalUrl) > 0 {
		_, err = task_attachment.AddUrl(req.Id, attachment.AddAttachment{
			UserId:      strconv.Itoa(claims.Id),
			Name:        req.Name,
			Description: req.Description,
			ExternalUrl: req.ExternalUrl,
		})
	} else {
		_, err = task_attachment.Add(req.Id, attachment.AddAttachment{
			UserId:      strconv.Itoa(claims.Id),
			Name:        req.Name,
			Description: req.Description,
			File:        fs,
		})
	}

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("上传附件失败")
	}
	return
}
