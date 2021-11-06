<template>
  <div class="register_body">
    <h3
      id="ji-chu-yong-fa"
      style="margin-top:20px;margin-left:20px;color:#49464a"
    >
      修改信息
    </h3>
    <el-form
      :model="ruleForm"
      :rules="rules"
      ref="ruleForm"
      label-width="80px"
      class="demo-dynamic"
      style="margin-top:25px"
    >
      <el-form-item label="邮箱" prop="consumer_email" style="width:100%">
        <el-tooltip
          class="item"
          effect="light"
          content="注册邮箱不能修改"
          placement="top"
        >
          <el-input
            v-model="ruleForm.consumer_email"
            :disabled="true"
            style="width:90%;margin-right:10px"
          ></el-input>
        </el-tooltip>
      </el-form-item>
      <el-form-item label="姓名" prop="consumer_name" style="width:100%">
        <el-input
          v-model="ruleForm.consumer_name"
          style="width:90%;margin-right:10px"
        ></el-input>
      </el-form-item>
      <el-form-item label="昵称" prop="consumer_nickname" style="width:100%">
        <el-input
          v-model="ruleForm.consumer_nickname"
          style="width:90%;margin-right:10px"
        ></el-input>
      </el-form-item>
      <el-form-item label="电话" prop="consumer_mobile" style="width:100%">
        <el-input
          v-model="ruleForm.consumer_mobile"
          style="width:45%;margin-right:5px"
        ></el-input> <span style="font-size:14px;color:#e54847"> *微信绑定电话号码</span>
      </el-form-item>
      <el-form-item label="性别" prop="consumer_gender" style="width:100%">
        <el-radio v-model="ruleForm.consumer_gender" label="先生"
          >先生</el-radio
        >
        <el-radio v-model="ruleForm.consumer_gender" label="女士"
          >女士</el-radio
        >
      </el-form-item>
      <el-form-item label="上传头像" prop="consumer_icon">
        <el-upload
          class="avatar-uploader"
          action="http://192.168.1.102:8001/api/consumer/avatar/upload"
          name="consumer_icon"
          :data="upParamData"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
          multiple="multiple"
        >
          <img
            v-if="ruleForm.consumer_icon"
            :src="static_url + ruleForm.consumer_icon"
            class="avatar"
          />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
      <el-form-item label="收款码" prop="consumer_wechar_icon">
        <el-upload
          class="avatar-uploader"
          action="http://192.168.1.102:8001/api/consumer/wechar/icon/upload"
          name="consumer_wechar_icon"
          :data="upWecharParamData"
          :show-file-list="false"
          :on-success="handleWecharSuccess"
          :before-upload="beforeWecharUpload"
          multiple="multiple"
        >
          <img
            v-if="ruleForm.consumer_wechar_icon"
            :src="static_url + ruleForm.consumer_wechar_icon"
            class="wecharIcon"
          />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
      <el-form-item>
        <el-button type="danger" @click="submitForm('ruleForm')"
          >立即修改</el-button
        >
        <el-button @click="handelToBack">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
// import { setInterval } from 'timers';
import { reqEmailCode, reqEditByEmail } from "@/api";
import global from "@/components/Global"

export default {
  name: "EmailSign",
  data() {
    return {
        static_url:global.static_url,
      id: Number,
      email: "",
      upParamData: {
        old_avatar_url: "",
        consumer_id: "",
      },
      upWecharParamData: {
        old_wechar_url:"",
        consumer_id:"",
        consumer_token:"",
      },
      ruleForm: {
        consumer_email: "",
        consumer_gender: "",
        consumer_mobile: "",
        consumer_name: "",
        consumer_nickname: "",
        consumer_icon: "",
        consumer_wechar_icon:"",
      },
      rules: {
        consumer_mobile: [
          { required: true, message: "请输入电话号码", trigger: "blur" },
        ],
      },
    };
  },
  mounted() {
    var consumerInfoJson = localStorage.getItem("consumerInfo");
    if (consumerInfoJson) {
      var consumerInfo = JSON.parse(consumerInfoJson);
      this.id = consumerInfo.id;
      this.email = consumerInfo.consumer_email;
      this.ruleForm.consumer_email = consumerInfo.consumer_email;
      this.ruleForm.consumer_gender = consumerInfo.consumer_gender;
      this.ruleForm.consumer_mobile = consumerInfo.consumer_mobile;
      this.ruleForm.consumer_name = consumerInfo.consumer_name;
      this.ruleForm.consumer_nickname = consumerInfo.consumer_nickname;
      this.ruleForm.consumer_icon = consumerInfo.consumer_icon;
      this.ruleForm.consumer_wechar_icon = consumerInfo.consumer_wechar_icon;
      this.upParamData.old_avatar_url = consumerInfo.consumer_icon;
      this.upParamData.consumer_id = consumerInfo.id;
      // this.upWecharParamData.old_wechar_url = consumer.consumer_wechar_icon;
      this.upWecharParamData.consumer_id = consumerInfo.id;
      this.upWecharParamData.consumer_token = consumerInfo.consumer_password;
    }
  },
  methods: {
    // 注册请求
    async editConsumerInfo() {
      const result = await reqEditByEmail(
        this.id,
        this.ruleForm.consumer_email,
        this.ruleForm.consumer_name,
        this.ruleForm.consumer_nickname,
        this.ruleForm.consumer_mobile,
        this.ruleForm.consumer_gender
      );
      if (result.code == 1) {
        this.$message.error(result.msg);
      } else if (result.code === 0) {
        var consumerResInfo = JSON.stringify(result.data);
        localStorage.removeItem("consumerInfo");
        localStorage.setItem("consumerInfo", consumerResInfo);
        this.$message({
          message: "修改成功!",
          type: "success",
        });
        setTimeout(() => {
          this.$router.push("/about/center");
        }, 500);
      }
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.editConsumerInfo();
        } else {
          this.$message.error("信息输入有误");
          return false;
        }
      });
    },
    // resetForm(formName) {
    //   this.$refs[formName].resetFields();
    //   this.ruleForm.consumer_email = this.email;
    //   this.ruleForm.consumer_icon = this.upParamData.old_avatar_url;
    //   this.ruleForm.consumer_wechar_icon = this.upWecharParamData.old_wechar_url
    // },
    handelToBack() {
      this.$router.back()
    },
    // 头像上传成功
    handleAvatarSuccess(res, file) {
      this.ruleForm.consumer_icon = res.path;
      if (res != "") {
        this.$message({
          message: "上传成功",
          type: "success",
        });
      } else {
        this.$message.error("上传失败");
      }
    },
    // 上传验证
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      if (isJPG && isLt2M) {
        this.upParamData.old_avatar_url = this.ruleForm.consumer_icon;
        this.upParamData.consumer_id = this.id;
      } else {
        this.upParamData.old_avatar_url = "";
      }
      //console.log(this.upParamData.old_avatar_url);
      //console.log(this.upParamData.admin_id);
      return isJPG && isLt2M;
    },
    // 收款码上传成功
    handleWecharSuccess(res, file) {
      this.ruleForm.consumer_wechar_icon = res.path;
      if (res != "") {
        this.$message({
          message: "上传成功",
          type: "success",
        });
      } else {
        this.$message.error("上传失败");
      }
    },
    // 收款码上传验证
    beforeWecharUpload(file) {
      const isPNG = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isPNG) {
        this.$message.error("上传图片只能是 PNG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      if (isPNG && isLt2M) {
        this.upWecharParamData.old_wechar_url = this.ruleForm.consumer_wechar_icon;
        this.upWecharParamData.consumer_id = this.id;
      } else {
        this.upWecharParamData.old_wechar_url = "";
      }
      //console.log(this.upParamData.old_avatar_url);
      //console.log(this.upParamData.admin_id);
      return isPNG && isLt2M;
    },
  },
};
</script>

<style scoped>
#content .register_body {
  width: 100%;
}
.el-upload {
    display: inline-block;
    text-align: center;
    cursor: pointer;
    outline: none;
}
  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 98px;
    height: 98px;
    line-height: 98px;
    text-align: center;
  }
  .avatar {
    width: 98px;
    height: 98px;
    display: block;
  }
    .wecharIcon {
    width: 98px;
    height: 138px;
    display: block;
  }
</style>
