package upload

import (
	"path"
	"strings"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/pkg/util"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// GetFileName 获取文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) // 原始文件名
	fileName = util.EncodeMD5(fileName)       // 将文件名进行 md5 加密

	return fileName + ext // 经过 md5 加密处理后的文件名
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}
