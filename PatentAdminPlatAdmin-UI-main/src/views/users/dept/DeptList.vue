<template>

  <div class="container">
    <div class="my-card" style="width: 100% ;">

      <div class="filter-container">
        <div style="display: flex;flex-direction: row;align-items: center">
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
          <el-button icon="el-icon-plus" size="small" type="primary">
            <router-link to="/dept/create">
              创建团队
            </router-link>
          </el-button>
        </div>
        <el-button icon="el-icon-refresh" size="mini" @click="getList" />

      </div>
      <el-table
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        fit
        highlight-current-row
        style="width: 100%;"
      >

        <el-table-column label="ID" prop="团队id" sortable="custom" width="100">
          <template slot-scope="{row}">
            <span>{{ row.deptId }}</span>
          </template>
        </el-table-column>
        <el-table-column label="团队名称" min-width="150px">
          <template slot-scope="{row}">
            <span>{{ row.deptName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="团队简介" min-width="200px">
          <template slot-scope="{row}">
            <span>{{ row.deptProperties }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="团队人数" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.members }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="260">
          <template slot-scope="{row}">
            <el-button
              size="mini"
              @click="$router.push({
                path:`/dept/update/${row.deptId}`,
                query:{
                  deptId:row.deptId,
                  deptName:row.deptName,
                  deptProperties:row.deptProperties}
              })"
            >
              更新团队信息
            </el-button>
            <el-button size="mini" type="danger" @click="deleteDept(row.deptId)">
              删除团队
            </el-button>

          </template>
        </el-table-column>

      </el-table>
      <div style="display:flex;flex-direction: row;justify-content: center;margin-top: 15px;">
        <el-pagination
          :current-page="listQuery.pageIndex"
          :hide-on-single-page="false"
          :page-size="listQuery.pageSize"
          :page-sizes="[10,20,40]"
          :total="results.count"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>

    </div>

  </div>

</template>
<script>

import { deleteDept, getDeptList } from '@/api/dept'

export default {
  name: 'TeamList',

  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        pageSize: 10,
        pageIndex: 1
      },
      query: '',
      results: {}

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
      getDeptList(this.listQuery).then(response => {
        this.list = response.data.data.list
        this.results = response.data.data
        this.listLoading = false
      })
    },
    deleteDept(deptId) {
      // show confirm
      this.$confirm('此操作将永久删除该团队, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteDept(deptId).then(response => {
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
    }

  }
}
</script>
<style scoped>

</style>
