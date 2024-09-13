<template>
  <div>

    <el-button
      v-if="patent.isClaimed"
      size="mini"
      type="danger"
      @click="claim"
    >取消认领
    </el-button>
    <el-button

      v-else
      size="mini"
      type="light"
      @click="showPopover()"
    >
      认领专利
    </el-button>
    <el-dialog :modal-append-to-body="false" :visible.sync="dialogTableVisible" title="认领专利">
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
    </el-dialog>

  </div>

</template>
<script>
import { claimPatent, unClaimPatent } from '@/api/patent'

export default {
  props: {
    afterConfirm: {
      type: Function,
      default: () => {
      }
    },
    patent: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      dialogTableVisible: false,
      desc: ''
    }
  },
  methods: {
    showPopover() {
      this.dialogTableVisible = true
      this.desc = ''
    },
    claim() {
      if (this.patent.isClaimed) {
        this.$confirm('确认取消认领该专利吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          unClaimPatent(this.patent.PNM).then(res => {
            if (res.data.code === 200) {
              this.$message.success('取消认领成功')
              this.patent.isClaimed = false
              this.afterConfirm()
            } else {
              this.$message.error(res.data.msg)
            }
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消取消认领'
          })
        })
      } else {
        this.patent.desc = this.desc
        claimPatent(this.patent).then(res => {
          if (res.data.code === 200) {
            this.$message.success('认领成功')
            this.dialogTableVisible = false
            this.patent.isClaimed = true
            this.afterConfirm()
          } else {
            this.$message.error(res.data.msg)
          }
        })
      }
    }
  }
}
</script>
