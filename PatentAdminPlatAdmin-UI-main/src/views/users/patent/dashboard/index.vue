<template>
  <div class="dashboard-editor-container">
    <panel-group :dashboard-data="dashboardData" @handleSetLineChartData="handleSetLineChartData" />
    <el-row :gutter="32">
      <!--      <el-col :lg="8" :sm="24" :xs="24">-->
      <!--        <div class="chart-wrapper">-->
      <!--          <raddar-chart />-->
      <!--        </div>-->
      <!--      </el-col>-->
      <el-col :lg="12" :sm="24" :xs="24">
        <div class="chart-wrapper">
          <pie-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
      <el-col :lg="12" :sm="24" :xs="24">
        <div class="chart-wrapper">
          <bar-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="8">
      <el-col
        :lg="{span: 24}"
        :md="{span: 24}"
        :sm="{span: 24}"
        :xl="{span: 24}"
        :xs="{span: 24}"
        style="padding-right:8px;margin-bottom:30px;"
      />

    </el-row>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import PieChart from './components/PieChart'
import BarChart from './components/BarChart'
import { getDashboardData } from '@/api/dashboard'

const lineChartData = {
  newVisitis: {
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    PieChart,
    BarChart

  },
  data() {
    return {
      dashboardData: {},
      lineChartData: lineChartData.newVisitis
    }
  },
  mounted() {
    getDashboardData().then(res => {
      this.dashboardData = res.data.data
      console.log(this.dashboardData)
    })
  },
  methods: {
    handleSetLineChartData(type) {
      this.lineChartData = lineChartData[type]
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 10px 20px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
