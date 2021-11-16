package api

import "encoding/json"

type Response struct {
	rd responseData
}

type responseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewMDResponse(msg string, obj interface{}) *Response {
	return &Response{
		rd: responseData{
			Code: CodeOk,
			Msg:  msg,
			Data: obj,
		},
	}
}

func NewDataResponse(obj interface{}) *Response {
	return &Response{
		rd: responseData{
			Code: CodeOk,
			Msg:  "",
			Data: obj,
		},
	}
}

func NewSuccessResponse(msg string) *Response {
	return &Response{
		rd: responseData{
			Code: CodeOk,
			Msg:  msg,
		},
	}
}

func NewErrorResponse(msg string) *Response {
	return &Response{
		rd: responseData{
			Code: CodeErr,
			Msg:  msg,
		},
	}
}

func NewResponse(code int, msg string, obj interface{}) *Response {
	return &Response{
		rd: responseData{
			Code: code,
			Msg:  msg,
			Data: obj,
		},
	}
}

func (r *Response) Json() (ret []byte) {
	ret, err := json.Marshal(r.rd)
	if err != nil {
		ret, _ = json.Marshal(responseData{
			Code: CodeErr,
			Msg:  "encode json error",
		})
	}
	return
}

func (r *Response) Code() int {
	return r.rd.Code
}

func (r *Response) Msg() string {
	return r.rd.Msg
}
