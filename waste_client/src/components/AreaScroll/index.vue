<template>
  <div>
    <el-drawer
      :visible.sync="showMap"
      :with-header="false"
      size="35%"
      direction="btt"
      ref="drawer"
      :destroy-on-close="true"
    >
      <template>
        <ul class="infinite-list" style="overflow:auto">
          <li
            v-for="(item, index) in areaList"
            :key="index"
            class="infinite-list-item"
          >
            <div @click="$emit('selectItem', item)">
              <el-tag type="danger" effect="dark">{{ item.name }}</el-tag>
              <el-tag type="danger">{{ item.district }}</el-tag
              ><br />
              {{ item.address }}
            </div>
            <el-divider></el-divider>
          </li>
        </ul>
      </template>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "areaScroll",
  props: {
    areaList: {
      type: Array,
      default: null,
    },
    showMap: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
  },
};
</script>

<style scoped></style>
