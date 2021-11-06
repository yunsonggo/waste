<template>
  <div class="city_body">
    <div class="city_list">
      <div class="city_hot">
        <h2>
          <el-badge :value="dingweiAddr" class="item">
            服务地区
          </el-badge>
        </h2>
        <ul class="clearfix">
          <li
            v-for="item in citiesInfo"
            :key="item.id"
            @click="selectArea(item.name)"
          >
            {{ item.name }}
          </li>
        </ul>
      </div>
      <div>
        <VMap :areaName="areaName"></VMap>
      </div>
    </div>
  </div>
</template>

<script>
import VMap from "@/components/VMap";
import { reqCities } from "@/api";
export default {
  name: "City",
  components: {
    VMap,
  },
  data() {
    return {
      citiesInfo: [],
      areaName: "",
    };
  },
  computed: {
    city() {
      if (this.$store.getters.location.addressComponent) {
        return this.$store.getters.location.addressComponent.city;
      } else {
        return "定位中";
      }
    },
    dingweiAddr() {
      var dingAddr = this.$store.getters.address;
      if (dingAddr) {
        return "定位:" + dingAddr
      }
    },
  },
  mounted() {
    var citiesData = localStorage.getItem("citiesData");
    var localAreaName = localStorage.getItem("areaName");
    if (citiesData) {
      this.citiesInfo = JSON.parse(citiesData);
    } else {
      this.getCitiesData();
    }
    if (localAreaName) {
      this.areaName = localAreaName;
    }
  },
  methods: {
    async getCitiesData() {
      const result = await reqCities();
      if (result.code == 0) {
        this.citiesInfo = result.data.cities;
        // 城市数据 本地存储
        localStorage.setItem("citiesData", JSON.stringify(this.citiesInfo));
      }
    },
    selectArea(item) {
      this.areaName = item;
      localStorage.removeItem("areaName");
      localStorage.setItem("areaName", item);
    },
  },
};
</script>
<style scoped>
#content .city_body {
  margin-top: 45px;
  display: flex;
  width: 100%;
  position: absolute;
  top: 0;
  bottom: 0;
}
.city_body .city_list {
  flex: 1;
  overflow: auto;
  background: #fff5f0;
}
.city_body .city_list::-webkit-scrollbar {
  background-color: transparent;
  width: 0;
}
.city_body .city_hot {
  margin-top: 20px;
}
.city_body .city_hot h2 {
  padding-left: 15px;
  line-height: 30px;
  font-size: 14px;
  background: #f0f0f0;
  font-weight: normal;
}
.city_body .city_hot ul li {
  float: left;
  background: #fff;
  width: 29%;
  height: 33px;
  margin-top: 15px;
  margin-left: 3%;
  padding: 0 4px;
  border: 1px solid #e6e6e6;
  border-radius: 3px;
  line-height: 33px;
  text-align: center;
  box-sizing: border-box;
}
.city_body .city_sort div {
  margin-top: 20px;
}
.city_body .city_sort h2 {
  padding-left: 15px;
  line-height: 30px;
  font-size: 14px;
  background: #f0f0f0;
  font-weight: normal;
}
.city_body .city_sort ul {
  padding-left: 10px;
  margin-top: 10px;
}
.city_body .city_sort ul li {
  line-height: 30px;
  line-height: 30px;
}
.city_body .city_index {
  width: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;
  border-left: 1px #e6e6e6 solid;
}
</style>
