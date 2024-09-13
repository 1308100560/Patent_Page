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
          <router-link style="margin-left: 10px" to="/search/advanced">
            <el-button class="filter-item" icon="el-icon-plus" size="small" type="primary">
              新增搜索式
            </el-button>
          </router-link>
          <el-button
            class="filter-item"
            icon="el-icon-upload2"
            size="small"
            style="margin-left: 10px"
            @click="exportExcel"
          >导出
          </el-button>
        </div>

      </div>

      <el-table
        :key="tableKey"
        v-loading="listLoading"
        :data="list"
        fit
        highlight-current-row
        style="width: 100%; "
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="50"
        />
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
        <el-table-column label="名称" width="200">
          <template slot-scope="{row}">
            <el-button type="text" @click="queryDetail(row)">{{ row.name }}</el-button>
          </template>
        </el-table-column>

        <el-table-column label="表达式" min-width="80">
          <template slot-scope="{row}">
            <span>{{ row.expression }}</span>
          </template>
        </el-table-column>
        <el-table-column label="备注" width="400">
          <template slot-scope="{row}">
            <span>{{ row.desc ? row.desc : '暂无' }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="small-padding fixed-width" label="操作" width="80">
          <template slot-scope="row">
            <div class="actions">

              <el-tooltip class="item" content="查看统计分析" effect="dark" placement="top">
                <el-button icon="el-icon-pie-chart" size="mini" type="light" @click="graphDetail(row)" />
              </el-tooltip>
              <el-button icon="el-icon-delete" size="mini" type="danger" @click="deleteQuery(row)" />
            </div>
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
import XLSX from 'xlsx'

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
    handleSelectionChange(val) {
      console.log(val)
      const exportData = []
      val.forEach(item => {
        exportData.push({
          'ID': item.id,
          '名称': item.name,
          '表达式': item.expression,
          '备注': item.desc
        })
      })
      this.multipleSelection = exportData
    },
    exportExcel() {
      if (!this.multipleSelection || this.multipleSelection.length === 0) {
        this.$message({
          message: '请先选择要导出的内容',
          type: 'warning'
        })
        return
      }
      const workbook = XLSX.utils.book_new()
      const worksheet = XLSX.utils.json_to_sheet(this.multipleSelection)
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1')
      XLSX.writeFile(workbook, '搜索式.xlsx')
    },
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
        this.list.forEach((item, index) => {
          item.id = index + 1
        })
        this.results = response.data.data
        this.listLoading = false
      })
    },
    queryDetail(row) {
      this.$router.push({ path: '/search/results', query: { q: row.expression, action: 'query' }})
    },
    graphDetail(row) {
      this.$router.push({ path: '/search/results', query: { q: row.row.expression, action: 'graph' }})
    },
    deleteQuery(row) {
      this.$confirm('此操作将永久删除该查询, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteQuery(row.row.queryID).then(response => {
          this.$message({
            message: '删除成功',
            type: 'success'
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

.space-between {
  display: flex;
  width: 100%;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.actions {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.actions >>> .el-button {
  padding: 0 !important;
  min-width: 34px !important;
  height: 30px !important;
  margin: 0 !important;
}
</style>

