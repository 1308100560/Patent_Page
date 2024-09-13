<template>
  <div>
    <el-tabs :tab-position="tabPosition" class="tabs" style="height: 200px;" type="border-card">
      <el-tab-pane>
        <span slot="label">标签1<i
          class="el-icon-edit"
          style="margin-left: 20px"
        /> </span>
        标签1
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">标签2<i
          class="el-icon-edit"
          style="margin-left: 20px"
        /> </span>
        标签1
        标签2
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">标签3<i
          class="el-icon-edit"
          style="margin-left: 20px"
        /> </span>
        标签3
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">
          <el-button size="mini" style="width: 100%">添加</el-button>
        </span>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>

export default {
  data() {
    return {
      tabPosition: 'left'
    }
  },

  methods: {}
}
</script>
<style scoped>
.tabs {
  margin: 10px;
}
</style>
