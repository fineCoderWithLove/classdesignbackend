package enum

// 错误码规则:
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//
//	----------------------------------------------------------
//	    第1位               2、3位                  4、5位
//	----------------------------------------------------------
//	  服务级错误码          模块级错误码	         具体错误码
//	----------------------------------------------------------
type Code int64
type Msg string

var (
	// Code
	OK               int64 = 200 // 成功
	Error            int64 = 500 // 错误
	NoAuthentication int64 = 403 // 未鉴权
	NetWorkError     int64 = 501 // 网络错误

	// Msg
	Success     string = "success"
	Fail        string = "fail"
	NoUserID    string = "用户id不正确"
	SystemError string = "系统异常"
)
