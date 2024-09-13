<template>
  <div v-loading.fullscreen.lock="loading" class="app-container">
    <div>
      <div class="my-card">
        <span>选择工艺包：</span>
        <el-select v-model="currentPackage" placeholder="请选择">
          <el-option
            v-for="item in packageList"
            :key="item.packageName"
            :label="item.packageName"
            :value="item.packageId"
          />
        </el-select>
      </div>
      <div v-show="tableData.length>0" class="my-card" style="position: relative;margin-top: 15px">
        <div id="myChart" style="height:800px;width: 800px" />
        <div v-show="tableData.length>0" class="table">
          <el-table
            :data="tableData"
            stripe
            style="width: 100%"
          >
            <el-table-column
              align="center"
              label="排名"
              prop="id"
              width="60"
            >
              <template slot-scope="{row}">
                <slot>{{ parseInt(row.id) + 1 }}</slot>
              </template>
            </el-table-column>
            <el-table-column
              align="center"
              label="发明人"
              prop="name"
            />
            <el-table-column
              align="center"
              label="关联指数"
              prop="value"
            />
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import echarts from 'echarts'
import { getPackageList, getRelationGraphByPackageId } from '@/api/package'

const option = {
  title: {
    text: '关联图谱',
    textAlign: 'auto',
    left: '10%', // '5' | '5%'，title 组件离容器左侧的距离
    right: 'auto', // 'title 组件离容器右侧的距离
    top: '10%', // title 组件离容器上侧的距离
    bottom: 'auto'
  },
  tooltip: {},
  legend: [
    {
      data: []
    }
  ],
  series: [
    {
      name: '专利发明人',
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
      animation: false,
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
  name: 'AssociationGraph',
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
    currentPackage: function(val) {
      this.getChartData(val)
    }
  },
  mounted() {
    this.chart = echarts.init(document.getElementById('myChart'))
    this.loadPackageList()
  },
  methods: {
    loadPackageList() {
      getPackageList().then(res => {
        this.packageList = res.data.data
      })
    },
    getChartData(packageId) {
      this.loading = true
      getRelationGraphByPackageId(packageId).then(res => {
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
<style>
.table {
  position: absolute;
  z-index: 99;
  top: 5px;
  right: 5px;
  overflow-y: auto;
  width: 250px;
  border: 1px solid #ccc;
}
</style>
