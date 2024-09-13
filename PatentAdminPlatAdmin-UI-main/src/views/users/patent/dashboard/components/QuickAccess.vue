<template>
  <div>
    <p>快捷访问</p>
    <el-row :gutter="40" class="panel-group">
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link :to="{path:'/report/generate/novelty',query:{type:'novelty'}}">
          <div class="card-panel">
            <div class="card-panel-icon-wrapper icon-people">
              <svg-icon class-name="card-panel-icon" icon-class="new" />
            </div>
            <div class="card-panel-description">
              <div class="card-panel-text">
                查新报告
              </div>
            </div>
          </div>
        </router-link>
      </el-col>
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <router-link :to="{path:'/report/generate/infringement',query:{type:'infringement'}}">
          <div class="card-panel">
            <div class="card-panel-icon-wrapper icon-people">
              <svg-icon class-name="card-panel-icon" icon-class="infringement" />
            </div>
            <div class="card-panel-description">
              <div class="card-panel-text">
                侵权报告
              </div>
            </div>
          </div>
        </router-link>
      </el-col>
      <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
        <download-able name="交底书模版" url="http://10.112.138.178:8000/static/docs/patent_confession_template.doc">
          <div class="card-panel">
            <div class="card-panel-icon-wrapper icon-message">
              <svg-icon class-name="card-panel-icon" icon-class="pack" />
            </div>
            <div class="card-panel-description">
              <div class="card-panel-text">
                交底书模版下载
              </div>

            </div>
          </div>
        </download-able>

      </el-col>
    </el-row>
  </div>
</template>

<script>

import DownloadAble from '@/views/users/components/DownloadAble'

export default {
  components: { DownloadAble }
}
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-people {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3
      }
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      height: 100%;
      font-weight: bold;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
      }

    }
  }
}

@media (max-width: 550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>
