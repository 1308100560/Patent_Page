<template>
  <div class="my-card" style="padding: 0!important;margin: 0!important;  background: #f2f2f2;">
    <div class="box-container">
      <div
        style="position: relative;width: 100%;height: 100%;overflow: hidden;display: flex;flex-direction: row;align-items: center;justify-content: center"
      >
        <img class="box-bg" src="https://as.zbjimg.com/static/as-ipr-zt/pc/20180326zljy/img/bannar_f533020.png">
        <h1 style="font-size:40px;color: #64d9d6;z-index: 999;position: absolute;top: 60px">专利买卖</h1>
      </div>
      <img class="img2" src="https://as.zbjimg.com/static/as-ipr-zt/pc/20180326zljy/img/sanjiao_2b27f59.png">
      <div class="box">
        <div class="box-header">
          <div class="box-header-item box-header-item-active">
            <div class="box-header-text">
              买专利
            </div>
          </div>
          <div class="box-header-item">
            <div class="box-header-text">
              卖专利
            </div>
          </div>
        </div>
        <div class="box-content">
          <div class="box-input">
            <input placeholder="请输入您想要的专利关键词。如“机械专利，用于生产”">
            <div class="box-button">立即匹配专利</div>
          </div>

        </div>
      </div>
    </div>
    <div>
      <img
        src="http://xsgame2023.oss-cn-hangzhou.aliyuncs.com/ab44ab552918dc1273971c8019e7ed48.png"
        style="width: 80%;margin: 150px 10%"
      >
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      a: 1
    }
  }
}
</script>
<style scoped>
.box-container {
  height: 400px;
  margin-top: 0;
  position: relative;
  background: #f2f2f2;
}

.box-bg {
  position: absolute;
  top: -20px;
  left: -10px;
  width: calc(100% + 30px);
  height: 100%;
}

.box {
  width: calc(100% - 120px - 300px - 50px);
  border: 1px solid #64d9d6;
  border-radius: 10px;
  background: white;
  position: absolute;
  bottom: -80px;
  right: 60px;
}

.img2 {
  width: 300px;
  position: absolute;
  left: 60px;
  bottom: -100px;
}

.box-header {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.box-header-item {
  height: 80px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  font-size: 26px;
  font-weight: bold;
  width: 50%;
  cursor: pointer;
  border: 1px solid rgba(0, 0, 0, 0.1);
  color: #64d9d6;
}

.box-header-item:first-child {
  border-top-left-radius: 10px;
}

.box-header-item:last-child {
  border-top-right-radius: 10px;
}

.box-header-text {
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  width: 40%;

}

.box-header-item-active {
  background: #f3f7ff;
}

.box-header-item-active .box-header-text {
  border-bottom: 3px solid #64d9d6;
}

.box-header-item:first-child {
  border-right: none;
}

.box-content {
  padding: 50px 0;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.box-input {
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.box-input input {
  border: 1px solid #64d9d6;
  padding: 0 30px;
  font-size: 14px;
  height: 60px;
  width: 60%;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
}

.box-input input:focus {
  outline: none;
}

.box-button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 60px;
  background: #64d9d6;
  color: #fff;
  padding: 0 30px;
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
}
</style>
