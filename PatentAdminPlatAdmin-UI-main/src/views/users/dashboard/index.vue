<template>
  <div class="dashboard-editor-container">
    <panel-group :dashboard-data="dashboardData" />
    <el-row :gutter="20">
      <el-col :lg="9" :sm="24" :xs="24">
        <div class="chart-wrapper">
          <pie-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
      <el-col :lg="15" :sm="24" :xs="24">
        <div class="chart-wrapper">
          <bar-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
    </el-row>
    <list-data :dashboard-data="dashboardData" />
    <div style="margin-top: 30px">
      <graph-data />
    </div>
    <div style="margin-top: 30px">
      <property-classroom />
    </div>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import PieChart from './components/PieChart'
import BarChart from './components/BarChart'
import PropertyClassroom from '@/views/users/dashboard/components/PropertyClassroom'
import { getDashboardData } from '@/api/dashboard'
import ListData from '@/views/users/dashboard/components/ListData.vue'
import GraphData from '@/views/users/dashboard/components/GraphData.vue'

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    PieChart,
    PropertyClassroom,
    BarChart,
    ListData,
    GraphData

  },
  data() {
    return {
      dashboardData: {}
    }
  },
  mounted() {
    getDashboardData().then(res => {
      this.dashboardData = res.data.data
      console.log(this.dashboardData)
    })
  },
  methods: {}
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 10px 20px;
  background-color: rgb(245, 246, 248);
  position: relative;

  .chart-wrapper {
    padding: 16px 16px 0;
    margin-bottom: 32px;
    border-radius: .25rem;
    word-wrap: break-word;
    background-color: #fff;
    background-clip: border-box;
    border: 0 solid #f6f6f6;
    color: #666;
    box-shadow: 0 2px 4px rgb(15 34 58 / 12%);
  }
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
