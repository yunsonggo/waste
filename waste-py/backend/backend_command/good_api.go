package backend_command

import (
	"github.com/gin-gonic/gin"
	"waste-py/backend/backend_controller"
	"waste-py/common"
)

// 获取所有商品
func GetGoods(ctx *gin.Context) {
	res,resErr := backend_controller.GetGoodsAll()
	if resErr != nil {
		common.Failed(ctx,resErr)
		return
	}
	common.Success(ctx,res)
}