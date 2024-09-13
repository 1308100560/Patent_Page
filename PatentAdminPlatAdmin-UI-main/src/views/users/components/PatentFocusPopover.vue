<template>
  <div>

    <el-button
      v-if="patent.isFocused"
      icon="el-icon-star-off"
      size="mini"

      type="danger"
      @click="focus"
    >关注
    </el-button>
    <el-popover
      v-else
      placement="left-end"
      width="350"
    >
      <el-form
        label-position="left"
        label-width="60px"
        size="small"
        style="margin: 10px"
      >
        <el-form-item label="专利" size="small">
          <el-input v-model="patent.TI" size="small" />
        </el-form-item>
        <el-form-item label="备注" size="small">
          <el-input v-model="desc" size="small" type="textarea" />
        </el-form-item>
        <div style="display: flex;flex-direction: row;align-items: center;justify-content: flex-end">
          <el-button size="mini" type="primary" @click="focus">确认</el-button>

        </div>
      </el-form>
      <el-button
        slot="reference"
        icon="el-icon-star-off"
        size="mini"
        type="light"
        @click="showPopover()"
      >
        关注
      </el-button>
    </el-popover>

  </div>

</template>
<script>
import { focusPatent, unFocusPatent } from '@/api/patent'

export default {
  props: {
    patent: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      desc: ''
    }
  },
  methods: {
    showPopover() {
      this.desc = ''
    },
    focus() {
      if (this.patent.isFocused) {
        unFocusPatent(this.patent.PNM).then(res => {
          if (res.data.code === 200) {
            this.$message.success('取消关注成功')
            this.patent.isFocused = false
          } else {
            this.$message.error(res.data.msg)
          }
        })
      } else {
        this.patent.desc = this.desc
        focusPatent(this.patent).then(res => {
          if (res.data.code === 200) {
            this.$message.success('关注成功')
            this.patent.isFocused = true
          } else {
            this.$message.error(res.data.msg)
          }
        })
      }
    }
  }
}
</script>
