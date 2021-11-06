package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waste-py/backend/backend_command"
	"waste-py/common"
	"waste-py/middleware"
	"waste-py/web/command/api"
)

func NewRouter() *gin.Engine {
	// 初始化引擎
	app := gin.Default()
	// 设置静态文件
	app.StaticFS("/api/static", http.Dir("./public"))
	// 启用sessions
	common.InitSession(app)
	// 启用cors
	app.Use(middleware.Cors())
	// web接口
	webGroup := app.Group("/api")
	{
		// 请求 邮箱验证码
		webGroup.POST("/consumer/email/captcha", api.PostEmailCode)
		// 请求base64验证码
		webGroup.GET("/consumer/captcha", api.GetCaptcha)
		// 邮箱注册
		webGroup.POST("/consumer/email/signup", api.ConsumerSignUpByEmail)
		// 邮箱登录
		webGroup.POST("/consumer/email/login", api.PostLoginByEmail)
		// 邮箱退出登录
		webGroup.POST("/consumer/email/logout", api.PostLogoutByEmail)
		// 邮箱找回密码
		webGroup.POST("/consumer/email/reset/pass", api.ConsumerResetPassByEmail)
		// 修改个人信息
		webGroup.POST("/consumer/edit/info", api.PostConsumerEdit)
		// 获取IP定位
		webGroup.GET("/consumer/ip/address", api.GetAddrByIp)
		// 获取城市数据
		webGroup.GET("/consumer/cities", api.GetCityInfo)
		// 消费者 注册接口
		webGroup.POST("/consumer/signup", api.ConsumerSignUpApi)
		// 消费者 登录接口
		webGroup.POST("/consumer/login", api.ConsumerLoginApi)
		// 消费者更新头像
		webGroup.POST("/consumer/avatar/upload", api.PostUploadAvatar)
		// 消费者更新收款码
		webGroup.POST("/consumer/wechar/icon/upload", api.PostUploadWechar)
		// 消费者 添加地址
		webGroup.POST("/consumer/add/address", api.PostAddAddress)
		// 消费者 地址列表
		webGroup.POST("/consumer/address/list", api.PostAddressList)
		// 消费者 删除地址
		webGroup.POST("/consumer/address/remove", api.PostRemoveAddress)
		// 消费者 修改地址
		webGroup.POST("/consumer/address/edit", api.PostEditAddress)
		// 获取所有商品数据
		webGroup.GET("/good/list", backend_command.GetGoods)
		// 获取数据库城市数据
		webGroup.GET("/position/cities", api.GetPositionCities)
		// 获取数据库地区数据
		webGroup.GET("/position/area/list", api.GetAreaList)
		// 获取数据库站点数据
		webGroup.GET("/position/list", api.GetPositionList)
		// 生成订单
		webGroup.POST("/consumer/order/insert", api.PostOrder)
		// 获取订单列表
		webGroup.POST("/order/list", api.PostOrderList)
		// 获取未完成订单
		webGroup.POST("/order/notdone/list", api.PostOrderNotdoneList)
		// 获取已完成订单
		webGroup.POST("/order/done/list", api.PostOrderDoneList)
		// 上传excel
		webGroup.POST("/excel/upload", backend_command.UploadExcel)
	}
	return app
}
