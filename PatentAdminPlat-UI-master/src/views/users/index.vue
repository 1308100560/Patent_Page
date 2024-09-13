<template>
  <div>
    <div class="header">
      <div class="searchBox">
        <div class="searchHeader">
          <h1 style="color: black;">专利检索</h1>
          <div class="searchBoxInput">
            <input
              v-model="query"
              placeholder="可输入申请人、发明(设计）人、技术关键词、分类号、申请号、公开(公告)号等或直接输入检索式"
            >
            <div class="searchBtn" @click="search">检索</div>
          </div>
          <div style="display: flex;flex-direction: row;justify-content: flex-start;width: 80%">
            <div class="searchAction">
              <router-link to="/search/advanced"><span>高级检索</span></router-link>
              <router-link to="/search/form"><span>表格检索</span></router-link>
            </div>
          </div>
        </div>
      </div>
      <div
        v-if="showCourse"
        class="tipBox"
      >
        <h2>
          <i class="el-icon-warning-outline" />
          使用教程
        </h2>
        <div style="line-height: 50px;">1. 输入关键词搜索数据库中专利信息；<br>
          2. 点击搜索出专利列表中的认领按钮，将所选专利认领到
          <strong>我的专利</strong>
          -
          <router-link class="link-type" to="/patent/claim"><strong>专利管理</strong></router-link>
          模块；<br>
          3. 点击 <strong>我的专利</strong>-
          <router-link class="link-type" to="/patent/claim"><strong>专利管理</strong></router-link>
          模块中<strong>加入专利组合按钮</strong>将特定专利加入不同的专利组合中；
        </div>

      </div>
    </div>

  </div>
</template>
<script>
export default {
  name: 'Users',
  data() {
    return {
      query: '',
      showCourse: false
    }
  },
  mounted() {
    const h = this.$route.query.h
    if (h) {
      this.showCourse = true
    }
  },
  methods: {
    search() {
      const query = this.query
      if (query) {
        this.$router.push({ path: '/search/results', query: { q: query }})
      }
    }
  }
}
</script>
<style scoped>
.header {
  padding: 10px 20px;
  border-bottom: 1px solid #ebeef5;
  height: calc(100vh - 50px);
  display: flex;
  flex-direction: column;
  align-items: center;
  background-image: linear-gradient(to top, #ace0f9 0%, #ffffff 100%);
  /*  width: 100%;*/
  /*background: url('https://fanyi.youdao.com/img/back.bb0f134c.png') no-repeat;*/
  /*background-size: 100%, 100%;*/
}

.searchBox {
  padding-top: 100px;
  width: 85%;
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  background: transparent;
}

.searchBoxInput {
  width: 80%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  position: relative;
}

.searchBoxInput input {
  width: 100%;
  height: 50px;
  position: relative;
  box-sizing: border-box;
  border-radius: 4px;
  padding-right: 180px;
  padding-left: 22px;
  font-size: 14px;
  background-color: rgba(255, 255, 255, 0.85);
  color: rgba(0, 0, 0, 0.65);
  outline: none;
  border: 0;
  line-height: 50px;
}

.searchBoxInput input:focus {
  outline: none;
}

.searchBtn {
  position: absolute;
  display: block;
  width: 108px;
  border-radius: 4px;
  box-sizing: border-box;
  height: 42px;
  line-height: 42px;
  right: 5px;
  text-align: center;
  color: #fff;
  background-color: #17A1FF;
  cursor: pointer;
  transition: 0.2s;
}

.searchBtn:hover {
  opacity: 0.7;
}

.searchHeader {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.searchAction {
  display: flex;
  flex-direction: row;
  color: #17233d;
  font-weight: bold;
  font-size: 13px;
  justify-content: space-around;
  margin-top: 7px;
  cursor: pointer;

}

.searchAction span {
  margin-right: 10px;
}

.searchAction span:hover {
  color: #17A1FF;
}

.tipBox {
  display: flex;
  flex-direction: column;
  width: 68%;
  margin-top: 40px;
  backdrop-filter: blur(30px);
  padding: 15px;
  font-weight: 400;
  color: #050f2e;
  word-break: break-word;
  border-radius: 8px;
  -webkit-backdrop-filter: blur(30px);
  background: rgba(255, 255, 255, .5);
}
</style>
