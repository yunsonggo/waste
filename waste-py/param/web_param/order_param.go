package web_param

type OrderInsertParam struct {
	OrderInfoJson OrderInfo `json:"order_info"`
}

type OrderInfo struct {
	GoodId         int64   `json:"good_id"` // 商品ID
	GoodName       string  `json:"good_name"`
	GoodUnit       string  `json:"goods_unit"`
	GoodPrice      float64 `json:"good_price"`  // 商品成交单价
	ConsumerId     int64   `json:"consumer_id"` // 消费者ID
	CityId         int64   `json:"cityid"`
	AreaId         int64   `json:"areaid"`
	PositionId     int64   `json:"positionid"`
	AddressId      int64   `json:"address_id"`           // 地址ID
	ConsumerName   string  `json:"consumer_name"`        // 联系人
	ConsumerMobile string  `json:"consumer_mobile"`      // 联系电话
	ConsumerGender string  `json:"consumer_gender"`      // 联系人性别
	WechatId       string  `json:"consumer_wechar_icon"` // 微信
	Remark         string  `json:"remark"`               // 备注信息
}
