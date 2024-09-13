<template>
  <div>

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
      <el-button :loading="uploading" icon="el-icon-upload" size="small" type="primary" @click="handleUploadFile">
        {{ uploading ? '上传中' : '上传附件' }}
        <input v-show="false" ref="uploadInput" type="file">
      </el-button>
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
            <div style="font-size: 13px;color: white;margin: 10px;text-align: center">{{
              file.name || '未命名'
            }}
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
          <img v-if="isImage(file.FilePath)" :src="`${file.FilePath}`" alt="" class="image">
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
import { updateTicket } from '@/api/ticket'
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
      uploading: false,
      files: []
    }
  },
  watch: {
    ticket: {
      handler() {
        this.files = JSON.parse(this.ticket.files)
        console.log(this.files)
      },
      immediate: true
    }
  },
  mounted() {
    console.log(this.ticket)
  },
  methods: {
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
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
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
