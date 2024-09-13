<template>
  <div class="home">
    <div ref="file" />
  </div>
</template>
<script>
import axios from 'axios'

const docx = require('docx-preview')

// 如果 docx-preview 版本比较新需要 es6 导出
// import { renderAsync } from 'docx-preview'

window.JSZip = require('jszip')
export default {
  mounted() {
    axios({
      method: 'get',
      responseType: 'blob', // 设置响应文件格式
      url: 'http://www.sriptb.com/static/docs/novelty_report_template.docx'
    }).then(({ data }) => {
      docx.renderAsync(data, this.$refs.file) // 渲染到页面预览
    })
  }
}
</script>
