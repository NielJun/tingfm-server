package service

import (
	"fmt"
	"tingfm/dao"
)

func GetCategories() (result string, err error) {
	categoryListEntity, err := dao.GetCategories()

	if err != nil {
		fmt.Printf("err: %#v", err.Error())
	}

	bytes, err := categoryListEntity.Marshal()

	if err != nil {
		fmt.Printf("err: %#v", err.Error())
	}
	return string(bytes), nil
}

func NewCategory( categoryName string)(err error)  {
	return dao.NewCategory(categoryName)
}

func DelCategory( categoryName string)(err error)  {
	return dao.DelCategory(categoryName)
}

func ModCategory( categoryName string)(err error)  {
	return dao.ModCategory(categoryName)
}