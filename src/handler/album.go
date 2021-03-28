package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tingfm/do"
	"tingfm/service"
)

func SearchAlbumHandle(c *gin.Context) {

	var albumRequestEntity do.NewAlbumRequestEntity

	err := c.Bind(&albumRequestEntity)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}
	result, err := service.SearchAlbum(albumRequestEntity.AlbumName)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}

	c.JSON(http.StatusOK, result)

}



func NewAlbumHandle(c *gin.Context) {

	var album do.Album

	err := c.Bind(&album)
	if err != nil {
		fmt.Printf("album params is not value able, err: %#v", err.Error())
		return
	}

	err = service.NewAlbum(&album)
	if err != nil {
		fmt.Printf("album params is not value able, err: %#v", err.Error())
		return
	}

	c.JSON(http.StatusOK, "添加成功")

	return

}

func DelAlbumHandle(c *gin.Context) {
	var albumRequestEntity do.NewAlbumRequestEntity

	err := c.Bind(&albumRequestEntity)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}
	err = service.DelAlbum(albumRequestEntity.AlbumName)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}

	c.JSON(http.StatusOK, "删除成功")
}

///增加播放次数
func AddALbumPlayedTime(c *gin.Context) {
	var albumRequestEntity do.NewAlbumRequestEntity

	err := c.Bind(&albumRequestEntity)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}
	err = service.AddALbumPlayedTime(albumRequestEntity.AlbumName)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}
	c.JSON(http.StatusOK, "点击播放一次啦")
}
