<template>
  <div>
    <div v-if="consumerInfo" class="me">
      <div class="headInfo">
        <div
          v-if="consumerInfo.consumer_icon"
          class="head-img"
          :style="{
            'background-image':
              'url(' + static_url + consumerInfo.consumer_icon + ')',
          }"
        ></div>
        <div
          v-else
          class="head-img"
          :style="{
            'background-image':
              'url(' + static_url + '/upload/consumer/space_avatar_0.png)',
          }"
        ></div>
        <div class="head-profile">
          <p class="user-id">
            {{ consumerInfo.consumer_nickname }}
            {{ consumerInfo.consumer_gender }}
          </p>
          <p class="user-phone">
            <i class="el-icon-user-solid"></i>
            <span>{{ consumerInfo.consumer_name }}</span>
          </p>
          <p class="user-phone">
            <i class="el-icon-message"></i>
            <span>{{ consumerInfo.consumer_email }}</span>
          </p>
          <p class="user-phone">
            <i class="el-icon-phone-outline"></i>
            <span>{{ consumerInfo.consumer_mobile }}</span>
          </p>
        </div>
      </div>
      <div class="address-cell">
        <i class="el-icon-edit"></i>
        <div class="address-index" @click="handelToEditInfo">
          <!-- <router-link to="/about/edit/info" style="text-decoration-line: none">
            <span>修改信息</span>
          </router-link> -->
          <span>修改信息</span>
          <i class="el-icon-caret-right"></i>
        </div>
      </div>
      <div class="address-cell" @click="handleToAddress">
        <i class="el-icon-location-information"></i>
        <div class="address-index">
          <!-- <router-link to="/about/address" style="text-decoration-line: none">
            <span>我的地址</span>
          </router-link> -->
          <span>我的地址</span>
          <i class="el-icon-caret-right"></i>
        </div>
      </div>
      <button @click="handleSignOut" class="loginOut-btn">退出登录</button>
    </div>
    <loading v-else></loading>
  </div>
</template>

<script>
import global from "@/components/Global";
import { reqLogOutByEmail } from "@/api";
export default {
  name: "mycenter",
  data() {
    return {
      static_url: global.static_url,
      consumerInfo: null,
    };
  },
  beforeRouteEnter(to, from, next) {
    var isAuth = localStorage.getItem("consumerInfo");
    if (!isAuth) {
      next((vm) => {
        vm.$message.error("登录以后才能进入个人中心");
        vm.$router.push("/about/login");
      });
    } else {
      next((vm) => {
        vm.consumerInfo = JSON.parse(isAuth);
      });
    }
  },
  methods: {
    handleSignOut() {
      this.logOut();
    },
    async logOut() {
      var idStr = this.consumerInfo.id.toString();
      const result = await reqLogOutByEmail(idStr);
      if (result.code === 0) {
        this.$message({
          type: "success",
          message: result.data,
        });
      }
      localStorage.removeItem("consumerInfo");
      this.$router.push("/about/login");
    },
    handelToEditInfo() {
      this.$router.push("/about/edit/info")
    },
    handleToAddress() {
      this.$router.push("/about/address")
    }
  },
};
</script>

<style scoped>
.me {
  width: 100%;
  height: 100%;
  overflow: auto;
  box-sizing: border-box;
}
.headInfo {
  display: flex;
  background-image: linear-gradient(0deg, #fff, #e54847);
  padding: 6.666667vw 4vw;
  color: #fff;
  height: 180px;
}
.head-img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-position: 0px 0px;
  background-size: 60px;
}
.head-profile {
  overflow: hidden;
  margin-left: 4.8vw;
  flex-grow: 1;
}
.head-profile .user-id {
  max-width: 40vw;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  font-size: 1.3rem;
  margin-bottom: 2.133333vw;
  font-weight: 700;
  display: flex;
  align-items: center;
}
.head-profile .user-phone {
  display: flex;
  align-items: center;
  font-size: 0.8rem;
  margin-top: 8px;
}
.user-phone > i {
  margin-right: 0.666667vw;
  font-size: 1rem;
}
.headInfo > i {
  font-size: 1.2rem;
}

.address-cell {
  margin-top: 2.666667vw;
  border-bottom: 1px solid #ee8685;
  font-size: 1rem;
  line-height: 4.533333vw;
  background: #ffffff;
  display: flex;
  align-items: center;
  padding-left: 6.666667vw;
  color: #333;
}
.address-cell > i {
  font-size: 1.3rem;
  color: rgb(74, 165, 240);
  margin-right: 2.666667vw;
}
.address-index {
  display: flex;
  justify-content: space-between;
  width: 100%;
  padding: 3.733333vw 2.666667vw 3.733333vw 0;
  align-content: center;
}
.address-index > i {
  font-size: 1.3rem;
  color: #ccc;
}
.loginOut-btn {
  display: block;
  width: 100%;
  text-align: center;
  padding: 3.733333vw 0;
  margin: 8.333333vw 0;
  color: #ff5339;
  border-radius: 0.8vw;
  font-size: 1rem;
  font-weight: 700;
  background-color: #fff;
  border: 0;
}
</style>
