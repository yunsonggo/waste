<template>
  <div class="wrapper" ref="wrapper">
    <slot></slot>
  </div>
</template>

<script>
import BScroll from "better-scroll";
export default {
  name: "Scroller",
  data() {
      return {
          scroll:Object
      }
  },
  props: {
    // 滚动触发函数
    handleToScroll: {
      type: Function,
      default: function () {},
    },
    // 滚动结束触发
    handleToTouchEnd: {
      type: Function,
      default: function () {},
    },
  },
  mounted() {
      this.init()
  },
  methods: {
    init() {
      // 包裹容器 和  配置参数
      var scroll = new BScroll(this.$refs.wrapper, {
        tap: "tap",
        probeType: 1,
      });
      this.scroll = scroll
      // 滚动之前重新计算高度
      scroll.on('beforeScrollStart', () => {
          scroll.refresh()
      })
      // 滚动触发事件
      scroll.on("scroll", (pos) => {
        this.handleToScroll(pos);
      });
      // 滚动结束 触发事件
      scroll.on("touchEnd", (pos) => {
        this.handleToTouchEnd(pos);
      });
      // scroll.on('refresh', () => {
      //   console.log("reflelelelele")
      // })
    },
    // 坐标 跳转方法
    toScrollTop(y) {
        this.scroll.scrollTo(0,y)
    },
    toStarScroll() {
      this.scroll.refresh()
    }
  },
};
</script>

<style scoped>
.wrapper {
  height: 100%;
}
</style>