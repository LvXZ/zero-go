package model

// @Author: lvxiaozheng
// @Date: 2021/6/11 11:52
// @Description:

//Http
type HttpResponse struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    []KvData `json:"data"`
}

type KvData struct {
	K string `json:"k"`
	V string `json:"v"`
}
