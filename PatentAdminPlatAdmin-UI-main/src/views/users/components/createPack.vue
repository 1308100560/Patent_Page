<template>
  <el-dialog
    :visible.sync="dialogVisible"
    title="创建工艺包"
    width="40%"
  >
    <el-form
      ref="form"
      :model="packForm"
      :rules="rules"
      label-width="80px"
      size="small"
      style="margin: 10px;z-index: 9999!important;"
    >
      <el-form-item label="工艺包" prop="packageName" size="small">
        <el-input v-model="packForm.packageName" size="small" />
      </el-form-item>
      <el-form-item label="描述" size="small">

        <el-input
          v-model="packForm.desc"
          :rows="2"
          placeholder="请输入内容"
          type="textarea"
        />

      </el-form-item>
      <div style="text-align: right; margin: 0">
        <el-button size="mini" type="text" @click="dialogVisible=false">取消</el-button>
        <el-button size="mini" type="primary" @click="submit">创建</el-button>
      </div>
    </el-form>
  </el-dialog>
</template>
<script>

export default {
  name: 'CreatePack',
  data() {
    return {
      dialogVisible: false,
      packForm: {
        packageName: '',
        desc: ''
      },
      rules: {
        packageName: [
          { required: true, message: '请输入工艺包名称', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    show() {
      this.dialogVisible = true
    },
    submit() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          this.$store.dispatch('package/create', this.packForm).then(() => {
            this.dialogVisible = false
            this.$message.success('创建成功')
            this.$store.dispatch('package/getPackageList')
          }).catch(error => {
            this.$message.error(error)
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>
