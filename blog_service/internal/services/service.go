package services

import (
	"context"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/dao"
	otgorm "github.com/smacker/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//svc.dao = dao.New(global.DBEngine)
	svc.dao = dao.New(otgorm.SetSpanToGorm(ctx, global.DBEngine))
	return svc
}
