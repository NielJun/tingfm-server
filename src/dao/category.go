package dao

import (
	"tingfm/do"
)

func GetCategories() (categoryListEntity do.CategoryListEntity, err error) {

	DB.Find(&categoryListEntity.Categories)
	return
}

///新建一个Categpry
func NewCategory(categoryName string) (err error) {
	var category = do.Category{
		Category: categoryName,
	}
	DB.Create(&category)

	return nil
}

func DelCategory(categoryName string) (err error) {

	var category = do.Category{
		Category: categoryName,
	}
	DB.Delete(&category)
	return nil
}

func ModCategory(categoryName string) (err error) {

	var category = do.Category{}
	DB.Where("category = ?", categoryName).First(&category)

	category.Category = categoryName
	DB.Save(&category)
	return nil
}
