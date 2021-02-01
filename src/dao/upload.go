package dao

import "tingfm/do"

func UploadAlbum(album *do.Album) (err error) {

	var hadRecord = DB.NewRecord(&album)

	if hadRecord {

	}

	return nil
}
