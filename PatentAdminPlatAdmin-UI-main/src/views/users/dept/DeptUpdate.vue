<template>
  <div class="container">
    <div class="my-card" style="width: 100% ;overflow-y: auto">
      <el-page-header content="更新团队" @back="$router.go(-1)" />
      <div style="display: flex;flex-direction: row;margin-top: 30px;justify-content: center">
        <el-form
          ref="createDeptForm"
          v-loading="loading"
          :model="form"
          :rules="formRules"
          label-position="right"
          label-width="80px"
          style="width: 800px"
        >
          <el-form-item label="团队名称" prop="deptName">
            <el-input v-model="form.deptName" />
          </el-form-item>
          <el-form-item label="团队简介" prop="deptProperties">
            <el-input v-model="form.deptProperties" rows="4" type="textarea" />
          </el-form-item>

        </el-form>
      </div>
      <div style="display: flex;flex-direction: row;justify-content: center">
        <el-button type="primary" @click="handleUpdateUserInfo">保存</el-button>
      </div>

    </div>
  </div>
</template>
<script>

import { updateDept } from '@/api/dept'

export default {
  data() {
    return {
      loading: false,
      form: {
        deptName: '',
        deptProperties: ''
      },
      deptList: [],
      formRules: {
        deptName: [
          { required: true, message: '请团队名称', trigger: 'blur' }
        ],
        deptProperties: [
          { required: true, message: '请输入团队简介', trigger: 'blur' }
        ]
      },
      updateFormVisible: false,
      deptId: null
    }
  },
  mounted() {
    const deptId = this.$route.query.deptId
    const deptName = this.$route.query.deptName
    const deptProperties = this.$route.query.deptProperties
    this.form.deptName = deptName
    this.form.deptProperties = deptProperties
    this.deptId = deptId
  },
  methods: {
    handleUpdateUserInfo() {
      this.$refs.createDeptForm.validate((valid) => {
        if (valid) {
          updateDept(this.deptId, this.form).then(res => {
            this.$message({
              message: '修改成功',
              type: 'success'
            })
            this.$router.go(-1)
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>
