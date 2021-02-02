package do

import "encoding/json"

func UnmarshalSearchResponseAlbumListEntity(data []byte) (SearchResponseAlbumListEntity, error) {
	var r SearchResponseAlbumListEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchResponseAlbumListEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchResponseAlbumListEntity struct {
	Albums []Album `json:"albums,omitempty"`
}
