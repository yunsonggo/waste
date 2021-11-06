package config

import (
	"github.com/gin-gonic/gin"
	"github.com/yakaa/log4g"
	"waste-py/dal"
)

func InitRrsl() {
	// 配置生成模式
	appConf := Conf.App
	if appConf.AppMode == gin.ReleaseMode {
		log4g.Init(log4g.Config{Path: "logs"})
		gin.DefaultWriter = log4g.InfoLog
		gin.DefaultErrorWriter = log4g.ErrorLog
	}
	// 配置数据库
	mysqlConf := Conf.Mysql
	redisConf := Conf.Redis
	dal.SqlEngine(mysqlConf.MysqlAddr,mysqlConf.MysqlDriver,mysqlConf.MysqlShow)
	dal.InitRedisStore(redisConf.RedisAddr, redisConf.RedisPwd)
}
