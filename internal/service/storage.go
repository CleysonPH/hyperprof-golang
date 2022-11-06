package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/cleysonph/hyperprof/config"
	"github.com/google/uuid"
)

func uploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s-%s", uuid.New().String(), fileHeader.Filename)
	filepath := path.Join("uploads", filename)
	tempFile, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	tempFile.Write(fileBytes)

	return fmt.Sprintf("http://%s/%s", config.Addr(), filepath), nil
}
