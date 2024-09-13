<template>
  <div ref="chart" style="height: 120px;width: 100%!important;" />
</template>
<script>
import echarts from 'echarts'
import resize from '@/views/users/search/components/mixins/resize'

require('echarts/theme/macarons') // echarts theme
export default {
  mixins: [resize],
  props: {
    patent: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      loading: false
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.refresh()
    })
  },
  methods: {
    getSource() {
      return this.patent.price ? '算法估值' : '报告估值'
    },
    getPriceDisplayText() {
      const price = this.patent.price ? this.patent.price : this.patent.eval_result.evalPrice
      return parseFloat(price)
    },
    refresh() {
      this.initChart()
    },
    initChart() {
      this.loading = true
      const self = this

      this.echart = echarts.init(this.$refs.chart)
      const option = {
        series: [
          {
            type: 'gauge',
            startAngle: 180,
            endAngle: 0,
            center: ['50%', '75%'],
            radius: '90%',
            min: 0,
            max: 100000,
            splitNumber: 8,
            pointer: {
              show: false
            },
            axisLine: {
              lineStyle: {
                width: 10,
                color: [[0.2, '#37A2DA'], [0.8, '#37A2DA'], [1, '#37A2DA']]
              }
            },
            axisTick: {
              show: false
            },
            splitLine: {
              length: 10,
              lineStyle: {
                width: 2
              }
            },
            axisLabel: {
              show: false
            },
            detail: {
              valueAnimation: true,
              formatter: '{value}万',
              fontSize: 16,
              offsetCenter: [0, '-10%']
            },
            title: {
              offsetCenter: [0, '30%'],
              fontSize: 12
            },
            data: [
              {
                value: this.getPriceDisplayText(),
                name: '来源：' + this.getSource()
              }
            ]
          }
        ]
      }

      this.echart.setOption(option)
      let count = 0
      const interval = setInterval(() => {
        self.echart.resize()
        count++
        if (count > 10) {
          clearInterval(interval)
        }
      }, 500)
    }
  }
}
</script>

