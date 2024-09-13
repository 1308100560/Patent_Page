<template>
  <el-form>
    <el-form-item label="旧密码">
      <el-input v-model="passwordForm.oldPassword" placeholder="请输入旧密码" type="password" />
    </el-form-item>
    <el-form-item label="新密码">
      <el-input v-model="passwordForm.newPassword" placeholder="请设置新密码" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">更新</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { modifyPassword } from '@/api/user'

export default {
  props: {
    user: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data() {
    return {
      passwordForm: {
        oldPassword: '',
        newPassword: ''
      }
    }
  },
  methods: {
    submit() {
      if (this.passwordForm.oldPassword === '' || this.passwordForm.newPassword === '') {
        this.$message({
          message: '请输入密码',
          type: 'error',
          duration: 5 * 1000
        })
        return
      }
      modifyPassword(this.passwordForm).then(response => {
        if (response.data.code === 200) {
          this.$message({
            message: '密码修改成功',
            type: 'success',
            duration: 5 * 1000
          })
        } else {
          this.$message({
            message: '密码修改失败，请重试',
            type: 'error',
            duration: 5 * 1000
          })
        }
      })
    }
  }
}
</script>
