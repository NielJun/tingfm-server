package routers

import (
	"github.com/gin-gonic/gin"
	"tingfm/handler"
)

///	Get 查询
///	Post 创建
///	Put 更新
/// Delete 删除

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/upload", handler.UploadAlbumHandle)

	//路由组,处理请求

	///标签组
	category := r.Group("/categories")
	{
		category.DELETE("/category", handler.DelCategoryHandle)
		category.POST("/category", handler.NewCategoryHandle)
		category.GET("/category", handler.GetCategoriesHandle)
		category.PUT("/category", handler.ModCategoriesHandle)
	}

	///专辑组
	album := r.Group("/album")
	{
		album.POST("/search_album", handler.SearchAlbumHandle)
		album.POST("/album", handler.NewAlbumHandle)
		album.DELETE("/album", handler.DelAlbumHandle)
	}

	///列表组
	albumsList := r.Group("/albums")
	{
		albumsList.POST("/list", handler.GetAlbumListHandle)
	}
	return r
}
