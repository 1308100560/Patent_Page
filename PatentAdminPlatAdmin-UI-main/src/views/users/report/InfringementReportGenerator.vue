<template>
  <div class="container">
    <div class="my-card">
      <h2 style="text-align: center">侵权报告</h2>
      <el-form ref="infringementForm" :model="form" :rules="formRules">
        <el-form-item label="专利标题" prop="title">
          <el-input v-model="form.title" />
        </el-form-item>
        <el-form-item label="评估目的" prop="purpose">
          <el-input v-model="form.purpose" rows="6" type="textarea" />
        </el-form-item>

        <el-form-item label="申请人情况" prop="applicant">
          <el-input v-model="form.applicant" rows="6" type="textarea" />
        </el-form-item>
        <div style="text-align: center">
          <el-button type="primary" @click="onSubmit">创建工单</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
import { getClaimedPatents } from '@/api/patent'
import { reportTicketSubmit } from '@/api/report'

export default {
  data() {
    return {
      patents: [],
      patent: '',
      reportList: [],
      form: {
        // 专利标题
        title: '',
        // 评估目的
        purpose: '',
        // 申请人情况
        applicant: ''
      },
      formRules: {
        title: [
          { required: true, message: '请输入专利标题', trigger: 'blur' }
        ],
        purpose: [
          { required: true, message: '请输入评估目的', trigger: 'blur' }
        ],
        applicant: [
          { required: true, message: '请输入申请人', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    getClaimedPatents().then(res => {
      res.data.data.list.map(item => {
        item.patentProperties = JSON.parse(item.patentProperties)
        return item
      })
      this.patents = res.data.data.list
    })
  },
  methods: {
    onSubmit() {
      this.$refs.infringementForm.validate(valid => {
        if (valid) {
          this.loading = true
          const properties = {
            ...this.form, keys: [
              { key: 'title', value: '专利标题' },
              { key: 'purpose', value: '评估目的' },
              { key: 'applicant', value: '申请人信息' }
            ]
          }
          const relObj = {
            reportName: `${this.form.title}`,
            reportType: '侵权报告'
          }
          const ticketFrom = {
            name: `${this.form.title}`,
            properties,
            relObj
          }
          reportTicketSubmit(ticketFrom).then(res => {
            if (res.data.code === 500) {
              this.$message({
                message: '当前专利已经申请报告',
                type: 'error'
              })
            } else {
              this.$message({
                message: '提交成功',
                type: 'success'
              })
              this.$router.push({ path: '/ticket/index' })
            }
            this.loading = false
          })
        }
      })
    }
  }
}
</script>
