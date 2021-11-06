import ajax from "./ajax";
const BASE_URL = "http://192.168.1.102:8080/api";

/**
 * @description: 城市及定位 请求
 * @type: GET 
 * @param {} 
 * @return code:状态 data:数据 msg:消息
 */
// 请求城市数据
export const reqCities = () => ajax(BASE_URL + "/consumer/cities");
// 请求IP定位
export const reqAddress = () => ajax(BASE_URL + "/consumer/ip/address");
/**
 * @description: 验证码 及 注册 登录 修改 请求
 * @type: GET/POST
 * @param {} / {email} 
 * @return code:状态 data:数据 msg:消息
 */
// 请求 获取短信验证码
// export const reqSendCode = (phone,tpl_id,key) => ajax(BASE_URL + '/sms/send', {phone,tpl_id,key},'POST')
// 请求 获取图形验证码
export const reqBs64Code = () => ajax(BASE_URL + '/consumer/captcha')
// 请求 邮箱验证码
export const reqEmailCode = (consumer_email) => ajax(BASE_URL + '/consumer/email/captcha',{consumer_email},'POST')
// 邮箱注册请求
export const reqSignUpByEmail = (consumer_email,consumer_password,consumer_check_password,check_email_code) => ajax(BASE_URL + "/consumer/email/signup", {consumer_email,consumer_password,consumer_check_password,check_email_code}, "POST");
// 邮箱登录请求
export const reqLoginByEmail = (consumer_email,consumer_password,check_code_id,check_code) => ajax(BASE_URL + '/consumer/email/login',{consumer_email,consumer_password,check_code_id,check_code},'POST')
// 退出登录请求
export const reqLogOutByEmail = (consumer_id) => ajax(BASE_URL + '/consumer/email/logout',{consumer_id},'POST')
// 邮箱找回密码
export const reqResetPassByEmail = (consumer_email,consumer_password,consumer_check_password,check_email_code) => ajax(BASE_URL + '/consumer/email/reset/pass',{consumer_email,consumer_password,consumer_check_password,check_email_code},'POST')
// 修改个人信息请求
export const reqEditByEmail = (id,email,name,nickname,mobile,gender) => ajax(BASE_URL + '/consumer/edit/info',{id,email,name,nickname,mobile,gender},'POST')
// 添加地址信息
export const reqAddAddress = (id,address,token) => ajax(BASE_URL + '/consumer/add/address' , {id,address,token},'POST')
// 请求所有地址信息
export const reqAddressAll = (consumer_id) => ajax(BASE_URL + '/consumer/address/list',{consumer_id},'POST')
// 删除地址
export const reqDelAddress = (address_id) => ajax(BASE_URL + '/consumer/address/remove',{address_id},'POST')
// 修改地址
export const reqEditAddressOne = (address_id,consumer_id,consumer_address) => ajax(BASE_URL + '/consumer/address/edit',{address_id,consumer_id,consumer_address},'POST')
// 请求数据库城市数据
export const reqMysqlCitiesData = () => ajax(BASE_URL + '/position/cities')
// 请求地区数据
export const reqMysqlAreasData = () => ajax(BASE_URL + '/position/area/list')
// 请求站点数据
export const reqPositionsData = () => ajax(BASE_URL + '/position/list')
// 提交订单
export const reqAddOrder = (order_info) => ajax(BASE_URL + '/consumer/order/insert',{order_info},'POST')
// 获取所有订单
export const reqOrderList = (consumer_id) => ajax(BASE_URL + '/order/list',{consumer_id},'POST')
// 获取未完成订单
export const reqNotdoneOrderList = (consumer_id) => ajax(BASE_URL + '/order/notdone/list',{consumer_id},'POST')
// 获取已完成订单
export const reqDoneOrderList = (consumer_id) => ajax(BASE_URL + '/order/done/list',{consumer_id},'POST')
/**
 * @description: 商品数据 请求
 * @type: GET
 * @param {} 
 * @return code:状态 data:数据 msg:消息
 */
// 请求所有商品数据
export const reqGoods = () => ajax(BASE_URL + "/good/list");
