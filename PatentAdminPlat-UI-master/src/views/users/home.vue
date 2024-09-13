<template>
  <div>
    <div class="header">
      <div style="display: flex;flex-direction: row;justify-content: flex-end;width: 90%">
        <el-button type="text" @click="login">登录</el-button>
        <router-link style="margin-left: 15px" target="_blank" to="/register">
          <el-button type="text">注册</el-button>
        </router-link>
      </div>
      <LoginComponent ref="loginRef" />
      <div class="searchBox">
        <div class="searchHeader">
          <h1 style="color: black;text-shadow: 3px 3px 3px #ccc;">知识产权转化手册</h1>
          <div class="searchBoxInput">
            <input placeholder="可输入申请人、发明(设计）人、技术关键词、分类号、申请号、公开(公告)号等或直接输入检索式">
            <div class="searchBtn">检索</div>
          </div>
          <div style="display: flex;flex-direction: row;justify-content: flex-start;width: 80%">
            <div class="searchAction">
              <router-link to="/search/advanced"><span>高级检索</span></router-link>
              <router-link to="/search/form"><span>表单检索</span></router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="recommend" style="margin: 40px 100px">
      <span class="result-item title">今日推荐</span>
      <div v-for="index in 8" :key="index" class="result-item">
        <div class="result-title row-center">
          <span class="text-primary">基于条件标准化流模型的强鲁棒性图像隐写方法及系统[ZH]</span>
          <el-tag effect="dark" size="mini" style="margin: 0 5px" type="warning">审中</el-tag>
          中国发明申请
        </div>
        <div class="desc">
          申请号：CN202210754766.3 申请日：2022.06.30 公开(公告）号：CN114827381A 公开（公告）日：2022.07.29
        </div>
        <div class="desc">
          申请（专利权）人：广东]北京大学深圳研究生院发明(设计)人：张健；许佑民
        </div>
        <div class="desc">
          本发明涉及基于条件标准化流模型的强鲁棒性图像隐写方法，包括步骤：S1.构建流模型单元：对输入的宿主图像和隐藏图像进行分布变换，转化为高频冗余信息和容器图像；S2.构造分布映射模块：建模高频冗余信息在容器图像条件依赖下的分布，将其映射到标准正态分布以隐式保存有价值的隐藏信息；S3.构造失真模拟单元：在模型训练过程中仿真各类失真干扰的影响；S4.构造图像增强模块：对受干扰的容器图像进行初步增强去噪处理；以及S5.构造条件调制模块：以干扰强度和类型为条件，调制流模型单元的网络参数。本发明方法解决当前图像隐写术在鲁棒性和还原质量、隐写容量上的缺陷和不足，解决之前基于学习的隐写术在受到失真干扰时性能大幅下降的问题。
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import LoginComponent from '@/views/users/components/LoginComponent'

export default {
  name: 'Users',
  components: {
    LoginComponent
  },
  data() {
    return {
      showLogin: false
    }
  },
  created() {

  },
  methods: {
    login() {
      this.$refs.loginRef.show()
    }
  }
}
</script>
<style>
.header {
  padding: 10px 20px;
  border-bottom: 1px solid #ebeef5;
  height: 350px;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-image: linear-gradient(to top, #ace0f9 0%, #ffffff 100%);
}

.searchBox {
  width: 70%;
  padding-top: 70px;
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
  text-shadow: 3px 3px 3px #ccc;
  justify-content: space-around;
  margin-top: 7px;
  cursor: pointer;

}

.searchAction span {
  margin-right: 10px;
}

.recommend {
  margin: 20px 10px 0 10px;
}

.result-title {
  font-size: 0.9rem;
}

.text-primary {
  color: #00a9ed;
}

.desc {
  font-size: 0.75rem;
  margin: 5px 0;
  line-height: 1.3rem;
}

.result-item {
  margin-top: 10px;
  padding: 20px 10px;
  cursor: pointer;
  border-radius: 10px;
  margin-bottom: 10px;
  transition: 0.3s;
}

.result-item:hover {
  background-color: #e3f8fd;
}

.recommend .title {
  font-size: 1.2rem;
  font-weight: bold;
  color: #17233d;
  text-shadow: 3px 3px 3px #ccc;
}

</style>
