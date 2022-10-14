package constant

// 接口相关返回值
const (
	Success        int = iota // 成功
	Failed                    // 失败
	TooManyRequest            // 短时间内太多请求
)

// redis相关定义
const (
	IpInterceptorDb = 9
)

// 一些key的相关定义
const (
	IpInterceptorKey = "request_ip:%s"
	ErrorList        = "error_list"
	ErrorCode        = "error_code"
	ErrorMsg         = "error_message"
)
