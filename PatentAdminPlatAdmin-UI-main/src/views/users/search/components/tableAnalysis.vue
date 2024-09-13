<template>
  <div>
    <div style="margin-bottom: 10px;">
      <el-tag
        v-for="chart in charts"
        :key="chart.id"
        :type="chart.active?'primary':'info'"
        style="margin-right: 5px;margin-bottom: 5px;cursor: pointer"
        @click="chart.active=!chart.active"
      >
        {{ chart.name }}
      </el-tag>
    </div>
    <div>
      <Chart
        v-for="chart in charts.filter(item=>item.active)"
        :key="'chart'+chart.id"
        :chart="chart"
        :query="query"
        :title="chart.name"
        style="margin-bottom: 10px"
      />
    </div>

  </div>
</template>
<script>
import Chart from './Chart'

export default {
  components: {
    Chart
  },
  props: {
    query: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      charts: [
        {
          'id': '101',
          'name': '专利类型分布',
          'active': true
        },
        {
          'id': '106',
          'name': '技术生命周期分析',
          'active': false
        },
        {
          'id': '201',
          'name': '年度申请量分析',
          'active': false
        },
        {
          'id': '301',
          'name': '申请人排行榜',
          'active': false
        },
        {
          'id': '501',
          'name': '发明人排行榜',
          'active': false
        },
        {
          'id': '721',
          'name': '主分类年度申请趋势分析',
          'active': false
        },
        {
          'id': '1001',
          'name': '省市申请量分析',
          'active': false
        },
        {
          'id': '1302',
          'name': '年度专利转让趋势分析',
          'active': false
        },
        {
          'id': '1303',
          'name': '转让人排名分析',
          'active': false
        },
        {
          'id': '1315',
          'name': '年度专利许可趋势分析',
          'active': false
        },
        {
          'id': '1316',
          'name': '许可人排名分析',
          'active': false
        }
      ]
    }
  }
}
</script>
