package dao

import "github.com/go-tour/blog_service/internal/model"

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetArticleTagListByTID(tagId uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagId}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListBYAIDs(articleIds []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIds(d.engine, articleIds)
}

func (d *Dao) CreateArticleTag(articleID, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		ArticleId: articleID,
		TagID:     tagId,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleId, tagId uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleId: articleId}
	values := map[string]any{
		"article_id":  articleId,
		"tag_id":      tagId,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleId uint32) error {
	articleTag := model.ArticleTag{ArticleId: articleId}
	return articleTag.Delete(d.engine)
}
