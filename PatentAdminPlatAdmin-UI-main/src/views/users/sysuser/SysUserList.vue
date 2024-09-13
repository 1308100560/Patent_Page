<template>
  <div class="container">
    <el-dialog :visible.sync="createUserDialog" title="添加用户">
      <el-form>
        <el-form-item label="用户名">
          <el-input v-model="createUserForm.username" type="text" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="createUserForm.password" type="password" />
        </el-form-item>
        <div style="display: flex;flex-direction: row;align-items: center;justify-content: center">
          <el-form-item>
            <el-button type="primary" @click="handleCreateUser">添加</el-button>
          </el-form-item>
        </div>
      </el-form>
    </el-dialog>
    <div class="my-card" style="width: 100% ;overflow-y: auto">
      <div class="filter-container" style="display: flex;flex-direction: row;align-items: center">
        <div>
          <el-input
            v-model="listQuery.query"
            class="filter-item"
            placeholder="关键词"
            size="small"
            style="width: 200px;margin-right: 10px"
          />

          <el-button class="filter-item" icon="el-icon-search" size="small" type="primary" @click="getList">
            搜索
          </el-button>
        </div>
        <el-button icon="el-icon-plus" size="small" @click="createUserDialog=true">
          添加用户
        </el-button>

      </div>
      <el-table
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        fit
        highlight-current-row
        style="width: 100%; "
      >
        <el-table-column align="center" label="用户ID" prop="userId" width="140">
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="用户名" min-width="140" prop="phone">
          <template slot-scope="{row}">
            <span>{{ row.nickName }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="电话号码" prop="phone" width="200">
          <template slot-scope="{row}">
            <span>{{ row.phone }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="性别" prop="sex" width="200">
          <template slot-scope="{row}">
            <span>{{ row.sex === 1 ? '男' : '女' }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="角色" prop="roleId" width="200">
          <template slot-scope="{row}">
            <span>{{ row.roleId === 1 ? '管理员' : '普通用户' }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="small-padding fixed-width" fixed="right" label="操作" width="270">
          <template slot-scope="{row}">
            <el-button size="mini" @click="$router.push(`/user/update/${row.userId}`)">
              修改用户信息
            </el-button>

            <el-button size="mini" type="warning" @click="resetPasswordDialog=true;currentUser=row.userId">
              重置密码
            </el-button>
            <el-button size="mini" type="danger" @click="deleteAccount(row.userId)">
              删除用户
            </el-button>

          </template>
        </el-table-column>
      </el-table>
      <div style="display:flex;flex-direction: row;justify-content: center;margin-top: 15px;">
        <el-pagination
          :current-page="listQuery.pageIndex"
          :hide-on-single-page="false"
          :page-size="100"
          :page-sizes="[10,20,40]"
          :total="results.count"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>

    </div>
    <el-dialog :visible.sync="resetPasswordDialog" title="重置密码">
      <el-form>
        <el-form-item label="新密码">
          <el-input v-model="resetPasswordForm.password" type="password" />
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input v-model="resetPasswordForm.confirmPassword" type="password" />
        </el-form-item>
        <div style="display: flex;flex-direction: row;align-items: center;justify-content: center">
          <el-form-item>
            <el-button type="primary" @click="resetPassword">确定</el-button>
          </el-form-item>
        </div>
      </el-form>
    </el-dialog>

  </div>
</template>

<script>
import { createUser, deleteAccount, getAllUsers, resetPassword } from '@/api/user'

export default {
  name: 'SysUserList',

  data() {
    return {
      currentUser: null,
      list: null,
      listLoading: true,
      createUserDialog: false,
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: ''
      },
      createUserForm: {
        username: '',
        password: '',
        roleId: 2,
        status: 2// 2生效
      },
      resetPasswordForm: {
        password: '',
        confirmPassword: ''
      },
      resetPasswordDialog: false,
      form: {
        patentId: '',
        type: ''
      },
      results: { count: 0 }

    }
  },
  created() {
    this.getList()
  },
  methods: {
    handleSearch() {
      this.listQuery.pageIndex = 1
      this.getList()
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.pageIndex = val
      this.getList()
    },
    getList() {
      this.listLoading = true
      getAllUsers(this.listQuery).then(response => {
        this.results = response.data.data
        this.list = response.data.data.list
        this.list.forEach((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        this.listLoading = false
      })
    },
    deleteAccount(userId) {
      this.$confirm('此操作将删除该用户, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteAccount(userId).then(response => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          this.getList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleCreateUser() {
      this.createUserForm.nickName = this.createUserForm.username
      createUser(this.createUserForm).then(response => {
        this.$message({
          type: 'success',
          message: '创建成功!'
        })
        this.createUserDialog = false
        this.getList()
      })
    },
    resetPassword() {
      if (this.resetPasswordForm.password.length < 6) {
        this.$message({
          type: 'error',
          message: '密码长度不能小于6位!'
        })
        return
      }
      if (this.resetPasswordForm.password !== this.resetPasswordForm.confirmPassword) {
        this.$message({
          type: 'error',
          message: '两次输入的密码不一致!'
        })
        return
      }
      this.$confirm('此操作将重置密码, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        resetPassword({ userId: this.currentUser, password: this.resetPasswordForm.password }).then(response => {
          this.$message({
            type: 'success',
            message: '重置成功!'
          })
          this.resetPasswordDialog = false
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消重置'
        })
      })
    }

  }
}
</script>
<style scoped>

</style>
