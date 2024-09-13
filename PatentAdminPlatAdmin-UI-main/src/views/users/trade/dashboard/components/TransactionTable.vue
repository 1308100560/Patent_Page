<template>
  <div>
    <p>浏览历史</p>
    <el-table :data="list" style="width: 100%;">

      <el-table-column label="关键词" min-width="200">
        <template slot-scope="scope">
          {{ scope.row.order_no | orderNoFilter }}
        </template>
      </el-table-column>
      <el-table-column label="时间" width="195" align="center">
        <template>
          2022-07-21 12:00:00
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100" align="center">
        <template>
          <el-tag :type="'success'">
            快速跳转
          </el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { transactionList } from '@/api/remote-search'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        success: 'success',
        pending: 'danger'
      }
      return statusMap[status]
    },
    orderNoFilter(str) {
      return str.substring(0, 30)
    }
  },
  data() {
    return {
      list: null
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      transactionList().then(response => {
        this.list = response.data.items.slice(0, 8)
      })
    }
  }
}
</script>
