package dao

import (
	"tingfm/do"
)

func NewAlbum(album *do.Album) (err error) {
	DB.Create(album)
	return nil
}

func SearchAlbum(albumName string) (album *do.Album, err error) {

	album = &do.Album{}

	DB.Where("album_name =?", albumName).First(album)

	return
}

func DelAlbum(albumName string) (err error) {

	var album = do.Album{
		AlbumName: albumName,
	}
	DB.Delete(&album)
	return nil
}

