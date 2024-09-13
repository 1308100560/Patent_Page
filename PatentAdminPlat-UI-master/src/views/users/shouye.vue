<template>
  <div class="container">
    <div class="content-wrapper">
      <h1 class="title">专利智能分析系统</h1>
      <div class="main-content">
        <div class="button-section">
          <div class="button-container" @click="goToReportPage">
            <el-button type="primary" class="custom-button report-button">
              <i class="el-icon-data-analysis"></i>
              新创性分析报告
            </el-button>
          </div>
          <div class="button-container" @click="goToDisclosurePage">
            <el-button type="primary" class="custom-button disclosure-button">
              <i class="el-icon-document"></i>
              专利技术交底书
            </el-button>
          </div>
        </div>
        <div class="explanation-section">
          <div class="explanation-card">
            <div class="card-icon">
              <i class="el-icon-data-analysis"></i>
            </div>
            <div class="card-content">
              <h3>新创性分析报告</h3>
              <p>利用AI算法，全面评估发明创新，精确判定新颖性和创造性。</p>
            </div>
          </div>
          <div class="explanation-card">
            <div class="card-icon">
              <i class="el-icon-document"></i>
            </div>
            <div class="card-content">
              <h3>专利技术交底书</h3>
              <p>智能辅助写作，快速生成专业详实的技术交底书。</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <footer class="footer">
      <div class="footer-content">
        <div class="copyright">© 2024 版权所有：北京卫平科技有限公司</div>
        <div class="footer-links">
          <a href="#">备案号</a>
          <a href="#">个人资料</a>
          <a href="#">帮助与反馈</a>
          <a href="#">关于思蒙</a>
          <a href="#">法律声明</a>
        </div>
        <div class="browser-recommendation">推荐使用Chrome浏览器</div>
      </div>
    </footer>
    <div class="background-decoration"></div>
  </div>
</template>

<script>
export default {
  methods: {
    goToReportPage() {
      this.$router.push('./New_innovation_analysis_report');
    },
    goToDisclosurePage() {
      this.$router.push('./IPC');
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 90vh;
  background-color: #ffffff;
  position: relative;
  overflow: hidden;
}

.content-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding: 10px;
  box-sizing: border-box;
}

.title {
  font-size: 50px;
  color: #1e88e5;
  text-align: center;
  margin-bottom: 15px;
  font-weight: bold;
}

.main-content {
  display: flex;
  justify-content: space-between;
  align-items: stretch;
  background-color: #f5f5f5;
  border-radius: 15px;
  overflow: hidden;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 1000px;
  padding: 15px;
}

.button-section, .explanation-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.button-section {
  padding-right: 15px;
}

.explanation-section {
  padding-left: 15px;
}

.button-container {
  margin-bottom: 10px;
}

.custom-button {
  width: 100%;
  height: 80px;
  font-size: 18px;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  transition: all 0.3s ease;
  border: none;
}

.custom-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 3px 10px rgba(30, 136, 229, 0.2);
}

.report-button {
  background: linear-gradient(135deg, #42a5f5 0%, #1976d2 100%);
  color: #ffffff;
}

.disclosure-button {
  background: linear-gradient(135deg, #66bb6a 0%, #43a047 100%);
  color: #ffffff;
}

.custom-button i {
  font-size: 24px;
  margin-bottom: 4px;
}

.explanation-card {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 10px;
  margin-bottom: 10px;
  display: flex;
  align-items: center;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.explanation-card:hover {
  transform: translateX(2px);
  box-shadow: 0 3px 12px rgba(0, 0, 0, 0.12);
}

.card-icon {
  font-size: 24px;
  color: #1e88e5;
  margin-right: 10px;
}

.card-content h3 {
  color: #1e88e5;
  margin-bottom: 3px;
  font-size: 16px;
}

.card-content p {
  color: #333333;
  line-height: 1.3;
  font-size: 13px;
}

.footer {
  width: 100%;
  background-color: #f5f5f5;
  padding: 8px 0;
  text-align: center;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 15px;
}

.copyright, .browser-recommendation {
  font-size: 11px;
  color: #666;
}

.footer-links {
  display: flex;
  gap: 12px;
}

.footer-links a {
  font-size: 11px;
  color: #1e88e5;
  text-decoration: none;
}

.footer-links a:hover {
  text-decoration: underline;
}

.background-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 10% 20%, rgba(30, 136, 229, 0.05) 0%, transparent 20%),
  radial-gradient(circle at 90% 80%, rgba(102, 187, 106, 0.05) 0%, transparent 20%);
  z-index: -1;
}

@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
  }

  .button-section, .explanation-section {
    width: 100%;
    padding: 0;
  }

  .footer-content {
    flex-direction: column;
    gap: 5px;
  }
}
</style>
