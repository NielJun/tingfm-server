package dao

import (
	"time"
	"tingfm/do"
)

func PostTodayUserMoney(money *do.Money) (err error) {

	//先查询是否有今天的记录 如果有 则删除
	// 如果没有则进行插入新的

	moneyList := make([]do.Money, 0)

	DB.Debug().Where("user_name = ? and to_days(`time`) = to_days(?)", money.UserName, time.Now().Format("2006-01-02 15:04:05")).Find(&moneyList)

	if len(moneyList) > 0 {
		for _,v := range moneyList {
			DB.Delete(&v)
		}
	}

	DB.Create(money)
	return nil
}

///获取当天的所有资金情况
func GetTodayUserMoney() (moneyList []do.Money, err error) {
	moneyList, err = GetSomeDayUserMoney(time.Now().Format("2006-01-02 15:04:05"))
	return
}

func GetSomeDayUserMoney(timeF string) (moneys []do.Money, err error) {
	DB.Debug().Where("to_days(`time`) = to_days(?)", timeF).Find(&moneys)
	return
}
