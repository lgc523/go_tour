package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/services"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/convert"
	"github.com/go-tour/blog_service/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	req := services.ArticleGetRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF("[app.article.Get.BindValid err %v]", errs)
		resp.FailRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}
	svc := services.New(c.Request.Context())
	article, err := svc.GetArticle(&req)
	if err != nil {
		global.Logger.ErrorF("GetArticle err: %s", err.Error())
		resp.FailRespWithExtraMsg(errcode.ErrGetArticle, errs.Error())
		return
	}
	if article.ID == 0 {
		resp.Success(nil)
		return
	}
	resp.Success(article)

}
func (a Article) List(c *gin.Context) {
	req := services.ArticleListRequest{}
	resp := app.NewResp(c)

	inValid, errs := app.BindAndValid(c, &req)
	if inValid {
		global.Logger.ErrorF("[article.List err: %s]", errs.Error())
		resp.FailRespWithExtraMsg(errcode.InvalidParam, errs.Error())
		return
	}

	svc := services.New(c.Request.Context())
	pager := app.Pager{PageNo: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articleList, count, err := svc.GetArticleList(&req, &pager)
	if err != nil {
		global.Logger.ErrorF("getArticleList err: %s", err.Error())
		resp.FailRespWithExtraMsg(errcode.ErrGetArticles, err.Error())
		return
	}
	//if count == 0 {
	//	resp.Success(nil)
	//	return
	//}
	resp.SuccessPage(articleList, count)
}
func (a Article) Create(c *gin.Context) {

}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
