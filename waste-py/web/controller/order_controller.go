package controller

import (
	"fmt"
	"math/rand"
	"time"
	"waste-py/param/web_param"
	"waste-py/web/model"
	"waste-py/web/server"
)

var os = server.NewOrderService()

// 获取数据库城市数据
func GetMysqlCitiesData() (citiesData []model.Cities, err error) {
	return os.QueryMysqlCities()
}

// 获取地区数据
func GetAreasData() (areaData []model.Area, err error) {
	return os.QueryAreas()
}

// 获取站点数据
func GetPositionsData() (positionsData []model.Position, err error) {
	return os.QueryPositions()
}

// 添加订单
func AddOrder(orderInfoParam web_param.OrderInfo) (orderId string, id int64, err error) {
	timeStr := time.Now().Format("20060102150405")
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	orderId = timeStr + code
	orderData := new(model.Order)
	orderData.OrderId = orderId
	orderData.GoodId = orderInfoParam.GoodId
	orderData.GoodName = orderInfoParam.GoodName
	orderData.GoodUnit = orderInfoParam.GoodUnit
	orderData.GoodPrice = orderInfoParam.GoodPrice
	orderData.ConsumerId = orderInfoParam.ConsumerId
	orderData.AddressId = orderInfoParam.AddressId
	orderData.ConsumerName = orderInfoParam.ConsumerName
	orderData.ConsumerMobile = orderInfoParam.ConsumerMobile
	orderData.ConsumerGender = orderInfoParam.ConsumerGender
	orderData.WechatId = orderInfoParam.WechatId
	orderData.Remark = orderInfoParam.Remark
	orderData.State = 1
	res, resErr := os.InsertOrder(orderData)
	return orderId, res, resErr
}

// 获取订单数据
func GetOrderList(id int64) (orderData []model.Order, err error) {
	return os.QueryOrderList(id)
}

// 获取未完成订单
func GetOrderNotdoneListController(id int64) (orderData []model.Order, err error) {
	return os.QueryOrderNotDoneList(id)
}

// 获取已完成订单
func GetOrderDoneListController(id int64) (orderData []model.Order, err error) {
	return os.QueryOrderDoneList(id)
}
