<template>
  <div>
    <p>我的工单</p>
    <el-table :data="list" style="width: 100%;">

      <el-table-column label="专利名称" min-width="200">
        <template slot-scope="scope">
          {{ scope.row.order_no | orderNoFilter }}
        </template>
      </el-table-column>
      <el-table-column label="提交时间" width="195" align="center">
        <template>
          2022-07-21 12:00:00
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template  slot-scope="scope">
          <el-tag :type="scope.row.status">
            {{scope.row.status==='success'? '已完成':'审核中'}}
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
