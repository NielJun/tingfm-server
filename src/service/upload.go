package service

import (
	"tingfm/dao"
	"tingfm/do"
)

func UploadAlbum() {

	dao.UploadAlbum(&do.Album{
		AlbumName:     "三国演义",
		AlbumPlayer:   "袁阔成",
		AlbumImageUrl: "三国演义.png",
		Description:   "三国演义",
		TotalAudioCount:    365,
	})

}
