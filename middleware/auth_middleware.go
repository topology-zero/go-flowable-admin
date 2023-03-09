package middleware

import (
	"strconv"
	"strings"

	"go-flow-admin/internal/response"
	"go-flow-admin/model"
	"go-flow-admin/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(c *gin.Context) {
	user, _ := c.Get("userInfo")
	claims := user.(*jwt.Claims)
	roleStr := "role:" + strconv.Itoa(claims.RoleId)
	ok, _ := model.Enforcer.Enforce(roleStr, c.FullPath(), strings.ToLower(c.Request.Method))
	if !ok {
		logrus.Warning("没有权限")
		response.HandleAbortResponse(c, "没有权限", 403)
		return
	}
	c.Next()
}
