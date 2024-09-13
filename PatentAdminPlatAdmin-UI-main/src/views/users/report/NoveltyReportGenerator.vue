<template>
  <div class="container">
    <div class="my-card">
      <h2 style="text-align: center">查新报告</h2>
      <el-form ref="form" :model="form" :rules="formRules">
        <el-form-item label="专利标题" prop="title">
          <el-input v-model="form.title" />
        </el-form-item>
        <el-form-item label="专利描述" prop="CL">
          <el-input v-model="form.CL" rows="6" type="textarea" />
        </el-form-item>
        <el-form-item label="关键词" prop="keywords">
          <el-input v-model="form.keywords" placeholder="多个关键词请用空格隔开" type="textarea" />
        </el-form-item>
        <el-form-item label="申请人">
          <el-input v-model="form.applicant" />
        </el-form-item>
        <el-form-item label="申请机构">
          <el-input v-model="form.org" />
        </el-form-item>
        <el-form-item style="text-align: center">
          <el-button :loading="loading" type="primary" @click="onSubmit">创建工单</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import { reportTicketSubmit } from '@/api/report'

export default {
  name: 'NoveltyReportGenerator',
  data() {
    return {
      loading: false,
      form: {
        title: '使用无人机数据产生用于自主车辆导航的高清地图',
        CL: '1.一种方法，其包括：将供自主车辆用于规划包含第一地理位置的导航路线的数字地图存储在存储器中；通过所述车辆从无人机(UAV)实时接收由所述UAV的传感器在所述第一地理位置处收集的传感器数据；通过至少一个处理装置处理所述接收的传感器数据以产生所述第一地理位置的地图数据；以及使用所述所产生的地图数据更新所述数字地图。',
        keywords: '',
        applicant: '',
        org: ''
      },
      formRules: {
        title: [
          { required: true, message: '请输入专利标题', trigger: 'blur' }
        ],
        CL: [
          { required: true, message: '请输入专利描述', trigger: 'blur' }
        ],
        keywords: [
          { required: true, message: '请输入关键词', trigger: 'blur' }
        ],
        applicant: [
          { required: true, message: '请输入申请人', trigger: 'blur' }
        ],
        org: [
          { required: true, message: '请输入申请机构', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    console.log('mounted')
  },
  methods: {

    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.loading = true
          const properties = {
            ...this.form, keywords: this.form.keywords.split(' '), keys: [
              { key: 'title', value: '专利标题' },
              { key: 'CL', value: '专利描述' },
              { key: 'keywords', value: '关键词' },
              { key: 'applicant', value: '申请人' },
              { key: 'org', value: '申请机构' }

            ]
          }
          const relObj = {
            reportName: `${this.form.title}`,
            reportType: '查新报告'
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
        } else {
          this.loading = false
          this.$message({
            message: '提交失败',
            type: 'error'
          })
        }
      })
    }
  }
}
</script>
<style scoped>

</style>
