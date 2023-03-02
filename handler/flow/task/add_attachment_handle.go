package task

import (
	"io"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// AddAttachmentHandle 添加任务附件
func AddAttachmentHandle(c *gin.Context) {
	var req taskType.AddAttachmentRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	var fileFs io.Reader

	if len(req.ExternalUrl) == 0 {
		file, err := c.FormFile("file")
		if err != nil {
			logrus.Errorf("%+v", errors.WithStack(err))
			response.HandleResponse(c, nil, errors.New("获取文件错误"))
			return
		}

		fileFs, err = file.Open()
		if err != nil {
			logrus.Errorf("%+v", errors.WithStack(err))
			response.HandleResponse(c, nil, errors.New("打开文件错误"))
			return
		}
	}

	err := task.AddAttachment(fileFs, &req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}
