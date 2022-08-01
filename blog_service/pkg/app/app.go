package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/pkg/errcode"
	"net/http"
)

type Resp struct {
	Ctx *gin.Context
}

type RespTemplate struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type Pager struct {
	PageNo    int `json:"pageNo"`
	PageSize  int `json:"pageSize"`
	TotalSize int `json:"totalSize"`
}

func NewResp(ctx *gin.Context) *Resp {
	return &Resp{ctx}
}

func (r *Resp) Success(data any) {
	r.Ctx.JSON(http.StatusOK, &RespTemplate{Code: 200, Data: data, Msg: "success"})
}

func (r *Resp) SuccessPage(list any, totalSize int) {
	r.Ctx.JSON(http.StatusOK, &RespTemplate{Code: 200, Data: gin.H{
		"list": list,
		"pageInfo": Pager{
			PageNo:    GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalSize: totalSize,
		}}, Msg: "success"})
}

func (r *Resp) FailResp(err *errcode.Error) {
	template := RespTemplate{
		Code: err.StatusCode(),
		Msg:  err.Msg()}
	r.Ctx.JSON(err.StatusCode(), template)
}

func (r *Resp) FailRespWithExtraMsg(err *errcode.Error, msg string) {
	template := RespTemplate{
		Code: err.StatusCode(),
		Msg:  err.Msg()}

	if len(msg) > 0 {
		template.Msg += ", cause: " + msg
	}
	r.Ctx.JSON(err.StatusCode(), template)
}
