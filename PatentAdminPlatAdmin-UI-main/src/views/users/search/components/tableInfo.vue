<template>
  <el-card class="box-card" shadow="hover" style="height: 400px!important;">
    <div slot="header" class="clearfix">
      <span style="font-size: 0.9rem">{{ title }}</span>
      <el-button style="float: right; padding: 3px 0" type="text">查看详细报告</el-button>
    </div>
    <div>
      <el-table
        v-if="tableType === 'categorical'"
        :data="tableData"
        size="mini"
        style="width: 100%"
      >
        <el-table-column
          label="分类号"
          prop="date"
          width="200"
        />
        <el-table-column
          label="百分比"
          prop="percent"
        >
          <template slot-scope="{row}">
            <el-progress
              :format="function() {
                return `${row.number} (${row.percent}%)`
              }"
              :percentage="row.percent"
              :stroke-width="20"
              :text-inside="true"
              class="progress-bar"
            />
          </template>
        </el-table-column>
      </el-table>
      <div v-if="tableType === 'pieChart'">
        <PieChart />
      </div>
      <el-table
        v-if="tableType === 'application'"
        :data="tableData"
        size="mini"
        style="width: 100%"
      >
        <el-table-column
          label="申请人"
          prop="person"
          width="200"
        />
        <el-table-column
          label="专利数"
          prop="number"
        />
        <el-table-column
          label="百分比"
          prop="percent"
        >
          <template slot-scope="{row}">
            {{ row.percent }}%
          </template>
        </el-table-column>
      </el-table>
      <el-table
        v-if="tableType === 'inventor'"
        :data="tableData"
        size="mini"
        style="width: 100%"
      >
        <el-table-column
          label="发明人"
          prop="person"
          width="200"
        />
        <el-table-column
          label="专利数"
          prop="number"
        />
        <el-table-column
          label="百分比"
          prop="percent"
        >
          <template slot-scope="{row}">
            {{ row.percent }}%
          </template>
        </el-table-column>
      </el-table>
    </div>
  </el-card>
</template>
<script>
import PieChart from './PieChart'

export default {
  name: 'TableInfo',
  components: { PieChart },
  props: {
    title: {
      type: String,
      default: ''
    },
    tableType: {
      type: String,
      default: 'categorical'
    },
    tableData: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {}
  }
}
</script>
<style scoped>
/deep/ .el-progress-bar__innerText {
  color: black;
}
</style>
