<template>
  <el-dialog
    :title="`创建${type==='focus'?'关注':'专利'}组合`"
    :visible.sync="dialogVisible"
    width="40%"
  >
    <el-form
      ref="form"
      :model="packForm"
      :rules="rules"
      label-width="120px"
      size="small"
      style="margin: 10px;z-index: 9999!important;"
    >
      <el-form-item label="组合名称" prop="packageName" size="small">
        <el-input v-model="packForm.packageName" size="small" />
      </el-form-item>
      <el-form-item label="用途" prop="packageName" size="small">
        <el-select
          v-model="packForm.properties.usage"
          allow-create
          default-first-option
          filterable
          placeholder="请选择专利用途"
          style="width: 100%"
        >
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="描述" size="small">
        <el-input
          v-model="packForm.properties.desc"
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
  props: {
    type: {
      type: String,
      default: 'claim'
    },
    getList: {
      type: Function,
      default: () => {
      }
    }
  },
  data() {
    return {
      options: [{
        value: '工程授权',
        label: '工程授权'
      }, {
        value: '批量交易',
        label: '批量交易'
      }, {
        value: '研究定向',
        label: '研究定向'
      }],
      dialogVisible: false,
      packForm: {
        packageName: '',
        properties: {
          usage: '',
          desc: ''
        },
        type: this.type
      },
      rules: {
        packageName: [
          { required: true, message: '请输入专利组合名称', trigger: 'blur' }
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
            this.getList()
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
