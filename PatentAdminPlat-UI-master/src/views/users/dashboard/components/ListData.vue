<template>
  <el-row :gutter="20">
    <el-col :lg="currentType===-1?5:6" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          合作分析
        </div>
        <div class="list">
          <div v-for="(field,index) in collaborators" :key="field.name" class="list-item">
            <div style="width: 40px">
              <div :class="{'indicator-primary':index<=2}" class="indicator">{{ index + 1 }}</div>
            </div>
            <div style="width: calc(100% - 40px);display: flex;flex-direction: row;align-items: center">
              <div style="width: 30%">
                {{ field.name }}
              </div>
              <div style="width: 70%">
                <el-progress :percentage="field.times" :stroke-width="20" :text-inside="true" />
              </div>
            </div>
          </div>

        </div>
      </div>
    </el-col>
    <el-col v-if="currentType===-1" :lg="5" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          竞争分析
        </div>
        <div class="list">
          <div v-for="(field,index) in competitors" :key="field.name" class="list-item">
            <div style="width: 40px">
              <div :class="{'indicator-primary':index<=2}" class="indicator">{{ index + 1 }}</div>
            </div>
            <div style="width: calc(100% - 40px);display: flex;flex-direction: row;align-items: center">
              <div style="width: 100%">
                {{ field.name }}
              </div>
            </div>
          </div>

        </div>
      </div>
    </el-col>
    <el-col :lg="currentType===-1?5:6" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          申请人
        </div>
        <div class="list">
          <div v-for="(field,index) in claimApartments" :key="field.name" class="list-item">
            <div style="width: 40px">
              <div :class="{'indicator-primary':index<=2}" class="indicator">{{ index + 1 }}</div>
            </div>
            <div style="width: calc(100% - 40px);display: flex;flex-direction: row;align-items: center">
              {{ field.name }}
            </div>
          </div>

        </div>
      </div>
    </el-col>
    <el-col :lg="currentType===-1?4:6" :sm="24" :xs="24">
      <div class="my-card" style="height: 400px">
        <div class="card-title">
          发明人
        </div>
        <div class="list">
          <div v-for="(field,index) in claimInventors" :key="field.name" class="list-item">
            <div style="width: 40px">
              <div :class="{'indicator-primary':index<=2}" class="indicator">{{ index + 1 }}</div>
            </div>
            <div style="width: calc(100% - 40px);display: flex;flex-direction: row;align-items: center">
              {{ field.name }}
            </div>
          </div>

        </div>
      </div>
    </el-col>
    <el-col :lg="currentType===-1?5:6" :sm="24" :xs="24">
      <div style="display: flex;flex-direction: column;justify-content: space-between;height: 400px">
        <router-link :to="{path:'/report/generate/novelty'}">
          <div class="my-card v-card ">
            <div class="card-title ">
              <svg-icon class-name="card-panel-icon" icon-class="new" />
              查新报告入口
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>
        <router-link :to="{path:'/report/generate/infringement'}">
          <div class="my-card v-card">
            <div class="card-title">
              <svg-icon class-name="card-panel-icon" icon-class="infringement" />
              侵权报告入口
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>
        <router-link :to="{path:'/report/generate/valuation'}">
          <div class="my-card v-card">
            <div class="card-title">
              <svg-icon class-name="card-panel-icon" icon-class="valuation" />
              估值报告入口
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>
        <download-able
          name="交底书模版.docx"
          url="http://www.sriptb.com/static/docs/patent_confession_template.doc"
        >
          <div class="my-card v-card">
            <div class="card-title">
              <svg-icon class-name="card-panel-icon" icon-class="report" />
              交底书模版下载
            </div>
            <div style="font-size: 14px">
              点击下载
            </div>
          </div>
        </download-able>

      </div>
    </el-col>

  </el-row>
</template>
<script>

import DownloadAble from '@/views/users/components/DownloadAble.vue'

export default {
  name: 'ListData',
  components: { DownloadAble },
  props: {
    currentType: {
      type: Number,
      default: -1
    },
    dashboardData: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      collaborators: [],
      competitors: [],
      claimApartments: [],
      claimInventors: []

    }
  },
  watch: {
    dashboardData: {
      handler: function(val) {
        this.claimApartments = val.claimApartments.slice(0, 5)
        this.claimInventors = val.claimInventors.slice(0, 5)
        this.collaborators = val.collaborators.slice(0, 5)
        this.competitors = val.competitors.slice(0, 5)
        // 将collaborators和competitors数组里面的times求和并设置成计算百分比
        const collaboratorSum = this.collaborators.reduce((prev, cur) => {
          return prev + cur.times
        }, 0)
        const sum = this.competitors.reduce((prev, cur) => {
          return prev + cur.times
        }, 0)
        this.collaborators.forEach(item => {
          item.times = Math.round(item.times / collaboratorSum * 100)
        })
        this.competitors.forEach(item => {
          item.times = Math.round(item.times / sum * 100)
        })
      },
      deep: true
    }
  }
}
</script>
<style scoped>
.card-title {
  font-size: 16px;
  margin: 0 0 7px;
  font-weight: 600;
}

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
  height: 90px;
  cursor: pointer
}

</style>
