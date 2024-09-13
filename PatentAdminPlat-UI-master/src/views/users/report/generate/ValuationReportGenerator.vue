<template>
  <div class="container">
    <div class="my-card">
      <h2 style="text-align: center">估值报告</h2>
      <el-form ref="infringementForm" :model="form" :rules="formRules">
        <el-form-item label="选择专利" prop="patentId">
          <el-select v-model="form.patentId" placeholder="仅能选择已经认领的专利" style="width: 100%" value="请选择">
            <el-option
              v-for="option in patents"
              :key="option.PNM"
              :label="option.patentProperties.TI"
              :value="option.patentId"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="评估目的" prop="purpose">
          <el-input v-model="form.purpose" rows="6" type="textarea" />
        </el-form-item>
        <el-form-item label="目标估值" prop="targetValue">
          <el-input v-model="form.targetValue" type="number" />
        </el-form-item>
        <el-form-item label="申请人情况" prop="applicant">
          <el-input v-model="form.applicant" rows="6" type="textarea" />
        </el-form-item>
        <el-form-item label="联系客服" prop="contact">
          {{ form.contact }}
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
import { getContactInfo, reportTicketSubmit } from '@/api/report'

export default {
  data() {
    return {
      patents: [],
      patent: '',
      reportList: [],
      form: {
        contact: '12345678',
        // 专利标题
        patentId: '',
        // 评估目的
        purpose: '',
        // 目标估值
        targetValue: '',
        // 申请人情况
        applicant: ''
      },
      formRules: {
        patentId: [
          { required: true, message: '请选择专利', trigger: 'blur' }
        ],
        purpose: [
          { required: true, message: '请输入评估目的', trigger: 'blur' }
        ],
        targetValue: [
          { required: true, message: '请输目标估值', trigger: 'blur' }
        ],
        applicant: [
          { required: true, message: '请输入申请人', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    const patentId = this.$route.query.pId
    getContactInfo().then(res => {
      this.form.contact = res.data.data
    })
    if (patentId) {
      this.form.patentId = parseInt(patentId)
    }
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
          const patent = this.patents.find(item => item.patentId === this.form.patentId)
          const reportName = patent.patentProperties.TI
          const properties = {
            ...this.form, keys: [
              { key: 'title', value: '专利标题' },
              { key: 'purpose', value: '评估目的' },
              { key: 'contact', value: '联系客服' },
              { key: 'targetValue', value: '目标估值' },
              { key: 'applicant', value: '申请人信息' }
            ], type: '估值报告'
          }
          const relObj = {
            patentId: this.form.patentId,
            reportName: `${reportName}`,
            reportType: '估值报告'
          }
          const ticketFrom = {
            name: `${reportName}`,
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
