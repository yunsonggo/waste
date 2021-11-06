package server

import (
	"waste-py/dal"
	"waste-py/web/model"
)

type ConsumerServer interface {
	// 根据邮箱查询用户
	QueryConsumerByEmail(email string) (consumer model.Consumers, err error)
	// 根据手机号查询用户
	QueryConsumerByMobile(mobile string) (consumer model.Consumers, err error)
	// 新用户入库
	InsertConsumerByPassword(email string, password string) (err error)
	// 根据邮箱修改密码
	UpdateConsumerPassword(email string, password string) (err error)
	// 修改个人信息
	UpdateConsumerInfo(id int64, name string, nickname string, mobile string, gender string) (err error)
	// 根据ID 更新/上传头像
	UpdateConsumerAvatarById(id int64, fileName string) (err error)
	// 根据ID 更新/上传 收款码
	UpdateConsumerWecharById(id int64, fileName string) (err error)
	// 用户添加地址
	InsertAddress(id int64, address string) (err error)
	// 用户列表
	QueryAddressList(id int64) (addressList []model.Address, err error)
	// 删除地址
	DeleteAddressById(id int64) (err error)
	// 修改地址
	EditAddressById(id int64, consumerId int64, consumerAddress string) (err error)
}

type ConsumerService struct{}

func NewConsumerService() ConsumerServer {
	return &ConsumerService{}
}

// 根据手机号查询用户
func (cs *ConsumerService) QueryConsumerByMobile(mobile string) (consumer model.Consumers, err error) {
	_, err = dal.DB.Where("mobile = ?", mobile).Get(&consumer)
	return
}

// 新用户入库
func (cs *ConsumerService) InsertConsumerByPassword(email string, password string) (err error) {
	var consumer model.Consumers
	consumer.Email = email
	consumer.Password = password
	_, err = dal.DB.InsertOne(&consumer)
	return
}

// 根据电子邮箱查询用户
func (cs *ConsumerService) QueryConsumerByEmail(email string) (consumer model.Consumers, err error) {
	_, err = dal.DB.Where("email = ?", email).Get(&consumer)
	return
}

// 根据邮箱修改密码
func (cs *ConsumerService) UpdateConsumerPassword(email string, password string) (err error) {
	consumer := new(model.Consumers)
	consumer.Password = password
	_, err = dal.DB.Where("email = ?", email).Update(consumer)
	return
}

// 修改个人信息
func (cs *ConsumerService) UpdateConsumerInfo(id int64, name string, nickname string, mobile string, gender string) (err error) {
	consumer := new(model.Consumers)
	consumer.Name = name
	consumer.NickName = nickname
	consumer.Mobile = mobile
	consumer.Gender = gender
	_, err = dal.DB.Where("id = ?", id).Update(consumer)
	return
}

// 根据ID 更新/上传头像
func (cs *ConsumerService) UpdateConsumerAvatarById(id int64, fileName string) (err error) {
	consumer := new(model.Consumers)
	consumer.Icon = fileName
	_, err = dal.DB.Where("id = ?", id).Update(consumer)
	return
}

// 根据ID 更新/上传 收款码
func (cs *ConsumerService) UpdateConsumerWecharById(id int64, fileName string) (err error) {
	consumer := new(model.Consumers)
	consumer.WecharIcon = fileName
	_, err = dal.DB.Where("id = ?", id).Update(consumer)
	return
}

// 用户添加地址
func (cs *ConsumerService) InsertAddress(id int64, address string) (err error) {
	newAddress := new(model.Address)
	newAddress.ConsumerId = id
	newAddress.ConsumerAddress = address
	_, err = dal.DB.Insert(newAddress)
	return
}

// 用户列表
func (cs *ConsumerService) QueryAddressList(id int64) (addressList []model.Address, err error) {
	err = dal.DB.Where("consumer_id = ?", id).Find(&addressList)
	return
}

// 删除地址
func (cs *ConsumerService) DeleteAddressById(id int64) (err error) {
	address := new(model.Address)
	_, err = dal.DB.Where("id = ?", id).Delete(address)
	return
}

// 修改地址
func (cs *ConsumerService) EditAddressById(id int64, consumerId int64, consumerAddress string) (err error) {
	address := new(model.Address)
	address.ConsumerAddress = consumerAddress
	_, err = dal.DB.Where("id = ? and consumer_id = ?", id, consumerId).Update(address)
	return
}
