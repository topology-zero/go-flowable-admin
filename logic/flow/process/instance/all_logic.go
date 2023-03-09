package instance

import (
	"time"

	"go-flow-admin/logic/common"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process/instance"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/history/history_process_instance"
	"github.com/topology-zero/flowablesdk/pkg/timefmt"
	"github.com/topology-zero/flowablesdk/process_definition"
	"github.com/topology-zero/flowablesdk/task"
)

// All 全部流程
func All(req *instance.ProcessInstanceRequest, ctx *svc.ServiceContext) (resp instance.ProcessInstanceResponse, err error) {

	adminUser := common.GetAdminUser()

	param := history_process_instance.ListRequest{}
	param.Start = (req.Page - 1) * req.PageSize
	param.Size = req.PageSize
	param.Sort = "startTime"
	param.Order = "desc"

	if len(req.ProcessName) > 0 {
		param.ProcessDefinitionName = req.ProcessName
	}

	if len(req.StartTime) > 0 && len(req.EndTime) > 0 {
		start, _ := timefmt.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local)
		end, _ := timefmt.ParseInLocation("2006-01-02 15:04:05", req.EndTime, time.Local)
		param.StartedAfter = start
		param.StartedBefore = end
	}

	list, count, err := history_process_instance.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取全部流程失败")
	}

	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = count

	copier.Copy(&resp.Data, &list)

	for i := range list {
		detail, err := process_definition.Detail(list[i].ProcessDefinitionId)
		if err != nil {
			ctx.Log.Errorf("%+v", errors.WithStack(err))
			err = errors.New("获取流程定义失败")
		}

		resp.Data[i].StartTime = list[i].StartTime.Format("2006-01-02 15:04:05")
		if list[i].EndTime != nil {
			resp.Data[i].EndTime = list[i].EndTime.Format("2006-01-02 15:04:05")
		}

		resp.Data[i].ProcessName = detail.Name
		resp.Data[i].Duration = list[i].DurationInMillis

		// 当前节点
		if list[i].EndTime == nil {
			tasks, _, err := task.List(task.ListRequest{ProcessInstanceId: list[i].Id})
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				err = errors.New("获取任务失败")
			}
			if len(tasks) > 0 {
				resp.Data[i].CurrentTask = tasks[0].Name
				if len(tasks[0].Assignee) > 0 {
					resp.Data[i].CurrentUser = adminUser[tasks[0].Assignee].Realname
				}
			}
		}
	}

	return
}
