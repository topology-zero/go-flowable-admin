package instance

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/history/history_activity_instance"
	"github.com/topology-zero/flowablesdk/history/history_task_instance"
	"github.com/topology-zero/flowablesdk/task/task_attachment"
	"github.com/topology-zero/flowablesdk/task/task_comment"
	"go-flow-admin/logic/common"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process/instance"
)

// Detail 流程详情
func Detail(req *instance.ProcessInstanceDetailRequest, ctx *svc.ServiceContext) (resp instance.ProcessInstanceDetailResponse, err error) {

	param := history_activity_instance.ListRequest{}
	param.Size = 1000
	param.Sort = "startTime"
	param.Order = "asc"
	param.ProcessInstanceId = req.Id
	activity, _, err := history_activity_instance.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取流程流转错误")
		return
	}

	adminUsers := common.GetAdminUser()
	for _, v := range activity {
		if v.ActivityType == "sequenceFlow" {
			continue
		}

		userName := ""
		if len(v.Assignee) > 0 {
			userName = adminUsers[v.Assignee].Realname
		}

		endTime := ""
		if v.EndTime != nil {
			endTime = v.EndTime.Format("2006-01-02 15:04:05")
		}

		act := instance.TaskHistory{
			ActionName: v.ActivityName,
			HandelUser: userName,
			CreateTime: v.StartTime.Format("2006-01-02 15:04:05"),
			HandleTime: endTime,
			UseTime:    v.DurationInMillis,
		}

		if len(v.TaskId) > 0 {
			// 获取表单
			var (
				includeTaskLocalVariables = true
				includeProcessVariables   = true
			)
			taskDetail, _, _ := history_task_instance.List(history_task_instance.ListRequest{
				TaskId:                    v.TaskId,
				IncludeTaskLocalVariables: &includeTaskLocalVariables,
				IncludeProcessVariables:   &includeProcessVariables,
			})
			if len(taskDetail) != 1 {
				continue
			}
			//for _, variable := range taskDetail[0].Variables {
			//	act.Form = append(act.Form, instance.Form{
			//		Id:    variable.Name,
			//		Name:  variable.Name,
			//		Type:  variable.Type,
			//		Value: variable.Value,
			//	})
			//}

			// 获取批注
			comments, err := task_comment.List(v.TaskId)
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				return resp, errors.New("获取批注错误")
			}
			for _, comment := range comments {
				act.Comment = append(act.Comment, instance.Comment{
					Id:      comment.Id,
					TaskId:  comment.TaskId,
					Message: comment.Message,
					Author:  adminUsers[comment.Author].Realname,
					Time:    comment.Time.Format("2006-01-02 15:04:05"),
				})
			}

			// 获取附件
			attachments, err := task_attachment.List(v.TaskId)
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				return resp, errors.New("获取附件错误")
			}
			for _, attachment := range attachments {
				act.Attachment = append(act.Attachment, instance.Attachment{
					Id:          attachment.Id,
					TaskId:      attachment.TaskUrl[strings.LastIndex(attachment.TaskUrl, "/")+1:],
					Url:         attachment.Url,
					Author:      adminUsers[attachment.UserId].Realname,
					Name:        attachment.Name,
					Description: attachment.Description,
					Time:        attachment.Time.Format("2006-01-02 15:04:05"),
				})
			}

			// 获取操作日志
			//events, err := task_event.List(v.TaskId)
			//if err != nil {
			//	ctx.Log.Errorf("%+v", errors.WithStack(err))
			//	err = errors.New("获取操作日志错误")
			//}
		}
		resp.History = append(resp.History, act)
	}

	return
}
