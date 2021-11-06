<template>
  <div id="main">
    <Header title="旧物 - 订单 ..."></Header>
    <div id="content">
      <div class="cinema_menu">
        <div class="hot_swtich">
          <router-link tag="div" to="/order/allorder" class="hot_item active">全部订单</router-link>
          <router-link tag="div" to="/order/notdone" class="hot_item">未完成订单</router-link>
          <router-link tag="div" to="/order/doneorder" class="hot_item">已完成订单</router-link>
        </div>
      </div>
      <keep-alive>
        <router-view />
      </keep-alive>
    </div>
    <Tabbar></Tabbar>
  </div>
</template>

<script>
import Header from "@/components/Header";
import Tabbar from "@/components/Tabbar";
export default {
  name: "Order",
  components: {
    Header,
    Tabbar,
  },
  
};
</script>

<style scoped>
#content .cinema_menu {
  width: 100%;
  height: 45px;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  justify-content: space-around;
  align-items: center;
  background: white;
}

.cinema_menu .hot_swtich {
  display: flex;
  height: 100%;
  line-height: 45px;
}
.cinema_menu .hot_item {
  font-size: 15px;
  color: #666;
  width: 80px;
  text-align: center;
  margin: 0 12px;
  font-weight: 700;
}

.cinema_menu .hot_item.router-link-active {
  color: #ef4238;
  border-bottom: 2px #ef4238 solid;
}
</style>