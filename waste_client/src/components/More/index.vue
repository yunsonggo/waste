<template>
  <div class="movie_body">
    <Scroller :handleToScroll="handleToScroll" :handleToTouchEnd="handleToTouchEnd">
      <ul v-if="moreGoods">
        <li class="pullDown">{{ pullDownMsg }}</li>
        <li v-for="(moreGood, index) in moreGoods" :key="index">
          <div class="pic_show" @tap="handleToDetail(moreGood.Id)">
            <img :src="staticUrl + moreGood.icon" alt="" />
          </div>
          <div class="info_list" @tap="handleToDetail(moreGood.Id)">
            <h2>{{ moreGood.name }} <img v-if="moreGood.tage" src="@/assets/hotsale.png" alt=""></h2>
            <p>
              单价:
              <span class="grade"
                >{{ moreGood.price.toFixed(2)
                }}{{ "元/" + moreGood.goods_unit }}</span
              >
              <span v-if="moreGood.old_price" class="price-line-through"
                >原价:{{ moreGood.old_price.toFixed(2) + "元/" + moreGood.goods_unit }}</span
              >
            </p>
            <p>描述:{{ moreGood.description }}</p>
            <p>已回收:{{ moreGood.sell_count }}</p>
          </div>
          <div class="btn_pre">回收</div>
        </li>
      </ul>
    </Scroller>
  </div>
</template>

<script>
import global from "@/components/Global";
import { reqGoods } from "@/api";
import Scroller from "@/components/Scroller";
export default {
  name: "More",
  data() {
    return {
      hotGoods: [],
      moreGoods: [],
      staticUrl: global.static_url,
      pullDownMsg: "",
    };
  },
  mounted() {
    this.getMoreGoods();
  },
  methods: {
    getMoreGoods() {
      if (this.$store.getters.goodsMore.length) {
        this.moreGoods = this.$store.getters.goodsMore;
      } else {
        this.getGoodsData();
      }
    },
    // 获取商品数据
    async getGoodsData() {
      const result = await reqGoods();
      if (result.code === 0) {
        this.$store.dispatch("setGoodsAll", result.data);
        result.data.forEach((item) => {
          if (item.is_hot == 1) {
            this.hotGoods.push(item);
          } else {
            this.moreGoods.push(item);
          }
        });
      }
      if (this.hotGoods) {
        this.$store.dispatch("setGoodsHot", this.hotGoods);
      }
      if (this.moreGoods) {
        this.$store.dispatch("setGoodsMore", this.moreGoods);
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
    // bscroll tag  事件 跳转详情页
    handleToDetail(id) {
      this.$router.push("/detail/more/" + id)
    },
  },
  components: {
    Scroller
  }
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