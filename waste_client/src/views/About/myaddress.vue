<template>
  <div class="cinema_body">
    <div class="card">
      <h3 id="ji-chu-yong-fa" style="margin:20px;color:#49464a;">
        我的地址
      </h3>
      <div>
        <el-button
          type="success"
          icon="el-icon-plus"
          circle
          size="mini"
          @click="handleToCity"
        ></el-button>
      </div>
    </div>

    <ul>
      <li v-for="(item, index) in addressList" :key="index">
        <div class="address">
          <el-input
            v-if="item.Id === editItemId"
            type="text"
            v-model="editAddressInfo"
            :placeholder="editAddressInfo"
            style="width:75%;margin-right:10px"
          ></el-input>
          <span v-else>{{ item.consumer_Address }}</span>
          <span v-if="item.Id === editItemId">
            <el-button
              type="success"
              icon="el-icon-check"
              circle
              size="mini"
              @click="handleEditOne(item.Id,item.consumer_id)"
            ></el-button>
            <el-button
              type="info"
              icon="el-icon-close"
              circle
              size="mini"
              @click="handleOffEditOne"
            ></el-button>
          </span>
          <span v-else>
            <el-button
              type="primary"
              icon="el-icon-edit"
              circle
              size="mini"
              @click="handleEditItem(item)"
            ></el-button>
            <el-button
              @click="handleDelAddress(item)"
              type="danger"
              icon="el-icon-delete"
              circle
              size="mini"
            >
            </el-button
          ></span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import { reqAddressAll, reqDelAddress, reqEditAddressOne } from "@/api";
export default {
  name: "myaddress",
  data() {
    return {
      consumerId: "",
      addressList: [],
      editItemId: 0,
      editAddressInfo: "",
    };
  },
  mounted() {
    var consumerInfo = localStorage.getItem("consumerInfo");
    var consumer = JSON.parse(consumerInfo);
    this.consumerId = consumer.id.toString();
    this.getAddressList();
  },
  methods: {
    async getAddressList() {
      const result = await reqAddressAll(this.consumerId);
      if (result.code === 0) {
        this.addressList = result.data;
      } else {
        this.$message.error("获取地址列表失败");
      }
    },
    handleDelAddress(item) {
      const h = this.$createElement;
      this.$msgbox({
        title: "删除地址",
        message: h("p", null, [
          h("span", null, "地址是: "),
          h("i", { style: "color:teal" }, item.consumer_Address),
        ]),
        showCancelButton: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        center: true,
      })
        .then(() => {
          // 调用保存地址方法
          this.delAddress(item);
        })
        .catch((action) => {
          this.$message({
            type: "info",
            message: "放弃删除",
          });
        });
    },
    async delAddress(item) {
      const result = await reqDelAddress(item.Id);
      if (result.code === 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        this.$router.push("/about/center");
      } else {
        this.$message.error(result.msg);
      }
    },
    handleToCity() {
      this.$router.push("/city");
    },
    handleEditItem(item) {
      this.editItemId = item.Id;
      this.editAddressInfo = item.consumer_Address;
    },
    async handleEditOne(itemId,consumerId) {
      const result = await reqEditAddressOne(itemId,consumerId,this.editAddressInfo)
      if (result.code === 0) {
        this.$message({
          type:"success",
          message:"修改地址成功"
        })
        this.handleOffEditOne()
        this.getAddressList()
      } else {
        this.$message.error("修改地址失败")
        this.handleOffEditOne()
      }
    },

    handleOffEditOne() {
      this.editItemId = 0;
      this.editAddressInfo = "";
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
#content .cinema_body {
  flex: 1;
  overflow: auto;
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
  font-size: 14px;
  color: #666;
}
.cinema_body .address span:nth-of-type(2) {
  float: right;
}
.cinema_body .card {
  display: flex;
}
.cinema_body .card div {
  margin-top: 20px;
  padding: 0 3px;
  border-radius: 2px;
  color: #f90;
  font-size: 13px;
  margin-right: 5px;
}
</style>
