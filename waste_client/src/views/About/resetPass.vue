<template>
  <div class="register_body">
    <h3
      id="ji-chu-yong-fa"
      style="margin-top:20px;margin-left:20px;color:#49464a"
    >
      找回密码
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
        <el-input
          v-model="ruleForm.consumer_email"
          style="width:90%;margin-right:10px"
        ></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="consumer_password">
        <el-input
          type="password"
          v-model="ruleForm.consumer_password"
          autocomplete="off"
          style="width:90%;"
          placeholder="6位以上大小写字母+数字"
        ></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="consumer_check_password">
        <el-input
          type="password"
          v-model="ruleForm.consumer_check_password"
          autocomplete="off"
          style="width:90%;"
        ></el-input>
      </el-form-item>
      <el-form-item label="验证码" prop="check_email_code" style="width:100%;">
        <el-input
          v-model="ruleForm.check_email_code"
          autocomplete="off"
          placeholder="三分钟内有效"
          style="width:35%;margin-right:10px"
        ></el-input>
        <!-- 请求邮箱验证码 -->
          <el-button
            type="danger"
            plain
            size="small"
            @click="checkAndPostEmailCode"
            :disabled="BtnDisable"
          >
            获取验证码
          </el-button>
          <span v-if="BtnDisable" style="margin-left:5px;color:#ccc;font-size:14px">({{DisableTime}})</span>
        <!-- 图形验证码 -->
        <!-- <el-image :src="CaptchaImg" fit="cover" style="height:45px;margin:0;padding:0"></el-image> -->
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')"
          >立即提交</el-button
        >
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
// import { setInterval } from 'timers';
import { reqEmailCode, reqResetPassByEmail } from "@/api";

export default {
  name: "EmailSign",
  data() {
    var validatePassRule = (value) => {
      const reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{6,20}$/;
      return reg.test(value);
    };
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else if (!validatePassRule(value)) {
        callback(new Error("密码由至少6位大写字母、小写字母和数字组成"));
      } else {
        if (this.ruleForm.consumer_check_password !== "") {
          this.$refs.ruleForm.validateField("consumer_check_password");
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.consumer_password) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        consumer_email: "",
        consumer_password: "",
        consumer_check_password: "",
        check_email_code: "",
      },
      rules: {
        consumer_email: [
          { required: true, message: "请输入邮箱地址", trigger: "blur" },
          {
            type: "email",
            message: "请输入正确的邮箱地址",
            trigger: ["blur", "change"],
          },
        ],
        consumer_password: [
          { required: true, validator: validatePass, trigger: "blur" },
        ],
        consumer_check_password: [
          { required: true, validator: validatePass2, trigger: "blur" },
        ],
        check_email_code: [
          { required: true, message: "请输入验证码", trigger: "blur" },
        ],
      },
      CaptchaId: "",
      CaptchaImg: "",
      BtnDisable:false,
      DisableTime:60,
    };
  },
  mounted() {
    // this.getBase64Code();
  },
  methods: {
    // 验证邮箱 并获取 验证码
    checkAndPostEmailCode() {
      console.log("checkemail");
      this.$refs.ruleForm.validateField("consumer_email", (callback) => {
        if (callback) {
          this.$message.error("邮箱地址错了哦,检查一下吧");
          return;
        }
        this.getEmailCode();
      });
    },
    // 请求邮箱验证码
    async getEmailCode() {
      console.log(this.ruleForm.consumer_email);
      const result = await reqEmailCode(this.ruleForm.consumer_email);
      console.log(result);
      if (result.code === 0) {
        this.$message({
          message: "验证码发送成功,去邮箱看看吧",
          type: "success",
        });
        this.BtnDisable = true
        this.timer = setInterval(() => {
          this.DisableTime --
          if (this.DisableTime === 0) {
            this.DisableTime = 60
            this.BtnDisable = false
            clearInterval(this.timer)
          }
        },1000)
      } else {
        this.$message.error("发送验证码失败");
      }
    },
    // 注册请求
    async resetPassByEmail() {
      const result = await reqResetPassByEmail(
        this.ruleForm.consumer_email,
        this.ruleForm.consumer_password,
        this.ruleForm.consumer_check_password,
        this.ruleForm.check_email_code
      );
      if (result.code == 1) {
        this.$message.error(result.msg)
      } else if (result.code == 0) {
        this.$message({
          message:"密码重置成功!",
          type:"success"
        })
        setTimeout(()=>{
          this.$router.push('/about/login')
        },500)
      }
      console.log(result);
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.resetPassByEmail();
        } else {
          this.$message.error("注册信息输入有误");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>

<style scoped>
#content .register_body {
  width: 100%;
}
</style>
