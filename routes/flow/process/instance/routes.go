// Code generated by goctl. DO NOT EDIT.
package instance

import (
	"go-flow-admin/handler/flow/process/instance"
	"go-flow-admin/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterFlowProcessInstanceRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware, middleware.AuthMiddleware)
	g.GET("/flow/proc_inst", instance.AllHandle)
	g.GET("/flow/proc_inst/:id", instance.DetailHandle)

}
