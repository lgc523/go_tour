package model

import (
	"github.com/jinzhu/gorm"
)

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tagId"`
	ArticleId uint32 `json:"articleId"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id=? AND is_del =?", a.ArticleId, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id=? AND is_del=?", a.TagID, 0).Find(&articleTags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return articleTags, err
	}
	return articleTags, nil
}

func (a ArticleTag) ListByAIds(db *gorm.DB, articleIds []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("article_id IN (?) AND is_del=?", articleIds, 0).Find(&articleTags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return articleTags, err
	}
	return articleTags, nil
}

func (a ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values any) error {
	if err := db.Model(&a).Where("article_id =? AND is_del=?", a.ArticleId, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del=?", a.ArticleId, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
