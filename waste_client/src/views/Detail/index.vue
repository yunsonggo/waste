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
          <div class="detail_list">
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
                <p>描述: {{ goodInfo.description }}</p>
                <p>已回收: {{ goodInfo.sell_count }} 单</p>
              </div>
            </div>
          </div>
          <div class="detail_intro">
            <el-divider content-position="left">回收价值</el-divider>
            <div class="good_content">{{ goodInfo.content }}</div>
          </div>
          <div style="height: 350px;margin-left:20px;">
            <el-divider content-position="left">回收步骤</el-divider>
            <el-steps :active="2" direction="vertical">
              <el-step
                title="1. 分类打包,准备下单"
                icon="el-icon-box"
                description="推荐垃圾分类,文明濮阳更有魅力"
              ></el-step>
              <el-step
                title="2. 网上下单,生成单号"
                icon="el-icon-mobile-phone"
                description="推荐分类下单,畅享健康生活"
              ></el-step>
              <el-step
                title="3. 给旧物标记单号"
                icon="el-icon-edit"
                description="切记标记单号,收钱全靠它"
              ></el-step>
              <el-step
                title="4. 投入回收箱"
                icon="el-icon-data-board"
                description="顺路投入回收箱,节省时间更好生活"
              ></el-step>
              <el-step
                title="5. 统一收回"
                icon="el-icon-shopping-cart-full"
                description="回收员统一收回,中心仓库称重计量"
              ></el-step>
              <el-step
                title="6. 平台支付"
                icon="el-icon-bank-card"
                description="添加好友,确认单号,支付到账,欢迎再来"
              ></el-step>
            </el-steps>
          </div>
          <div style="height:50px;margin-top:50px;margin-left:35%">
            <el-button type="danger" plain @click="handelAddOrder(goodInfo.Id)">立即回收</el-button>
          </div>
        </div>
      </div>
    </Scroller>
  </div>
</template>

<script>
import Header from "@/components/Header";
import { reqGoods } from "@/api";
import global from "@/components/Global";
import Scroller from "@/components/Scroller";

export default {
  name: "Detail",
  components: {
    Header,
    Scroller,
  },
  props: {
    detailId: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      title:"详情",
      staticUrl: global.static_url,
      isLoading: true,
      allGoodsData: [],
      goodInfo: [],
      pullDownMsg: "",
    };
  },
  mounted() {
    this.getAllGoodsData();
  },
  methods: {
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
      if (this.detailId) {
        this.allGoodsData.forEach((item) => {
          if (item.Id == this.detailId) {
            this.goodInfo = item;
            this.isLoading = false;
            console.log(this.goodInfo);
            this.$nextTick(() => {
              this.$refs.scro.scroll.refresh();
              // refScroll()
            });
          }
        });
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
    // 提交订单
    handelAddOrder(id) {
      this.$router.push("/order/add/" + id)
    },
  },
};
</script>

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
</style>
