<template>
  <div class="container ">
    <div class="my-card">
      <el-dialog :visible.sync="packageDetailDialogVisible" title="编辑专利组合">
        <el-form
          ref="form"
          :model="packageDetail"
          label-width="120px"
          size="small"
          style="margin: 10px"
        >
          <el-form-item label="名称" prop="packageName" size="small">
            <el-input v-model="packageDetail.packageName" size="small" />
          </el-form-item>
          <el-form-item label="用途" prop="packageName" size="small">
            <el-select
              v-model="packageDetail.properties.usage"
              allow-create
              default-first-option
              filterable
              placeholder="请选择专利用途"
              style="width: 100%"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="备注" size="small">
            <el-input
              v-model="packageDetail.properties.desc"
              :rows="2"
              placeholder="请输入内容"
              type="textarea"
            />

          </el-form-item>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="packageDetailDialogVisible=false">取消</el-button>
            <el-button size="mini" type="primary" @click="updatePackageSubmit">保存</el-button>
          </div>
        </el-form>
      </el-dialog>
      <el-page-header :content="packageDetail.packageName" @back="$router.go(-1)" />
      <div style="margin-top: 10px">
        <div class="text" style="display: flex;flex-direction: row;align-items: center">

          <!--          <div v-if="packageDetail.properties.desc">-->
          <!--            <el-divider direction="vertical" />-->
          <!--            {{ packageDetail.properties.desc }}-->
          <!--            <el-divider direction="vertical" />-->
          <!--          </div>-->
          <div>

            附件数量：{{ packageDetail.files ? JSON.parse(packageDetail.files).length : 0 }}
            <el-divider direction="vertical" />
            专利数量：{{ patentList && patentList.list.length }}

            <el-divider direction="vertical" />
            估值： ¥ {{ (packageDetail.totalPrice).toFixed(2) }}
            <el-divider direction="vertical" />

            创建时间：{{ packageDetail.CreatedAt|localTime }}
          </div>

        </div>
      </div>
      <div
        style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;margin-top: 10px"
      >
        <div style="display: flex;flex-direction: row;align-items: center">
          <!--        <el-button icon="el-icon-download" size="small" type="primary">打包下载</el-button>-->
          <el-button icon="el-icon-edit" size="small" type="primary" @click="packageDetailDialogVisible=true">编辑
          </el-button>
          <el-button icon="el-icon-delete" size="small" type="danger" @click="handleDeletePackage()">删除专利组合
          </el-button>
        </div>
        <div />
      </div>
    </div>
    <div
      style="display: flex;overflow-y:auto;flex-direction: row;justify-content: space-between;padding-top: 10px;height: calc(100vh - 220px)"
    >
      <div
        class="my-card"
        style="width: 69%;padding-top: 0;"
      >
        <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
          <p>专利列表</p>
          <div>
            <router-link :to="`${packageDetail.type==='focus'?'/explore/follow':'/patent/claim'}`">
              <el-button size="small" style="margin-left: 10px" type="primary">添加专利</el-button>
            </router-link>
            <el-button
              class="filter-item"
              icon="el-icon-upload2"
              size="small"
              style="margin-left: 10px"
              @click="exportExcel"
            >导出
            </el-button>
          </div>

        </div>
        <div style="overflow-y: auto;height: calc(100% - 30px)">
          <el-table
            :data="patentList.list"
            style="width: 100%"
            @selection-change="handleSelectionChange"
          >
            <el-table-column
              type="selection"
              width="50"
            />
            <el-table-column
              label="名称"
              prop="date"
              width="180"
            >

              <template slot-scope="scope">
                {{ JSON.parse(scope.row.patentProperties).PNM }}
              </template>
            </el-table-column>
            <el-table-column
              label="专利名称"
              min-width="180"
              prop="name"
            >
              <template slot-scope="scope">
                <router-link :to="`/search/detail/${JSON.parse(scope.row.patentProperties).PNM }`" class="link-type">
                  {{ JSON.parse(scope.row.patentProperties).TI }}
                </router-link>
              </template>
            </el-table-column>
            <el-table-column
              label="备注"
              prop="address"
            >
              <template>
                无
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              width="90"
            >
              <template slot-scope="scope">
                <div style="display: flex;flex-direction: row">
                  <el-button
                    size="mini"
                    style="margin-right: 10px"
                    type="danger"
                    @click="handleDeletePatentFromPackage(scope.row)"
                  >
                    删除
                  </el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>

      </div>
      <div style="width: 30%;padding-top: 0!important;margin-top: 0!important;">
        <div
          :style="{height:packageDetail.type!=='demo'?'50%':'100%'}"
          class="my-card"
          style="height: 50%;padding-top: 0!important;"
        >
          <div style="display: flex;flex-direction: row;justify-content: space-between;align-items: center">
            <p>文档浏览区</p>
            <el-button
              :loading="uploading"
              icon="el-icon-upload"
              size="small"
              style="height: 35px"
              type="primary"
              @click="handleUploadFile"
            >
              {{ uploading ? '上传中' : '上传文件' }}
              <input v-show="false" ref="uploadInput" type="file">
            </el-button>
          </div>
          <div class="cards">
            <el-table
              :data="files"
              style="width: 100%"
            >
              <el-table-column
                label="文件名称"
                prop="date"
                width="180"
              >

                <template slot-scope="scope">
                  {{ scope.row.name }}
                </template>
              </el-table-column>
              <el-table-column
                label="操作"
                min-width="180"
                prop="name"
              >
                <template slot-scope="scope">
                  <div style="display: flex;flex-direction: row">
                    <el-button
                      size="mini"
                      style="margin-right: 10px"
                      type="danger"
                      @click="handleDeleteFile(scope.row)"
                    >
                      删除
                    </el-button>
                    <download-able :name="scope.row.name" :url="`${scope.row.full_path}`">
                      <el-button size="mini" type="primary">下载</el-button>
                    </download-able>
                  </div>
                </template>
              </el-table-column>
            </el-table>

          </div>
        </div>
        <div
          v-if="packageDetail.type!=='demo'"
          v-loading="loadingGraph"
          :style="{height:packageDetail.type!=='demo'?'calc(50% - 10px)':'100%'}"
          class="my-card"
          style="margin-top: 10px;padding-top: 10px"
        >
          <div
            style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;margin-bottom: 10px"
          >
            <div>
              <span style="font-size: 14px">选择图谱：</span>
              <el-select v-model="graphType" filterable placeholder="请选择" size="mini">
                <el-option
                  label="发明人图谱"
                  value="rela"
                />
                <el-option
                  label="技术路线图谱"
                  value="tech"
                />
              </el-select>
            </div>
            <router-link
              :to="{path:`/${packageDetail.type==='focus'?'explore':'patent'}/graph`,query:{packageId:packageDetail.packageId,graphType:graphType}}"
              class="link-type"
              style="font-size: 0.9rem"
            >查看大图
            </router-link>
          </div>
          <div id="chartDetail" style="height: calc(100% - 10px);width: calc(100% - 20px)" />
        </div>
      </div>

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
import XLSX from 'xlsx'
import { getGraphData } from '@/api/graph'
import echarts from 'echarts'

require('echarts/theme/macarons') // echarts theme

const option = {
  title: {
    text: '技术图谱',
    textAlign: 'auto',
    left: '10%', // '5' | '5%'，title 组件离容器左侧的距离
    right: 'auto', // 'title 组件离容器右侧的距离
    top: '20%', // title 组件离容器上侧的距离
    bottom: 'auto'
  },
  tooltip: {},
  legend: [
    {
      data: []
    }
  ],
  series: [
    {
      name: '专利关键词',
      type: 'graph',
      layout: 'force',
      data: null,
      links: null,
      categories: [],
      roam: true,
      label: {
        show: true,
        formatter: '{b}',
        position: 'right'
      },
      focusNodeAdjacency: true,
      legendHoverLink: true,
      animation: false,
      force: {
        initLayout: 'circular',
        layoutAnimation: false,
        repulsion: 100
      },
      lineStyle: {
        color: 'source',
        opacity: 0.2,
        curveness: 0.3
      },
      emphasis: {
        focus: 'adjacency',
        itemStyle: {
          shadowColor: 'rgba(0, 0, 0, 0.4)',
          shadowBlur: 15
        },
        lineStyle: {
          width: 3
        },
        label: {
          textBorderColor: 'rgba(255, 255, 255, 0.8)',
          textBorderWidth: 2
        }
      }
    }
  ]
}
export default {
  name: 'PackageDetail',
  components: { DownloadAble },
  data() {
    return {
      loadingGraph: false,
      options: [{
        value: '工程授权',
        label: '工程授权'
      }, {
        value: '批量交易',
        label: '批量交易'
      }, {
        value: '研究定向',
        label: '研究定向'
      }],
      value: [],
      uploading: false,
      graphType: 'tech',
      packageDetail: { packageName: '', properties: { usage: '', desc: '' }},
      patentList: [],
      files: [],
      packageDetailDialogVisible: false
    }
  },
  watch: {
    graphType: function(val) {
      this.getChartData()
    }
  },

  mounted() {
    this.loadPackageDetail()
    const self = this
    const interval = setInterval(() => {
      const chart = document.getElementById('chartDetail')
      if (chart !== null) {
        self.chart = echarts.init(document.getElementById('chartDetail'))
        this.getChartData()
        clearInterval(interval)
      }
    }, 1000)
  },
  methods: {
    handleSelectionChange(val) {
      const exportData = []
      let cnt = 1
      val.forEach(item => {
        exportData.push({
          'ID': cnt++,
          '专利名称': JSON.parse(item.patentProperties).TI,
          'PNM': item.PNM,
          '法律状态': JSON.parse(item.patentProperties).CLS,
          '交易状态': JSON.parse(item.patentProperties).CLS
        })
      })
      this.multipleSelection = exportData
    },
    exportExcel() {
      if (!this.multipleSelection || this.multipleSelection.length === 0) {
        this.$message({
          message: '请先选择要导出的专利',
          type: 'warning'
        })
        return
      }
      const workbook = XLSX.utils.book_new()
      const worksheet = XLSX.utils.json_to_sheet(this.multipleSelection)
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1')
      XLSX.writeFile(workbook, `${this.packageDetail.packageName}-专利列表.xlsx`)
    },
    getChartData() {
      this.loadingGraph = true
      const params = {
        scope: 'package',
        type: 'rela',
        packageId: this.packageDetail.packageId
      }
      this.loading = true
      if (this.graphType === 'rela') {
        option.title.text = '发明人图谱'
      } else {
        option.title.text = '技术路线图谱'
      }
      getGraphData(params).then(res => {
        this.loadingGraph = false
        const results = res.data.data
        if (results == null) {
          option.series[0].data = null
          option.series[0].links = null
        } else {
          option.series[0].data = results.nodes
          option.series[0].links = results.links
          this.tableData = results.nodes.slice(0, 10)
          option.legend[0].data = this.tableData.map(item => item.name)
          option.series[0].categories = this.tableData.map(item => {
            return { name: item.name }
          })
          this.chart.setOption(option)
        }
        this.loading = false
      })
    },
    updatePackageSubmit() {
      if (this.packageDetail.packageName === '') {
        this.$message.error('请输入包名')
        return
      }
      const data = {
        packageName: this.packageDetail.packageName,
        properties: {
          desc: this.packageDetail.properties.desc,
          usage: this.packageDetail.properties.usage
        }
      }
      updatePackage(this.packageDetail.packageId, data).then(res => {
        this.packageDetailDialogVisible = false
        this.$message.success('保存成功')
      })
    },
    loadPackageDetail() {
      const packageId = this.$route.params.id
      getPackage(packageId).then(res => {
        res.data.data.properties = res.data.data.properties ? JSON.parse(res.data.data.properties) : {}

        this.packageDetail = res.data.data
        this.files = JSON.parse(this.packageDetail.files === '' ? '[]' : this.packageDetail.files)
        getPatentListByPackageId(packageId).then(res => {
          this.patentList = res.data.data
        })
      })
    },
    handleDeletePackage() {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deletePackage(this.packageDetail.packageId).then(res => {
          this.$message.success('删除成功')
          this.$router.push('/patent/package')
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        })
      })
    },
    handleDeletePatentFromPackage(patent) {
      // confirm
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removePatentFromPackage(this.packageDetail.packageId, patent.PNM).then(res => {
          this.$message.success('删除成功')
          this.loadPackageDetail()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        })
      })
    },
    handleUploadFile() {
      this.$refs.uploadInput.click()
      this.$refs.uploadInput.onchange = e => {
        const formData = new FormData()
        formData.append('file', e.target.files[0])
        this.uploading = true
        uploadFile(formData).then(res => {
          const file = res.data.data
          updatePackage(this.packageDetail.packageId, {
            filesOpt: 'add',
            files: [file]
          }).then(res => {
            this.uploading = false
            this.loadPackageDetail()
            this.$message.success('上传成功')
          })
        })
      }
    },
    handleDeleteFile(file) {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        updatePackage(this.packageDetail.packageId, {
          filesOpt: 'del',
          files: [file]
        }).then(res => {
          this.loadPackageDetail()
          this.$message.success('删除成功')
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        })
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

.imageField img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

</style>
