package dao

import (
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/pkg/app"
)

func (d *Dao) GetTag(tagId uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: tagId}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, pageNo, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	offset := app.GetPageOffset(pageNo, pageSize)
	return tag.List(d.engine, offset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy}}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string, isDel uint8) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	updates := map[string]any{
		"state":      state,
		"modifiedBy": modifiedBy,
		"isDel":      isDel,
	}
	if name != "" {
		updates["name"] = name
	}
	return tag.Update(d.engine, updates)
}
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}
