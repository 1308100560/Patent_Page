<template>
  <div class="container">
    <el-row :gutter="20" class="panel-group">
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link to="/report/generate/novelty">
          <div class="card-panel">
            <div class="card-info">
              <div class="card-panel-description">
                <div class="card-panel-num">
                  查新报告
                </div>
                <div class="card-panel-text">
                  立即生成
                </div>

              </div>
              <div class="card-panel-icon-wrapper icon-people">
                <svg-icon class-name="card-panel-icon" icon-class="new" />
              </div>

            </div>
          </div>
        </router-link>
      </el-col>
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link to="/report/generate/infringement">
          <div class="card-panel">
            <div class="card-info">

              <div class="card-panel-description">
                <div class="card-panel-num">
                  侵权报告
                </div>
                <div class="card-panel-text">
                  立即生成
                </div>
              </div>
              <div class="card-panel-icon-wrapper icon-message">
                <svg-icon class-name="card-panel-icon" icon-class="infringement" />
              </div>
            </div>
            <!--          <div class="growth"> 相比上周 <i class="el-icon-caret-top" /> 2.65%</div>-->
          </div>
        </router-link>

      </el-col>
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link to="/report/generate/valuation">
          <div class="card-panel">
            <div class="card-info">
              <div class="card-panel-description">

                <div class="card-panel-num">
                  估值报告
                </div>
                <div class="card-panel-text">
                  立即生成
                </div>

              </div>
              <div class="card-panel-icon-wrapper icon-money">
                <svg-icon class-name="card-panel-icon" icon-class="valuation" />
              </div>
            </div>
            <!--          <div class="growth"> 相比上周 <i class="el-icon-caret-top" /> 2.65%</div>-->
          </div>
        </router-link>
      </el-col>
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link to="/report/list">
          <div class="card-panel">
            <div class="card-info">

              <div class="card-panel-description">
                <div class="card-panel-num">其他报告</div>
                <div class="card-panel-text">
                  报告数量
                </div>

              </div>
              <div class="card-panel-icon-wrapper icon-shopping">
                <svg-icon class-name="card-panel-icon" icon-class="report" />
              </div>
            </div>
            <!--          <div class="growth">相比上周 <i class="el-icon-caret-top" /> 2.65%-->
          </div>
        </router-link>
      </el-col>
    </el-row>
    <div class="my-card" style="width: 100%;">
      <el-tabs v-model="activeName">

        <el-tab-pane :key="'first'" label="专利维度" name="first">
          <div v-if="activeName === 'first'">
            <patent-tab :data-list="patentResults" :get-and-show-report-by-ids="getAndShowReportByIds" />
          </div>
        </el-tab-pane>
        <el-tab-pane :key="'second'" label="报告维度" name="second">
          <!--解决切换闪一下-->
          <div v-if="activeName === 'second'">
            <report-tab :data-list="reportResults" />
          </div>

        </el-tab-pane>
      </el-tabs>
    </div>

  </div>

</template>
<script>

import { getReportListByIds, userReportList } from '@/api/report'
import ReportTab from '@/views/users/report/tabs/ReportTab.vue'
import PatentTab from '@/views/users/report/tabs/PatentTab.vue'
import { getClaimedPatents } from '@/api/patent'

export default {
  name: 'ReportList',
  components: {
    ReportTab, PatentTab
  },
  data() {
    return {
      activeName: 'first',
      reportResults: null,
      patentResults: null
    }
  },

  created() {
    const ids = this.$route.query.ids
    if (ids) {
      this.getAndShowReportByIds(ids.split(',').map(item => parseInt(item)))
    } else {
      this.getReportList()
    }
    this.getPatentList()
  },
  methods: {
    getAndShowReportByIds(ids) {
      this.activeName = 'second'
      getReportListByIds(JSON.stringify(ids)).then(response => {
        this.reportResults = {
          list: response.data.data
        }
        console.log(this.reportResults)
      })
    },
    getReportList() {
      userReportList(this.listQuery).then(response => {
        this.reportResults = response.data.data
      })
    },
    getPatentList() {
      getClaimedPatents({ needEvalResult: true }).then(response => {
        const results = response.data.data
        console.log(results)
        results.list.map(item => {
          item.patentProperties = item.patentProperties ? JSON.parse(item.patentProperties) : {}
        })
        results.list.forEach((item, index) => {
          item.id = index + 1
        })
        this.patentResults = results
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.panel-group {
  margin-top: 10px;
}

.growth {
  font-size: 12px;
  color: #999;
  padding: 0 20px;
  margin-bottom: 14px;
}

.card-info {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin: 14px 20px 0 20px;
}

.card-panel-col {
  margin-bottom: 30px;
}

.card-panel {
  display: flex;
  flex-direction: column;
  font-size: 12px;
  padding-bottom: 10px;
  border-radius: .25rem;
  background-color: #fff;
  background-clip: border-box;
  border: 0 solid #f6f6f6;
  color: #666;
  box-shadow: 0 2px 4px rgb(15 34 58 / 12%);
}

.card-panel:hover {
  .card-panel-icon-wrapper {
    color: #fff;
  }

  .icon-people {
    background: #40c9c6;
  }

  .icon-message {
    background: #36a3f7;
  }

  .icon-money {
    background: #f4516c;
  }

  .icon-shopping {
    background: #34bfa3
  }
}

.icon-people {
  color: #40c9c6;
}

.icon-message {
  color: #36a3f7;
}

.icon-money {
  color: #f4516c;
}

.icon-shopping {
  color: #34bfa3
}

.card-panel-icon-wrapper {
  padding: 16px;
  transition: all 0.38s ease-out;
  border-radius: 6px;
}

.card-panel-icon {
  float: left;
  font-size: 40px;
}

.card-panel-description {
  font-weight: bold;
}

.card-panel-text {
  line-height: 16px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 15px;
  margin-top: 12px;
}

.card-panel-num {
  font-size: 23px;
}

</style>
