<template>
  <el-card ref="chartCard" class="box-card" style="min-height: 500px!important;width: 100%!important;">
    <div slot="header" class="clearfix">
      <span>{{ title }}</span>
      <el-button style="float: right; padding: 3px 0" type="text" />
    </div>
    <div ref="chart" style="height: 400px;width: 100%!important;" />
  </el-card>

</template>
<script>
import echarts from 'echarts'
import resize from './mixins/resize'
import { getChartOption } from '@/api/chart'

require('echarts/theme/macarons') // echarts theme
export default {
  mixins: [resize],
  props: {
    title: {
      type: String,
      default: '图谱'
    },
    query: {
      type: String,
      default: ''
    },
    chart: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data() {
    return {}
  },
  watch: {
    query: {
      handler: function(val) {
        this.query = val
      },
      deep: true
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart(this.chart.id, this.query)
    })
  },
  methods: {

    initChart(chartId, query) {
      this.echart = echarts.init(this.$refs.chart, 'macarons')
      console.log(chartId, query)
      getChartOption(chartId, { query }).then(({ data }) => {
        const { option } = data.data
        this.echart.setOption(JSON.parse(option))

        // timer to resize chart
        let count = 0
        const interval = setInterval(() => {
          this.echart.resize()
          count++
          if (count > 10) {
            clearInterval(interval)
          }
        }, 500)
      })
    }
  }
}
</script>

<style scoped>
/deep/ .el-progress-bar__innerText {
  color: black;
}
</style>
