package dao

import (
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/pkg/app"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagId         uint32 `json:"tagId"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"coverImageUrl"`
	CreatedBy     string `json:"createdBy"`
	ModifiedBy    string `json:"modifiedBy"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
	values := map[string]any{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)

}

func (d *Dao) GetArticle(articleId uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: articleId}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(articleId uint32) error {
	article := model.Article{Model: &model.Model{ID: articleId}}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagID(articleId uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, articleId)
}

func (d *Dao) GetArticleListByTagID(articleId uint32, state uint8, pageNo, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, articleId, app.GetPageOffset(pageNo, pageSize), pageSize)
}
