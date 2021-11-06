package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"waste-py/common"
	"waste-py/param/web_param"
	"waste-py/web/controller"
)

// 发送邮箱验证码
func PostEmailCode(ctx *gin.Context) {
	var emailCodeParam web_param.EmailCodeParam
	err := ctx.ShouldBind(&emailCodeParam)
	if err != nil {
		common.Failed(ctx, "获取注册邮箱失败")
		return
	}
	var toEmail = emailCodeParam.Email
	// 发送邮箱验证码
	theCode, err := common.SendEmailCode(toEmail)
	if err != nil {
		common.Failed(ctx, "发送验证码错误,请重试")
		return
	}
	// 设置session
	ssErr := common.SetSess(ctx, "email_code_"+toEmail, theCode, 180)
	if ssErr != nil {
		common.Failed(ctx, "设置session错误")
		return
	}
	common.Success(ctx, "发送邮箱验证码成功!")
	return
}

// 消费者 邮箱注册
func ConsumerSignUpByEmail(ctx *gin.Context) {
	var signUpParam web_param.SignParam
	err := ctx.ShouldBind(&signUpParam)
	if err != nil {
		common.Failed(ctx, "获取注册参数错误,请重试!")
		return
	}
	// 用户验证 并注册
	toEmail := signUpParam.Email
	ssCodeData := common.GetSess(ctx, "email_code_"+toEmail)
	if ssCodeData == nil {
		common.Failed(ctx, "验证码已失效,请重新获取验证码!")
		return
	}
	ssCode := ssCodeData.(string)
	resErr := controller.ConsumerSignUpController(signUpParam, ssCode)
	if resErr != nil {
		common.Failed(ctx, resErr.Error())
		return
	}
	common.Success(ctx, "注册成功")
	return
}

// 消费者 邮箱找回密码
func ConsumerResetPassByEmail(ctx *gin.Context) {
	var signUpParam web_param.SignParam
	err := ctx.ShouldBind(&signUpParam)
	if err != nil {
		common.Failed(ctx, "获取表单参数错误,请重试!")
		return
	}
	// 用户验证 并注册
	toEmail := signUpParam.Email
	ssCodeData := common.GetSess(ctx, "email_code_"+toEmail)
	if ssCodeData == nil {
		common.Failed(ctx, "验证码已失效,请重新获取验证码!")
		return
	}
	ssCode := ssCodeData.(string)
	resErr := controller.ConsumerResetPassController(signUpParam, ssCode)
	if resErr != nil {
		common.Failed(ctx, resErr.Error())
		return
	}
	common.Success(ctx, "重置密码成功")
	return
}

// 发送base64验证码
func GetCaptcha(ctx *gin.Context) {
	res, err := common.GenerateCaptchaStr()
	if err != nil {
		common.Failed(ctx, "生成验证码错误")
		return
	}
	common.Success(ctx, res)
	return
}

// 消费者 邮箱登录接口
func PostLoginByEmail(ctx *gin.Context) {
	var loginParam web_param.LoginParam
	err := ctx.ShouldBind(&loginParam)
	if err != nil {
		common.Failed(ctx, "获取登录参数失败")
		return
	}
	// 验证
	res, resErr := controller.ConsumerEmailLoginController(loginParam)
	if resErr != nil {
		common.Failed(ctx, resErr.Error())
		return
	}
	res.Password = ""
	// 生成JWT TOKEN
	jwtToken, jwtErr := common.ReleaseToken(res.Id)
	if jwtErr != nil {
		common.Failed(ctx, "发放TOKEN失败,请重新登录")
		return
	}
	// 设置SESSION
	idStr := strconv.FormatInt(res.Id, 10)
	sessErr := common.SetSess(ctx, "consumer_"+idStr, res.Id, 3600*24*7)
	if sessErr != nil {
		common.Failed(ctx, "设置SESSION错误")
		return
	}
	res.Password = jwtToken
	common.Success(ctx, res)
	return
}

// 消费者 邮箱 退出登录
func PostLogoutByEmail(ctx *gin.Context) {
	var consumerId web_param.ConsumerId
	_ = ctx.ShouldBind(&consumerId)
	idStr := consumerId.IdStr
	_ = common.SetSess(ctx, "consumer_"+idStr, 0, -1)
	common.Success(ctx, "退出登录成功")
}

// 消费者 修改个人信息
func PostConsumerEdit(ctx *gin.Context) {
	var editParam web_param.ConsumerEditParam
	err := ctx.ShouldBind(&editParam)
	if err != nil {
		common.Failed(ctx, "获取修改参数失败")
		return
	}
	err = controller.ConsumerEditInfoController(editParam)
	if err != nil {
		common.Failed(ctx, "修改出现错误")
		return
	}
	res, err := controller.GetConsumerByEmail(editParam.Email)
	if err != nil {
		common.Failed(ctx, "获取用户信息失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 消费者 头像上传
func PostUploadAvatar(ctx *gin.Context) {
	_, consumerAvatarFileHeader, aErr := ctx.Request.FormFile("consumer_icon")
	if consumerAvatarFileHeader != nil {
		if aErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		//设置 更新 头像 旧数据 用于删除 旧图片
		oldAvatarUrl := ctx.Request.PostFormValue("old_avatar_url")
		if oldAvatarUrl != "" {
			oldFilePath := "public/" + oldAvatarUrl
			_ = os.Remove(oldFilePath)
		}
		//随机数
		code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		extString := filepath.Ext(consumerAvatarFileHeader.Filename)
		fileName := "upload/consumer/" + "avatar" + strconv.FormatInt(time.Now().Unix(), 10) + code + extString
		filePath := "public/" + fileName
		err := ctx.SaveUploadedFile(consumerAvatarFileHeader, filePath)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		// 设置 更新头像 ID
		editAvatarConsumerId := ctx.Request.PostFormValue("consumer_id")
		id, _ := strconv.ParseInt(editAvatarConsumerId, 10, 64)
		editErr := controller.UpdateConsumerAvatarController(id, fileName)
		if editErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"path": fileName,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"path": "",
	})
	return
}

// 消费者 收款码上传
func PostUploadWechar(ctx *gin.Context) {
	// 验证TOKEN
	consumerToken := ctx.Request.PostFormValue("consumer_token")
	// 解析token
	_, Claims, parseErr := common.ParseToken(consumerToken)
	if parseErr != nil {
		err := errors.New("解析TOKEN失败")
		common.Failed(ctx, err.Error())
		return
	}
	consumerParseId := Claims.UserId
	// 设置 更新头像 ID
	editWecharConsumerId := ctx.Request.PostFormValue("consumer_id")
	id, _ := strconv.ParseInt(editWecharConsumerId, 10, 64)
	if consumerParseId != id {
		parseErr := errors.New("验证登录用户失败")
		common.Failed(ctx, parseErr.Error())
		return
	}

	// 解析并保存收款码
	_, consumerWecharFileHeader, aErr := ctx.Request.FormFile("consumer_wechar_icon")
	if consumerWecharFileHeader != nil {
		if aErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		//设置 更新 收款码 旧数据 用于删除 旧图片
		oldWecharUrl := ctx.Request.PostFormValue("old_wechar_url")
		if oldWecharUrl != "" {
			oldFilePath := "public/" + oldWecharUrl
			_ = os.Remove(oldFilePath)
		}
		//随机数
		code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		extString := filepath.Ext(consumerWecharFileHeader.Filename)
		fileName := "upload/consumer/wechar/" + "wechar" + strconv.FormatInt(time.Now().Unix(), 10) + code + extString
		filePath := "public/" + fileName
		err := ctx.SaveUploadedFile(consumerWecharFileHeader, filePath)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		// 上面解析的ID
		editErr := controller.UpdateConsumerWecharController(id, fileName)
		if editErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"path": fileName,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"path": "",
	})
	return
}

// 消费者 添加地址
func PostAddAddress(ctx *gin.Context) {
	var addressParam web_param.AddressParam
	err := ctx.ShouldBind(&addressParam)
	if err != nil {
		common.Failed(ctx, "获取地址参数失败")
		return
	}
	// 调用添加地址控制器
	resErr := controller.AddAddressController(addressParam)
	if resErr != nil {
		common.Failed(ctx, resErr.Error())
		return
	}
	common.Success(ctx, "添加地址成功!")
	return
}

// 消费者 地址列表
func PostAddressList(ctx *gin.Context) {
	var idParam web_param.ConsumerId
	err := ctx.ShouldBind(&idParam)
	if err != nil {
		common.Failed(ctx, "获取用户ID失败")
		return
	}
	id, _ := strconv.ParseInt(idParam.IdStr, 10, 64)
	res, resErr := controller.AddressListController(id)
	if resErr != nil {
		common.Failed(ctx, "获取地址列表失败")
		return
	}
	common.Success(ctx, res)
	return
}

// 消费者删除地址
func PostRemoveAddress(ctx *gin.Context) {
	var addressId web_param.AddressId
	err := ctx.ShouldBind(&addressId)
	if err != nil {
		common.Failed(ctx, "获取地址参数失败")
		return
	}
	// 调用删除地址
	id := addressId.Id
	resErr := controller.RemoveAddressController(id)
	if resErr != nil {
		err = errors.New("删除地址失败")
		common.Failed(ctx, err.Error())
		return
	}
	common.Success(ctx, "删除成功")
	return
}

// 消费者 修改地址
func PostEditAddress(ctx *gin.Context) {
	var editAddressParam web_param.EditAddressParam
	err := ctx.ShouldBind(&editAddressParam)
	if err != nil {
		common.Failed(ctx, "获取地址参数失败")
		return
	}
	// 调用修改方法
	id := editAddressParam.AddressId
	consumerId := editAddressParam.ConsumerId
	consumerAddress := editAddressParam.ConsumerAddress
	resErr := controller.EditAddressController(id, consumerId, consumerAddress)
	if resErr != nil {
		common.Failed(ctx, "修改地址失败")
		return
	}
	common.Success(ctx, "修改地址成功")
	return
}

// 消费者 手机号 注册接口
func ConsumerSignUpApi(ctx *gin.Context) {
	// 解析注册参数
	var signUpParam web_param.SignParam
	err := ctx.ShouldBind(&signUpParam)
	if err != nil {
		common.Failed(ctx, "手机号码、密码和验证码都要填写哈!")
		return
	}
	// 用户注册
	//err = controller.ConsumerSignUpController(&signUpParam)
	//if err != nil {
	//	common.Failed(ctx, err.Error())
	//	return
	//}
	common.Success(ctx, "注册成功")
	return
}

// 消费者 手机号 登录接口
func ConsumerLoginApi(ctx *gin.Context) {
	// 解析登录参数
	var loginParam web_param.LoginParam
	err := ctx.ShouldBind(&loginParam)
	if err != nil {
		common.Failed(ctx, "手机号码、密码和验证码都要填写哈!")
		return
	}
	// 用户登录
	resConsumer, err := controller.ConsumerLoginController(&loginParam)
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	// 设置session
	stringId := strconv.FormatInt(resConsumer.Id, 10)
	sessionErr := common.SetSess(ctx, "consumer_"+stringId, resConsumer.Id, 1000*60*3)
	if sessionErr != nil {
		common.Failed(ctx, "设置session错误,登录失败!")
		return
	}
	// 生成token
	tokenString, tokenErr := common.ReleaseToken(resConsumer.Id)
	if tokenErr != nil {
		common.Failed(ctx, "生成密钥错误,登录失败")
		return
	}
	// 登录成功 返回token
	common.Success(ctx, tokenString)
	return
}
