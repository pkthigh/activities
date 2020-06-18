package gateway

// ResultInfo 返回信息
type ResultInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
