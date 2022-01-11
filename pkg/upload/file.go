package upload

import (
	"io"
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

// 创建在上传文件时所使用的保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	// 该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，
	// 若涉及的目录均已存在，则不会进行任何操作，直接返回 nil
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// 保存所上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open() // 打开源地址的文件
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst) // 创建目标地址的文件
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
