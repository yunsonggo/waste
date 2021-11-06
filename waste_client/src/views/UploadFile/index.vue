<template>
  <div>
    <el-form>
      <el-upload
        class="upload-demo"
        :show-file-list="false"
        :before-upload="beforeUpload"
        :on-success="onSuccess"
        :on-error="onError"
        action="http://192.168.1.102:8080/api/excel/upload"
        multiple="multiple"
        name="excel_file"
      >
        <el-button size="small" type="primary">点击上传</el-button>
        <div slot="tip" class="el-upload__tip">只能上传xlsx文件</div>
      </el-upload>
    </el-form>

    <div v-if="filePath">{{ filePath }}</div>
  </div>
</template>

<script>
export default {
  name: "uploadfile",
  data() {
    return {
      filePath: "",
    };
  },
  methods: {
    beforeUpload(file) {
      const extXls = file.name.split(".")[1] === "xls";
      const extXlsx = file.name.split(".")[1] === "xlsx";
      if (!extXls && !extXlsx) {
        this.$message.error("上传失败,上传只能是 xls、xlsx格式!");
      }
      return extXls || extXlsx;
    },
    onSuccess(res, file) {
      if (res != "") {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        console.log(res);
        this.filePath = res.path;
      } else {
        this.$message.error("上传失败");
      }
    },
    onError(res, file) {
      this.$message.error("上传失败");
    },
  },
};
</script>

<style scoped>
</style>