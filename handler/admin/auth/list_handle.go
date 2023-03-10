package auth

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/auth"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle ๆ้ๅ่กจ
func ListHandle(c *gin.Context) {
	resp, err := auth.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}
