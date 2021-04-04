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
		album.POST("/add_played_time", handler.AddALbumPlayedTime)
	}

	///列表组
	albumsList := r.Group("/albums")
	{
		albumsList.POST("/list", handler.GetAlbumListHandle)
	}

	///演播者组
	playerList := r.Group("/players")
	{
		///根据类型取得对应的列表
		playerList.POST("players", handler.GetPlayerListByCategoriesHandle)
	}

	///赚钱分组
	money := r.Group("/money")
	{
		money.POST("post_money", handler.PostTodayUserMoneyHandle)
		money.POST("get_today_money", handler.GetTodayUserMoneyHandle)
		money.POST("get_someday_money", handler.GetSomeDayUserMoneyHandle)
	}
	return r
}
