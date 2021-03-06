package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/internal/service"
	"github.com/pudongping/gin-blog-service/pkg/app"
	"github.com/pudongping/gin-blog-service/pkg/convert"
	"github.com/pudongping/gin-blog-service/pkg/errcode"
	"github.com/pudongping/gin-blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// 上传文件
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file") // 读取入参 file 字段的上传文件信息
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt() // 利用入参 type 字段作为所上传文件类型的确认依据
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
