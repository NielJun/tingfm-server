package service

import (
	"fmt"
	"tingfm/dao"
	"tingfm/do"
)

func NewAlbum(album *do.Album) (err error) {
	return dao.NewAlbum(album)
}

func SearchAlbum(albumName string) (result string, err error) {

	album, err := dao.SearchAlbum(albumName)
	if err != nil {
		fmt.Printf("album  is not value able, err: %#v", err.Error())
		return
	}

	bytes, err := album.Marshal()
	return string(bytes), err
}

func DelAlbum(album string) (err error) {

	return dao.DelAlbum(album)
}

func AddALbumPlayedTime(albumName string) (err error) {

	return dao.AddALbumPlayedTime(albumName)
}
