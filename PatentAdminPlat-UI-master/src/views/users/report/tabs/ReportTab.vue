<template>
  <div v-if="dataList" v-loading="listLoading">
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

    </div>
    <el-table
      :data="results.list"
      :stripe="true"
      fit
      highlight-current-row
      style="width: 100%;"
    >

      <el-table-column align="center" label="ID" prop="id" width="100">
        <template slot-scope="{row}">
          <span>{{ row.reportId }}</span>
        </template>
      </el-table-column>
      <el-table-column label="报告名称" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.reportName }}</span>
        </template>

      </el-table-column>
      <el-table-column label="报告类型" min-width="60">
        <template slot-scope="{row}">
          <span> {{ row.reportType }} </span>
        </template>

      </el-table-column>

      <el-table-column label="申请时间" width="180px">
        <template slot-scope="{row}">
          <span>{{ row.CreatedAt|localTime }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" class-name="small-padding fixed-width" label="报告文件" width="85">
        <template slot-scope="{row}">
          <div class="actions">
            <el-button
              icon="el-icon-view"
              size="mini"
              type="light"
              @click="preview"
            />

            <el-button
              icon="el-icon-download"
              size="mini"
              type="primary"
              @click="handleReportDownload(row)"
            />

          </div>
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
</template>
<script>
import { userReportList } from '@/api/report'
import { downloadFile, generateNoveltyReport } from '@/views/users/utils'

export default {
  name: 'ReportList',
  props: {
    dataList: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      listLoading: false,
      activeName: 'first',
      results: null,
      listQuery: {
        pageSize: 10,
        pageIndex: 1,
        query: ''
      }
    }
  },
  watch: {
    dataList: {
      handler(val) {
        console.log(val)
        this.results = val
      },
      deep: true
    }
  },
  created() {
    this.results = JSON.parse(JSON.stringify(this.dataList))
  },
  methods: {
    preview() {
      this.$message.info('暂未开放')
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
    isImage(filePath) {
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
    },
    getList() {
      this.listLoading = true
      userReportList(this.listQuery).then(response => {
        this.results = response.data.data
        this.listLoading = false
      })
    },
    handleNoveltyReportDownload(report) {
      generateNoveltyReport(report.reportName, report.reportProperties).then(file => {
        console.log(file)
      })
      this.$message({
        message: '下载成功',
        type: 'success'
      })
    },
    handleOtherReportDownload(report) {
      const files = JSON.parse(report.files)
      const reportPath = `${files[0].full_path}`
      const reportName = files[0].name
      downloadFile(reportPath, reportName)
      this.$message({
        message: '下载成功',
        type: 'success'
      })
    },
    handleReportDownload(row) {
      if (row.reportType === '查新报告') {
        this.handleNoveltyReportDownload(row)
      } else {
        this.handleOtherReportDownload(row)
      }
    }

  }
}
</script>
<style scoped>
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
