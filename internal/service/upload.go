package service

import (
	"errors"
	"mime/multipart"
	"os"
	"time"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename) // 获取文件名称
	if !upload.CheckContainExt(fileType, fileName) {    // 检查文件后缀是否包含在约定的后缀配置项中
		return nil, errors.New("file suffix is not supported")
	}

	if upload.CheckMaxSize(fileType, file) { // 检查文件大小是否超出最大大小限制
		return nil, errors.New("exceeded maximum file limit")
	}

	timeFolder := time.Now().Format("20060102")               // 增加时间目录
	uploadSavePath := upload.GetSavePath() + "/" + timeFolder // 获取文件保存地址
	if upload.CheckSavePath(uploadSavePath) {                 // 检查保存目录是否存在
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil { // 创建在上传文件时所使用的保存目录
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckPermission(uploadSavePath) { // 检查文件权限是否足够
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil { // 保存所上传的文件
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + timeFolder + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
