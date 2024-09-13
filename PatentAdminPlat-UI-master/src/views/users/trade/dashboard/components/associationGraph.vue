<template>
  <div id="myChart2" style="height:100%;width: 100%" />
</template>
<script>
import echarts from 'echarts'
import { getFocusRelationGraph } from '@/api/patent'

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
  data() {
    return {
      chart: null,
      packageList: [],
      currentPackage: '请选择',
      loading: false,
      tableData: []
    }
  },
  mounted() {
    this.chart = echarts.init(document.getElementById('myChart2'))
    this.getChartData()
  },
  methods: {
    getChartData() {
      this.loading = true
      getFocusRelationGraph().then(res => {
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
