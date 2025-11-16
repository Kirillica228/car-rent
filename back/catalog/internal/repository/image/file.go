package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"catalog/pkg/utils"
)

type FileRepository struct {
	storagePath string
	baseURL     string
}

func NewFileRepository(storagePath, baseURL string) *FileRepository {
	return &FileRepository{storagePath: storagePath, baseURL: baseURL}
}

func (r *FileRepository) Save(file multipart.File, header *multipart.FileHeader, carID uint) (string, error) {
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("car_%d_%s%s", carID, utils.GenerateRandomString(8), ext)
	fullPath := filepath.Join(r.storagePath, fileName)

	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", r.baseURL, fileName), nil
}

func (r *FileRepository) Delete(fileName string) error {
	fullPath := filepath.Join(r.storagePath, fileName)
	return os.Remove(fullPath)
}
