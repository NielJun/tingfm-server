package do

import "encoding/json"

//专辑相关的数据结构
type Album struct {
	ID              uint
	AlbumName       string `json:"album_name" gorm:"not null;unique"`
	AlbumPlayer     string `json:"album_player"`
	AlbumImageUrl   string `json:"album_image_url"`
	Description     string `json:"description"`
	TotalAudioCount uint   `json:"total_audio_count" gorm:"default:0"`
	Category        string `json:"category"`
}

func UnmarshalAlbum(data []byte) (Album, error) {
	var r Album
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Album) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//////
