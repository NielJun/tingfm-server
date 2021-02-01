package dao

import (
	"tingfm/do"
)

func GetAlbumList(params *do.GetAlbumListRequestEntity) (albumList do.AlbumList, err error) {

	DB.Where("category=?", params.CategoryName).Offset(params.Offset).Limit(params.Limit).Find(&albumList.Items)
	return

}
