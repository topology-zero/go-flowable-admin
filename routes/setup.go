// Code generated by goctl. DO NOT EDIT.
package routes

import (
	adminAuth "go-flow-admin/routes/admin/auth"
	adminBase "go-flow-admin/routes/admin/base"
	adminGenerate "go-flow-admin/routes/admin/generate"
	adminLogin "go-flow-admin/routes/admin/login"
	adminRole "go-flow-admin/routes/admin/role"
	adminUser "go-flow-admin/routes/admin/user"
	flowBase "go-flow-admin/routes/flow/base"
	flowCate "go-flow-admin/routes/flow/cate"
	FlowForm "go-flow-admin/routes/flow/form"
	flowProcess "go-flow-admin/routes/flow/process"
	flowProcessInstance "go-flow-admin/routes/flow/process/instance"
	flowTask "go-flow-admin/routes/flow/task"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) {
	adminLogin.RegisterAdminLoginRoute(e)
	adminUser.RegisterAdminUserRoute(e)
	adminRole.RegisterAdminRoleRoute(e)
	adminAuth.RegisterAdminAuthRoute(e)
	adminBase.RegisterAdminBaseRoute(e)
	adminGenerate.RegisterAdminGenerateRoute(e)
	flowCate.RegisterFlowCateRoute(e)
	flowBase.RegisterFlowBaseRoute(e)
	flowProcess.RegisterFlowProcessRoute(e)
	flowProcessInstance.RegisterFlowProcessInstanceRoute(e)
	flowTask.RegisterFlowTaskRoute(e)
	FlowForm.RegisterFlowFormRoute(e)
}
