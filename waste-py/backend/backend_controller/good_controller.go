package backend_controller

import (
	"errors"
	"waste-py/backend/backend_model"
	"waste-py/backend/backend_server"
)

var gs = backend_server.NewGoodServer()

// 获取所有商品
func GetGoodsAll() (goodsData []backend_model.Good, err error) {
	goodsData,err = gs.QueryGoodAll()
	if err != nil {
		resErr := errors.New("查询所有商品数据失败")
		return nil,resErr
	}
	return
}
