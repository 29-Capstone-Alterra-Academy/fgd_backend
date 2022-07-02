package storage

import (
	"fgd/app/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type StorageHelper struct {
	conf config.Config
}

func (h *StorageHelper) InitializeStaticDirectory() error {
	return os.MkdirAll(filepath.Join(h.conf.STATIC_ROOT, h.conf.STATIC_PATH), os.ModePerm)
}

func (h *StorageHelper) StoreFile(file *multipart.FileHeader) (string, error) {
	fileName := uuid.New().String()

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	dst, err := os.Create(filepath.Join(h.conf.STATIC_ROOT, h.conf.STATIC_PATH, fileName))
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}

func NewStorageHelper(c config.Config) *StorageHelper {
	return &StorageHelper{
		conf: c,
	}
}
