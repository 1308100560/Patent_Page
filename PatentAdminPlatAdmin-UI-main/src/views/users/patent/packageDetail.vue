<template>
  <div class="container ">
    <div class="my-card">
      <el-dialog :visible.sync="packageDetailDialogVisible" title="编辑工艺包">
        <el-form
          ref="form"
          :model="packageDetail"
          label-width="80px"
          size="small"
          style="margin: 10px"
        >
          <el-form-item label="工艺包" prop="packageName" size="small">
            <el-input v-model="packageDetail.packageName" size="small" />
          </el-form-item>
          <el-form-item label="描述" size="small">

            <el-input
              v-model="packageDetail.desc"
              :rows="2"
              placeholder="请输入内容"
              type="textarea"
            />

          </el-form-item>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="packageDetailDialogVisible=false">取消</el-button>
            <el-button size="mini" type="primary" @click="updatePackageSubmit">更新</el-button>
          </div>
        </el-form>
      </el-dialog>
      <el-page-header :content="packageDetail.packageName" @back="$router.go(-1)" />
      <div style="margin-top: 10px">
        <div class="text">
          {{ packageDetail.desc }}
        </div>
      </div>
      <div
        style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;margin-top: 10px"
      >
        <div style="display: flex;flex-direction: row;align-items: center">
          <el-button :loading="uploading" icon="el-icon-upload" size="small" type="primary" @click="handleUploadFile">
            {{ uploading ? '上传中' : '上传文件' }}
            <input v-show="false" ref="uploadInput" type="file">
          </el-button>
          <!--        <el-button icon="el-icon-download" size="small" type="primary">打包下载</el-button>-->
          <el-button icon="el-icon-edit" size="small" type="primary" @click="packageDetailDialogVisible=true">编辑
          </el-button>
          <el-button icon="el-icon-delete" size="small" type="danger" @click="handleDeletePackage()">删除专利包
          </el-button>
        </div>
        <div>
          <el-input
            class="filter-item"
            placeholder="搜索"
            size="small"
            style="width: 200px;margin-right: 10px"
          />

          <el-button class="filter-item" disabled icon="el-icon-search" size="small" type="primary">
            搜索
          </el-button>
        </div>
      </div>
    </div>
    <div v-if="files.length>0 || patentList.list.length>0" class="my-card" style="margin-top: 10px">
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
                file.FileName || '未命名'
              }}
              </div>
              <div style="display: flex;flex-direction: row">
                <el-button
                  size="mini"
                  style="margin-right: 10px"
                  type="danger"
                  @click="handleDeleteFile(file.FilePath)"
                >
                  删除
                </el-button>
                <download-able :name="file.FileName" :url="`http://${file.FilePath}`">
                  <el-button size="mini" type="primary">下载</el-button>
                </download-able>
              </div>
            </div>
          </div>

          <div class="imageField">
            <img v-if="isImage(file.FilePath)" :src="`http://${file.FilePath}`" alt="" class="image">
            <svg v-else aria-hidden="true" class="image">
              <use :xlink:href="`#icon-files2`" />
            </svg>
          </div>

          <div style="font-size: 0.8rem;text-align: center;margin: 0 5px">
            {{ (file.FileName || '未命名') | textCut(20) }}
          </div>
        </el-card>
      </div>
      <div class="cards">

        <el-card
          v-for="(patent) in patentList.list"
          :key="patent.patentId"
          :body-style="{padding: 0}"
          :style="{ padding: '0px',margin:'10px 10px 10px 0'}"
          class="my-card2"
          shadow="hover"
        >
          <div class="card-actions">
            <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 100%">
              <div style="font-size: 13px;color: white;margin: 10px;text-align: center" />
              <div style="display: flex;flex-direction: row">
                <el-button size="mini" type="danger" @click="handleDeletePatentFromPackage(patent)">
                  删除
                </el-button>
              </div>
            </div>
          </div>
          <router-link :to="`/search/detail/${ JSON.parse(patent.patentProperties).PNM}`">
            <svg aria-hidden="true" class="image">
              <use xlink:href="#icon-patentF" />
            </svg>
          </router-link>

          <div style="font-size: 0.8rem;text-align: center;margin: 0 5px">
            {{ JSON.parse(patent.patentProperties).TI | textCut(20) }}<br>
          </div>
        </el-card>
      </div>
    </div>
    <div
      v-else
      style="text-align: center;width: 100%;margin: 20px 0;"
    >
      <el-empty description="还没有专利哦" />
    </div>
  </div>

</template>

<script>
import {
  deletePackage,
  getPackage,
  getPatentListByPackageId,
  removePatentFromPackage,
  updatePackage
} from '@/api/package'
import { uploadFile } from '@/api/upload'
import DownloadAble from '@/views/users/components/DownloadAble'
import router from '@/router'

export default {
  name: 'PackageDetail',
  components: { DownloadAble },
  data() {
    return {
      uploading: false,
      packageDetail: { packageName: '' },
      patentList: [],
      files: [],
      packageDetailDialogVisible: false
    }
  },

  mounted() {
    this.loadPackageDetail()
  },
  methods: {
    router() {
      return router
    },
    updatePackageSubmit() {
      updatePackage(this.packageDetail.packageId, {
        packageName: this.packageDetail.packageName,
        desc: this.packageDetail.desc
      }).then(res => {
        this.packageDetailDialogVisible = false
        this.$message.success('更新成功')
      })
    },
    loadPackageDetail() {
      const packageId = this.$route.params.id
      getPackage(packageId).then(res => {
        this.packageDetail = res.data.data
        this.files = JSON.parse(this.packageDetail.files === '' ? '[]' : this.packageDetail.files)
        getPatentListByPackageId(packageId).then(res => {
          this.patentList = res.data.data
        })
      })
    },
    handleDeletePackage() {
      deletePackage(this.packageDetail.packageId).then(res => {
        this.$message.success('删除成功')
        this.$router.push('/patent/package')
      })
    },
    handleDeletePatentFromPackage(patent) {
      console.log(patent)
      removePatentFromPackage(this.packageDetail.packageId, patent.PNM).then(res => {
        this.$message.success('删除成功')
        this.loadPackageDetail()
      })
    },
    handleUploadFile() {
      this.$refs.uploadInput.click()
      this.$refs.uploadInput.onchange = e => {
        const formData = new FormData()
        formData.append('file', e.target.files[0])
        this.uploading = true
        uploadFile(formData).then(res => {
          const path = res.data.data.path
          updatePackage(this.packageDetail.packageId, {
            filesOpt: 'add',
            files: [path]
          }).then(res => {
            this.uploading = false
            this.loadPackageDetail()
            this.$message.success('上传成功')
          })
        })
      }
    },
    handleDeleteFile(filePath) {
      updatePackage(this.packageDetail.packageId, {
        filesOpt: 'del',
        files: [filePath]
      }).then(res => {
        this.loadPackageDetail()
        this.$message.success('删除成功')
      })
    },
    isImage(filePath) {
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
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

.my-card2 {
  width: 160px;
  height: 190px;
}

.imageField {
  height: 150px;
  padding: 10px;
}

.imageField img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.image {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.my-card2 {
  position: relative;
}

.my-card2:hover .card-actions {
  display: flex;
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

</style>
