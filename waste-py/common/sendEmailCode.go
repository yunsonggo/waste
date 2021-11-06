package common

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"waste-py/config"
	"waste-py/tools"
)

func SendEmailCode(toEmail string) (theCode string, err error) {
	// 生成验证码随机数
	code := tools.GetRandomCode()
	mailHeader := map[string][]string{
		"From":    {config.Conf.Email.FromEmail},
		"To":      {toEmail},
		"Subject": {"旧物Online注册验证码"},
	}
	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	// text/html html格式  text/plain 纯文本
	m.SetBody("text/html", "尊敬的用户,欢迎您注册'旧物Online...<br>您的注册验证码为:"+code+"<br>畅享绿色健康生活,感谢您支持可再生资源回收事业!")
	//m.SetAddressHeader()
	smtpAddr := config.Conf.Email.SmtpAddr
	smtpPort, _ := strconv.Atoi(config.Conf.Email.SmtpPort)
	smtpUser := config.Conf.Email.FromEmail
	smtpPass := config.Conf.Email.SmtpPass
	d := gomail.NewDialer(smtpAddr, smtpPort, smtpUser, smtpPass)
	err = d.DialAndSend(m)
	return code, err
}
