// Code generated by goctl. DO NOT EDIT.
package generate

import (
	"go-flow-admin/handler/admin/generate"
	"go-flow-admin/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminGenerateRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware)
	g.GET("/generate", generate.ListHandle)
	g.POST("/generate", generate.GenerateHandle)

}
