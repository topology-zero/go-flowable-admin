package task

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/comment"
	"github.com/topology-zero/flowablesdk/history/history_activity_instance"
	flowTask "github.com/topology-zero/flowablesdk/task"
	"github.com/topology-zero/flowablesdk/task/task_attachment"
	"github.com/topology-zero/flowablesdk/task/task_comment"
	"github.com/topology-zero/flowablesdk/task/task_form"
	"go-flow-admin/logic/common"
	"go-flow-admin/model"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

type processDetail struct {
	ctx          *svc.ServiceContext
	processId    string
	processDefId string
	adminUsers   map[string]*model.AdminUserModel
	detail       *task.TaskDetailResponse
}

func newProcessDetail(id string, ctx *svc.ServiceContext) *processDetail {
	p := &processDetail{
		ctx:        ctx,
		processId:  id,
		adminUsers: common.GetAdminUser(),
		detail:     new(task.TaskDetailResponse),
	}
	return p
}

// Detail 任务详情
func Detail(req *task.TaskDetailRequest, ctx *svc.ServiceContext) (task.TaskDetailResponse, error) {
	p := newProcessDetail(req.Id, ctx)
	err := p.getProcessHistory()
	return *p.detail, err
}

// getProcessHistory 获取流程实例历史流转
func (p *processDetail) getProcessHistory() error {
	var data []task.Action

	req := history_activity_instance.ListRequest{}
	req.Size = 1000
	req.Sort = "startTime"
	req.Order = "asc"
	req.ProcessInstanceId = p.processId
	activity, _, err := history_activity_instance.List(req)
	if err != nil {
		p.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return errors.New("获取历史流程错误")
	}

	for _, v := range activity {
		p.processDefId = v.ProcessDefinitionId
		if v.ActivityType == "sequenceFlow" {
			continue
		}

		userName := ""
		if len(v.Assignee) > 0 {
			userName = p.adminUsers[v.Assignee].Realname
		}

		endTime := ""
		if v.EndTime != nil {
			endTime = v.EndTime.Format("2006-01-02 15:04:05")
		}

		act := task.Action{
			ActionName: v.ActivityName,
			HandelUser: userName,
			CreateTime: v.StartTime.Format("2006-01-02 15:04:05"),
			HandleTime: endTime,
			UseTime:    v.DurationInMillis,
		}

		if len(v.TaskId) > 0 {
			taskDetail, _ := flowTask.Detail(v.TaskId)

			// 获取表单
			if v.EndTime == nil && len(taskDetail.FormKey) > 0 {
				if err = p.getForm(v.TaskId); err != nil {
					return err
				}
			}

			// 获取备注
			act.Comment, err = p.getTaskComment(v.TaskId)
			if err != nil {
				return err
			}

			// 获取附件
			act.Attachment, err = p.getTaskAttachment(v.TaskId)
			if err != nil {
				return err
			}
		}

		data = append(data, act)
	}
	p.detail.History = data
	return nil
}

// 获取任务的备注
func (p *processDetail) getTaskComment(id string) (data []task.Comment, err error) {
	var comments []comment.Comment
	comments, err = task_comment.List(id)
	if err != nil {
		p.ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取任务表单错误")
		return
	}

	for _, c := range comments {
		data = append(data, task.Comment{
			Id:      c.Id,
			TaskId:  c.TaskId,
			Author:  p.adminUsers[c.Author].Realname,
			Message: c.Message,
			Time:    c.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// 获取任务的附件
func (p *processDetail) getTaskAttachment(id string) (data []task.Attachment, err error) {
	attachments, err := task_attachment.List(id)
	if err != nil {
		p.ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取流程附件错误")
		return
	}
	for _, attachment := range attachments {
		data = append(data, task.Attachment{
			Id:          attachment.Id,
			TaskId:      id,
			Url:         attachment.TaskUrl[strings.LastIndex(attachment.TaskUrl, "/")+1:],
			Author:      p.adminUsers[attachment.UserId].Realname,
			Name:        attachment.Name,
			Description: attachment.Description,
			Time:        attachment.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// 获取任务表单
func (p *processDetail) getForm(id string) error {
	// 内置表单用这个
	//data, err := form.GetFrom(form.GetFromRequest{
	//	TaskId: id,
	//})

	// 外置表单用这个
	getForm, err := task_form.GetForm(id)
	p.ctx.Log.Errorf("%+v", getForm)

	if err != nil {
		p.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return errors.New("获取任务表单错误")
	}

	if len(getForm.Fields) != 0 {
		p.detail.FormRule = getForm.Fields[0].Value.(string)
		p.detail.FormOption = getForm.Fields[1].Value.(string)
	}
	return nil
}
