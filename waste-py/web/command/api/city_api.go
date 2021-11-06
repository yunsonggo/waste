package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"os"
	"waste-py/common"
)

func GetCityInfo(ctx *gin.Context) {
	filePath, _ := os.Getwd()
	fileName := "/web/model/city_data.json"
	path := filePath + fileName
	jsonData := gojsonq.New().File(path).Get()
	common.Success(ctx, jsonData)
	return
}

// 根据IP 获取定位
func GetAddrByIp(ctx *gin.Context) {
	// 获取请求IP
	ip := ctx.ClientIP()
	fmt.Println(ip)
	//ip := exnet.ClientPublicIP(ctx.Request)
	//if ip == "" {
	//	ip = exnet.ClientIP(r)
	//}
	//res, err := http.Get("https://restapi.amap.com/v5/ip?ip=" + ip + "&key=0ec5cb053e95f73c8a96cdf0c3b943a9&type=4")
	res, err := http.Get("https://restapi.amap.com/v5/ip?ip=123.52.203.143&key=0ec5cb053e95f73c8a96cdf0c3b943a9&type=4")
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	fmt.Printf("%v+", string(body))
	common.Success(ctx, string(body))
	return
}
