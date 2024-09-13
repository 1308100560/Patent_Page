<template>
  <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
    <div class="my-card" style="width: 49%">
      <div class="card-title">知产课堂</div>
      <el-row :gutter="20">
        <el-col v-for="i in 4" :key="i" :lg="12" :sm="24" :xs="24">
          <div class="d-flex flex-column">
            <video controls height="270">
              <source src="http://www.daweisoft.com//upload/file/gaoxiaodisijiang.mp4" type="video/mp4">
            </video>
            <div style="text-align: center;margin-top: 10px"> 第{{ i }}讲：高校优势学科的专利特色库建设</div>
          </div>
        </el-col>
      </el-row>
      <div class="link-type" style="text-align: center;margin-top: 20px;cursor: pointer">查看更多</div>

    </div>
    <div class="my-card" style="width: 49%">
      <div class="card-title">专利推荐</div>
      <el-row :gutter="20">
        <el-col v-for="i in 4" :key="i" :lg="12" :sm="24" :xs="24">
          <div class="d-flex flex-column">
            <video controls height="270">
              <source src="http://www.daweisoft.com//upload/file/gaoxiaodisijiang.mp4" type="video/mp4">
            </video>
            <div style="text-align: center;margin-top: 10px"> 第{{ i }}讲：高校优势学科的专利特色库建设</div>
          </div>
        </el-col>
      </el-row>
      <div class="link-type" style="text-align: center;margin-top: 20px;cursor: pointer">查看更多</div>

    </div>
  </div>

</template>
<script>
export default {
  name: 'PropertyClassroom',
  data() {
    return {
      query: ''
    }
  },
  created() {

  },
  methods: {}
}
</script>
<style scoped>

</style>
