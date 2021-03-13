package dao

import "tingfm/do"

func GetPlayerListByCategories(categoryName string) (playerListResponse *do.GetPlayersByCategoryResponseEntity, err error) {

	DB.Where("category=?", categoryName).Offset(0).Limit(30).Find(&playerListResponse)


	return nil, err
}
