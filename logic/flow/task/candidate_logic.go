package task

import (
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	processTask "github.com/topology-zero/flowablesdk/task"
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// Candidate 我的任务
func Candidate(req *task.TaskListRequest, ctx *svc.ServiceContext) (resp task.TaskListResponse, err error) {
	userInfo, _ := ctx.GinContext.Get("userInfo")
	claims := userInfo.(*jwt.Claims)

	query := processTask.ListRequest{}
	query.CandidateUser = strconv.Itoa(claims.Id)
	query.Sort = "name"
	query.Start = (req.Page - 1) * req.PageSize
	query.Size = req.PageSize

	list, count, err := processTask.List(query)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("获取我的候选任务错误")
		return
	}

	resp.Total = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	_ = copier.Copy(&resp.Data, &list)
	for i := range list {
		resp.Data[i].TaskName = list[i].Name
		resp.Data[i].CreateTime = list[i].CreateTime.Format("2006-01-02 15:04:05")
	}
	return
}
