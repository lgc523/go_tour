package errcode

var (
	Success          = NewError(0, "success")
	ServerErr        = NewError(10001, "system error")
	InvalidParam     = NewError(10002, "param illegal")
	NotFound         = NewError(10003, "not found")
	AuthNotFoundFail = NewError(10004, "auth pair not found")
	AuthFail         = NewError(10005, "auth failed")
	AuthTimeout      = NewError(10006, "jwt timeout")
	TokenErr         = NewError(10007, "token err")
	OpTimeOut        = NewError(10008, "operator time out")
	TooManyReq       = NewError(10009, "too many requests")
)
