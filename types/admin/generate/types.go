// Code generated by goctl. DO NOT EDIT.
package generate

type GenerateListResponse struct {
	Tables []string `json:"tables"`
}

type GenerateRequest struct {
	Table string `json:"table" binding:"required" label:"表名"` // 表名
}

type GenerateResponse struct {
	Js  string `json:"js"`  // js 文件内容
	Vue string `json:"vue"` // vue 文件内容
	Api string `json:"api"` // api 文件内容
}
