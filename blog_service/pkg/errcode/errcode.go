package errcode

import (
	"fmt"
	"github.com/go-tour/blog_service/global"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		global.Logger.ErrorF("[errCode %d already set, please change another.. ]", code)
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("errCode: %d, errMsg: %s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) ErrMsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerErr.code:
		return http.StatusInternalServerError
	case InvalidParam.code:
		return http.StatusForbidden
	case NotFound.code:
		return http.StatusNotFound
	case AuthFail.code, AuthNotFoundFail.code, TokenErr.code:
		return http.StatusNonAuthoritativeInfo
	case TooManyReq.code:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
