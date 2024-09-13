<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
import resize from './mixins/resize'

require('echarts/theme/macarons') // echarts theme

const animationDuration = 6000

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '300px'
    },
    dashboardData: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data() {
    return {
      chart: null,
      year: [],
      count: []
    }
  },
  watch: {
    dashboardData: {
      handler: function(val) {
        this.year = val.publicationDates.map(item => item.year).reverse()
        this.count = val.publicationDates.map(item => item.count).reverse()
        this.initChart(this.year, this.count)
      },
      deep: true
    }
  },

  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart(year, count) {
      this.chart = echarts.init(this.$el, 'macarons')

      this.chart.setOption(
        {
          title: {
            show: true,
            text: '我的专利发表年份'
          },

          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            }
          },
          grid: {
            top: 40,
            left: '2%',
            right: '2%',
            bottom: '3%',
            containLabel: true
          },
          xAxis: [{
            type: 'category',
            data: year,
            axisTick: {
              alignWithLabel: true
            }
          }],
          yAxis: [{
            type: 'value',
            axisTick: {
              show: false
            }
          }],
          series: [{
            name: '专利发表数量',
            type: 'line',
            data: count,
            animationDuration,
            itemStyle: {
              color: '#ffb980'
            },
            label: {
              show: true,
              position: 'top'
            }
          }, {
            name: '专利发表数量',
            type: 'bar',
            data: count,
            animationDuration

          }]
        }
      )
    }
  }
}
</script>
