<template>
  <div class="dashboard-editor-container">
    <div>
      <el-select v-model="currentType" filterable placeholder="选择专利组合" style="width: 300px" value="">
        <el-option :value="-1" label="所有记录" />
        <el-option v-for="p in packageList.list" :key="p.packageId" :label="p.packageName" :value="p.packageId" />
      </el-select>
    </div>
    <panel-group :dashboard-data="dashboardData" />
    <el-row :gutter="20">
      <el-col :lg="9" :sm="24" :xs="24">
        <div class="chart-wrapper" @click="$router.push('/patent/claim')">
          <pie-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
      <el-col :lg="15" :sm="24" :xs="24">
        <div class="chart-wrapper" @click="$router.push('/patent/claim')">
          <bar-chart :dashboard-data="dashboardData" />
        </div>
      </el-col>
    </el-row>
    <list-data :current-type="currentType" :dashboard-data="dashboardData" />
    <div style="margin-top: 30px">
      <graph-data :current-type="currentType" />
    </div>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import PieChart from './components/PieChart'
import BarChart from './components/BarChart'
import { getDashboardData, getDashboardDataByPackage } from '@/api/dashboard'
import ListData from '@/views/users/dashboard/components/ListData.vue'
import GraphData from '@/views/users/dashboard/components/GraphData.vue'
import { getPackageList } from '@/api/package'

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    PieChart,
    BarChart,
    ListData,
    GraphData

  },
  data() {
    return {
      dashboardData: {},
      currentType: -1,
      packageList: []
    }
  },
  watch: {
    currentType() {
      if (this.currentType === -1) {
        getDashboardData().then(res => {
          this.dashboardData = res.data.data
        })
      } else {
        getDashboardDataByPackage(this.currentType).then(res => {
          const data = res.data.data
          data.patentClaimCount = data.patentCount
          data.claimApartments = data.packageApartments
          data.claimInventors = data.packageInventors
          data.competitors = []
          this.dashboardData = data
        })
      }
    }
  },
  mounted() {
    getDashboardData().then(res => {
      this.dashboardData = res.data.data
      console.log(this.dashboardData)
    })
    getPackageList({ pageIndex: 1, pageSize: 999999 }).then(res => {
      this.packageList = res.data.data
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
