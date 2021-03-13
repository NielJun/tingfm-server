package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tingfm/do"
	"tingfm/service"
)

// 通过给定类型查询对应的播放作者
func GetPlayerListByCategoriesHandle(c *gin.Context) {


	var params do.GetPlayersByCategoryRequestEntity

	err := c.Bind(&params)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}

	result,err:= service.GetPlayerListByCategories(params.Properties.CategoryName.Type)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
