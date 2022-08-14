package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/services"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/errcode"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) GetAuth(c *gin.Context) {
	req := services.AuthRequest{}
	resp := app.NewResp(c)
	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF(c, "app.bindAndValid errs:%v", errs)
		resp.ErrResp(errcode.InvalidParam.WithDetails(errs.Error()))
		return
	} else {
		global.Logger.InfoF(c, "app.binding check valid", req)
	}
	svc := services.New(c.Request.Context())
	err := svc.CheckAuth(&req)
	if err != nil {
		global.Logger.ErrorF(c, "svc.checkAuth err: %v", err)
		resp.ErrResp(errcode.AuthNotFoundFail)
		return
	}
	token, err := app.GenerateToken(req.AppKey, req.AppSecret)
	if err != nil {
		global.Logger.ErrorF(c, "app.generateToken err: %v", err)
		resp.ErrResp(errcode.AuthFail.WithDetails(err.Error()))
		return
	}
	resp.Success(gin.H{"token": token})
}
