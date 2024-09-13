<template>
  <div>
    <el-dialog :modal-append-to-body="false" :visible.sync="dialogTableVisible" title="关注专利">
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
    </el-dialog>
    <el-button
      v-if="patent.isFocused"
      size="mini"
      type="danger"
      @click="focus"
    >取消关注
    </el-button>
    <el-button
      v-else
      size="mini"
      type="light"
      @click="showPopover()"
    >
      关注专利
    </el-button>
  </div>

</template>
<script>
import { focusPatent, unFocusPatent } from '@/api/patent'

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
      desc: '',
      dialogTableVisible: false
    }
  },
  methods: {
    showPopover() {
      this.dialogTableVisible = true
      this.desc = ''
    },
    focus() {
      if (this.patent.isFocused) {
        this.$confirm('确定取消关注吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          unFocusPatent(this.patent.PNM).then(res => {
            if (res.data.code === 200) {
              this.afterConfirm()
              this.dialogTableVisible = false
              this.$message.success('取消关注成功')
              this.patent.isFocused = false
            } else {
              this.$message.error(res.data.msg)
            }
          })
        }).catch(() => {
        })
      } else {
        this.patent.desc = this.desc
        focusPatent(this.patent).then(res => {
          if (res.data.code === 200) {
            this.afterConfirm()
            this.dialogTableVisible = false
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
