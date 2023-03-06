package task

import (
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/process_definition"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"

	processTask "github.com/topology-zero/flowablesdk/task"
)

// Assignee 我的任务
func Assignee(req *task.TaskListRequest, ctx *svc.ServiceContext) (resp task.TaskListResponse, err error) {
	userInfo, _ := ctx.GinContext.Get("userInfo")
	claims := userInfo.(*jwt.Claims)

	query := processTask.ListRequest{}
	query.Assignee = strconv.Itoa(claims.Id)
	query.Sort = "createTime"
	query.Order = "desc"
	query.Start = (req.Page - 1) * req.PageSize
	query.Size = req.PageSize

	list, count, err := processTask.List(query)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取我的任务错误")
		return
	}

	resp.Total = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	_ = copier.Copy(&resp.Data, &list)

	var detail process_definition.ProcessDefinition
	for i := range list {
		detail, err = process_definition.Detail(list[i].ProcessDefinitionId)
		if err != nil {
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			err = errors.New("获取流程定义错误")
			return
		}
		resp.Data[i].ProcessName = detail.Name
		resp.Data[i].TaskName = list[i].Name
		resp.Data[i].CreateTime = list[i].CreateTime.Format("2006-01-02 15:04:05")
	}
	return
}
