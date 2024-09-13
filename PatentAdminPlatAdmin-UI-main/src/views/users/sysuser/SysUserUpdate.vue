<template>
  <div class="container">
    <div class="my-card" style="width: 100% ;overflow-y: auto">
      <el-page-header content="修改用户信息" @back="$router.go(-1)" />
      <div style="display: flex;flex-direction: row;margin-top: 30px;justify-content: center">
        <el-form
          ref="updateUserForm"
          v-loading="loading"
          :model="form"
          :rules="formRules"
          label-position="right"
          label-width="80px"
          style="width: 800px"
        >
          <el-form-item label="头像" prop="avatar">
            <el-input v-model="form.avatar" />
          </el-form-item>
          <el-form-item label="昵称" prop="nickName">
            <el-input v-model="form.nickName" />
          </el-form-item>
          <el-form-item label="个人简介" prop="bio">
            <el-input v-model="form.bio" rows="4" type="textarea" />
          </el-form-item>
          <el-form-item label="性别" prop="sex">
            <el-select v-model="form.sex" style="width: 100%">
              <el-option label="男" value="1" />
              <el-option label="女" value="2" />
            </el-select>
          </el-form-item>
          <el-form-item label="所属团队" prop="deptId">
            <el-select v-model="form.deptId" style="width: 100%">
              <el-option v-for="dept in deptList" :key="dept.deptId" :label="dept.deptName" :value="dept.deptId" />
            </el-select>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="form.phone" />
          </el-form-item>
          <el-form-item label="兴趣" prop="interest">
            <el-input v-model="form.interest" rows="4" type="textarea" />
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
import { getUserInfoByUserId, updateUserInfo } from '@/api/user'
import { getDeptList } from '@/api/dept'

export default {
  data() {
    return {
      loading: false,
      activeName: 'basic',
      form: {
        nickName: '',
        bio: '',
        avatar: '',
        deptId: '',
        password: '',
        email: '',
        phone: '',
        interest: '',
        sex: '',
        status: ''
      },
      deptList: [],
      formRules: {
        nickName: [
          { required: true, message: '请输入昵称', trigger: 'blur' }
        ],
        bio: [
          { required: true, message: '请输入个人简介', trigger: 'blur' }
        ],
        avatar: [
          { required: true, message: '请输入头像', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '请输入手机号', trigger: 'blur' }
        ]
      },
      updateFormVisible: false
    }
  },
  mounted() {
    this.loading = true
    getDeptList().then(res => {
      this.deptList = res.data.data.list
      getUserInfoByUserId(this.$route.params.id).then(res => {
        res.data.data.deptId = res.data.data.dept ? res.data.data.dept.deptId : ''
        for (const key in res.data.data) {
          this.form[key] = res.data.data[key]
        }
        this.loading = false
      })
    })
  },
  methods: {
    handleUpdateUserInfo() {
      this.$refs.updateUserForm.validate((valid) => {
        if (valid) {
          updateUserInfo(this.form).then(res => {
            this.$message({
              type: 'success',
              message: '修改成功'
            })
            this.updateFormVisible = false
            getUserInfoByUserId(this.$route.params.id).then(res => {
              res.data.data.deptId = res.data.data.dept.deptId
              for (const key in res.data.data) {
                this.form[key] = res.data.data[key]
              }
            })
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>
