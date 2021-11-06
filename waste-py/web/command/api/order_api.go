package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"waste-py/common"
	"waste-py/param/web_param"
	"waste-py/web/controller"
)

// 获取数据库城市数据
func GetPositionCities(ctx *gin.Context) {
	res, resErr := controller.GetMysqlCitiesData()
	if resErr != nil {
		common.Failed(ctx, "获取城市数据失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 获取地区数据
func GetAreaList(ctx *gin.Context) {
	res, resErr := controller.GetAreasData()
	if resErr != nil {
		common.Failed(ctx, "获取地区数据失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 获取站点数据
func GetPositionList(ctx *gin.Context) {
	res, resErr := controller.GetPositionsData()
	if resErr != nil {
		common.Failed(ctx, "获取站点数据失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 生成订单
func PostOrder(ctx *gin.Context) {
	var orderParamJson web_param.OrderInsertParam
	err := ctx.ShouldBindJSON(&orderParamJson)
	if err != nil {
		common.Failed(ctx, "获取订单数据失败")
		return
	}
	orderInfoParam := orderParamJson.OrderInfoJson
	orderId, id, resErr := controller.AddOrder(orderInfoParam)
	if resErr != nil {
		common.Failed(ctx, "插入订单失败")
		return
	}
	fmt.Println("id", id)
	common.Success(ctx, orderId)
	return
}

// 订单列表
func PostOrderList(ctx *gin.Context) {
	var consumerParam web_param.ConsumerId
	err := ctx.ShouldBind(&consumerParam)
	if err != nil {
		common.Failed(ctx, "获取用户ID失败")
		return
	}
	idStr := consumerParam.IdStr
	id, _ := strconv.ParseInt(idStr, 10, 64)
	res, resErr := controller.GetOrderList(id)
	if resErr != nil {
		common.Failed(ctx, "获取订单数据失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 未完成订单
func PostOrderNotdoneList(ctx *gin.Context) {
	var consumerParam web_param.ConsumerId
	err := ctx.ShouldBind(&consumerParam)
	if err != nil {
		common.Failed(ctx, "获取用户ID失败")
		return
	}
	idStr := consumerParam.IdStr
	id, _ := strconv.ParseInt(idStr, 10, 64)
	res, resErr := controller.GetOrderNotdoneListController(id)
	if resErr != nil {
		common.Failed(ctx, "获取订单数据失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 已完成订单
func PostOrderDoneList(ctx *gin.Context) {
	var consumerParam web_param.ConsumerId
	err := ctx.ShouldBind(&consumerParam)
	if err != nil {
		common.Failed(ctx, "获取用户ID失败")
		return
	}
	idStr := consumerParam.IdStr
	id, _ := strconv.ParseInt(idStr, 10, 64)
	res, resErr := controller.GetOrderDoneListController(id)
	if resErr != nil {
		common.Failed(ctx, "获取订单数据失败")
		return
	}
	common.Success(ctx, res)
	return
}
