<template>
  <div class="container">
    <div class="my-card">
      <el-page-header content="搜索列表" style="margin-bottom: 15px" @back="$router.go(-1)" />
      <div>
        <el-input v-model="searchForm.query" class="input-with-select" placeholder="请输入内容">
          <el-button slot="append" icon="el-icon-search" plain type="primary" @click="doSearch(false)" />
        </el-input>
        <div class="advancedFilter">
          <div>
            <router-link to="/search/advanced"><span style="margin-right: 10px">高级搜索</span></router-link>
            <router-link to="/search/form"><span style="margin-right: 10px">表单搜索</span></router-link>
          </div>
        </div>
      </div>
      <el-tabs v-model="currentTab" style="margin-top: 10px" type="border-card">
        <el-tab-pane :key="'searchList'" name="searchList">
          <span slot="label">搜索列表</span>
          <keep-alive>
            <search-list ref="search" :query="searchForm" />
          </keep-alive>
        </el-tab-pane>
        <el-tab-pane :key="'ana'" label="搜索分析" name="ana">
          <keep-alive>
            <table-analysis ref="chart" :query="searchForm.query" />
          </keep-alive>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>

import tableAnalysis from '@/views/users/search/components/tableAnalysis'
import SearchList from '@/views/users/components/SearchList'

export default {
  components: { tableAnalysis, SearchList },
  data() {
    return {
      currentTab: 'searchList',
      searchForm: {
        query: ''
      },
      searchLoading: false
    }
  },
  beforeRouteEnter(to, from, next) {
    to.params.showCacheData = from && from.name === 'SearchDetail'
    next()
  },
  mounted() {
    const { q, action } = this.$route.query
    console.log(this.$route.params.showCacheData)
    if (q) {
      this.searchForm.query = q
      if (action === 'graph') {
        this.currentTab = 'ana'
      } else {
        this.doSearch(this.$route.params.showCacheData)
      }
    }
  },
  methods: {
    doSearch(showCache) {
      this.$refs.search.search(showCache)
    }
  }
}
</script>
<style scoped>
.container {
  padding: 0 15px;
}

.advancedFilter {
  font-size: 13px;
  margin: 5px 0;
  color: gray;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

</style>
