<template>
  <div>
    <el-dialog :visible.sync="finishDialog" title="完结工单">
      <div style="display: flex;flex-direction: column;align-items: center;justify-content: center">
        <div
          style="cursor:pointer;height: 100px;width: 100px;border: 1px dashed lightgray;border-radius: 10px;line-height: 100px;text-align: center"
          @click="uploadReport"
        >
          点击上传
          <input v-show="false" ref="reportUpload" type="file">
        </div>
        <div style="margin: 15px 0">
          {{ currentFile ? currentFile.name : '' }}
        </div>
        <div v-if="reportType==='估值报告'" style="margin-bottom: 20px">
          <el-input
            v-model="price"
            controls-position="right"
            placeholder="请输入报告估值"
          />
        </div>
        <el-button type="primary" @click="handleFinish">立即完结</el-button>
      </div>
    </el-dialog>
    <div
      style="max-height: 500px;overflow-y: auto;margin: 20px 0;border:1px solid lightgray;padding: 10px;border-radius: 0.15rem"
    >
      <el-steps :active="ticket.optLogs.length" align-center direction="vertical">
        <el-step
          v-for="(opt,i) in JSON.parse(ticket.optLogs)"
          :key="'opt'+i"
          :description="'操作时间：'+new Date(opt.timestamp).toLocaleString() "
          :title="opt.operation"
        />

      </el-steps>
    </div>
    <div>
      <div>
        <el-input
          v-model="optMsg"
          :rows="2"
          placeholder="请输入内容"
          type="textarea"
        />
      </div>
      <div style="margin-top: 15px">
        <el-button @click="levelMsg">留言</el-button>
        <el-button :loading="uploading" size="small" type="primary" @click="handleUploadFile">
          {{ uploading ? '上传中' : '上传附件' }}
          <input v-show="false" ref="uploadInput" type="file">
        </el-button>
        <el-button
          :disabled="isClosed(ticket)"
          :loading="uploading"
          size="small"
          type="success"
          @click="finishDialog=true"
        >
          {{ isClosed(ticket) ? '已完结' : '完结工单' }}
        </el-button>
      </div>
    </div>
    <div class="cards">
      <el-card
        v-for="(file,index) in files"
        :key="`file-${index}`"
        :body-style="{padding: 0}"
        :style="{ padding: '0px',margin:'10px 10px 10px 0'}"
        class="my-card2"
        shadow="hover"
      >
        <div class="card-actions">
          <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 100%">
            <div
              style="font-size: 13px;color: white;margin: 10px;
            text-align: center;flex-direction: row;display: flex;justify-content: center;flex-wrap: wrap"
            >
              {{ file.name || '未命名' }}
            </div>
            <div style="display: flex;flex-direction: row">
              <el-button
                size="mini"
                style="margin-right: 10px"
                type="danger"
                @click="handleDeleteFile(file)"
              >
                删除
              </el-button>
              <download-able :name="file.name" :url="`${file.full_path}`">
                <el-button size="mini" type="primary">下载</el-button>
              </download-able>
            </div>
          </div>
        </div>

        <div class="imageField">
          <img v-if="isImage(file.name)" :src="`${file.full_path}`" alt="" class="image">
          <svg v-else aria-hidden="true" class="image">
            <use :xlink:href="`#icon-files2`" />
          </svg>
        </div>

        <div style="font-size: 0.8rem;text-align: center;margin: 0 5px 10px 5px">
          {{ (file.name || '未命名') | textCut(15) }}
        </div>
      </el-card>
    </div>

  </div>
</template>
<script>

import { uploadFile } from '@/api/upload'
import { finishTicket, getTicket, updateTicket } from '@/api/ticket'
import DownloadAble from '@/views/users/components/DownloadAble.vue'

export default {
  components: {
    DownloadAble
  },
  props: {
    switchPage: {
      type: Function,
      default: () => {
      }
    },

    ticket: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      price: '',
      uploading: false,
      files: [],
      optMsg: '',
      finishDialog: false,
      currentFile: '',
      reportType: ''
    }
  },
  watch: {
    ticket: {
      handler() {
        this.reportType = this.ticket.relObj.reportType
        console.log(this.reportType)
        this.files = JSON.parse(this.ticket.files)
      },
      immediate: true
    }
  },
  mounted() {
    console.log(this.ticket)
  },
  methods: {
    isClosed(ticket) {
      return ticket.status === 'closed' || ticket.status === 'finished'
    },
    handleDeleteFile(file) {
      updateTicket(this.ticket.id, {
        filesOpt: 'del',
        files: [file]
      }).then(res => {
        this.files = this.files.filter(f => f.full_path !== file.full_path)
        this.ticket.files = JSON.stringify(this.files)
        this.$message.success('删除成功')
      })
    },
    isImage(filePath) {
      console.log(filePath)
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
    },
    levelMsg() {
      updateTicket(this.ticket.id, {
        optMsg: this.optMsg
      }).then(res => {
        this.$message.success('留言成功')
        getTicket(this.ticket.id).then(res => {
          this.ticket = res.data.data
          this.optMsg = ''
        })
      })
    },
    uploadReport() {
      this.$refs.reportUpload.click()
      this.$refs.reportUpload.onchange = e => {
        this.currentFile = e.target.files[0]
      }
    },
    handleFinish() {
      if (!this.currentFile) {
        this.$message.error('请上传报告')
        return
      }
      if (this.reportType === '估值报告' && !this.price) {
        this.$message.error('请填写估值价格')
        return
      }
      const formData = new FormData()
      formData.append('file', this.currentFile)
      this.uploading = true
      uploadFile(formData).then(fileResp => {
        const file = fileResp.data.data
        const data = {
          relObj: {
            filesOpt: 'add',
            files: [file]
          }
        }
        if (this.reportType === '估值报告') {
          data.relObj.reportProperties = {
            evalPrice: this.price
          }
        }
        updateTicket(this.ticket.id, data).then(res => {
          finishTicket(this.ticket.id).then(res => {
            console.log(res)
            this.uploading = false
            this.$message.success('完结成功')
            this.finishDialog = false
          })
        })
      })
    },
    handleUploadFile() {
      this.$refs.uploadInput.click()
      this.$refs.uploadInput.onchange = e => {
        const formData = new FormData()
        formData.append('file', e.target.files[0])
        this.uploading = true
        uploadFile(formData).then(fileResp => {
          const file = fileResp.data.data
          updateTicket(this.ticket.id, {
            filesOpt: 'add',
            files: [file]
          }).then(res => {
            console.log(res)
            this.uploading = false
            this.files.push(fileResp.data.data)
            this.ticket.files = JSON.stringify(this.files)
            this.$message.success('上传成功')
          })
        })
      }
    }
  }
}
</script>
<style scoped>
.cards {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin: 0 auto;
  flex-wrap: wrap;
}

.card-actions {
  display: none;
  position: absolute;
  height: 100%;
  width: 100%;
  padding: 0 5%;
  top: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
  cursor: pointer;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.my-card2:hover .card-actions {
  display: flex;
}

.image {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.my-card2 {
  position: relative;
}

.my-card2 {
  width: 130px;
  min-height: 150px;
}

.imageField {
  height: 120px;
  padding: 10px;
}
</style>
