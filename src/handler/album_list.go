package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tingfm/do"
	"tingfm/service"
)


func GetAlbumListHandle(c *gin.Context) {
	var params do.GetAlbumListRequestEntity

	err := c.Bind(&params)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}
	result, err := service.GetAlbumList(&params)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
