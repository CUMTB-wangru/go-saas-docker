<template>
  <div id="register">
    <el-radio-group id="register-top">
      <el-radio-button :label="false">&ensp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;用&ensp;户&ensp;注&ensp;册&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;</el-radio-button>
    </el-radio-group>
    <el-form
      ref="formRegisterRef"
      :model="formRegister"
      :rules="rules"
      label-position="top"
      label-width="100px"
      size="mini"
    >
      <el-form-item label="用户名" prop="name" class="item">
        <el-input v-model.trim="formRegister.name" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email" class="item">
        <el-input v-model.trim="formRegister.email" placeholder="请输入邮箱"></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="phone" class="item">
        <el-input v-model.trim="formRegister.phone" placeholder="请输入手机号"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" class="item">
        <el-input v-model.trim="formRegister.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="sure_password" class="item">
        <el-input v-model.trim="formRegister.sure_password" placeholder="请输入确认密码"></el-input>
      </el-form-item>
      <el-form-item label="获取验证码" prop="code" class="item">
        <el-col :span="24">
          <el-input v-model.trim="formRegister.code" placeholder="请输入验证码" style="width: 90%;"></el-input>
        </el-col>
        <el-button type="info" @click="send_sms">{{ get_sms_text }}</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" style="margin-top: 8px;" @click="userRegister" prop="rule_all">注册</el-button>
        <!-- @click：注册事件 -->
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { register, sendSms } from "/@/api/basic/index";
import { ElMessage } from 'element-plus';
export default defineComponent({
  name: "Register",
  setup() {
    // 邮箱校验
    let checkEmail = (rule: any, value: string, callback: any) => {
      const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
      if (!value) {
        return callback(new Error('邮箱不能为空'))
      }
      if (mailReg.test(value)) {
        callback()
      } else {
        callback(new Error('请输入正确的邮箱格式'))
      }
    };
    // 确认密码校验
    let checkPassword = (rule: any, value: string, callback: any) => {
      if (value !== formRegister.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    };
    const formRegister = reactive({
      name: null,
      email: null,
      phone: null,
      password: null,
      sure_password: null,
      code: null,
    });
    const formRegisterRef = ref()
    const rules = {
      name: [
        {
          required: true,
          message: '用户名不能为空',
          trigger: 'blur',  // 失去焦点时验证
        }, {
          min: 4,
          max: 10,
          message: '请输入4-10个字符',
          trigger: 'blur',  // 失去焦点时验证
        },
      ],
      email: [
        {
          required: true,
          validator: checkEmail,
          trigger: 'blur'
        }
      ],
      phone: [
        { required: true, message: '请输入手机号码', trigger: 'blur' },
        { min: 11, max: 11, message: '请输入11位手机号码', trigger: 'blur' },
        {
          pattern: /^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$/,
          message: '请输入正确的手机号码'
        }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 18, message: '请输入6-18密码', trigger: 'blur' },
      ],
      sure_password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 18, message: '请输入6-18密码', trigger: 'blur' },
        {
          validator: checkPassword,
          trigger: 'blur'
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
    };
    let is_send = ref(false);
    let get_sms_text = ref("点击发送短信");
    // 短信验证码
    let send_sms = () => {
      // 判断状态
      if (is_send.value) {
        return;
      }
      // 需要判断一下手机号是否为空
      if (!formRegister.phone) {
        return
      }
      // 短信发送状态
      is_send.value = true;
      sendSms({ phone: formRegister.phone }).then((response: any) => {
        ElMessage({
          showClose: true,
          message: response.message,
          type: 'success',
        });
        let t = 60;
        let timer = setInterval(() => {
          if (t > 0) {
            t--;
            get_sms_text.value = `${t}秒后重新发送`;
          } else {
            get_sms_text.value = "点击发送短信";
            is_send.value = false;
            clearInterval(timer);
          }
        }, 1000)
      }).catch((err: any) => {
        is_send.value = false;
        ElMessage({
          showClose: true,
          message: err.response.message,
          type: 'error',
        })
      })
    };
    // 注册
    let userRegister = () => {
      formRegisterRef.value.validate((valid: any) => {
        if (valid) {
          register(formRegister).then((response: any) => {
            ElMessage({
              showClose: true,
              message: response.message,
              type: 'success',
            })
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
    };
    return {
      formRegister,
      rules,
      formRegisterRef,
      send_sms,
      is_send,
      userRegister,
      get_sms_text,

    };
  },
});
</script>

<style scoped>
#register {
  width: 300px;
  top: 0;
  bottom: 0px;
  left: 0;
  right: 0;
  margin: 80px auto 199px auto;
  padding: 30px 25px;
  border: 1px solid #c0c0c0;
  border-radius: 8px;
  box-shadow: 5px 5px 4px #c0c0c0;
}
#register-top {
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
