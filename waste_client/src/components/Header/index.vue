<template>
  <header id="header">
    <i class="iconfont icon-right" @click="handleToBack"></i>
    <h1>{{ title }}</h1>
  </header>
</template>

<script>
export default {
  name: "Header",
  props: {
    title: {
      type: String,
      default: "首页",
    },
  },
  methods: {
    handleToBack() {
      // console.log('ontouchstart' in document)//打印是否支持touch事件。
      this.$router.back();
    },
  },
};
</script>

<style scoped>
#header {
  width: 100%;
  height: 50px;
  color: #fff;
  background: #e54847;
  border-bottom: 1px solid #e54847;
  position: relative;
}
#header h1 {
  font-size: 18px;
  text-align: center;
  line-height: 50px;
  font-weight: normal;
}
#header i {
  position: absolute;
  left: 5px;
  top: 50%;
  margin-top: -13px;
  font-size: 26px;
}
</style>
