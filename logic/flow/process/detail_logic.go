package process

import (
	"github.com/MasterJoyHunan/flowablesdk/deployment"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"
)

// Detail 流程详情
func Detail(req *process.DetailProcessRequest, ctx *svc.ServiceContext) (resp process.DetailProcessResponse, err error) {
	detail, err := deployment.Detail(req.Id)

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取详情失败")
		return
	}

	xml, err := deployment.ResourceContent(req.Id, detail.Name+".bpmn20.xml")
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取详情xml文件失败")
		return
	}

	resp.Id = detail.Id
	resp.Name = detail.Name
	resp.Xml = xml
	return
}
