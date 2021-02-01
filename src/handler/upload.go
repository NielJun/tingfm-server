package handler

import (
	"github.com/gin-gonic/gin"
	"tingfm/service"
)


///上传文件
func UploadAlbumHandle(context *gin.Context) {
	service.UploadAlbum()
}
