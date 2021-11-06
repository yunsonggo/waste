<template>
  <div id="detailContrainer" class="slide-enter-active">
    <Scroller
      :handleToScroll="handleToScroll"
      :handleToTouchEnd="handleToTouchEnd"
      ref="scro"
    >
      <div>
        <div class="pullDown">{{ pullDownMsg }}</div>
        <loading v-if="isLoading"></loading>
        <div v-else id="content" class="contentDetail">
          <!-- <div class="detail_list">
            <div
              class="detail_list_bg"
              :style="{
                'background-image': 'url(' + staticUrl + goodInfo.icon + ')',
              }"
            ></div>
            <div class="detail_list_filter"></div>
            <div class="detail_list_content">
              <div class="detail_list_img">
                <img :src="staticUrl + goodInfo.icon" alt="" />
              </div>
              <div class="detail_list_info">
                <h2 v-if="goodInfo.is_hot == 1">
                  <el-badge value="hot" class="item">
                    {{ goodInfo.name }}
                  </el-badge>
                </h2>
                <h2 v-else>{{ goodInfo.name }}</h2>
                <h4>单价: {{ goodInfo.price }}元 /{{ goodInfo.goods_unit }}</h4>
                <p class="line_through">
                  原价: {{ goodInfo.old_price }}元 /{{ goodInfo.goods_unit }}
                </p>
                <p>描述: {{ goodInfo.remarkription }}</p>
                <p>已回收: {{ goodInfo.sell_count }} 单</p>
              </div>
            </div>
          </div> -->
          <div class="detail_intro">
            <el-divider content-position="left">生成订单</el-divider>
          </div>
          <div>
            <el-form
              :model="ruleForm"
              :rules="rules"
              ref="ruleForm"
              label-width="100px"
              class="demo-ruleForm"
            >
              <el-form-item label="旧物名称" prop="name">
                <el-input
                  v-model="ruleForm.name"
                  :disabled="true"
                  style="width:80%"
                ></el-input>
              </el-form-item>
              <el-form-item label="地区" prop="cityid">
                <el-select
                  v-model="ruleForm.cityid"
                  placeholder="请选择地区"
                  @change="remoteCityMethod"
                >
                  <el-option
                    v-for="city in citiesInfo"
                    :key="city.city_id"
                    :label="city.city_name"
                    :value="city.Id"
                  ></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="街道" prop="areaid">
                <el-select
                  v-model="ruleForm.areaid"
                  placeholder="请选择街道"
                  @change="remoteAreaMethod"
                >
                  <el-option
                    v-for="(area, index) in areasInfo"
                    :key="index"
                    :label="area.area_name"
                    :value="area.Id"
                  ></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="站点" prop="positionid">
                <el-select
                  v-model="ruleForm.positionid"
                  placeholder="请选择站点"
                >
                  <el-option
                    v-for="(position, index) in positionsInfo"
                    :key="index"
                    :label="position.position_name"
                    :value="position.Id"
                  ></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="联系人" prop="consumer_name">
                <el-input
                  v-model="ruleForm.consumer_name"
                  :disabled="true"
                  style="width:80%"
                ></el-input>
              </el-form-item>
              <el-form-item label="微信电话" prop="consumer_mobile">
                <el-input
                  v-model="ruleForm.consumer_mobile"
                  :disabled="true"
                  style="width:80%"
                ></el-input>
              </el-form-item>
              <el-form-item label="地址" prop="address_id">
                <el-select
                  v-model="ruleForm.address_id"
                  placeholder="请选择地址"
                >
                  <el-option
                    v-for="(address, index) in myAddressData"
                    :key="index"
                    :label="address.consumer_Address"
                    :value="address.Id"
                  ></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="收款码" prop="consumer_wechar_icon">
                <img
                  v-if="ruleForm.consumer_wechar_icon"
                  :src="staticUrl + ruleForm.consumer_wechar_icon"
                  class="wecharIcon"
                />
              </el-form-item>
              <el-form-item label="备注" prop="remark">
                <el-input
                  type="textarea"
                  v-model="ruleForm.remark"
                  style="width:80%"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="danger" @click="submitForm('ruleForm')"
                  >提交订单</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </Scroller>
  </div>
</template>

<script>
import Header from "@/components/Header";
import {
  reqGoods,
  reqMysqlCitiesData,
  reqMysqlAreasData,
  reqPositionsData,
  reqAddressAll,
  reqAddOrder
} from "@/api";
import global from "@/components/Global";
import Scroller from "@/components/Scroller";

export default {
  name: "addorder",
   beforeRouteEnter(to, from, next) {
    var isAuth = localStorage.getItem("consumerInfo");
    if (!isAuth) {
      next((vm) => {
        vm.$message.error("登录以后才能提交订单哦");
        vm.$router.push("/about/login");
      });
    } else {
      next();
    }
  },
  components: {
    Header,
    Scroller,
  },
  props: {
    goodId: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      title: "详情",
      staticUrl: global.static_url,
      isLoading: true,
      consumerInfo: null,
      allGoodsData: [],
      goodInfo: [],
      pullDownMsg: "",
      localCityName: "",
      localAreaName: null,
      localAddrName: null,
      citiesInfo: [],
      areasInfo: [],
      areasData: [],
      positionsInfo: [],
      positionsData: [],
      myAddressData: [],
      orderId:"",
      ruleForm: {
        name: "",
        good_id:"",
        good_name:"",
        good_price:"",
        goods_unit:"",
        cityid: "",
        areaid: "",
        positionid: "",
        consumer_id:"",
        consumer_name: "",
        consumer_mobile: "",
        consumer_gender:"",
        address_id: "",
        consumer_wechar_icon: "",
        remark: "",
      },
      rules: {
        name: [
          { required: true, message: "旧物名称不能为空", trigger: "blur" },
        ],
        cityid: [{ required: true, message: "请选择地区", trigger: "change" }],
        consumer_mobile: [
          {
            required: true,
            message: "请在个人中心添加微信绑定电话",
            trigger: "blur",
          },
        ],
      },
    };
  },
  mounted() {
    this.getConsumerInfo();
    this.getAllGoodsData();
    this.getAreasInfo();
    this.getPositionsInfo();
    this.getAddressList();
    this.getCitiesInfo();
  },
  methods: {
    // 获取登录用户信息
    getConsumerInfo() {
      var consumerStr = localStorage.getItem("consumerInfo");
      this.consumerInfo = JSON.parse(consumerStr);
      console.log(this.consumerInfo);
      this.ruleForm.consumer_name = this.consumerInfo.consumer_name;
      this.ruleForm.consumer_mobile = this.consumerInfo.consumer_mobile;
      this.ruleForm.consumer_wechar_icon = this.consumerInfo.consumer_wechar_icon;
      this.ruleForm.consumer_id = this.consumerInfo.id
      this.ruleForm.consumer_gender = this.consumerInfo.consumer_gender
      this.localCityName = localStorage.getItem("areaName");
      var localAddrStr = localStorage.getItem("addr");
      this.localAddrName = JSON.parse(localAddrStr);
    },
    async getAddressList() {
      const result = await reqAddressAll(this.consumerInfo.id.toString());
      if (result.code === 0) {
        this.myAddressData = result.data;
      } else {
        this.$message.error("获取地址列表失败");
      }
    },
    // 获取本地商品数据
    getAllGoodsData() {
      var goodsData = this.$store.getters.goodsAll;
      if (goodsData.length) {
        this.allGoodsData = goodsData;
        this.getDetailInfo();
      } else {
        this.getBackGoodsData();
      }
    },
    // 获取商品数据
    async getBackGoodsData() {
      const result = await reqGoods();
      if (result.code === 0) {
        this.$store.dispatch("setGoodsAll", result.data);
        this.allGoodsData = result.data;
        this.getDetailInfo();
      }
    },
    // 获取当前商品信息
    getDetailInfo() {
      if (this.goodId) {
        this.allGoodsData.forEach((item) => {
          if (item.Id == this.goodId) {
            this.goodInfo = item;
            console.log(this.goodInfo);
            this.ruleForm.name = this.goodInfo.name;
            this.ruleForm.good_id = this.goodInfo.Id
            this.ruleForm.good_name = this.goodInfo.name
            this.ruleForm.good_price = this.goodInfo.price
            this.ruleForm.goods_unit = this.goodInfo.goods_unit
            this.isLoading = false;
            this.$nextTick(() => {
              this.$refs.scro.scroll.refresh();
              // refScroll()
            });
          }
        });
      }
    },
    // 获取城市数据
    async getCitiesInfo() {
      const result = await reqMysqlCitiesData();
      if (result.code === 0) {
        this.citiesInfo = result.data;
        this.citiesInfo.forEach((city) => {
          if (city.city_name == this.localCityName) {
            this.ruleForm.cityid = city.Id;
            this.remoteCityMethod(city.Id)
          }
        });
      } else {
        this.$message.error("获取数据库城市数据错误");
      }
    },
    // 获取地区数据
    async getAreasInfo() {
      const result = await reqMysqlAreasData();
      if (result.code === 0) {
        this.areasInfo = result.data;
        this.areasData = result.data;
      } else {
        this.$message.error("获取地区数据错误");
      }
    },
    // 获取站点数据
    async getPositionsInfo() {
      const result = await reqPositionsData();
      if (result.code === 0) {
        this.positionsInfo = result.data;
        this.positionsData = result.data;
      } else {
        this.$message.error("获取站点数据错误");
      }
    },
    // 滚动触发事件
    handleToScroll(pos) {
      if (pos.y > 30) {
        this.pullDownMsg = "正在更新...";
      }
    },
    // 滚动结束事件
    handleToTouchEnd(pos) {
      if (pos.y > 30) {
        // 可以调用请求 更新数据
        // const result = .....
        this.pullDownMsg = "更新成功...";
      }
      setTimeout(() => {
        this.pullDownMsg = "";
      }, 1000);
    },
    remoteCityMethod(data) {
      this.areasInfo = [];
      this.ruleForm.areaid = "";
      var resAreaArr = [];
      this.areasData.forEach((area) => {
        if (area.cities_id == data) {
          resAreaArr.push(area);
        }
      });
      resAreaArr.forEach((item) => {
        this.areasInfo.push(item);
      });
    },
    remoteAreaMethod(data) {
      this.positionsInfo = [];
      this.ruleForm.positionid = "";
      var resPositionArr = [];
      this.positionsData.forEach((position) => {
        if (position.area_id == data) {
          resPositionArr.push(position);
        }
      });
      resPositionArr.forEach((item) => {
        this.positionsInfo.push(item);
      });
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.addThisOrder()
        } else {
          this.$message.error("检查填写内容")
          return false;
        }
      });
    },
    // 提交订单
    async addThisOrder() {
      const result = await reqAddOrder(this.ruleForm)
      console.log(result);
      if (result.code === 0) {
        this.orderId = result.data
        this.handleMessageBox(this.orderId )
      } else {
        this.$message.error("提交订单失败,请重试")
      }
    },
     handleMessageBox(orderId) {
      const h = this.$createElement;
      this.$msgbox({
        title: "请将订单号标注在旧物上",
        message: h("p", null, [
          h("span", null, "标注订单号后投入回收站 "),
          h("i", { style: "color:teal;font-size:18px" }, orderId),
        ]),
        showCancelButton: true,
        confirmButtonText: "确定",
        showCancelButton:false,
        showClose:false,
        center: true,
      }).then(() => {
         this.$router.push("/order")
        })
    },
  },
};
</script>

<style>
.el-message-box {
  width: 330px;
}
</style>
<style scoped>
#detailContrainer {
  position: absolute;
  left: 0;
  top: 0;
  z-index: 100;
  width: 100%;
  /* min-height: 100%; */
  height: 500px;
  background: white;
}
#detailContrainer.slide-enter-active {
  animation: 0.3s slideMove;
}
@keyframes slideMove {
  0% {
    transform: translateX(100%);
  }
  100% {
    transform: translateX(0);
  }
}
#detailContrainer.pullDown {
  margin: 0;
  padding: 0;
  border: none;
}
#content.contentDetail {
  display: block;
  margin-bottom: 0;
}
#content .detail_list {
  height: 200px;
  width: 100%;
  position: relative;
  overflow: hidden;
}
.detail_list .detail_list_bg {
  width: 100%;
  height: 100%;
  background: 0 40%;
  filter: blur(20px);
  background-size: cover;
  position: absolute;
  left: 0;
  top: 0;
}
.detail_list .detail_list_filter {
  width: 100%;
  height: 100%;
  position: absolute;
  background-color: #40454d;
  opacity: 0.55;
  position: absolute;
  left: 0;
  top: 0;
  z-index: 1;
}
.detail_list .detail_list_content {
  display: flex;
  width: 100%;
  height: 100%;
  position: relative;
  left: 0;
  top: 0;
  z-index: 2;
}
.detail_list .detail_list_img {
  width: 108px;
  height: 150px;
  border: solid 1px #f0f2f3;
  margin: 20px;
}
.detail_list .detail_list_img img {
  width: 100%;
  height: 100%;
}
.detail_list .detail_list_info {
  margin-top: 20px;
}
.detail_list .detail_list_info h2 {
  font-size: 20px;
  color: white;
  font-weight: normal;
  line-height: 40px;
}
.detail_list .detail_list_info h4 {
  font-size: 16px;
  color: white;
  font-weight: normal;
  line-height: 20px;
}
.detail_list .detail_list_info p {
  color: white;
  line-height: 20px;
  font-size: 14px;
}
.detail_list .detail_list_info .line_through {
  font-style: italic;
  color: #ccc;
  line-height: 20px;
  font-size: 14px;
  text-decoration: line-through;
}

#content .detail_intro {
  padding: 10px;
}
#content .detail_intro .good_content {
  line-height: 22px;
}
#content .detail_player {
  margin-top: 20px;
}
.detail_player .swiper-slide {
  width: 70px;
  margin-right: 20px;
  text-align: center;
  font-size: 14px;
}
.detail_player .swiper-slide img {
  width: 100%;
  margin-bottom: 5px;
}
.detail_player .swiper-slide p:nth-of-type(2) {
  color: #999;
}
.wecharIcon {
  width: 98px;
  height: 138px;
  display: block;
}
</style>
