<template>
  <div class="my-card">
    <el-skeleton v-if="fullscreenLoading" :rows="30" animated class="info" />
    <div v-if="!fullscreenLoading" v-loading.fullscreen.lock="fullscreenLoading" class="container">
      <div
        style="width:90%;margin: 30px auto;display: flex;flex-direction: row;align-items: center;justify-content: space-between"
      >
        <div class="link-type" @click="goBack"><i class="el-icon-arrow-left" />返回</div>
        <h3 style="text-align: center">{{ patent.TI }}</h3>
        <div />
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
            <span>{{ patent.PINN }}</span>
          </div>
        </div>
        <div class="desc">
          <span>摘要：</span>
          <p>{{ patent.ABST }}</p>
        </div>
        <div class="desc">
          <span>权利要求书：</span>
          <p>{{ patent.CLM }}</p>
        </div>
        <div class="desc">
          <span>说明书：</span>
          <p>{{ patent.DESCR }}</p>
        </div>

      </div>
    </div>
  </div>

</template>
<script>
import { getPatentDetail } from '@/api/patent'

export default {
  name: 'Detail',
  data() {
    return {
      patent: {},
      fullscreenLoading: false
    }
  },
  mounted() {
    this.fullscreenLoading = true
    const PNM = this.$route.params.id
    getPatentDetail(PNM).then(({ data }) => {
      this.patent = data.data
      this.fullscreenLoading = false
    })
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    }
  }

}
</script>
<style scoped>

.info {
  width: 90%;
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
</style>
