package do

import "encoding/json"

//专辑相关的数据结构
type Album struct {
	ID              uint   `json:"id"`                                 //ID名字
	AlbumName       string `json:"album_name" gorm:"not null;unique"`  //专辑名字
	AlbumPlayer     string `json:"album_player"`                       //专辑的演播者
	AlbumImageUrl   string `json:"album_image_url"`                    //图片的地址
	Description     string `json:"description"`                        //描述
	TotalAudioCount uint   `json:"total_audio_count" gorm:"default:0"` //总的音频数目
	Category        string `json:"category"`                           //当前的播放类型
	StateType       uint   `json:"state_type"`                         //专辑的状态 0: 免费 1：vip 2：精品 3:推荐
	PlayedTime      int64  `json:"played_time"`                        //播放的次数
	StarNum         uint   `json:"star_num"`                           //评分星数
}

func UnmarshalAlbum(data []byte) (Album, error) {
	var r Album
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Album) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NewAlbumRequestEntity struct {
	AlbumName string `json:"album_name"`
}
