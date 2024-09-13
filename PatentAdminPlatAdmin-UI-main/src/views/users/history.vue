<template>
  <div class="container">
    <div class="my-card" style="width: 100% ;overflow-y: auto">
      <div class="filter-container">
        <div>
          <el-select
            v-model="listQuery.userId"
            filterable
            placeholder="筛选用户"
            size="small"
            style="margin-right: 10px"
          >
            <el-option
              v-for="item in userList"
              :key="item.userId"
              :label="item.username"
              :value="item.userId"
            />
          </el-select>
          <el-select
            v-model="listQuery.action"
            filterable
            placeholder="日志类型"
            size="small"
            style="margin-right: 10px"
          >
            <el-option
              v-for="item in ops"
              :key="item.name"
              :label="item.name"
              :value="item.query"
            />
          </el-select>
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

      </div>
      <div style="margin: 5px 0">
        共有 <strong>{{ results.count }}</strong> 条记录
      </div>
      <el-table
        v-loading="listLoading"
        :data="list"
        fit
        highlight-current-row
        style="width: 100%;"
      >
        <el-table-column
          align="center"
          label="ID"
          prop="id"
          width="60"
        >
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作用户" width="120">
          <template slot-scope="{row}">
            <span>{{ getUserName(row.CreateBy) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="400">
          <template slot-scope="{row}">
            <span>{{ row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="详细信息" min-width="120px">
          <template slot-scope="{row}">
            <span>{{ row.request }}</span>
          </template>
        </el-table-column>
        <el-table-column label="请求时间" width="300px">
          <template slot-scope="{row}">
            <span>{{ row.CreatedAt|localTime }}</span>
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
  </div>
</template>

<script>
import { getTraceList } from '@/api/trace'
import { getAllUsers } from '@/api/user'

export default {
  name: 'ComplexTable',

  data() {
    return {
      userList: [],
      ops: [{
        name: '搜索日志',
        query: 'Search'
      }, {
        name: '报告日志',
        query: 'Report'
      }, {
        name: '图表日志',
        query: 'Graph'
      }],
      list: null,
      listLoading: true,
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: '',
        userId: '',
        action: ''
      }

    }
  },
  created() {
    this.getList()
    this.getAllUserList()
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
    getUserName(userId) {
      const find = this.userList.find(item => item.userId === parseInt(userId))
      if (find) {
        return find.username
      }
      return 'NA'
    },
    getAllUserList() {
      getAllUsers({ pageIndex: 1, pageSize: 999999 }).then(response => {
        this.userList = response.data.data.list
      })
    },
    getList() {
      this.listLoading = true
      getTraceList(this.listQuery).then(response => {
        this.results = response.data.data
        this.list = response.data.data.list
        this.list.forEach((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        this.listLoading = false
      })
    }
  }
}
</script>
<style>

</style>
