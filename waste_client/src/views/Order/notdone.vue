<template>
  <div class="cinema_body">
    <Scroller
      :handleToScroll="handleToScroll"
      :handleToTouchEnd="handleToTouchEnd"
      ref="scro"
    >
      <ul v-if="notDoneOrderList">
        <li class="pullDown">{{ pullDownMsg }}</li>
        <li v-for="(order, index) in notDoneOrderList" :key="index">
          <div>
            <span>{{ order.good_name }}</span>
            <span class="q"
              ><span class="price">单价: {{ order.good_price.toFixed(2) }}</span
              >元/{{ order.goods_unit }}</span
            >
          </div>
          <div class="address">
            <span style="color:teal;font-size:16px"
              >订单号: {{ order.order_id }}</span
            >
            <span v-if="order.state == 1" style="color:#00c78c;font-size:16px"
              >未完成</span
            >
            <span v-if="order.state == 0">状态:已完成</span>
          </div>
          <div class="address">
            <span>提交时间: {{ order.created_time }}</span>
          </div>
          <div class="address">
            <span>更新时间: {{ order.update_time }}</span>
          </div>
          <div class="card">
            <div>
              联系人: {{ order.ConsumerName }} {{ order.consumer_gender }}
            </div>
            <div>联系电话: {{ order.consumer_mobile }}</div>
          </div>
        </li>
      </ul>
      <div v-else>
        <p style="font-size:16px;margin:10px;color:teal;text-align:center;">
          尊敬的用户,目前没有订单
        </p>
        <p style="text-align:center;" @tap="handleToHome">
          <el-button type="danger" size="mini" style="font-size:16px;"
            >点击这里去下单吧</el-button
          >
        </p>
      </div>
    </Scroller>
  </div>
</template>

<script>
import { reqNotdoneOrderList } from "@/api";
import Scroller from "@/components/Scroller";
export default {
  name: "notdone",
  components: {
    Scroller,
  },
  data() {
    return {
      notDoneOrderList: [],
      consumerInfo: {},
      pullDownMsg: "",
    };
  },
  beforeRouteEnter(to, from, next) {
    var isAuth = localStorage.getItem("consumerInfo");
    if (!isAuth) {
      next((vm) => {
        vm.$message.error("登录以后才能查看订单哦");
        vm.$router.push("/about/login");
      });
    } else {
      next((vm) => {
        var consumerInfoStr = localStorage.getItem("consumerInfo");
        vm.consumerInfo = JSON.parse(consumerInfoStr);
        vm.getNotdoneOrderList(vm.consumerInfo.id.toString());
      });
    }
  },
  methods: {
    async getNotdoneOrderList(id) {
      const result = await reqNotdoneOrderList(id);
      if (result.code === 0) {
        this.notDoneOrderList = result.data;
      } else {
        this.$message.error("获取订单数据失败");
      }
    },
    handleToHome() {
      this.$router.push("/hot");
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
  },
};
</script>

<style scoped>
#content .cinema_body {
  flex: 1;
  overflow: auto;
  height: 300px;
}
.cinema_body ul {
  padding: 20px;
}
.cinema_body li {
  border-bottom: 1px solid #e6e6e6;
  margin-bottom: 20px;
}
.cinema_body div {
  margin-bottom: 10px;
}
.cinema_body .q {
  font-size: 11px;
  color: #f03d37;
}
.cinema_body .price {
  font-size: 18px;
}
.cinema_body .address {
  font-size: 13px;
  color: #666;
}
.cinema_body .address span:nth-of-type(2) {
  float: right;
}
.cinema_body .card {
  display: flex;
}
.cinema_body .card div {
  padding: 0 3px;
  height: 15px;
  line-height: 15px;
  border-radius: 2px;
  color: #f90;
  border: 1px solid #f90;
  font-size: 13px;
  margin-right: 5px;
}
.cinema_body .card div.or {
  color: #f90;
  border: 1px solid #f90;
}
.cinema_body .card div.bl {
  color: #589daf;
  border: 1px solid #589daf;
}
.cinema_body .pullDown {
  margin: 0;
  padding: 0;
  border: none;
}
</style>
