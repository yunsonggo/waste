<template>
  <div id="main">
    <Header :title="title"></Header>
    <div id="content">
      <div class="movie_menu">
        <router-link tag="div" to="/city" class="city_name">
          <span @click="handleShowCity">
              {{ city }}
          </span>
          <i class="iconfont icon-lower-triangle"></i>
        </router-link>
        <div class="hot_swtich">
          <router-link tag="div" to="/hot" class="hot_item active"
            >热门</router-link
          >
          <router-link tag="div" to="/more" class="hot_item">更多</router-link>
        </div>
        <router-link tag="div" to="/search" class="search_entry">
          <i class="iconfont icon-sousuo"></i>
        </router-link>
        <!-- <router-link tag="div" to="/upload" class="hot_item">上传</router-link> -->
      </div>
      <!-- 二级路由 -->
        <router-view />
    </div>
    <Tabbar></Tabbar>
  </div>
</template>

<script>
import Header from "@/components/Header";
import Tabbar from "@/components/Tabbar";

export default {
  name: "Home",
  data() {
    return {
      showCities: false,
      title:"旧物 - Online ...",
      areaName: "",
    };
  },
  computed: {
    city() {
      if (this.$store.getters.location.district) {
        return this.$store.getters.location.district;
      } else {
        return "地区";
      }
    },
  },
  methods: {
    // 显示隐藏城市组件
    handleShowCity() {
      this.showCities = !this.showCities;
      if (!this.showCities) {
        this.$router.go(-1);
      }
    },
  },
  components: {
    Header,
    Tabbar
  },
};
</script>

<style scoped>
#content .movie_menu {
  width: 100%;
  height: 45px;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  z-index: 10;
}
.movie_menu .city_name {
  margin-left: 20px;
  height: 100%;
  line-height: 45px;
}

.movie_menu .city_name.router-link-avtive {
  color: #ef4238;
  border-bottom: 2px #ef4238 solid;
}
.movie_menu .hot_swtich {
  display: flex;
  height: 100%;
  line-height: 45px;
}
.movie_menu .hot_item {
  font-size: 15px;
  color: #666;
  width: 80px;
  text-align: center;
  margin: 0 12px;
  font-weight: 700;
}

.movie_menu .hot_item.router-link-active {
  color: #ef4238;
  border-bottom: 2px #ef4238 solid;
}
.movie_menu .search_entry {
  margin-right: 20px;
  height: 100%;
  line-height: 45px;
}

.movie_menu .search_entry.router-link-active {
  color: #ef4238;
  border-bottom: 2px #ef4238 solid;
}
.movie_menu .search_entry i {
  font-size: 24px;
  color: red;
}

.slide-enter-active{ animation : 13s detailMove;}
@keyframes detailMove{
	0%{
		transform : translateX(100%);
	}
	100%{
		transform : translateX(0);
	}
}
</style>
