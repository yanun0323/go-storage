package internal

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yanun0323/pkg/logs"
	"gorm.io/gorm"
)

type FileManager struct {
	DB *gorm.DB
}

func NewFileManager(db *gorm.DB) *FileManager {
	return &FileManager{
		DB: db,
	}
}

func (fm *FileManager) GetFile(c echo.Context) error {
	id := c.Param("filename")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, Response("filename is required"))
	}

	var file File
	if err := fm.DB.Table(file.TableName()).Where("id = ?", id).Take(&file).Error; err != nil {
		return c.JSON(http.StatusNotFound, Response("file not found, err: %+v", err))
	}

	return c.Blob(http.StatusOK, file.Type, file.Content)
}

func (fm *FileManager) PostFile(c echo.Context) error {
	fh, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response("get file from 'file' field, err: %+v", err))
	}

	f, err := fh.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response("open file '%s', err: %+v", fh.Filename, err))
	}
	defer f.Close()

	logs.Infof("file name: %s, size: %d, header: %+v", fh.Filename, fh.Size, fh.Header)

	buf, err := io.ReadAll(f)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	extIdx := strings.LastIndex(fh.Filename, ".")
	if extIdx == -1 {
		return c.JSON(http.StatusBadRequest, Response("file name must have extension"))
	}

	ext := fh.Filename[extIdx+1:]

	h := md5.New()
	_, _ = h.Write(buf)
	filename := fmt.Sprintf("%x.%s", h.Sum(nil), ext)
	file := File{
		ID:      filename,
		Type:    http.DetectContentType(buf),
		Content: buf,
	}
	if err := fm.DB.Create(&file).Error; err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusInternalServerError, Response("create file, err: %+v", err))
		}
	}

	return c.JSON(http.StatusOK, struct {
		Filename string `json:"filename"`
	}{
		Filename: filename,
	})
}
