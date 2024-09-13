<template>
  <div class="container">
    <div style="display: flex;flex-direction: row;width: 100%;justify-content: center">
      <div style="width: 45%;display: flex;flex-direction: column;align-items:center;">
        <div style="display: flex;flex-direction: column;align-items: flex-start;">
          <img src="./SRIPPM.jpg" alt="SRIPPM Logo" style="width: 250px; height: auto; margin-bottom: 20px;" />
          <strong style="color: white;font-size: 54px;letter-spacing: 0;font-weight: 700;">
            科研人员知识产权手册
          </strong>
          <strong style="color: white;font-size: 28px;letter-spacing: 3.81px;font-weight: 700;margin-top: 20px">
            science researcher intellectual property professional manual
          </strong>
          <div style="color: white;font-size: 32px;letter-spacing: 3.81px;font-weight: 300;margin-top: 60px">
            科研工作者的IP AI
          </div>
        </div>
      </div>
      <div style="width: 40%;display: flex;flex-direction: row;align-items: center;justify-content: center">
        <el-form ref="ruleForm" :model="ruleForm" :rules="rules" class="form">
          <div style="font-size: 35px;font-weight: bold;margin-bottom: 15px;">欢迎登录
          </div>
          <el-form-item class="form-item" label="用户名" prop="username">
            <el-input v-model="ruleForm.username" class="form-input" placeholder="请输入用户名" />
          </el-form-item>

          <el-form-item class="form-item" label="密码" prop="password">
            <el-input
              v-model="ruleForm.password"
              class="form-input"
              placeholder="请输入登录密码"
              type="password"
            />
          </el-form-item>
          <el-form-item prop="code">
            <el-row align="middle" type="flex">
              <el-col :span="15">
                <el-input v-model="ruleForm.code" autocomplete="off" name="code" placeholder="验证码" type="text" />
              </el-col>
              <el-col :offset="1" :span="8">
                <img :src="captchaBase64" alt="cap" style="width: 100%;" @click="initCaptcha">
              </el-col>
            </el-row>
          </el-form-item>

          <div style="margin-top: 40px">
            <el-button
              style="width: 100%;height: 50px;border-radius: 20px;font-size: 16px;font-weight: bold"
              type="primary"
              @click="submitForm('ruleForm')"
            >
              登录
            </el-button>
          </div>
        </el-form>
      </div>
    </div>
  </div>
</template>
<script>

import { getCaptcha } from '@/api/user'

export default {

  data() {
    const notEmptyValidator = (rule, value, callback) => {
      console.log(rule)
      if (value === '') {
        callback(new Error(`${rule.fullField}不能为空`))
      } else {
        callback()
      }
    }
    return {
      captchaBase64: '',
      ruleForm: {
        username: '',
        password: '',
        code: '',
        uuid: ''
      },
      rules: {
        username: [
          { validator: notEmptyValidator, trigger: 'blur', fullField: '用户名' }
        ],
        password: [
          { validator: notEmptyValidator, trigger: 'blur', fullField: '密码' }
        ]
      }
    }
  },
  mounted() {
    this.initCaptcha()
  },
  methods: {
    initCaptcha() {
      const self = this
      getCaptcha().then(res => {
        self.captchaBase64 = res.data.data
        self.ruleForm.uuid = res.data.id
      })
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        this.initCaptcha()
        if (valid) {
          this.$store.dispatch('user/login', this.ruleForm).then(() => {
            this.$message({
              message: '登录成功',
              type: 'success'
            })
            this.$router.push({ path: this.redirect || '/', query: this.otherQuery })
          }).catch(error => {
            this.$message({
              message: error,
              type: 'error'
            })
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    login() {
      this.$refs.loginRef.show()
    }
  }
}
</script>
<style scoped>
.container {
  height: 100vh;
  background: url('https://passport.baidu.com/static/passpc-account/img/reg_bg_min.jpg') no-repeat;
  background-size: cover;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.form {
  width: 400px;
  background-color: rgba(255, 255, 255, .9);
  padding: 60px 25px;
  border-radius: 10px;
}

.form-item {
  margin: 30px 0;
}

.form-item >>> .el-input__inner {
  height: 40px;

}

</style>
