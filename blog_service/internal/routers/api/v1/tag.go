package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/services"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/convert"
	"github.com/go-tour/blog_service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(c *gin.Context) {
	req := services.TagListReq{}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF(c, "app.tag.List.BindAndValid errs: %v", errs)
		resp.ErrRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}

	svc := services.New(c.Request.Context())
	tagCap, err := svc.CountTag(&services.CountTagReq{Name: req.Name, State: req.State})
	if err != nil {
		global.Logger.ErrorF(c, "[svc.CountTag err: 5v]", err)
		resp.ErrRespWithExtraMsg(errcode.CountTagFail, err.Error())
		return
	}

	pager := app.Pager{PageNo: app.GetPage(c), PageSize: app.GetPageSize(c)}
	tags, err := svc.GetTagList(&req, &pager)
	if err != nil {
		global.Logger.ErrorF(c, "[svc.GetTagList err: 5v]", err)
		resp.ErrRespWithExtraMsg(errcode.GetListFail, err.Error())
		return
	}

	resp.SuccessPage(tags, tagCap)
}
func (t Tag) Create(c *gin.Context) {
	req := services.CreateTagReq{}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF(c, "app.tag.Create.BindAndValid errs: %v", errs)
		resp.ErrRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}

	svc := services.New(c.Request.Context())
	err := svc.CreateTag(&req)
	if err != nil {
		global.Logger.ErrorF(c, "svc.tag.Create err: %v", err)
		resp.ErrResp(errcode.CreateTagFail)
		return
	}
	resp.Success(gin.H{})
}

func (t Tag) Update(c *gin.Context) {
	req := services.UpdateTagReq{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF(c, "app.tag.Update.BindAndValid errs: %v", errs)
		resp.ErrRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}

	svc := services.New(c.Request.Context())
	err := svc.UpdateTag(&req)
	if err != nil {
		global.Logger.ErrorF(c, "svc.tag.Update err: %v", err)
		resp.ErrResp(errcode.UpdateTagFail)
		return
	}
	resp.Success(gin.H{})
}

func (t Tag) Delete(c *gin.Context) {
	req := services.DeleteTagReq{ID: convert.StrTo(c.Param("id")).MustUint32()}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF(c, "app.tag.Delete.BindAndValid errs: %v", errs)
		resp.ErrRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}

	svc := services.New(c.Request.Context())
	err := svc.DeleteTag(&req)
	if err != nil {
		global.Logger.ErrorF(c, "svc.tag.Delete err: %v", err)
		resp.ErrRespWithExtraMsg(errcode.DeleteTagFail, err.Error())
		return
	}
	resp.Success(gin.H{})
}
