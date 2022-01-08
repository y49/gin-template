package service

import (
	"errors"
	"gin-template/global"
	"gin-template/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetFileSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContentExt(fileType, fileName) {
		return nil, errors.New("file suffix not is not supported.")
	}

	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("error creating")
		}
	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("expected maximum size.")
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
