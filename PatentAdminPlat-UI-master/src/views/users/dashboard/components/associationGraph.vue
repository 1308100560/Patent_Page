<template>
  <div id="myChart2" style="height:100%;width: 100%" />
</template>
<script>
import echarts from 'echarts'
import { getGraphData } from '@/api/graph'

const option = {

  tooltip: {},
  legend: [
    {
      data: [],
      show: false
    }
  ],
  series: [
    {
      name: '专利关键词',
      type: 'graph',
      layout: 'force',
      data: null,
      links: null,
      categories: [],
      roam: true,
      label: {
        show: true,
        formatter: '{b}',
        position: 'right'
      },
      focusNodeAdjacency: true,
      legendHoverLink: true,
      animation: true,
      force: {
        initLayout: 'circular',
        layoutAnimation: false,
        repulsion: 100
      },
      lineStyle: {
        color: 'source',
        opacity: 0.2,
        curveness: 0.3
      },
      emphasis: {
        focus: 'adjacency',
        itemStyle: {
          shadowColor: 'rgba(0, 0, 0, 0.4)',
          shadowBlur: 15
        },
        lineStyle: {
          width: 3
        },
        label: {
          textBorderColor: 'rgba(255, 255, 255, 0.8)',
          textBorderWidth: 2
        }
      }
    }
  ]
}
export default {
  name: 'TechGraph',
  props: {
    currentType: {
      type: Number,
      default: -1
    }
  },
  data() {
    return {
      chart: null,
      packageList: [],
      currentPackage: '请选择',
      loading: false,
      tableData: []
    }
  },
  watch: {
    currentType() {
      this.getChartData()
    }
  },
  mounted() {
    this.chart = echarts.init(document.getElementById('myChart2'))
    this.getChartData()
  },
  methods: {
    getChartData() {
      this.loading = true
      const params = {
        scope: this.currentType === -1 ? 'all-claimed' : 'package',
        type: 'rela',
        packageId: this.currentType === -1 ? null : this.currentType
      }
      getGraphData(params).then(res => {
        const results = res.data.data
        if (results == null) {
          option.series[0].data = null
          option.series[0].links = null
        } else {
          option.series[0].data = results.nodes
          option.series[0].links = results.links
          this.tableData = results.nodes.slice(0, 10)
          option.legend[0].data = this.tableData.map(item => item.name)
          option.series[0].categories = this.tableData.map(item => {
            return { name: item.name }
          })
          console.log(option)
          this.chart.setOption(option)
        }
        this.loading = false
      })
    }
  }
}
</script>
