package backend_server

import (
	"waste-py/backend/backend_model"
	"waste-py/dal"
)

type Gooder interface {
	// 获取所有商品
	QueryGoodAll() (goodsData []backend_model.Good, err error)
}

type GoodServer struct {}

func NewGoodServer() Gooder {
	return &GoodServer{}
}

// 获取所有商品
func (gs *GoodServer) QueryGoodAll() (goodsData []backend_model.Good, err error) {
	err = dal.DB.Find(&goodsData)
	return
}