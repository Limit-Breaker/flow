package common

type Response struct {
	// 代码
	Code int `json:"code" example:"10000"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func (r *Response) ResponseOK() *Response {
	r.Code = 0
	return r
}

func (r *Response) ResponseErr(code int) *Response {
	r.Code = code
	return r
}
