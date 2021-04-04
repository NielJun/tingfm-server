package do

import "encoding/json"

func UnmarshalMoney(data []byte) (Money, error) {
	var r Money
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Money) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Money struct {
	UserName    string  `json:"user_name"`
	RewardTime  int64   `json:"reward_time"`
	ExpectMoney float64 `json:"expect_money"`
	ID          int64   `json:"id"`
}

type MoneyList struct {
	Moneys []Money `json:"moneys"`
	TotalMoney float64 `json:"total_money"`
}

func (ml *MoneyList) Marshal() ([]byte, error) {
	return json.Marshal(ml)
}

func UnMarshalMoneyList(data []byte) (MoneyList, error) {
	var r MoneyList
	err := json.Unmarshal(data, r)
	return r, err
}

type SomeDayParams struct {
	Params string `json:"time"`
}

func (ml *SomeDayParams) Marshal() ([]byte, error) {
	return json.Marshal(ml)
}

func UnMarshalSomeDayParam(data []byte) (SomeDayParams, error) {
	var r SomeDayParams
	err := json.Unmarshal(data, r)
	return r, err
}
