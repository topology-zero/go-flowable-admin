package process

import (
	"fmt"

	"github.com/MasterJoyHunan/flowablesdk/process_definition"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"
)

// List 流程列表
func List(req *process.ListProcessRequest, ctx *svc.ServiceContext) (resp process.ListProcessResponse, err error) {
	listReq := process_definition.ListRequest{}
	listReq.Name = req.Name
	listReq.Latest = true
	listReq.Sort = "name"
	listReq.Start = req.Page - 1
	listReq.Size = req.PageSize

	list, count, err := process_definition.List(listReq)

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取列表失败")
		return
	}

	resp.Total = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	_ = copier.Copy(&resp.Data, &list)
	for i := range list {
		resp.Data[i].Version = fmt.Sprintf("v%d", list[i].Version)
	}
	return
}
