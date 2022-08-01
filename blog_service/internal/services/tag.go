package services

import (
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/pkg/app"
)

//valid req

type CountTagReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagReq struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"createdBy" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagReq struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modifiedBy" binding:"required,min=3,max=100"`
	IsDel      uint8  `form:"isDel" binding:"oneof=0 1"`
}

type DeleteTagReq struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//dao func

func (svc *Service) CountTag(param *CountTagReq) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListReq, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.PageNo, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagReq) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagReq) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy, param.IsDel)
}

func (svc *Service) DeleteTag(param *DeleteTagReq) error {
	return svc.dao.DeleteTag(param.ID)
}
