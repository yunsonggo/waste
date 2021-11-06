package model

import (
	"waste-py/tools"
)

type Consumers struct {
	// 如果此处 `xorm:"id"`，则插入数据的时候，会默认为0，插入成功后不会把新插入的id返回，如果想得到插入后的主键id，则id不需要写`xorm:"id"`
	Id          int64          `xorm:"pk autoincr" json:"id"`                // id
	Name        string         `xorm:"varchar(20)" json:"consumer_name"`     // 真实姓名
	NickName    string         `xorm:"varchar(30)" json:"consumer_nickname"` // 昵称
	Mobile      string         `xorm:"varchar(20)" json:"consumer_mobile"`   // 手机号码
	Email       string         `xorm:"varchar(50)" json:"consumer_email"`
	Password    string         `xorm:"varchar(255)" json:"consumer_password"` // 密码
	Icon        string         `xorm:"varchar(255)" json:"consumer_icon"`     // 头像
	WecharIcon  string         `xorm:"varchar(255)" json:"consumer_wechar_icon"`
	Gender      string         `xorm:"varchar(10)" json:"consumer_gender"` // 性别
	CreatedTime tools.JsonTime `xorm:"created" json:"created_time"`        // 创建时间
}
