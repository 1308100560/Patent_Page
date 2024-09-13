<template>
  <div>

    <el-button
      v-if="patent.isClaimed"
      icon="el-icon-check"
      size="mini"
      type="success"
      @click="claim"
    >认领
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
          <el-button size="mini" type="primary" @click="claim">确认</el-button>

        </div>
      </el-form>
      <el-button
        slot="reference"
        icon="el-icon-check"
        size="mini"
        type="light"
        @click="showPopover()"
      >
        认领
      </el-button>
    </el-popover>

  </div>

</template>
<script>
import { claimPatent, unClaimPatent } from '@/api/patent'

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
    claim() {
      if (this.patent.isClaimed) {
        unClaimPatent(this.patent.PNM).then(res => {
          if (res.data.code === 200) {
            this.$message.success('取消认领成功')
            this.patent.isClaimed = false
          } else {
            this.$message.error(res.data.msg)
          }
        })
      } else {
        this.patent.desc = this.desc
        claimPatent(this.patent).then(res => {
          if (res.data.code === 200) {
            this.$message.success('认领成功')
            this.patent.isClaimed = true
          } else {
            this.$message.error(res.data.msg)
          }
        })
      }
    }
  }
}
</script>
