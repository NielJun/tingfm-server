package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tingfm/service"
)

type NewCategoryRequestEntity struct {
	CategoryName string `json:"category_name"`
}

///获取所有的分类
func GetCategoriesHandle(c *gin.Context) {
	result, err := service.GetCategories()

	if err != nil {
		fmt.Printf("err: %#v", err.Error())
	}

	c.JSON(http.StatusOK, result)

}

///新建一个Category
func NewCategoryHandle(c *gin.Context) {
	var newCategory NewCategoryRequestEntity
	err := c.Bind(&newCategory)

	if err != nil {
		fmt.Printf("bind json faild. err is %#v", err.Error())
	}

	service.NewCategory(newCategory.CategoryName)

	c.JSON(http.StatusOK,&newCategory)

}

///删除一个Category
func DelCategoryHandle(c *gin.Context) {

	var newCategory NewCategoryRequestEntity
	err := c.Bind(&newCategory)

	if err != nil {
		fmt.Printf("bind json faild. err is %#v", err.Error())
	}

	service.DelCategory(newCategory.CategoryName)

	c.JSON(http.StatusOK,&newCategory)


}

func ModCategoriesHandle(c* gin.Context){
	var newCategory NewCategoryRequestEntity
	err := c.Bind(&newCategory)

	if err != nil {
		fmt.Printf("bind json faild. err is %#v", err.Error())
	}

	service.ModCategory(newCategory.CategoryName)

	c.JSON(http.StatusOK,&newCategory)
}
