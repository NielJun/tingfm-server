package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tingfm/do"
	"tingfm/service"
)

/// 发送今天的赚钱的状态
func PostTodayUserMoneyHandle(c *gin.Context) {

	var params do.Money

	err := c.Bind(&params)
	if err != nil {
		fmt.Printf("album name params is not value able, err: %#v", err.Error())
		return
	}

	err = service.PostTodayUserMoney(&params)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, "post field")
		return
	}

	c.JSON(http.StatusOK, "post success")

}

///获取当天的所有资金情况
func GetTodayUserMoneyHandle(c *gin.Context) {
	result, err := service.GetTodayUserMoney()
	if err != nil {
		c.JSON(fmt.Printf("get faild,err is %#v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func GetSomeDayUserMoneyHandle(c *gin.Context) {

	var param do.SomeDayParams

	err := c.Bind(&param)
	if err != nil {
		return
	}

	result, err := service.GetSomeDayUserMoney(param.Params)
	if err != nil {
		c.JSON(fmt.Printf("get faild,err is %#v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
	return
}
