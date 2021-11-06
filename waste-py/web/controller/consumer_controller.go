package controller

import (
	"errors"
	"waste-py/common"
	"waste-py/param/web_param"
	"waste-py/tools"
	"waste-py/web/model"
	"waste-py/web/server"
)

var cs = server.NewConsumerService()

// 用户注册控制器
func ConsumerSignUpController(param web_param.SignParam, ssCode string) (err error) {
	// 验证 验证码
	if param.CheckEmailCode != ssCode {
		err = errors.New("验证码错了")
		return err
	}
	// 验证密码格式
	checkRes := tools.VerifyPassword(param.Password)
	if checkRes == false {
		err = errors.New("密码由至少8位大写字母、小写字母和数字组成")
		return err
	}
	// 验证 确认密码
	if param.CheckPassword != param.Password {
		err = errors.New("两次密码对不上,再来再来")
		return err
	}
	// 查询用户是否已经注册
	resData, queryErr := cs.QueryConsumerByEmail(param.Email)
	if queryErr != nil {
		err = errors.New("注册服务器忙碌,请重试一下")
		return err
	}
	if resData.Id > 0 {
		err = errors.New("电子邮箱已经注册了")
		return err
	}
	// 加密密码
	password := tools.EncoderSha256(param.Password)
	// 新用户入库
	insertErr := cs.InsertConsumerByPassword(param.Email, password)
	if insertErr != nil {
		err = errors.New("数据库闹情绪了,请稍后重试")
		return err
	}
	return nil
}

// 用户找回密码控制器
func ConsumerResetPassController(param web_param.SignParam, ssCode string) (err error) {
	// 验证 验证码
	if param.CheckEmailCode != ssCode {
		err = errors.New("验证码错了")
		return err
	}
	// 验证密码格式
	checkRes := tools.VerifyPassword(param.Password)
	if checkRes == false {
		err = errors.New("密码由至少8位大写字母、小写字母和数字组成")
		return err
	}
	// 验证 确认密码
	if param.CheckPassword != param.Password {
		err = errors.New("两次密码对不上,再来再来")
		return err
	}
	// 查询用户是否已经注册
	resData, queryErr := cs.QueryConsumerByEmail(param.Email)
	if queryErr != nil {
		err = errors.New("注册服务器忙碌,请重试一下")
		return err
	}
	if resData.Id > 0 {
		// 加密密码
		password := tools.EncoderSha256(param.Password)
		// 修改用户密码
		updateErr := cs.UpdateConsumerPassword(param.Email, password)
		if updateErr != nil {
			err = errors.New("数据库闹情绪了,请稍后重试")
			return err
		}
		return nil
	}
	err = errors.New("该邮箱尚未注册,请确认邮箱是否正确")
	return err
}

// 用户邮箱登录控制器
func ConsumerEmailLoginController(loginParam web_param.LoginParam) (resConsumer *model.Consumers, err error) {
	// 验证验证码
	var code common.VerifyCaptchaBody
	code.Id = loginParam.CheckCodeId
	code.VerifyValue = loginParam.CheckCode
	checkRes := common.VerifyCaptchaStr(code)
	if !checkRes {
		err = errors.New("验证码错误,请重新输入")
		return nil, err
	}
	// 验证用户名密码
	verifyEmailRes := tools.VerifyEmailFormat(loginParam.Email)
	if !verifyEmailRes {
		err = errors.New("邮箱格式不正确")
		return nil, err
	}
	// 验证密码格式
	verifyPasswordRes := tools.VerifyPassword(loginParam.Password)
	if !verifyPasswordRes {
		err = errors.New("密码以字母/数字/下划线_开头,长度8-20位")
		return nil, err
	}
	// 加密密码
	password := tools.EncoderSha256(loginParam.Password)
	// 验证用户密码是否正确
	resData, queryErr := cs.QueryConsumerByEmail(loginParam.Email)
	if queryErr != nil {
		err = errors.New("登录服务器忙碌,请重试一下")
		return nil, err
	}
	if resData.Id == 0 {
		err = errors.New("邮箱暂未注册,请先注册吧")
		return nil, err
	}
	if password != resData.Password {
		err = errors.New("密码错误,登录失败")
		return nil, err
	}
	resConsumer = &resData
	return resConsumer, nil
}

// 用户手机登录控制器
func ConsumerLoginController(param *web_param.LoginParam) (resConsumer *model.Consumers, err error) {
	// 验证 验证码
	if param.CheckCode != "111" {
		err = errors.New("验证码错了,刷新一下")
		return nil, err
	}
	// 验证手机号码格式
	verifyMobileRes := tools.VerifyMobileFormat(param.Mobile)
	if !verifyMobileRes {
		err = errors.New("手机号码错了")
		return nil, err
	}
	// 验证密码格式
	verifyPasswordRes := tools.VerifyPassword(param.Password)
	if !verifyPasswordRes {
		err = errors.New("密码以字母/数字/下划线_开头,长度6-18位")
		return nil, err
	}
	// 加密密码
	password := tools.EncoderSha256(param.Password)
	// 验证用户密码是否正确
	resData, queryErr := cs.QueryConsumerByMobile(param.Mobile)
	if queryErr != nil {
		err = errors.New("登录服务器忙碌,请重试一下")
		return nil, err
	}
	if resData.Id == 0 {
		err = errors.New("手机号码未注册,请先注册吧")
		return nil, err
	}
	if password != resData.Password {
		err = errors.New("密码错误,登录失败")
		return nil, err
	}
	resConsumer = &resData
	return
}

// 修改个人信息
func ConsumerEditInfoController(param web_param.ConsumerEditParam) (err error) {
	return cs.UpdateConsumerInfo(param.Id, param.Name, param.NickName, param.Mobile, param.Gender)
}

// 根据邮箱查询用户
func GetConsumerByEmail(email string) (consumer model.Consumers, err error) {
	return cs.QueryConsumerByEmail(email)
}

// 根据ID 更新/上传头像
func UpdateConsumerAvatarController(id int64, fileName string) (err error) {
	return cs.UpdateConsumerAvatarById(id, fileName)
}

// 根据ID 更新/上传 收款码
func UpdateConsumerWecharController(id int64, fileName string) (err error) {
	return cs.UpdateConsumerWecharById(id, fileName)
}

// 添加地址控制器
func AddAddressController(addressParam web_param.AddressParam) (err error) {
	var id int64 = addressParam.Id
	var address string = addressParam.Address
	var token string = addressParam.Token
	// 解析token
	_, Claims, parseErr := common.ParseToken(token)
	if parseErr != nil {
		err = errors.New("解析TOKEN失败")
		return err
	}
	consumerId := Claims.UserId
	if consumerId != id {
		err = errors.New("验证登录用户失败")
		return err
	}
	resErr := cs.InsertAddress(id, address)
	if resErr != nil {
		err = errors.New("地址入库失败")
		return err
	}
	return nil
}

// 地址列表
func AddressListController(id int64) (addressList []model.Address, err error) {
	return cs.QueryAddressList(id)
}

// 删除地址
func RemoveAddressController(id int64) (err error) {
	return cs.DeleteAddressById(id)
}

// 修改地址
func EditAddressController(id int64, consumerId int64, consumerAddress string) (err error) {
	return cs.EditAddressById(id, consumerId, consumerAddress)
}
