package storage

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

func InitializeStaticDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func StoreFile(file *multipart.FileHeader) (string, error) {
	fileName := uuid.New().String()

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	dst, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}
