<template>
  <div @click="download">
    <slot />
  </div>
</template>
<script>
import { saveAs } from 'file-saver'

export default {
  name: 'DownloadAble',
  props: {
    url: {
      type: String,
      default: () => ('')
    },
    name: {
      type: String,
      default: () => ('')
    }
  },
  data() {
    return {}
  },
  methods: {
    download() {
      saveAs(this.url, this.name)
      this.$message({
        message: '下载中',
        type: 'success'
      })
    }
  }
}
</script>
