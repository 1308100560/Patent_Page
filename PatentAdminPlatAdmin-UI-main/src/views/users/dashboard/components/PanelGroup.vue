<template>
  <el-row :gutter="20" class="panel-group">
    <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
      <router-link to="/patent/claim">
        <div class="card-panel">
          <div class="card-info">
            <div class="card-panel-description">
              <count-to
                :duration="2600"
                :end-val="dashboardData.patentClaimCount"
                :start-val="0"
                class="card-panel-num"
              />
              <div class="card-panel-text">
                专利数量
              </div>

            </div>
            <div class="card-panel-icon-wrapper icon-people">
              <svg-icon class-name="card-panel-icon" icon-class="patent" />
            </div>

          </div>
          <!--          <div class="growth"> 相比上周 <i class="el-icon-caret-top" /> 2.65%</div>-->
        </div>
      </router-link>
    </el-col>
    <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
      <router-link to="/patent/package">
        <div class="card-panel">
          <div class="card-info">

            <div class="card-panel-description">
              <count-to :duration="3000" :end-val="dashboardData.packageCount" :start-val="0" class="card-panel-num" />

              <div class="card-panel-text">
                工艺包数量
              </div>
            </div>
            <div class="card-panel-icon-wrapper icon-message">
              <svg-icon class-name="card-panel-icon" icon-class="pack" />
            </div>
          </div>
          <!--          <div class="growth"> 相比上周 <i class="el-icon-caret-top" /> 2.65%</div>-->
        </div>
      </router-link>

    </el-col>
    <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
      <router-link to="/trade/estimate">
        <div class="card-panel">
          <div class="card-info">
            <div class="card-panel-description">
              <count-to
                :duration="3200"
                :end-val="dashboardData.patentTotalPrice"
                :start-val="0"
                class="card-panel-num"
              />
              <div class="card-panel-text">
                专利估值 (RMB)
              </div>

            </div>
            <div class="card-panel-icon-wrapper icon-money">
              <svg-icon class-name="card-panel-icon" icon-class="money" />
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
              <div class="card-panel-num">
                {{ dashboardData.patentStatus && getNumber(dashboardData) }}
              </div>
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
</template>

<script>
import CountTo from 'vue-count-to'

export default {
  components: {
    CountTo
  },
  props: {
    dashboardData: {
      type: Object,
      default: () => {
      }
    }
  },
  methods: {

    getNumber(dashboardData) {
      console.log(dashboardData.reportStatus)
      const novelty = parseInt(dashboardData.reportStatus['查新报告'] || 0)
      const infringement = parseInt(dashboardData.reportStatus['侵权报告'] || 0)
      const valuation = parseInt(dashboardData.reportStatus['估值报告'] || 0)
      return novelty + infringement + valuation
    }
  }
}
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;
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
