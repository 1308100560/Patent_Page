<template>
  <div class="container">
    <div class="my-card" style="width: 100%">
      <create-pack ref="createPack" />
      <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
        <div>
          <el-input
            class="filter-item"
            placeholder="工艺包名称名称"
            size="small"
            style="width: 180px;margin-right: 10px"
          />

          <el-button class="filter-item" disabled icon="el-icon-search" size="small" type="primary">
            搜索
          </el-button>
          <el-button icon="el-icon-folder-add" size="small" type="primary" @click="showCreatePack">
            创建工艺包
          </el-button>
        </div>
        <el-button icon="el-icon-refresh" size="small" @click="getPackageList" />
      </div>
    </div>
    <div v-if="packageList.length>0" class="my-card" style="margin-top: 10px">
      <div class="cards">

        <el-card
          v-for="o in packageList"
          :key="o.packageId"
          :style="{ padding: '0',margin:'20px 10px 10px 0' }"
          class="my-card-2"
        >
          <svg aria-hidden="true" class="image">
            <use xlink:href="#icon-folder" />
          </svg>
          <div>
            <div style="font-size:1rem;margin-bottom: 8px">
              {{ o.packageName }}
            </div>
            <el-row justify="space-between" type="flex">
              <router-link :to="{path:`package/${o.packageId}`,params:'1212'}" style="margin-right: 5px">
                <el-button icon="el-icon-view" size="mini" type="light">查看</el-button>
              </router-link>
              <!--            <el-button icon="el-icon-download" size="mini" type="primary">打包下载</el-button>-->
            </el-row>
          </div>
        </el-card>
      </div>
    </div>
    <div
      v-if="packageList.length===0"
      style="text-align: center;width: 100%;margin: 20px 0;"
    >
      <el-empty description="没有工艺包哦" />
    </div>
  </div>
</template>
<script>
import createPack from '@/views/users/components/createPack'
import { mapGetters } from 'vuex'

export default {
  name: 'TechPack',
  components: {
    createPack
  },
  computed: {
    ...mapGetters([
      'packageList'
    ])
  },
  mounted() {
    this.getPackageList()
  },
  methods: {
    getPackageList() {
      this.$store.dispatch('package/getPackageList', this.form).then(() => {
        console.log('yes')
      })
    },
    showCreatePack() {
      this.$refs.createPack.show()
    }
  }
}
</script>
<style scoped>

.cards {
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: wrap;
}

.my-card-2 {
  width: 200px;
  height: 250px;
}

.image {
  width: 100%;
  display: block;
}
</style>
