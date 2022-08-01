package errcode

var (
	Success          = NewError(0, "success")
	ServerErr        = NewError(10001, "system error")
	InvalidParam     = NewError(10002, "param illegal")
	NotFound         = NewError(10003, "not found")
	AuthNotFoundFail = NewError(10004, "auth pair not found failed")
	AuthFail         = NewError(10005, "auth failed")
	TokenErr         = NewError(10006, "token err")
	TimeOut          = NewError(10007, "operator time out")
	TooManyReq       = NewError(10008, "too many requests")
)
