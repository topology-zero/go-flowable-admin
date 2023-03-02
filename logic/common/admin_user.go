package common

import (
	"strconv"

	"go-flow-admin/model"
	"go-flow-admin/query"
)

func GetAdminUser() map[string]*model.AdminUserModel {
	adminUser, _ := query.AdminUserModel.Find()
	user := make(map[string]*model.AdminUserModel)

	for _, userModel := range adminUser {
		user[strconv.Itoa(userModel.ID)] = userModel
	}

	return user
}
