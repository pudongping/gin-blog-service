package upload

import (
	"io/ioutil"
	"mime/multipart"
	"os"
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

// 获取文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) // 原始文件名
	fileName = util.EncodeMD5(fileName)       // 将文件名进行 md5 加密

	return fileName + ext // 经过 md5 加密处理后的文件名
}

// 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst) // 获取文件的描述信息 FileInfo
	return os.IsNotExist(err)
}

// 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

// 检查文件大小是否超出最大大小限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}

	}

	return false
}

// 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}
