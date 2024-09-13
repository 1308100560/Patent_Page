<template>
  <el-row :gutter="20">

    <el-col :lg="5" :sm="24" :xs="24">
      <div style="display: flex;flex-direction: column;justify-content: space-between;height: 340px">
        <router-link :to="{path:'/report/generate/novelty'}">
          <div class="my-card v-card">
            <div class="card-title">
              竞争对手/行业人才
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>
        <router-link :to="{path:'/report/generate/infringement'}">
          <div class="my-card v-card">
            <div class="card-title">
              技术前延
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>
        <router-link :to="{path:'/report/generate/valuation'}">
          <div class="my-card v-card">
            <div class="card-title">
              专利监测
            </div>
            <div style="font-size: 14px">
              快捷跳转
            </div>
          </div>
        </router-link>

      </div>
    </el-col>
    <el-col :lg="9" :sm="24" :xs="24">
      <div class="my-card" style="height: 340px">
        <div class="card-title">
          相关专利图谱
        </div>
        <div class="list">
          <div v-for="(field,index) in fieldRanking" :key="field.name" class="list-item">
            <div style="width: 40px">
              <div :class="{'indicator-primary':index<=2}" class="indicator">{{ index + 1 }}</div>
            </div>
            <div style="width: calc(100% - 40px);display: flex;flex-direction: row;align-items: center">
              <div style="width: 20%">
                {{ field.name }}
              </div>
              <div style="width: 75%">
                <el-progress :percentage="field.value" :stroke-width="20" :text-inside="true" />
              </div>
            </div>
          </div>

        </div>
      </div>
    </el-col>
    <el-col :lg="10" :sm="24" :xs="24">
      <div class="my-card" style="height: 340px">
        <div class="card-title">
          相关人员图谱
        </div>
        <div class="list">
          <div v-for="(field,index) in inventorRanking" :key="field.name" class="list-item">
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

  </el-row>
</template>
<script>

export default {
  name: 'ListData',
  data() {
    return {
      fieldRanking: [
        { name: '计算机', value: 70 },
        { name: '神经网络', value: 20 },
        { name: '网络安全', value: 10 },
        { name: '信息隐藏', value: 7 },
        { name: 'APT攻击', value: 5 }
      ],
      inventorRanking: [
        { name: '张三', value: 70 },
        { name: '李四', value: 20 },
        { name: '王五', value: 10 },
        { name: '赵六', value: 7 },
        { name: '田七', value: 5 }
      ]
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
  height: 105px;
  cursor: pointer
}

</style>
