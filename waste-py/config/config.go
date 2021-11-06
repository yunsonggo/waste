package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	App struct {
		AppName string `json:"app_name"`
		AppUrl  string `json:"app_url"`
		AppMode string `json:"app_mode"`
		AppPort string `json:"app_port"`
	}
	Session struct {
		SessionName   string `json:"session_name"`
		SessionSecret string `json:"session_secret"`
		SessionMode   string `json:"session_mode"`
	}
	Mysql struct {
		MysqlDriver string `json:"mysql_driver"`
		MysqlAddr   string `json:"mysql_addr"`
		MysqlShow   bool   `json:"mysql_show"`
	}
	Redis struct {
		RedisAddr string `json:"redis_addr"`
		RedisPwd  string `json:"redis_pwd"`
		RedisHold int    `json:"redis_hold"`
	}
	RabbitMq struct {
		RabbitMqAddr string `json:"rabbitmq_addr"`
	}
	Jwt struct {
		JwtKey string `json:"jwt_key"`
		Issuer string `json:"issuer"`
	}
	Email struct {
		FromEmail string `json:"from_email"`
		SmtpAddr  string `json:"smtp_addr"`
		SmtpPort  string `json:"smtp_port"`
		SmtpPass  string `json:"smtp_pass"`
	}
}

var Conf Config

func ConfigInit() {
	var cfg = new(Config)
	file, err := os.Open("config/config.json")
	if err != nil {
		err := errors.New("读取配置文件失败")
		panic(err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&cfg)
	if err != nil {
		err := errors.New("解析配置文件失败")
		panic(err.Error())
	}
	Conf = *cfg
	return
}
