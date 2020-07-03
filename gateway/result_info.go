package gateway

import "activities/common/errs"

// ResultInfo 返回信息
type ResultInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Successful 快捷生成
func Successful(data interface{}) *ResultInfo {
	return &ResultInfo{
		Code: errs.DatabaseError.Int(),
		Msg:  "successful",
		Data: data,
	}
}
