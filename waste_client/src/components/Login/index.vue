<template>
  <div class="login_body">
    <div>
      <h3
        id="ji-chu-yong-fa"
        style="margin-top:20px;margin-left:20px;color:#49464a"
      >
        用户登录
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
        <el-form-item label="验证码" prop="check_bs64_code" style="width:100%;height:40px; line-height:40px;">
          <el-input
            v-model="ruleForm.check_bs64_code"
            autocomplete="off"
            placeholder="三分钟内有效"
            style="width:45%;margin-right:5px"
          ></el-input>
          <!-- 图形验证码 -->
            <el-image
              style="margin:0;padding:0;height:40px;width:130px ;position: absolute;"
              :src="CaptchaImg"
              fit="none"
              @click="handleToVerifyImg"
            ></el-image>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            style="width:240px"
            @click="submitForm('ruleForm')"
            >登录</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="login_link">
      <router-link to="/about/register">
        <el-button type="text" style="color:#e54847;margin-right:15px;"
          >立即注册</el-button
        >
      </router-link>
      <router-link to="/about/reset/pass">
        <el-button type="text" style="color:#e54847;margin-left:15px"
          >忘记密码</el-button
        >
      </router-link>
    </div>
  </div>
</template>

<script>
import { reqBs64Code,reqLoginByEmail } from "@/api";
export default {
  name: "login",
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
        callback()
      }
    };
    return {
      ruleForm: {
        consumer_email: "",
        consumer_password: "",
        check_bs64_code: "",
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
        check_bs64_code: [
          { required: true, message: "请输入验证码", trigger: "blur" },
        ],
      },
      CaptchaId: "",
      CaptchaImg: "",
    };
  },
  beforeRouteEnter(to, from, next) {
    var isAuth = localStorage.getItem("consumerInfo");
    if (isAuth) {
      next((vm) => {
        vm.$router.push("/about/center");
      });
    } else {
      next();
    }
  },
  mounted() {
    this.getBs64Captcha();
  },
  methods: {
    //获取图形验证码
    async getBs64Captcha() {
      const result = await reqBs64Code();
      if (result.code === 0) {
        console.log(result);
        this.CaptchaId = result.data.Id;
        this.CaptchaImg = result.data.B64s;
      }
    },
    async LoginByEmail() {
      const result = await reqLoginByEmail(this.ruleForm.consumer_email,this.ruleForm.consumer_password,this.CaptchaId,this.ruleForm.check_bs64_code)
      if (result.code === 1) {
        this.$message.error(result.msg)
        return
      } else {
        var consumerInfo = JSON.stringify(result.data)
        localStorage.setItem("consumerInfo",consumerInfo)
        this.$message({
          type:"success",
          message:"登录成功!"
        })
        setTimeout(() => {
          this.$router.push('/about/center')
        },300)
      }
    },
    // 登录
    submitForm(formName) {
      console.log(formName);
      this.$refs[formName].validate((valid) => {
        console.log(valid);
        if (valid) {
          console.log("tologin");
          this.LoginByEmail();
        } else {
          this.$message.error("登录信息输入有误");
          return false;
        }
      });
    },
    handleToVerifyImg() {
      this.getBs64Captcha();
    },
  },
};
</script>

<style scoped>
.login_body .login_link {
  display: flex;
  justify-content: center;
}
</style>
