<template>
  <div class="my-card" style="display: flex;flex-direction: row;justify-content: space-between">
    <create-package ref="createPack" :get-list="()=>{loadInfoPanelData(patent.PNM)}" />
    <create-package ref="createFocusPack" :get-list="()=>{loadInfoPanelData(patent.PNM)}" type="focus" />
    <div :style="{width:`calc(100% - ${actionWidth}px)`}" class="my-card" style="overflow-y: auto">
      <div style="height: 30px">
        <el-page-header content="专利详情" style="" @back="goBack" />
      </div>
      <div
        v-loading="fullscreenLoading"
        style="
        height: calc(100% - 30px);
        overflow-y: auto;;
        max-width: 1140px;
        margin:0 auto;display: flex;flex-direction: column;position: relative!important;"
      >

        <div
          style="width:80%;margin: 10px auto;display: flex;flex-direction: row;align-items: center;justify-content: center"
        >
          <h3 style="text-align: center">{{ patent.TI }}</h3>
        </div>

        <div class="info">

          <div class="row-desc">
            <div>
              <span>申请(专利)号：	</span>
              <span>{{ patent.AN }}</span>
            </div>
            <div>
              <span>申请日：	</span>
              <span>{{ patent.AD }}</span>
            </div>
          </div>
          <div class="row-desc">
            <div>
              <span>申请公布号：	</span>
              <span>	{{ patent.PNM }}</span>
            </div>
            <div>
              <span>公开公告日：</span>
              <span>	{{ patent.PD }}</span>
            </div>
          </div>
          <div class="row-desc">
            <div>
              <span>申请人：	</span>
              <span>		{{ patent.PA }}</span>
            </div>
          </div>
          <div class="row-desc">
            <div>
              <span>地址：</span>
              <span>	{{ patent.AR }}</span>
            </div>
          </div>
          <div class="row-desc">
            <div>
              <span>  发明人：</span>
              <span>{{ patent.INN }}</span>
            </div>
          </div>
          <div class="desc">
            <span>摘要：</span>
            <pre>{{ patent.ABST }}</pre>
          </div>
          <div class="desc">
            <span>权利要求书：</span>
            <pre>{{ patent.CLM }}</pre>
          </div>
          <div class="desc">
            <span>说明书：</span>
            <pre>{{ patent.DESCR }}</pre>
          </div>

        </div>
      </div>
    </div>
    <div
      v-loading="fullscreenLoading"
      :style="{width:`${actionWidth-10}px`}"
      class="my-card"
      style="padding: 10px!important;padding-left: 15px;overflow-y: auto"
    >
      <div style="display: flex;flex-direction: column;align-items: flex-start">
        <div style="width: 100%">
          <div
            class="d-flex flex-row align-items-center"
            style="justify-content: space-between;width: 100%"
            @click="showAction"
          >
            <div style="display: flex;flex-direction: row;align-items: center;margin-top: 10px">
              <div v-show="actionVisible"><strong>相关信息及操作</strong></div>
            </div>
            <!--            <div v-show="actionVisible" style="margin-left: 5px;font-size:14px;cursor: pointer">收起</div>-->
          </div>
          <div v-show="actionVisible" style="margin-top: 10px">
            <div style="display: flex;flex-direction: row">
              <patent-claim-popover
                :after-confirm="()=>{loadInfoPanelData(patent.PNM)}"
                :patent="patent"
                style="margin-right: 5px"
              />
              <patent-focus-popover
                :after-confirm="()=>{loadInfoPanelData(patent.PNM)}"
                :patent="patent"
                style="margin-right: 5px"
              />
            </div>
          </div>
        </div>
        <div v-if="patentInDb" style="margin-top: 10px;width: 100%">
          <div class="d-flex flex-row align-items-center" style="width: 100%;justify-content: space-between">
            <div style="display: flex;flex-direction: row;align-items: center">
              <i class="el-icon-money" style="font-size: 28px" />
              <div v-show="actionVisible" style="margin-left: 5px"><strong>专利估值</strong></div>
            </div>
            <div v-show="actionVisible">
              <router-link to="/report/generate/valuation">
                <el-button size="mini" style="margin-left: 10px">去估值</el-button>
              </router-link>
            </div>
          </div>

          <div v-if="patentInDb.price||patentInDb.eval_result" v-show="actionVisible" style="width: 100%">
            <gauge-chart :patent="patentInDb" />
          </div>
          <div v-else v-show="actionVisible" style="padding: 10px 0;font-size: 12px">
            暂无估值数据
            <router-link
              :to="`/report/generate/valuation?pId=${patentInDb.patentId}`"
              class="link-type"
              style="margin-left: 5px"
            >去估值
            </router-link>
          </div>
        </div>
        <div v-if="patent&&patent.isClaimed" style="width: 100%;margin-top: 10px">
          <div style="width: 100%">
            <div class="d-flex flex-row align-items-center" style="width: 100%;justify-content: space-between">
              <div style="display: flex;flex-direction: row;align-items: center">
                <svg-icon icon-class="patent" style="font-size: 25px" />
                <div v-show="actionVisible" style="margin-left: 5px"><strong>专利组合</strong></div>
              </div>
              <div v-show="actionVisible">
                <el-button size="mini" style="margin-left: 10px" @click="$refs.createPack.show()">创建</el-button>
              </div>
            </div>
            <div v-show="actionVisible" style="margin-top: 10px">
              <div style="display: flex;flex-direction: row;flex-wrap: wrap;max-height: 400px;overflow-y: auto">
                <el-table
                  :data="packageList"
                  size="mini"
                  stripe
                  style="width: 100%"
                >
                  <el-table-column
                    label="专利组合"
                    width="180"
                  >
                    <template scope="{row}">
                      <div class="actions">
                        <router-link
                          :to="`/patent/package/${row.packageId}`"
                          class="link-type"
                        >{{ row.packageName }}
                        </router-link>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column
                    align="right"
                    label="操作"
                  >
                    <template scope="{row}">
                      <div class="actions">
                        <el-button
                          :icon="`el-icon-${row.selected?'delete':'plus'}`"
                          :type="`${row.selected?'danger':''}`"
                          size="mini"
                          @click="handleCheckBoxChange(row)"
                        />
                      </div>
                    </template>
                  </el-table-column>

                </el-table>

              </div>
            </div>
          </div>
        </div>
        <div v-if="patent&&patent.isFocused" style="width: 100%;margin-top: 10px">
          <div style="width: 100%">
            <div class="d-flex flex-row align-items-center" style="width: 100%;justify-content: space-between">
              <div style="display: flex;flex-direction: row;align-items: center">
                <svg-icon icon-class="guide" style="font-size: 25px" />
                <div v-show="actionVisible" style="margin-left: 5px"><strong>关注组合</strong></div>
              </div>
              <div v-show="actionVisible">
                <el-button size="mini" style="margin-left: 10px" @click="$refs.createFocusPack.show()">创建</el-button>
              </div>
            </div>
            <div v-show="actionVisible" style="margin-top: 10px">
              <div style="display: flex;flex-direction: row;flex-wrap: wrap;max-height: 400px;overflow-y: auto">
                <el-table
                  :data="focusPackageList"
                  size="mini"
                  stripe
                  style="width: 100%"
                >
                  <el-table-column
                    label="专利组合"
                    width="180"
                  >
                    <template scope="{row}">
                      <div class="actions">
                        <router-link
                          :to="`/patent/package/${row.packageId}`"
                          class="link-type"
                        >{{ row.packageName }}
                        </router-link>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column
                    align="right"
                    label="操作"
                  >
                    <template scope="{row}">
                      <div class="actions">
                        <el-button
                          :icon="`el-icon-${row.selected?'delete':'plus'}`"
                          :type="`${row.selected?'danger':''}`"
                          size="mini"
                          @click="handleCheckBoxChange(row)"
                        />
                      </div>
                    </template>
                  </el-table-column>

                </el-table>

              </div>
            </div>
          </div>
        </div>
        <div v-if="patentInDb" style="margin: 10px 0;width: 100%">
          <div class="d-flex flex-row align-items-center" style="width: 100%;justify-content: space-between">
            <div style="display: flex;flex-direction: row;align-items: center">
              <i class="el-icon-tickets" style="font-size: 28px" />
              <div v-show="actionVisible" style="margin-left: 5px"><strong>相关报告</strong></div>
            </div>
            <div v-show="actionVisible">
              <router-link to="/report/list">
                <el-button size="mini" style="margin-left: 10px">去申请报告</el-button>
              </router-link>
            </div>
          </div>

          <div v-show="actionVisible" style="width: 100%">
            <div style="margin-top: 10px;display: flex;flex-direction: row;align-items:center;flex-wrap: wrap">
              <el-tag
                :type="activeReport==='infringementReport'?'primary':'info'"
                size="small"
                style="margin-right: 5px;cursor: pointer"
                @click="activeReport='infringementReport'"
              >
                侵权报告
              </el-tag>
              <el-tag
                :type="activeReport==='valuationReport'?'primary':'info'"
                size="small"
                style="cursor: pointer"
                @click="activeReport='valuationReport'"
              >
                估值报告
              </el-tag>
            </div>
            <div
              v-if="reports[activeReport]&&reports[activeReport].length===0"
              style="font-size: 12px;margin-top: 10px"
            >
              暂无
            </div>
            <div v-else>
              <div
                v-for="report in reports[activeReport]"
                :key="`report-${report.reportId}`"
                style="margin-top: 5px;width: 100%"
              >
                <router-link :to="`/report/list?ids=${[report.reportId]}`" class="link-type" style="font-size: 12px">
                  {{ report.reportName }}
                </router-link>
                <div
                  style="
                         margin-top: 5px;
                            display: flex;
                            flex-direction: row;align-items: center;justify-content: space-between;width: 100%"
                >
                  <div style="font-size: 12px">
                    生成时间：{{ report.CreatedAt|localTime }}
                  </div>
                </div>

              </div>

            </div>
          </div>
        </div>
        <div v-if="patent&&patent.isClaimed" style="width:100%">
          <div class="d-flex flex-row align-items-center" style="width: 100%;justify-content: space-between">
            <div style="display: flex;flex-direction: row;align-items: center">
              <i class="el-icon-video-camera" style="font-size: 28px" />
              <div v-show="actionVisible" style="margin-left: 5px"><strong>关联视频</strong></div>
            </div>
            <div v-show="actionVisible">
              <el-button size="mini" style="margin-left: 10px" @click="$message.warning('尚未开放')">添加</el-button>
            </div>
          </div>
          <div v-show="actionVisible" style="margin-top: 10px;font-size: 12px">
            <div>暂无</div>
            <!--            <video controls style="width: 100%">-->
            <!--              <source src="http://www.daweisoft.com//upload/file/gaoxiaodisijiang.mp4" type="video/mp4">-->
            <!--            </video>-->
          </div>
        </div>

      </div>
    </div>
  </div>

</template>
<script>
import { claimPatent, getFocusByPNM, getPatentByPNM, getPatentDetail } from '@/api/patent'
import GaugeChart from '@/views/users/search/detailPannel/GaugeChart.vue'
import { addPatentToPackage, getPackageList, removePatentFromPackage } from '@/api/package'
import PatentClaimPopover from '@/views/users/components/PatentClaimPopover.vue'
import CreatePackage from '@/views/users/components/CreatePackage.vue'
import PatentFocusPopover from '@/views/users/components/PatentFocusPopover.vue'
import { getReportListByIds } from '@/api/report'

export default {
  name: 'Detail',
  components: { PatentFocusPopover, PatentClaimPopover, GaugeChart, CreatePackage },
  data() {
    return {
      patent: {},
      packageList: [],
      focusPackageList: [],
      currentPackage: null,
      patentInDb: null,
      actionWidth: 65,
      fullscreenLoading: false,
      actionVisible: false,
      activeReport: 'infringementReport',
      reports: {
        infringementReport: [],
        valuationReport: []
      }
    }
  },
  mounted() {
    this.fullscreenLoading = true
    const PNM = this.$route.params.id
    getPatentDetail(PNM).then(({ data }) => {
      this.patent = data.data
      this.fullscreenLoading = false
      this.loadInfoPanelData(PNM)
    })

    this.showAction()
  },
  methods: {
    loadInfoPanelData(PNM) {
      getPatentByPNM(PNM).then(({ data }) => {
        let packageIDs = []
        let focusPackageIDs = []
        if (data.data.count !== 0) {
          this.patent.isClaimed = true
          packageIDs = data.data.list[0].packageIDs ? data.data.list[0].packageIDs : []
          getPackageList({}).then(({ data }) => {
            data.data.list.map(item => {
              item.selected = packageIDs.indexOf(item.packageId) > -1
            })
            this.packageList = data.data.list
          })
        } else {
          this.patent.isClaimed = false
        }
        if (data.data.list.length > 0) {
          if (data.data.list[0].relaTypedReports) {
            getReportListByIds(JSON.stringify(data.data.list[0].relaTypedReports['侵权报告'])).then(res => {
              this.reports.infringementReport = res.data.data
            })
            getReportListByIds(JSON.stringify(data.data.list[0].relaTypedReports['估值报告'])).then(res => {
              this.reports.valuationReport = res.data.data
            })
          }
        }
        this.patentInDb = data.data.list[0]
        getFocusByPNM(PNM).then(({ data }) => {
          this.patent.isFocused = data.data.count !== 0
          if (data.data.count !== 0) {
            focusPackageIDs = data.data.list[0].packageIDs ? data.data.list[0].packageIDs : []
          }
          getPackageList({ type: 'focus' }).then(({ data }) => {
            data.data.list.map(item => {
              console.log('sasa', focusPackageIDs, item.packageId, focusPackageIDs.indexOf(item.packageId))
              item.selected = focusPackageIDs.indexOf(item.packageId) > -1
            })
            this.focusPackageList = data.data.list
            console.log(this.focusPackageList)
          })
        })
      })
    },
    handleCheckBoxChange(e) {
      this.currentPackage = e
      if (e.selected) {
        this.$confirm('确定将专利移除该专利组合中吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          removePatentFromPackage(e.packageId, this.patentInDb.PNM).then(({ data }) => {
            e.selected = false
            this.$message({
              type: 'success',
              message: '移除成功!'
            })
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消移除'
          })
        })
      } else {
        this.$confirm('确定将专利添加到该专利组合中吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          claimPatent(this.patent).then(({ data }) => {
            addPatentToPackage(e.packageId, this.patent.PNM, this.patent).then(({ data }) => {
              e.selected = true
              this.$message({
                type: 'success',
                message: '添加成功!'
              })
            })
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消添加'
          })
        })
      }
    },
    goBack() {
      this.$router.go(-1)
    },
    showAction() {
      if (this.actionVisible) {
        this.actionWidth = 65
        this.actionVisible = false
      } else {
        this.actionWidth = 300
        this.actionVisible = true
      }
    }
  }

}
</script>
<style scoped>

.info {
  width: 80%;
  min-width: 800px;
  margin: 10px auto;
}

.desc span {
  font-size: 1rem;
  font-weight: bold;
}

.desc p {
  font-size: 0.8rem;
  color: #666;
  line-height: 1.5;
}

.row-desc {
  margin: 20px 0;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.row-desc div {
  width: 50%;
}

.row-desc {
  font-size: 0.9rem;
}

.row-desc span:first-child {
  font-weight: bold;
}

pre {
  line-height: 33px;
  font-size: 14px !important;
  white-space: pre-wrap;
}

.my-checkbox {
  margin: 0 5px 5px 0 !important;
}
</style>
