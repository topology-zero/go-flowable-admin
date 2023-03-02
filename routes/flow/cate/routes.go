// Code generated by goctl. DO NOT EDIT.
package cate

import (
	"go-flow-admin/handler/flow/cate"
	"go-flow-admin/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterFlowCateRoute(e *gin.Engine) {
	g := e.Group("")
	g.Use(middleware.JwtMiddleware, middleware.AuthMiddleware)
	g.GET("/flow/cate", cate.ListHandle)
	g.POST("/flow/cate", cate.AddHandle)
	g.PUT("/flow/cate/:id", cate.EditHandle)
	g.DELETE("/flow/cate/:id", cate.DelHandle)

}
