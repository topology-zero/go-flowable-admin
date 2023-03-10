package form

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_definition"
)

// List 流程外置表单列表
func List(req *form.FlowFormListRequest, ctx *svc.ServiceContext) (resp form.FlowFormListResponse, err error) {
	param := form_definition.ListRequest{}
	param.Start = (req.Page - 1) * req.PageSize
	param.Size = req.PageSize
	param.Sort = "version"
	param.Order = "desc"
	param.Latest = true
	if len(req.Name) > 0 {
		param.Name = req.Name
	}
	list, count, err := form_definition.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取流程外置表单列表错误")
		return
	}

	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = count

	copier.Copy(&resp.Data, &list)

	return
}
