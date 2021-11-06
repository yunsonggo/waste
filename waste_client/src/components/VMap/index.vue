<template>
  <div class="amap-wrapper">
    <h2 class="title">
      <el-tag v-if="areaName" type="danger" effect="dark" size="mini">
        {{ areaName }}
      </el-tag>
      {{ msg }}
    </h2>
    <div class="search-wrapper">
      <el-input
        v-model="keyword"
        placeholder="请输入地址"
        @change="handleBlur"
        style="width:90%"
      ></el-input>

      <el-tooltip
        class="item"
        effect="light"
        content="添加地址"
        placement="top"
        style="font-size:20px"
      >
        <i
          class="el-icon-plus"
          style="margin-left:10px;color:#e54847"
          @click="handleAddAddress"
        ></i>
      </el-tooltip>
    </div>
    <div>
      <div id="container"></div>
    </div>
    <div>
      <AreaScroll
        :areaList="areaList"
        :showMap="showMap"
        @selectItem="selectAddr"
      ></AreaScroll>
    </div>
  </div>
</template>

<script>
import AreaScroll from "@/components/AreaScroll";
import {reqAddAddress} from "@/api"
export default {
  name: "vmap",
  props: {
    areaName: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      MyMap: Object,
      keyword: "",
      msg: "",
      lng: 0,
      lat: 0,
      zoom: 13,
      center: [115.05863, 35.76028],
      markers: [],
      areaList: [],
      showMap: false,
    };
  },
  components: {
    AreaScroll,
  },
  mounted() {
    this.init();
    this.getLocalAddr();
  },
  watch: {
    keyword() {
      this.searchPlace();
    },
  },
  methods: {
    init() {
      var self = this;
      var map = new AMap.Map("container", {
        resizeEnable: true, //是否监控地图容器尺寸变化
        zoom: self.zoom, //初始化地图层级
        center: self.center, //初始化地图中心点
        showIndoorMap: false, //关闭室内地图
      });
      self.MyMap = map;
      map.on("complete", () => {
        // console.log("地图加载完成！");
        AMap.plugin("AMap.Geolocation", () => {
          var geolocation = new AMap.Geolocation({
            enableHighAccuracy: true, //是否使用高精度定位，默认:true
            maximumAge: 0, //定位结果缓存0毫秒，默认：0
            convert: true, //自动偏移坐标，偏移后的坐标为高德坐标，默认：true
            showButton: true, //显示定位按钮，默认：true
            buttonPosition: "RB", //定位按钮停靠位置，默认：'LB'，左下角
            showMarker: true, //定位成功后在定位到的位置显示点标记，默认：true
            showCircle: true, //定位成功后用圆圈表示定位精度范围，默认：true
            panToLocation: true, //定位成功后将定位到的位置作为地图中心点，默认：true
            zoomToAccuracy: true, //定位成功后调整地图视野范围使定位位置及精度范围视野内可见，默认：f
            timeout: 10000, //超过10秒后停止定位，默认：5s
            buttonOffset: new AMap.Pixel(10, 20), //定位按钮与设置的停靠位置的偏移量，默认：Pixel(10, 20)
          });
          geolocation.getCurrentPosition();
          AMap.event.addListener(geolocation, "complete", onComplete);
          AMap.event.addListener(geolocation, "error", onError);
          function onComplete(data) {
            // console.log(data);
            self.$store.dispatch("setLocation", data.addressComponent);
            self.$store.dispatch("setAddress", data.formattedAddress);
            self.$store.dispatch("setPosition", data.position);
            var position = new AMap.LngLat(
              data.position.lng,
              data.position.lat
            );
            self.setMarkerByPosition(position, "目前定位");
          }
          function onError(data) {
            // console.log(data);
            // console.log("定位失败");
            self.getLnglatLocation();
          }
        });
      });
      //   console.log("center:" + self.center)
      //   map.setCenter(self.center);
    },
    getLnglatLocation() {
      console.log("开始经纬度定位");
      var self = this;
      AMap.plugin("AMap.CitySearch", () => {
        var citySearch = new AMap.CitySearch();
        citySearch.getLocalCity(function(status, result) {
          if (status === "complete" && result.info === "OK") {
            // 查询成功，result即为当前所在城市信息
            // console.log(result);
            // -----------------------
            AMap.plugin("AMap.Geocoder", function() {
              var geocoder = new AMap.Geocoder({
                // city 指定进行编码查询的城市，支持传入城市名、adcode 和 citycode
                city: result.adcode,
              });
              var lnglat = result.rectangle.split(";")[0].split(",");
              geocoder.getAddress(lnglat, function(status, data) {
                if (status === "complete" && data.info === "OK") {
                  // console.log(data)
                  // result为对应的地理位置详细信息
                  self.$store.dispatch(
                    "setLocation",
                    data.regeocode.addressComponent
                  );
                  self.$store.dispatch(
                    "setAddress",
                    data.regeocode.formattedAddress
                  );
                  self.$store.dispatch("setPosition", result.bounds.nc);
                  //console.log(data);
                  //   var lnglat = result.rectangle.split(";")[0].split(",");
                  var position = new AMap.LngLat(
                    result.bounds.nc.lng,
                    result.bounds.nc.lat
                  );
                  // if (position) {
                  //   this.center = "[" + result.bounds.nc.lng + "," + result.bounds.nc.lat + "]";
                  // }
                  self.setMarkerByPosition(
                    position,
                    data.regeocode.formattedAddress
                  );
                }
              });
            });
          }
        });
      });
    },
    setMarkerByPosition(position, title) {
      this.MyMap.setCenter(position);
      var marker = new AMap.Marker({
        position: position,
        title: title,
      });
      // 将创建的点标记添加到已有的地图实例：
      this.MyMap.add(marker);
    },
    searchPlace() {
      const self = this;
      if (self.keyword) {
        AMap.plugin("AMap.Autocomplete", function() {
          // 实例化Autocomplete
          var autoOptions = {
            //city 限定城市，默认全国
            //   city: self.city,
            city: "濮阳市",
          };
          var autoComplete = new AMap.Autocomplete(autoOptions);
          autoComplete.search(self.keyword, function(status, result) {
            // 搜索成功时，result即是对应的匹配数据
            // console.log(result);
            self.areaList = result.tips;
            self.showMap = true;
            // console.log(self.areaList);
          });
        });
      } else {
        self.showMap = false;
        self.areaList = [];
        self.init();
      }
    },
    getLocalAddr() {
      var addrString = localStorage.getItem("addr");
      if (addrString) {
        var addrData = JSON.parse(addrString);
        this.msg = ""
        this.msg += addrData.name;
        var position = new AMap.LngLat(
          addrData.location.lng,
          addrData.location.lat
        );
        this.setMarkerByPosition(position, addrData.name);
      }
    },
    saveLocalAddr(data) {
      if (data) {
        localStorage.removeItem("addr");
        localStorage.setItem("addr", JSON.stringify(data));
        this.msg = ""
        this.msg += data.name;
        var position = new AMap.LngLat(data.location.lng, data.location.lat);
        this.setMarkerByPosition(position, data.name);
      }
    },
    handleBlur() {
      console.log("change");
      this.showMap = false;
    },
    clearInput() {
      this.showMap = false;
      this.areaList = [];
      this.keyword = "";
    },
    selectAddr(data) {
      console.log(data);
      if (data) {
        this.getAddrMsg(data.name);
        this.saveLocalAddr(data);
      } else {
        this.clearInput();
      }
    },
    getAddrMsg(msg) {
      this.msg = ""
      this.msg += msg;
      this.clearInput();
    },
    handleAddAddress() {
      const h = this.$createElement;
      var msgBody = "";
      var dingAddr = this.$store.getters.address;
      if (this.areaName) {
        msgBody += this.areaName;
      }
      if (this.msg != "") {
        msgBody += this.msg;
      } else if (dingAddr) {
        msgBody += "定位地址:";
        msgBody += dingAddr;
      } else {
        this.$message.error("没有地址可以添加");
        return;
      }
      this.$msgbox({
        title: "添加地址",
        message: h("p", null, [
          h("span", null, "地址是: "),
          h("i", { style: "color:teal" }, msgBody),
        ]),
        showCancelButton: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        center: true,
      }).then(() => {
          // 调用保存地址方法
        this.addAddress(msgBody)
        })
        .catch((action) => {
          this.$message({
            type: "info",
            message: "放弃保存",
          });
        });
    },
    async addAddress(address) {
       var consumerInfo = localStorage.getItem("consumerInfo")
         var consumer = JSON.parse(consumerInfo)
         var id = consumer.id
         var tokenWord = consumer.consumer_password
      const result = await reqAddAddress(id,address,tokenWord)
      console.log(result);
      if (result.code === 0) {
        this.$message({
            type: "success",
            message: "地址添加成功!",
          });
          this.$router.push('/about/center')
      }else {
        this.$message.error("地址添加失败!");
      }
    }
  },
};
</script>

<style>
.el-message-box {
  width: 330px;
}
</style>

<style scoped>
.amap-wrapper {
  margin: 15px;
  padding: 0;
  width: 100%;
  height: 100%;
}
.amap-wrapper .title {
  font-size: 14px;
}
.amap-wrapper .search-wrapper {
  margin-top: 10px;
  width: 300px;
}
.amap-box {
  margin: 10px;
  padding: 0;
  height: 400px;
  width: 400px;
}
.amap-wrapper .area {
  margin-top: 16px;
  background: #fff;
}
.amap-wrapper .area li {
  border-bottom: 1px solid #eee;
  padding: 8px 16px;
  color: #aaa;
}
.amap-wrapper .area li h4 {
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}

#container {
  margin-top: 10px;
  width: 100%;
  height: 350px;
}
</style>
