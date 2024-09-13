<template>
  <el-form ref="form" :model="user" :rules="rules">
    <el-form-item label="姓名" prop="username">
      <el-input v-model="user.username" disabled />
    </el-form-item>
    <el-form-item label="个人简介" prop="bio">
      <el-input v-model="user.bio" name="bio" type="textarea" />
    </el-form-item>
    <el-form-item label="邮箱" prop="email">
      <el-input v-model="user.email" name="email" />
    </el-form-item>
    <el-form-item label="单位" prop="departure">
      <el-input v-model="user.departure.deptName" disabled name="departure" />
    </el-form-item>
    <el-form-item label="电话" prop="phone">
      <el-input v-model="user.phone" name="phone" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit('form')">更新</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  props: {
    user: {
      type: Object,
      default: () => {
        return {
          name: '',
          email: ''
        }
      }
    }
  },
  data() {
    const phoneValidator = (rule, value, callback) => {
      if (value === '') {
        callback()
      } else {
        const reg = /^1[3456789]\d{9}$/
        if (reg.test(value)) {
          callback()
        } else {
          callback(new Error('请输入正确的手机号'))
        }
      }
    }
    return {
      bio: this.user.bio,
      rules: {
        email: [
          {
            type: 'email',
            message: '请输入正确的邮箱地址',
            trigger: 'blur'
          }
        ],
        phone: [{
          validator: phoneValidator,
          trigger: 'blur'
        }]
      }
    }
  },
  methods: {
    submit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$store.dispatch('user/updateInfo', this.user).then(() => {
            this.$message.success('更新成功')
          })
        }
      })
    }
  }
}
</script>
