<template>
  <div class="container">
    <div class="my-card" style="width: 100% ;">
      <div class="filter-container">
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
        <el-button icon="el-icon-refresh" size="mini" @click="getList" />

      </div>

      <el-table
        :key="tableKey"
        v-loading="listLoading"
        :data="list"
        fit
        highlight-current-row
        style="width: 100%; "
      >
        <el-table-column
          align="center"
          label="ID"
          prop="id"
          width="60"
        >
          <template slot-scope="{row}">
            <span>{{ row.queryID }}</span>
          </template>
        </el-table-column>
        <el-table-column label="名称" min-width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="表达式" min-width="80">
          <template slot-scope="{row}">
            <span>{{ row.expression }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="描述" width="200">
          <template slot-scope="{row}">
            <span>{{ row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="200">
          <template slot-scope="row">
            <el-button size="mini" type="light" @click="queryDetail(row)">
              快速访问
            </el-button>
            <el-button size="mini" type="danger" @click="deleteQuery(row)">
              删除
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
  </div>
</template>

<script>
import { deleteQuery, getQueryList } from '@/api/search'

export default {
  name: 'QueryList',
  data() {
    return {
      tableKey: 0,
      list: null,
      listLoading: true,
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: ''
      },
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
      getQueryList(this.listQuery).then(response => {
        this.list = response.data.data.list
        this.results = response.data.data
        this.listLoading = false
      })
    },
    queryDetail(row) {
      this.$router.push({ path: '/search/results', query: { q: row.row.expression }})
    },
    deleteQuery(row) {
      deleteQuery(row.row.queryID).then(response => {
        this.$message({
          message: '删除成功',
          type: 'success'
        })
        this.getList()
      })
    }
  }
}
</script>
<style>

</style>
