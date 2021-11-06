<template>
  <div class="search_body">
    <div class="search_input">
      <div class="search_input_wrapper">
        <i class="iconfont icon-sousuo"></i>
        <input type="text" v-model="message"/>
      </div>
    </div>
    <div class="search_result">
      <h3>搜索回收商品</h3>
      <ul>
        <li v-for="(good,index) in goodList" :key="index">
          <div class="img" @touchstart="handleToDetail(good.Id)"><img :src="staticUrl + good.icon" /></div>
          <div class="info" @touchstart="handleToDetail(good.Id)">
            <p><span>{{good.name}}</span><span>现价:{{ good.price.toFixed(2)
              }}{{ "元/" + good.goods_unit }}</span></p>
            <p v-if="good.old_price" class="price-line-through">原价:{{ good.old_price.toFixed(2) + '元/' + good.goods_unit}}</p>
            <p>描述:{{good.description}}</p>
            <p>已回收:{{good.sell_count}}</p>
          </div>
          <div class="btn_mall">回收</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import global from "@/components/Global";
import { reqGoods } from "@/api";
export default {
  name: "Search",
  data() {
    return {
      message: "",
      goodList: [],
      goodsData: [],
      moreGoods: [],
      hotGoods: [],
      staticUrl: global.static_url,
    };
  },
  watch: {
    message(newVal) {
      this.goodList = []
      if (this.goodsData.length > 0 && newVal != "") {
        this.goodsData.forEach((item) => {
          if (item.name.indexOf(newVal) >= 0) {
            this.goodList.push(item);
          } else if (item.description.indexOf(newVal) >= 0) {
            this.goodList.push(item);
          }
        });
        if (this.goodList.length <= 0) {
          var keyArray = newVal.split("")
          keyArray.forEach((key) => {
            this.goodsData.forEach((good) => {
              if (good.name.indexOf(key) >= 0) {
                this.goodList.push(good);
              } else if (good.description.indexOf(key) >= 0) {
                this.goodList.push(good);
              }
            })
          })
        }
      }
    },
  },
  mounted() {
    this.goodList = []
    this.getGoodsAll();
  },
  methods: {
    getGoodsAll() {
      this.goodsData = [];
      if (this.$store.getters.goodsHot.length) {
        this.goodsData = this.$store.getters.goodsHot;
        this.moreGoods = this.$store.getters.goodsMore;
        this.moreGoods.forEach((item) => {
          this.goodsData.push(item);
        });
      } else {
        this.getGoodsData();
      }
    },
    // 获取商品数据
    async getGoodsData() {
      const result = await reqGoods();
      if (result.code === 0) {
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
    // bscroll tag  事件 跳转详情页
    handleToDetail(id) {
      this.$router.push("/detail/more/" + id)
    },
  },
};
</script>

<style scoped>
#content .search_body{ flex:1; overflow:auto;}
.search_body .search_input{ padding: 8px 10px; background-color: #f5f5f5; border-bottom: 1px solid #e5e5e5;}
.search_body .search_input_wrapper{ padding: 0 10px; border: 1px solid #e6e6e6; border-radius: 5px; background-color: #fff; display: flex; line-height: 20px;}
.search_body .search_input_wrapper i{font-size: 16px; padding: 4px 0;}
.search_body .search_input_wrapper input{ border: none; font-size: 13px; color: #333; padding: 4px 0; outline: none; margin-left: 5px; width:100%;}

.search_body .search_result h3{ font-size: 15px; color: #999; padding: 9px 15px; border-bottom: 1px solid #e6e6e6;}

.search_body .search_result li{ border-bottom:1px #c9c9c9 dashed; padding: 10px 15px; box-sizing:border-box; display: flex;}
.search_body .search_result .img{ width: 60px; float:left; }
.search_body .search_result .img img{ width: 100%; }
.search_body .search_result .info{ float:left; margin-left: 15px; flex:1;}
.search_body .search_result .info p{ height: 22px; display: flex; line-height: 22px; font-size: 12px;}
.search_body .search_result .info p:nth-of-type(1) span:nth-of-type(1){ font-size: 18px; flex:1; }
.search_body .search_result .info p:nth-of-type(1) span:nth-of-type(2){ font-size: 16px; color:#fc7103; margin-right: 20px;}
.search_body .search_result .info .price-line-through {
  margin-left: 15px;
  text-decoration: line-through;
}
.search_body .search_result  .btn_mall,
.search_body .search_result  .btn_pre {
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
.search_body .search_result  .btn_pre {
  background-color: #3c9fe6;
}
</style>