package dao

import (
	"tingfm/do"
)

func GetAlbumList(params *do.GetAlbumListRequestEntity) (albumList do.AlbumListResponseEntity, err error) {

	DB.Order("played_time desc").Where("category=?", params.CategoryName).Offset(params.Offset).Limit(params.Limit).Find(&albumList.Albums)
	DB.Order("played_time desc").Where("category=?", params.CategoryName).Offset(0).Limit(5).Find(&albumList.RecommendAlbums)

	return

}
