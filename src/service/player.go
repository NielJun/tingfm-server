package service

import (
	"fmt"
	"tingfm/dao"
)

func GetPlayerListByCategories(categoryName string) (result string, err error) {

	playerListResponse, err := dao.GetPlayerListByCategories(categoryName)
	if err != nil {
		fmt.Printf("playerList is not value able, err: %#v", err.Error())
		return
	}

	bytes, err := playerListResponse.Marshal()
	return string(bytes), err
}
