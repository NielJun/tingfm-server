package do

import "encoding/json"

func UnmarshalAlbumListResponseEntity(data []byte) (AlbumListResponseEntity, error) {
	var r AlbumListResponseEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlbumListResponseEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AlbumListResponseEntity struct {
	Albums          []Album `json:"albums,omitempty" `
	RecommendAlbums []Album `json:"recommend_albums,omitempty"`
}

type GetAlbumListRequestEntity struct {
	CategoryName string `json:"category_name"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}
