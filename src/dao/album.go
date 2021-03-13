package dao

import (
	"fmt"
	"math/rand"
	"tingfm/do"
)

func NewAlbum(album *do.Album) (err error) {
	DB.Create(album)
	return nil
}

///模糊搜寻
func SearchAlbum(albumName string) (searchResponse *do.SearchResponseAlbumListEntity, err error) {

	searchResponse = &do.SearchResponseAlbumListEntity{Albums: make([]do.Album, 0, 10)}

	var param = fmt.Sprintf("%%%s%%", albumName)

	DB.Where("album_name LIKE ?", param).Find(&searchResponse.Albums)

	return
}

func DelAlbum(albumName string) (err error) {

	var album = do.Album{
		AlbumName: albumName,
	}
	DB.Delete(&album)
	return nil
}

/// 增加播放次数
func AddALbumPlayedTime(albumName string) (err error) {

	album := do.Album{}

	DB.Where("album_name = ?", albumName).Take(&album)

	//修改food模型的值

	var randomValue = rand.Int63n(10)

	album.PlayedTime += randomValue

	DB.Save(&album)

	return nil
}
