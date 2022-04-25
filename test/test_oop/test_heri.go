package main

/**
测试继承
*/
import (
	"fmt"
	"strconv"
)

// 定义一个接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone3 struct {
	name     string
	quantity int
	price    int
}

func (phone Phone3) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Phone3) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gitft FreeGift) settleAccount() int {
	return 0
}

func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calALprice(goods []Good) int {
	var allPrice int
	for index, good := range goods {
		fmt.Println(index)
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
func main() {
	iPhone := Phone3{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calALprice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
