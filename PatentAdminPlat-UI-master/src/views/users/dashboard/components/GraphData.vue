<template>
  <el-row :gutter="20" style="margin-bottom: 50px">

    <el-col :lg="5" :sm="24" :xs="24">
      <div style="display: flex;flex-direction: column;justify-content: space-between;height: 400px">
        <div class="my-card v-card" @click="getCompetitorData">
          <div class="card-title">
            竞争对手/行业人才
          </div>
          <div style="font-size: 14px">
            快捷跳转
          </div>
        </div>
        <div class="my-card v-card" @click="getInfringementData">
          <div class="card-title">
            技术前延
          </div>
          <div style="font-size: 14px">
            快捷跳转
          </div>
        </div>
        <div class="my-card v-card" @click="getPatentData">
          <div class="card-title">
            专利监测
          </div>
          <div style="font-size: 14px">
            快捷跳转
          </div>
        </div>

      </div>
    </el-col>
    <el-col :lg="9" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          相关专利图谱
        </div>
        <tech-graph :current-type="currentType" />
      </div>
    </el-col>
    <el-col :lg="10" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          相关人员图谱
        </div>
        <association-graph :current-type="currentType" />
      </div>
    </el-col>

  </el-row>
</template>
<script>
import TechGraph from '@/views/users/dashboard/components/techGraph.vue'
import AssociationGraph from '@/views/users/dashboard/components/associationGraph.vue'
import { getICG } from '@/api/patent'

export default {
  name: 'ListData',
  components: {
    TechGraph,
    AssociationGraph
  },
  props: {
    currentType: {
      type: Number,
      default: -1
    }
  },
  data() {
    return {}
  },
  methods: {
    getCompetitorData() {
      getICG().then(res => {
        const icgInfo = res.data.data
        const queryList = []
        let i = 0
        for (const item of icgInfo) {
          queryList.push(
            {
              id: i,
              fieldName: 'ICG',
              value: item.ICG,
              operation: 'or'
            }
          )
          i++
        }
        this.$router.push({ path: '/search/competitor', query: { q: queryList }})
      })
    },
    // 技术前延
    getInfringementData() {
      getICG().then(res => {
        const icgInfo = res.data.data
        const queryList = []
        for (const item of icgInfo) {
          queryList.push(`ICG = '${item.ICG}'`)
        }
        this.$router.push({ path: '/search/results', query: { q: queryList.join(' or ') }})
      })
    },
    // 专利监测
    getPatentData() {
      this.$message({
        message: '暂未开放',
        type: 'warning'
      })
    }
  }
}
</script>
<style scoped>

.list {
  display: flex;
  font-size: 14px;
  flex-direction: column;
}

.indicator {
  margin-right: 15px;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  text-align: center;
  line-height: 20px;
  background-color: lightgray;
}

.indicator-primary {
  background-color: black;
  color: white;

}

.list-item {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-top: 30px;
}

.v-card {
  height: 125px;
  cursor: pointer
}

</style>
