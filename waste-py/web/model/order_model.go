package model

import "waste-py/tools"

type Order struct {
	Id             int64
	OrderId        string         `xorm:"varchar(255)" json:"order_id"` // 订单编号
	GoodId         int64          `xorm:"int" json:"good_id"`           // 商品ID
	GoodName       string         `xorm:"varchar(60)" json:"good_name"`
	GoodPrice      float64        `xorm:"float" json:"good_price"`                 // 商品成交单价
	GoodUnit       string         `xorm:"varchar(30)" json:"goods_unit"`           // 商品单位
	GoodNumber     int            `xorm:"int" json:"good_number"`                  // 商品数量
	NumberImage    string         `xorm:"varchar(255)" json:"number_image"`        // 数量照片
	TotalPrice     float64        `xorm:"float" json:"total_price"`                // 商品总价
	ConsumerId     int64          `xorm:"int" json:"consumer_id"`                  // 消费者ID
	AddressId      int64          `xorm:"int" json:"address_id"`                   // 地址ID
	ConsumerName   string         `xorm:"varchar(50) json:"consumer_name`          // 联系人
	ConsumerMobile string         `xorm:"varchar(11)" json:"consumer_mobile"`      // 联系电话
	ConsumerGender string         `xorm:"varchar(8)" json:"consumer_gender"`       // 联系人性别
	WechatId       string         `xorm:"varchar(50)" json:"consumer_wechar_icon"` // 微信
	Remark         string         `xorm:"varchar(255)" json:"remark"`              // 备注信息
	State          int            `xorm:"int" json:"state"`                        // 订单状态 0: 已完成 1：未完成
	CreatedTime    tools.JsonTime `xorm:"created" json:"created_time"`             // 创建时间
	UpdateTime     tools.JsonTime `xorm:"updated" json:"update_time"`              // 更新时间
}
