package do

import "encoding/json"

type Category struct {
	Category string  `json:"category" gorm:"not null;unique"`
	ID       float64 `json:"id"`
}

func UnmarshalCategoryListEntity(data []byte) (CategoryListEntity, error) {
	var r CategoryListEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CategoryListEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CategoryListEntity struct {
	Categories []Category `json:"categories,omitempty"`
}
