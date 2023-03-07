package task

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_definition"
	"github.com/topology-zero/flowablesdk/history/history_activity_instance"
	"github.com/topology-zero/flowablesdk/history/history_task_instance"
	"github.com/topology-zero/flowablesdk/task/task_attachment"
	"github.com/topology-zero/flowablesdk/task/task_comment"
	"go-flow-admin/logic/common"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// Detail 任务详情
func Detail(req *task.TaskDetailRequest, ctx *svc.ServiceContext) (resp task.TaskDetailResponse, err error) {
	param := history_activity_instance.ListRequest{}
	param.Size = 1000
	param.Sort = "startTime"
	param.Order = "asc"
	param.ProcessInstanceId = req.Id
	activity, _, err := history_activity_instance.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取历史流程错误")
		return
	}

	// 获取全部用户
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

		act := task.Action{
			TaskId:     v.TaskId,
			ActionName: v.ActivityName,
			HandelUser: userName,
			CreateTime: v.StartTime.Format("2006-01-02 15:04:05"),
			HandleTime: endTime,
			UseTime:    v.DurationInMillis,
		}

		if len(v.TaskId) > 0 {
			// 获取填写的表单
			var (
				includeTaskLocalVariables = true
				includeProcessVariables   = true
			)
			taskDetail, _, _ := history_task_instance.List(history_task_instance.ListRequest{
				TaskId:                    v.TaskId,
				IncludeTaskLocalVariables: &includeTaskLocalVariables,
				IncludeProcessVariables:   &includeProcessVariables,
			})
			detail := taskDetail[0]
			for _, variable := range detail.Variables {
				act.FormProperties = append(act.FormProperties, task.Properties{
					Id:    variable.Name,
					Value: variable.Value,
				})
			}

			// 获取表单
			if len(detail.FormKey) > 0 {
				formDef, _, _ := form_definition.List(form_definition.ListRequest{
					Key:    detail.FormKey,
					Latest: true,
				})
				model, err := form_definition.Model(formDef[0].Id)
				if err != nil {
					ctx.Log.Errorf("%+v", errors.WithStack(err))
					return resp, errors.New("获取表单错误")
				}
				act.FormRule = model.Fields[0].Value.(string)
				act.FormOption = model.Fields[1].Value.(string)
			}

			// 获取批注
			comments, err := task_comment.List(v.TaskId)
			if err != nil {
				ctx.Log.Errorf("%+v", errors.WithStack(err))
				return resp, errors.New("获取批注错误")
			}
			for _, comment := range comments {
				act.Comment = append(act.Comment, task.Comment{
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
				act.Attachment = append(act.Attachment, task.Attachment{
					Id:          attachment.Id,
					TaskId:      attachment.TaskUrl[strings.LastIndex(attachment.TaskUrl, "/")+1:],
					Url:         attachment.Url,
					Author:      adminUsers[attachment.UserId].Realname,
					Name:        attachment.Name,
					Description: attachment.Description,
					Time:        attachment.Time.Format("2006-01-02 15:04:05"),
				})
			}
		}
		resp.History = append(resp.History, act)
	}

	return
}
