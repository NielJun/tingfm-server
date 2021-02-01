package do

import "encoding/json"

type GetAlbumListRequestEntity struct {
	CategoryName string `json:"category_name"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}


func UnmarshalAlbumList(data []byte) (AlbumList, error) {
	var r AlbumList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlbumList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AlbumList struct {
	Items []Album `json:"items,omitempty"`
}

