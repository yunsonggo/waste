package server

import (
	"waste-py/dal"
	"waste-py/web/model"
)

type OrderInterface interface {
	// 获取数据库城市数据
	QueryMysqlCities() (citiesData []model.Cities, err error)
	// 获取地区数据
	QueryAreas() (areaData []model.Area, err error)
	// 获取站点数据
	QueryPositions() (positionsData []model.Position, err error)
	// 插入订单
	InsertOrder(orderData *model.Order) (id int64, err error)
	// 获取订单数据
	QueryOrderList(id int64) (orderData []model.Order, err error)
	// 获取未完成订单数据
	QueryOrderNotDoneList(id int64) (orderData []model.Order, err error)
	// 获取已完成订单数据
	QueryOrderDoneList(id int64) (orderData []model.Order, err error)
}

type OrderService struct{}

func NewOrderService() OrderInterface {
	return &OrderService{}
}

// 获取数据库城市数据
func (os *OrderService) QueryMysqlCities() (citiesData []model.Cities, err error) {
	err = dal.DB.Find(&citiesData)
	return
}

// 获取地区数据
func (os *OrderService) QueryAreas() (areaData []model.Area, err error) {
	err = dal.DB.Find(&areaData)
	return
}

// 获取站点数据
func (os *OrderService) QueryPositions() (positionsData []model.Position, err error) {
	err = dal.DB.Find(&positionsData)
	return
}

// 插入订单
func (os *OrderService) InsertOrder(orderData *model.Order) (id int64, err error) {
	return dal.DB.Insert(orderData)
}

// 获取订单数据
func (os *OrderService) QueryOrderList(id int64) (orderData []model.Order, err error) {
	err = dal.DB.Where("consumer_id = ?", id).Find(&orderData)
	return
}

// 获取未完成订单数据
func (os *OrderService) QueryOrderNotDoneList(id int64) (orderData []model.Order, err error) {
	err = dal.DB.Where("consumer_id = ? and state = 1", id).Find(&orderData)
	return
}

// 获取已完成订单数据
func (os *OrderService) QueryOrderDoneList(id int64) (orderData []model.Order, err error) {
	err = dal.DB.Where("consumer_id = ? and state = 0", id).Find(&orderData)
	return
}
