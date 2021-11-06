package model

type Address struct {
	Id              int64
	ConsumerId      int64  `xorm:"varchar(20)" json:"consumer_id"`
	ConsumerAddress string `xorm:"varchar(255)" json:"consumer_Address"`
}
