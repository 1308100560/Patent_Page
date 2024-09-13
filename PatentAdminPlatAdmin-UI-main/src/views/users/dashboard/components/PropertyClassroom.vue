<template>
  <div class="my-card">
    <div class="card-title">知产课堂</div>
    <el-row :gutter="20">
      <el-col v-for="i in 6" :key="i" :lg="8" :sm="24" :xs="24">
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
