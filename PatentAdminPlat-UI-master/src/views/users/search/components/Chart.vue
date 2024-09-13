<template>
  <el-card
    ref="chartCard"
    v-loading="loading"
    class="box-card"
    style="height: 450px!important;width: 100%!important;"
  >
    <div slot="header" class="clearfix">
      <span>{{ title }}</span>
      <el-button style="float: right; padding: 3px 0" type="text" @click="refresh">刷新</el-button>
    </div>
    <div v-show="chartLoaded" ref="chart" style="height: 350px;width: 100%!important;" />
    <div v-show="!chartLoaded&&chart.id" style="position: relative">
      <img
        :src="require(`@/assets/charts/${chart.id}.png`)"
        alt="chart"
        class="grey-image"
      >
      <div class="refresh">
        <el-button icon="el-icon-refresh" type="primary" @click="refresh">刷新</el-button>
      </div>

    </div>

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
    return {
      loading: false,
      chartLoaded: false
    }
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

  },
  methods: {
    refresh() {
      this.chartLoaded = true
      this.initChart(this.chart.id, this.query)
    },
    initChart(chartId, query) {
      const self = this
      this.loading = true
      this.echart = echarts.init(this.$refs.chart, 'macarons')
      getChartOption(chartId, { query }).then(({ data }) => {
        const { option } = data.data
        this.echart.setOption(JSON.parse(option))
        this.echart.on('dataViewChanged', function() {
          console.log('dataViewChanged')
        })
        self.loading = false
        // timer to resize chart
        let count = 0
        const interval = setInterval(() => {
          self.echart.resize()
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

.grey-image {
  height: 350px;
  width: 100% !important;
  object-fit: contain;
  opacity: 0.2;
}

.refresh {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
