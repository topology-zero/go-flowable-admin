// Code generated by goctl. DO NOT EDIT.
package base

type FlowcateResponse struct {
	List []FlowCateData `json:"list"` // 分类
}

type FlowCateData struct {
	Id   int    `json:"id"`   // ID
	Name string `json:"name"` // 分类名
}
