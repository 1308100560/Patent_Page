<template>
  <el-form>
    <el-form-item label="密码">
      <el-input v-model="password" placeholder="请设置密码" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">更新</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { resetPassword } from '@/api/user'

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
    return {
      password: ''
    }
  },
  methods: {
    submit() {
      resetPassword({ userId: 1, password: this.password }).then(res => {
        this.$message.success('更新成功')
      })
    }
  }
}
</script>
