package service

import (
	"fmt"
	"tingfm/dao"
	"tingfm/do"
)

func GetAlbumList( params * do.GetAlbumListRequestEntity)( result string ,err error){

	 albumList ,err:=  dao.GetAlbumList(params)
	if err != nil {
		fmt.Printf("err: %#v",err.Error())
		return
	}

	bytes,err :=  albumList.Marshal()
	if err != nil {
		fmt.Printf("err: %#v",err.Error())
		return
	}
	return string(bytes),err
}
