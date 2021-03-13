package do

import "encoding/json"


////请求


func UnmarshalGetPlayersByCategoryRequestEntity(data []byte) (GetPlayersByCategoryRequestEntity, error) {
	var r GetPlayersByCategoryRequestEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetPlayersByCategoryRequestEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetPlayersByCategoryRequestEntity struct {
	Schema     string     `json:"$schema"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	CategoryName CategoryName `json:"category_name"`
}

type CategoryName struct {
	Type string `json:"type"`
}




func UnmarshalGetPlayersByCategoryResponseEntity(data []byte) (GetPlayersByCategoryResponseEntity, error) {
	var r GetPlayersByCategoryResponseEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetPlayersByCategoryResponseEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetPlayersByCategoryResponseEntity struct {
	Schema     string        `json:"$schema"`
	Type       string        `json:"type"`
	Properties ResponseProperties    `json:"properties"`
	Required   []interface{} `json:"required"`
}

type ResponseProperties struct {
	PlayerNameArray PlayerNameArray `json:"player_name_array"`
}

type PlayerNameArray struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}

type Items struct {
	Type string `json:"type"`
}

