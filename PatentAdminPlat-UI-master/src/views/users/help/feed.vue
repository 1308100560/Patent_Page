<template>
  <div class="container">
    <div class="my-card">
      <div
        style="display: flex;flex-direction: row;align-items: center;justify-content: center;margin: 20px 0;margin-bottom: 40px"
      >
        <h2>问题反馈</h2>
      </div>
      <el-form
        ref="form"
        :model="form"
        :rules="rules"
        class="demo-ruleForm"
        label-width="100px"
        style="padding-bottom: 20px"
      >
        <el-form-item label="邮箱">
          <div style="display: flex;flex-direction: row">
            <div style="width: 150px">{{ form.email }}</div>
            <el-button size="mini" @click="copyText(form.email)">复制</el-button>
          </div>
        </el-form-item>
        <el-form-item label="手机号">
          <div style="display: flex;flex-direction: row">
            <div style="width: 150px">{{ form.phone }}</div>
            <el-button size="mini" @click="copyText(form.phone)">复制</el-button>
          </div>
        </el-form-item>
        <el-form-item prop="phone">
          如果您在使用过程中遇到问题，可以通过以下方式联系我们，我们会尽快解决您的问题。
        </el-form-item>

      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      form: {
        email: 'fuxu5241@163.com',
        phone: '13581995241'
      },
      rules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: ['blur', 'change'] }
        ],
        phone: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          { pattern: /^1[3456789]\d{9}$/, message: '请输入正确的手机号格式', trigger: ['blur', 'change'] }
        ],
        description: [{ required: true, message: '请输入问题描述', trigger: 'blur' }]
      }
    }
  },
  methods: {
    copyText(text) {
      const input = document.createElement('input')
      input.setAttribute('readonly', 'readonly')
      input.setAttribute('value', text)
      document.body.appendChild(input)
      input.select()
      if (document.execCommand('copy')) {
        document.execCommand('copy')
        this.$message({
          message: '复制成功',
          type: 'success'
        })
      }
      document.body.removeChild(input)
    },
    submitForm() {
      this.$refs.form.validate(valid => {
        if (valid) {
          // 表单验证通过，可以提交数据
          console.log(this.form)
        } else {
          console.log('表单验证失败')
          return false
        }
      })
    }
  }
}
</script>
