package _interface

type (
	ResponseStruct struct {
		Ret  int         `json:"ret"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)
