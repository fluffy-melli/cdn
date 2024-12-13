package ui

import (
	"cdn-module/config"
	_ "cdn-module/docs"
	"cdn-module/packages/cache"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var configs *config.CDN_CONFIG_JSON
var caches *cache.Safe

// UploadFile godoc
//	@Summary		Upload a file
//	@Description	Uploads a file to the server
//	@Tags			File
//	@Accept			multipart/form-data
//	@Param			file	formData	file	true	"File to upload"
//	@Success		200		{object}	gin.H	"File uploaded successfully"
//	@Failure		400		{object}	gin.H	"Bad request"
//	@Failure		500		{object}	gin.H	"Internal server error"
//	@Router			/upload [post]

func UploadFile(c *gin.Context) {
	start := time.Now()
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, configs.Server.MaxFSMB*1024*1024)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "File is required"})
		return
	}
	if c.Request.ContentLength > configs.Server.MaxFSMB*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "File size exceeds limit"})
		return
	}
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(configs.Cache_Dir, newFileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Unable to save file"})
		return
	}
	caches.Cache(configs, start, newFileName)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "File uploaded successfully", "filename": newFileName})
}

// SendFile godoc
//	@Summary		Send a file
//	@Description	Sends a file by filename
//	@Tags			File
//	@Param			filename	path		string	true	"Filename to retrieve"
//	@Success		200			{string}	string	"File content"
//	@Failure		404			{object}	gin.H	"File not found"
//	@Router			/file/{filename} [get]

func SendFile(c *gin.Context) {
	filename := strings.ReplaceAll(c.Param("filename"), "/", "")
	file, exists := caches.LoadFile(*configs, filename)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "no file"})
		return
	}
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)

	if mimeType == "video/mp4" {
		StreamVideo(c, file, filename)
	} else {
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		c.Data(http.StatusOK, mimeType, file)
	}
}

func StreamVideo(c *gin.Context, fileData []byte, filename string) {
	fileSize := int64(len(fileData))
	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", filename))
	c.Header("Accept-Ranges", "bytes")

	rangeHeader := c.GetHeader("Range")
	fmt.Printf("Range Header: %s\n", rangeHeader)

	if rangeHeader == "" {
		// Range가 없을 경우, 기본적으로 1MB 범위를 설정
		const defaultBytesToSend = 1 * 1024 * 1024 // 1MB
		bytesToSend := int64(defaultBytesToSend)

		if bytesToSend > fileSize {
			bytesToSend = fileSize
		}

		start := int64(0)
		end := start + bytesToSend - 1

		if end >= fileSize {
			end = fileSize - 1
		}

		c.Header("Content-Length", strconv.FormatInt(end-start+1, 10))
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		c.Status(http.StatusPartialContent)
		c.Data(http.StatusPartialContent, "video/mp4", fileData[start:end+1])
		return
	}

	// Range 헤더가 있는 경우
	var start, end int64
	if _, err := fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end); err != nil {
		c.String(http.StatusRequestedRangeNotSatisfiable, "Error: The requested range is invalid")
		return
	}

	// 끝 바이트가 지정되지 않았을 경우
	if end < 0 {
		end = fileSize - 1
	}

	// 유효한 범위인지 확인
	if start >= fileSize || end < 0 || start > end {
		c.String(http.StatusRequestedRangeNotSatisfiable, "Error: Requested range is not satisfiable")
		return
	}

	if end >= fileSize {
		end = fileSize - 1
	}

	c.Header("Content-Length", strconv.FormatInt(end-start+1, 10))
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	c.Status(http.StatusPartialContent)

	// 요청된 범위의 바이트를 클라이언트에 전송
	if _, err := c.Writer.Write(fileData[start : end+1]); err != nil {
		c.String(http.StatusInternalServerError, "Error: Failed to write to the client")
	}
}

func FileList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "content": caches.GetKeys()})
}

func INIT(cf *config.CDN_CONFIG_JSON, memory *cache.Safe) {
	configs = cf
	caches = memory
}
