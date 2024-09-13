<template>
  <div style="display: flex;flex-direction: row;align-items: center;font-size: 12px;justify-content: space-between">
    <div class="row-center" style="margin-right: 20px">
      <svg-icon class-name="card-panel-icon" icon-class="infringement" />
      侵权报告
      <router-link
        v-if="!infringementReport|| infringementReport.length===0"
        :to="`/report/generate/infringement?pId=${patentId}`"
        class="link-type "
        style="margin-left: 3px"
      >申请
      </router-link>
      <div v-else class="link-type" style="margin-left: 3px" @click="showReport(infringementReport)">
        ({{ infringementReport.length }})
      </div>
    </div>
    <div class="row-center">
      <svg-icon class-name="card-panel-icon" icon-class="valuation" />
      <span>估值报告</span>
      <router-link
        v-if="!valuationReport||valuationReport.length===0"
        :to="`/report/generate/valuation?pId=${patentId}`"
        class="link-type"
        style="margin-left: 3px"
      >申请
      </router-link>
      <div v-else class="link-type " style="margin-left: 3px" @click="showReport(valuationReport)">
        ({{ valuationReport.length }})
      </div>
    </div>
  </div>
</template>
<script>

export default {
  props: {
    getAndShowReportByIds: {
      type: Function,
      default: () => {
      }
    },
    patentId: {
      type: Number,
      default: 0
    },
    reportList: {
      type: Object,
      default: () => []
    }
  },
  data() {
    return {
      infringementReport: [],
      valuationReport: []
    }
  },
  mounted() {
    if (this.reportList) {
      this.infringementReport = this.reportList['侵权报告']
      this.valuationReport = this.reportList['估值报告']
    }
  },
  methods: {
    showReport(ids) {
      this.getAndShowReportByIds(ids)
    }
  }
}
</script>
<style scoped>
.row-center {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
