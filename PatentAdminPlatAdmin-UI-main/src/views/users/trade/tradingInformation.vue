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
        style="width: 100%;"
      >
        <el-table-column
          align="center"
          label="ID"
          prop="id"
          sortable="custom"
          width="60"
        >
          <template slot-scope="{row}">
            <span>{{ row.patentId }}</span>
          </template>
        </el-table-column>
        <el-table-column label="专利名称" min-width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentProperties.TI }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="交易时间" sortable="custom" width="120px">
          <template slot-scope="{row}">
            <span>{{ row.patentProperties.AD }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="交易金额" width="110px">
          <template>
            <span>100000</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="状态" width="100">
          <template slot-scope="{row}">
            <el-tag :type="row.status ">
              已完成
            </el-tag>
          </template>
        </el-table-column>
        <!--        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="130">-->
        <!--          <template slot-scope="row">-->
        <!--            <el-button size="mini" type="danger" @click="unClaimClick(row)">-->
        <!--              取消认领-->
        <!--            </el-button>-->
        <!--          </template>-->
        <!--        </el-table-column>-->
      </el-table>

    </div>
    <!--    <div style="width: 300px">-->
    <!--      <PatentRecommend />-->
    <!--    </div>-->
  </div>
</template>

<script>
import { getClaimedPatents, unClaimPatent } from '@/api/patent'
import waves from '@/directive/waves' // waves directive

export default {
  name: 'ComplexTable',
  directives: { waves },

  data() {
    return {
      tableKey: 0,
      list: null,
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      }

    }
  },
  created() {
    // this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getClaimedPatents(this.listQuery).then(response => {
        const results = response.data.data
        results.map(item => {
          item.patentProperties = JSON.parse(item.patentProperties)
        })
        this.list = results
        console.log(this.list)
        this.listLoading = false
      })
    },
    unClaimClick(row) {
      unClaimPatent(row.row.patentId).then(response => {
        this.$message({
          message: '取消认领成功',
          type: 'success',
          duration: 5 * 1000
        })
        this.getList()
      })
    }
  }
}
</script>
<style>

</style>
