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
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        fit
        highlight-current-row
        style="width: 100%;"
      >

        <el-table-column align="center" label="ID" prop="id" width="60">
          <template slot-scope="{row}">
            <span>{{ row.reportId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="报告类型" min-width="60">
          <template slot-scope="{row}">
            <span> {{ row.reportType }} </span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="报告名称" min-width="150px">
          <template slot-scope="{row}">
            <span>{{ row.reportName }}</span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="申请时间" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.CreatedAt|localTime }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" class-name="small-padding fixed-width" label="报告文件" width="200">
          <template slot-scope="{row}">
            <el-button
              size="mini"
              type="light"
            >
              预览
            </el-button>
            <el-button
              size="mini"
              type="primary"
              @click="handleReportDownload(row)"
            >
              下载
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

import { userReportList } from '@/api/report'
import { generateNoveltyReport } from '@/views/users/utils'

export default {
  name: 'ReportList',

  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        pageSize: 1,
        pageIndex: 20,
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
    isImage(filePath) {
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
    },
    getList() {
      this.listLoading = true
      userReportList().then(response => {
        this.list = response.data.data.list
        this.results = response.data.data
        this.listLoading = false
      })
    },
    handleNoveltyReportDownload(report) {
      generateNoveltyReport(report.reportName, report.reportProperties)
    },
    handleOtherReportDownload(report) {
      console.log(report)
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

</style>
