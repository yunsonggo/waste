package web_param

// 消费者 获取邮箱验证码参数
type EmailCodeParam struct {
	Email string `json:"consumer_email" binding:"required"` // 邮箱
}

// 消费者 注册参数
type SignParam struct {
	Mobile         string `json:"consumer_mobile" `                           // 手机号码
	Email          string `json:"consumer_email" binding:"required"`          // 邮箱
	Password       string `json:"consumer_password" binding:"required"`       // 密码
	CheckPassword  string `json:"consumer_check_password" binding:"required"` // 确认密码
	CheckCode      string `json:"check_code" `                                // 验证码
	CheckEmailCode string `json:"check_email_code" binding:"required"`        // 邮箱验证码
}

// 消费者 登录参数
type LoginParam struct {
	Mobile         string `json:"consumer_mobile"`                      // 手机号码
	Password       string `json:"consumer_password" binding:"required"` // 密码
	Email          string `json:"consumer_email" binding:"required"`    // 邮箱
	CheckCodeId    string `json:"check_code_id" binding:"required"`     // 验证码ID
	CheckCode      string `json:"check_code" binding:"required"`        // 验证码
	CheckEmailCode string `json:"check_email_code"`                     // 邮箱验证码
}

// 消费者ID参数
type ConsumerId struct {
	IdStr string `json:"consumer_id"`
}

// 消费者修改个人信息参数
type ConsumerEditParam struct {
	Id       int64  `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile" binding:"required"`
	Gender   string `json:"gender"`
}

// 消费者添加地址
type AddressParam struct {
	Id      int64  `json:"id" binding:"required"`
	Address string `json:"address" binding:"required"`
	Token   string `json:"token" binding:"required"`
}

// 地址ID
type AddressId struct {
	Id int64 `json:"address_id" binding:"required"`
}

// 修改地址参数
type EditAddressParam struct {
	AddressId       int64  `json:"address_id" binding:"required"`
	ConsumerId      int64  `json:"consumer_id" binding:"required"`
	ConsumerAddress string `json:"consumer_address" binding:"required"`
}
