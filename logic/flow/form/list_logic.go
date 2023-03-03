package form

import (
	"github.com/MasterJoyHunan/flowablesdk/external_form/form_definition"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// List 流程外置表单列表
func List(req *form.FlowFormListRequest, ctx *svc.ServiceContext) (resp form.FlowFormListResponse, err error) {
	param := form_definition.ListRequest{}
	param.Start = (req.Page - 1) * req.PageSize
	param.Size = req.PageSize
	param.Sort = "version"
	param.Order = "desc"
	param.Latest = true
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
