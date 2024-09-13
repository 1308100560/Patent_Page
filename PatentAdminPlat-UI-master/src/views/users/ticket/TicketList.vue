<template>
  <div class="container">
    <div v-if="currentPage==='index'" class="my-card" style="width: 100% ;">
      <div class="filter-container">
        <div>
          <el-input
            v-model="listQuery.query"
            class="filter-item"
            placeholder="工单信息"
            size="small"
            style="width: 200px;margin-right: 10px"
          />

          <el-button class="filter-item" icon="el-icon-search" size="small" type="primary" @click="getList">
            搜索
          </el-button>
        </div>

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
        <el-table-column
          label="标题"
          min-width="300"
          prop="name"
        >
          <template slot-scope="{row}">
            <span>{{ row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="当前状态"
          prop="status"
          width="100"
        >
          <template slot-scope="{row}">
            <el-tag :type="getStatusColor(row.status)">{{ row.status|statusFormat }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          prop="status"
        >
          <template slot-scope="{row}">
            {{ row.CreatedAt|localTime }}
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          prop="op"
        >
          <template slot-scope="{row}">
            <el-button size="mini" @click="showPreview(row)">预览</el-button>
            <!--            <el-button icon="el-icon-switch-button" size="mini">结单</el-button>-->
            <el-button :disabled="isClosed(row.status)" size="mini" @click="handleCloseTicket(row.id)">撤销</el-button>
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
    <div v-if="currentPage==='progress'" class="my-card" style="width: 100% ;">
      <ticket-progress :switch-page="switchPage" :ticket="currentTicket" />
    </div>
    <div v-if="currentPage==='preview'" class="my-card" style="width: 100% ;">
      <ticket-preview :switch-page="switchPage" :ticket="currentTicket" />
    </div>
  </div>
</template>

<script>
import { closeTicket, getTicketList } from '@/api/ticket'
import TicketProgress from '@/views/users/ticket/components/TicketProgress'
import TicketPreview from '@/views/users/ticket/components/TicketPreview.vue'

export default {
  name: 'TicketList',
  components: { TicketProgress, TicketPreview },
  data() {
    return {
      list: [],
      listLoading: false,
      currentPage: 'index',
      currentTicket: {},
      listQuery: {
        type: 'report',
        pageIndex: 1,
        pageSize: 10,
        query: ''
      },
      results: {
        count: 1
      }
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
    switchPage(page) {
      this.currentPage = page
    },
    showPreview(row) {
      this.currentTicket = row
      this.currentPage = 'preview'
    },
    showProgress(row) {
      this.currentTicket = row
      this.currentPage = 'progress'
    },
    isClosed(status) {
      return status === 'closed' || status === 'finished'
    },
    getStatusColor(status) {
      switch (status) {
        case 'open':
          return 'warning'
        case 'closed':
          return 'danger'
        case 'finished':
          return 'success'
        default:
          return 'info'
      }
    },
    handleCloseTicket(ticketId) {
      closeTicket(ticketId).then(response => {
        this.$message({
          message: '撤销成功',
          type: 'success'
        })
        this.getList()
      })
    },
    getList() {
      this.listLoading = true
      getTicketList(this.listQuery).then(response => {
        const results = response.data.data
        this.list = results.list
        this.listLoading = false
      })
    }
  }
}
</script>
<style>

</style>
