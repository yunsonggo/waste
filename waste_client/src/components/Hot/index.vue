<template>
  <div class="movie_body" ref="movie_body">
    <Scroller :handleToScroll="handleToScroll" :handleToTouchEnd="handleToTouchEnd" ref="scro">
      <ul>
        <li class="pullDown">{{ pullDownMsg }}</li>
        <li v-for="(hotGood, index) in hotGoods" :key="index">
          <!-- 添加点击事件 使用 bscroll 提供的 tag事件 跳转详情页-->
          <div class="pic_show" @tap="handleToDetail(hotGood.Id)">
            <img :src="staticUrl + hotGood.icon" alt="" />
          </div>
          <div class="info_list" @tap="handleToDetail(hotGood.Id)">
            <h2>
              {{ hotGood.name }}
              <img v-if="hotGood.tage" src="@/assets/hotsale.png" alt="" />
            </h2>
            <p>
              单价:
              <span class="grade"
                >{{ hotGood.price.toFixed(2)
                }}{{ "元/" + hotGood.goods_unit }}</span
              >
              <span v-if="hotGood.old_price" class="price-line-through"
                >原价:{{
                  hotGood.old_price.toFixed(2) + "元/" + hotGood.goods_unit
                }}</span
              >
            </p>
            <p>描述:{{ hotGood.description }}</p>
            <p>已回收:{{ hotGood.sell_count }}</p>
          </div>
          <div class="btn_mall" @tap="handelAddOrder(hotGood.Id)">回收</div>
        </li>
      </ul>
    </Scroller>
  </div>
</template>

<script>
import { reqGoods } from "@/api";
import global from "@/components/Global";
// import BScroll from "better-scroll";
import Scroller from "@/components/Scroller";


export default {
  name: "Hot",
  data() {
    return {
      hotGoods: [],
      moreGoods: [],
      staticUrl: global.static_url,
      pullDownMsg: "",
      toStartScroll:false
    };
  },
  mounted() {
    this.getGoodsData();
  },
  methods: {
    // 获取商品数据
    async getGoodsData() {
      this.refScroll = false;
      const result = await reqGoods();
      if (result.code === 0) {
        this.$store.dispatch("setGoodsAll",result.data);
        result.data.forEach((item) => {
          if (item.is_hot == 1) {
            this.hotGoods.push(item);
          } else {
            this.moreGoods.push(item);
          }
        });
      }
      this.$nextTick(() => {
          this.$refs.scro.scroll.refresh()
          // refScroll() 
          });
      if (this.hotGoods) {
        this.$store.dispatch("setGoodsHot", this.hotGoods);
        // this.$refs.scro.toStartScroll()
      }
      if (this.moreGoods) {
        this.$store.dispatch("setGoodsMore", this.moreGoods);
      }
    },
    // bscroll tag  事件 跳转详情页
    handleToDetail(id) {
      this.$router.push("/detail/hot/" + id)
    },
    // 提交订单
    handelAddOrder(id) {
      this.$router.push("/order/add/" + id)
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
  components: {
    Scroller,
  },
};
</script>

<style scoped>
#content .movie_body {
  flex: 1;
  overflow: auto;
}
.movie_body ul {
  margin: 0 12px;
  overflow: hidden;
}
.movie_body ul li {
  margin-top: 12px;
  display: flex;
  align-items: center;
  border-bottom: 1px #e6e6e6 solid;
  padding-bottom: 10px;
}
.movie_body .pic_show {
  width: 64px;
  height: 90px;
}
.movie_body .pic_show img {
  width: 100%;
}
.movie_body .info_list {
  margin-left: 10px;
  flex: 1;
  position: relative;
}
.movie_body .info_list h2 {
  font-size: 17px;
  line-height: 24px;
  width: 150px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.movie_body .info_list p {
  font-size: 13px;
  color: #666;
  line-height: 22px;
  width: 200px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.movie_body .info_list .grade {
  font-weight: 700;
  color: #faaf00;
  font-size: 15px;
}
.movie_body .info_list .price-line-through {
  margin-left: 15px;
  text-decoration: line-through;
}
.movie_body .info_list img {
  width: 50px;
  position: absolute;
  right: 10px;
  top: 5px;
}
.movie_body .btn_mall,
.movie_body .btn_pre {
  width: 47px;
  height: 27px;
  line-height: 28px;
  text-align: center;
  background-color: #f03d37;
  color: #fff;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
}
.movie_body .btn_pre {
  background-color: #3c9fe6;
}
.movie_body .pullDown {
  margin: 0;
  padding: 0;
  border: none;
}
</style>