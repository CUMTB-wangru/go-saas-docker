<template>
  <div id="login">
    <el-radio-group v-model="isCollapse" id="login-top">
      <el-radio-button :label="false">&emsp;密&ensp;码&ensp;登&ensp;录&emsp;</el-radio-button>
      <el-radio-button :label="true">&emsp;短&ensp;信&ensp;登&ensp;录&emsp;</el-radio-button>
    </el-radio-group>
    <el-form
      ref="formLoginRef"
      :model="formLogin"
      :rules="rules"
      label-position="top"
      label-width="100px"
      size="mini"
      v-if="isCollapse === false"
    >
      <el-form-item label="邮箱或手机号" prop="name" class="item">
        <el-input v-model.trim="formLogin.name" placeholder="请输入邮箱或手机号"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" class="item">
        <el-input v-model.trim="formLogin.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-checkbox v-model="remember">记住密码</el-checkbox>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="userLogin">登录</el-button>
      </el-form-item>
    </el-form>
    <el-form
      ref="formSmsLoginRef"
      :model="formSmsLogin"
      :rules="rulesSms"
      label-position="top"
      label-width="100px"
      size="mini"
      v-if="isCollapse === true"
    >
      <el-form-item label="手机号" prop="name" class="item">
        <el-input v-model.trim="formSmsLogin.phone" placeholder="请输入手机号"></el-input>
      </el-form-item>
      <el-form-item label="短信验证码" prop="code" class="item">
        <el-col :span="24">
          <el-input v-model.trim="formSmsLogin.code" placeholder="请输入短信验证码" style="width: 90%;"></el-input>
        </el-col>
        <el-button type="info">获取验证码</el-button>
      </el-form-item>
      <el-form-item>
        <router-link :to="{ path: 'register' }">没有账号？立即注册...</router-link>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="userSmsLogin">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { login, loginSms } from "/@/api/basic";
import { ElMessage } from 'element-plus';
// 引入腾讯防水墙
import '/@/assets/js/TCaptcha.js';
export default defineComponent({
  name: "Login",
  setup() {
    // 状态选择
    const isCollapse = ref(false);
    // 记住密码
    const remember = ref(false);
    // 密码登录
    let formLogin = reactive({
      name: null,
      password: null,
      ticket: null,
      randstr: null,
    })
    const formLoginRef = ref();
    const rules = {
      name: [
        {
          required: true,
          message: '用户名不能为空',
          trigger: 'blur',  // 失去焦点时验证
        }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 18, message: '请输入6-18密码', trigger: 'blur' },
      ]
    }
    // 短信登录
    const formSmsLogin = reactive({
      phone: null,
      code: null,
    })
    const formSmsLoginRef = ref(null)
    const rulesSms = {
      phone: [
        { required: true, message: '请输入手机号码', trigger: 'blur' },
        { min: 11, max: 11, message: '请输入11位手机号码', trigger: 'blur' },
        {
          pattern: /^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$/,
          message: '请输入正确的手机号码'
        }
      ],
      code: [
        {
          required: true,
          message: '验证码不能为空',
          trigger: 'blur',  // 失去焦点时验证
        }, {
          min: 6,
          max: 6,
          message: '请输入验证码',
          trigger: 'blur',  // 失去焦点时验证
        },
      ],
    }


    // 防水墙
    let captcha = () => {
      try {
        // 这里要配置成常量文件，不然容易被攻击
        let captcha1 = new TencentCaptcha(import.meta.env.VITE_CAPTCHA_APP_ID, (res: any) => {
          if (res.ret === 0) {
            formLogin.ticket = res.ticket;
            formLogin.randstr = res.randstr;
            login(formLogin).then((response: any) => {
              ElMessage({
                showClose: true,
                message: response.message,
                type: 'success',
              })
              if (remember) {
                // 记住密码
                localStorage.token = response.data.token;
                sessionStorage.removeItem("token");
              } else {
                // 不记住密码
                sessionStorage.token = response.data.token;
                localStorage.removeItem("token");
              }
              // 用户基本信息
              localStorage.user_avatar = response.data.avatar;
              localStorage.user_name = response.data.nickname;
            }).catch((err: any) => {
              ElMessage({
                showClose: true,
                message: err.message,
                type: 'error',
              })
            })
          }
        });
        captcha1.show();  // 显示验证码
      } catch (error) {
        loadErrorCallback;
      }
    }

    let callback = (res: any) => {
      /* res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
         res（客户端出现异常错误 仍返回可用票据） = {ret: 0, ticket: "String", randstr: "String", errorCode: Number, errorMessage: "String"}
         res（用户主动关闭验证码）= {ret: 2}
      */
      if (res.ticket) {
        // 上传票据 可根据errorCode和errorMessage做特殊处理或统计
      }
    }
    // 验证码js加载错误处理
    let loadErrorCallback = () => {
      var CaptchaAppId = ''
      /* 生成票据或自行做其它处理 */
      var ticket = 'terror_1001_' + CaptchaAppId + Math.floor(new Date().getTime() / 1000);
      callback({
        ret: 0,
        randstr: '@' + Math.random().toString(36).substr(2),
        ticket: ticket,
        errorCode: 1001,
        errorMessage: 'jsload_error',
      });
    }

    // 登录
    let userLogin = () => {
      formLoginRef.value.validate((valid: any) => {
        if (valid) {
          captcha();
        } else {
          return false
        }
      })
    };
    // 短信登录
    let userSmsLogin = () => {
      formLoginRef.value.validate((valid: any) => {
        if (valid) {
          loginSms(formLogin).then((response: any) => {
            ElMessage({
              showClose: true,
              message: response.message,
              type: 'success',
            })
            // 用户基本信息
            localStorage.user_avatar = response.data.avatar;
            localStorage.user_name = response.data.nickname;
          }).catch((err: any) => {
            ElMessage({
              showClose: true,
              message: err.message,
              type: 'error',
            })
          })
        } else {
          return false
        }
      })
    }
    return {
      remember,
      isCollapse,
      // 密码登录
      formLogin,
      rules,
      formLoginRef,
      userLogin,
      // 短信登陆
      formSmsLogin,
      formSmsLoginRef,
      rulesSms,
      userSmsLogin,
    };
  },
});
</script>

<style scoped>
#login {
  width: 300px;
  top: 0;
  bottom: 0px;
  left: 0;
  right: 0;
  margin: 80px auto 415px auto;
  padding: 30px 25px;
  border: 1px solid #c0c0c0;
  border-radius: 8px;
  box-shadow: 5px 5px 4px #c0c0c0;
}
#login-top {
  width: 300px;
  text-align: center;
  margin-bottom: 20px;
}
:deep(.el-form-item__label) {
  /* vue3.0深度解析，vue的scope属性会禁止第三方的组件修改，可用deep()深度解析来执行操作 */
  padding: 0 !important;
}
:deep(.el-form-item) {
  margin-bottom: 8px !important;
}
</style>
