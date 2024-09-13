<template>
  <el-dialog :visible.sync="formVisible" center class="diag" title="用户名密码登录">
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item prop="username">
        <el-input v-model="form.username" autocomplete="off" name="username" placeholder="用户名" />
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="form.password" autocomplete="off" name="password" placeholder="密码" type="password" />
      </el-form-item>
      <el-form-item prop="code">
        <el-row align="middle" type="flex">
          <el-col :span="15">
            <el-input v-model="form.code" autocomplete="off" name="code" placeholder="验证码" type="text" />
          </el-col>
          <el-col :offset="1" :span="8">
            <img :src="captchaBase64" alt="cap" style="width: 100%;" @click="changeCaptcha">
          </el-col>
        </el-row>
      </el-form-item>
      <el-form-item style="margin-top: 20px">
        <el-button style="width: 100%;height: 40px" type="primary" @click="submitForm('form')">登录</el-button>
      </el-form-item>

      <!--      <el-form-item>-->
      <!--        <router-link target="_blank" to="/register">-->
      <!--          <span style="color: #0085fa;font-size: 14px;cursor: pointer">注册账号</span>-->
      <!--        </router-link>-->
      <!--      </el-form-item>-->
    </el-form>

  </el-dialog>
</template>
<script>
import { getCaptcha } from '@/api/user'

export default {
  name: 'LoginComponent',
  data() {
    const notEmptyValidator = (rule, value, callback) => {
      if (value === '') {
        callback(new Error(`${rule.fullField}不能为空`))
      } else {
        callback()
      }
    }
    return {
      formVisible: false,
      loginLoading: false,
      captchaBase64: '',
      rules: {
        username: [
          {
            validator: notEmptyValidator, trigger: 'blur', fullField: '账号'
          }
        ],
        password: [
          {
            validator: notEmptyValidator, trigger: 'blur', fullField: '密码'
          }
        ],
        code: [
          {
            validator: notEmptyValidator, trigger: 'blur', fullField: '验证码'
          }
        ]
      },
      form: {
        username: '',
        password: '',
        code: ''
      }
    }
  },
  mounted() {
  },
  methods: {
    show() {
      this.form = {
        username: '',
        password: '',
        code: ''
      }
      this.initCaptcha()
    },
    initCaptcha() {
      const self = this
      getCaptcha().then(res => {
        self.captchaBase64 = res.data.data
        self.formVisible = true
        self.form.uuid = res.data.id
      })
    },
    changeCaptcha() {
      this.initCaptcha()
    },
    submitForm(formName) {
      const self = this
      this.loginLoading = true
      this.$refs[formName].validate((valid) => {
        self.loginLoading = false
        if (valid) {
          this.$store.dispatch('user/login', this.form).then(() => {
            this.$message({
              message: '登录成功',
              type: 'success'
            })
            self.formVisible = false
            this.$router.push({ path: this.redirect || '/', query: this.otherQuery })
            this.loading = false
          })
            .catch(error => {
              this.$message({
                message: error,
                type: 'error'
              })
            })
        } else {
          this.initCaptcha()
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
.diag {
  width: 100%;
}

.diag >>> .el-dialog {
  width: 350px;
  margin-top: 25vh !important;
}
</style>
