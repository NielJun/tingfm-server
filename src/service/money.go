package service

import (
	"fmt"
	"tingfm/dao"
	"tingfm/do"
)

func PostTodayUserMoney(postExpectMoney *do.Money) (err error) {

	err = dao.PostTodayUserMoney(postExpectMoney)
	if err != nil {
		fmt.Printf("playerList is not value able, err: %#v", err.Error())
		return
	}
	return nil
}

///获取当天的所有资金情况
func GetTodayUserMoney() (result string, err error) {

	moneys, err := dao.GetTodayUserMoney()
	if err != nil {
		return
	}

	var totalMoney float64

	for _, v := range moneys {
		totalMoney += v.ExpectMoney
	}

	moneyList := do.MoneyList{
		Moneys:     moneys,
		TotalMoney: totalMoney,
	}

	bytes, err := moneyList.Marshal()

	if err != nil {
		return
	}

	return string(bytes), err
}

func GetSomeDayUserMoney(timeF string) (result string, err error) {
	moneys, err := dao.GetSomeDayUserMoney(timeF)
	if err != nil {
		return
	}

	moneyList := do.MoneyList{
		Moneys: moneys,
	}
	bytes, err := moneyList.Marshal()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
