package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/services"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/convert"
	"github.com/go-tour/blog_service/pkg/errcode"
	"github.com/go-tour/blog_service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (Upload) UploadFile(c *gin.Context) {
	resp := app.NewResp(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustUint32()
	if err != nil {
		resp.ErrResp(errcode.InvalidParam.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		resp.ErrResp(errcode.InvalidParam)
		return
	}
	svc := services.New(c.Request.Context())
	uploadFile, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.ErrorF("svc.uploadFile err: %v", err)
		resp.ErrResp(errcode.ErrUploadFile.WithDetails(err.Error()))
		return
	}
	resp.Success(gin.H{"fileAccessUrl": uploadFile.AccessUrl})

}
