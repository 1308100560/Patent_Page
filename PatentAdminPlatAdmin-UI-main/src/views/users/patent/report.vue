<template>

  <div class="app-container" style="display: flex;flex-direction: row">
    <div style="width: 100% ;">
      <div style="display: flex;flex-direction: row;align-items: center;margin-bottom: 20px">
        <el-select v-model="value" placeholder="请选择" size="small" style="width: 200px;margin-right: 10px">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <el-button
          class="filter-item"
          icon="el-icon-search"
          size="small"
          type="primary"
          @click="getReport(queryid)"
        >
          搜索
        </el-button>
        <!--
        <el-button v-waves class="filter-item" icon="el-icon-refresh" size="small" type="primary" style="position:absolute;
        right:20px;" @click="flash()">
          刷新
        </el-button> -->
        <el-button
          v-if="ifquery || ifpatent || ifupload"
          v-waves
          class="filter-item"
          icon="el-icon-d-arrow-left"
          size="small"
          type="primary"
          @click="back"
        >
          返回
        </el-button>

      </div>
      <input v-show="false" ref="uploadInput" type="file">

      <el-table
        v-if="ifpatent === false && ifupload === false"
        :key="tableKey"
        :data="list"
        :loading="listLoading"
        border
        fit
        highlight-current-row
        style="width: 100%;
      border-radius: 10px!important;"
      >

        <el-table-column align="center" label="ID" prop="id" sortable="custom" width="60">
          <template slot-scope="{row}">
            <span>{{ row.reportId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="报告名称" min-width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.reportName }}</span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="申请时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.CreatedAt|localTime }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="报告类型" min-width="60">
          <template slot-scope="{row}">
            <span> {{ row.Type }} </span>
          </template>

        </el-table-column>

        <el-table-column align="center" label="状态" width="110px">

          <template slot-scope="{row}">
            <span>{{ row.rejectTag }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" class-name="small-padding fixed-width" label="报告文件" width="130">
          <template slot-scope="{row}">
            <el-button
              :disabled="row.rejectTag !== '已上传'"
              icon="el-icon-view"
              size="mini"
              type="primary"
              @click="getFiles(row.reportId)"
            >
              详情
            </el-button>
          </template>
        </el-table-column>

      </el-table>

      <el-table
        v-if="ifpatent"
        :key="tableKey"
        :data="patentlist"
        :loading="listLoading"
        border
        fit
        n
        style="width: 100%;  border-radius: 10px!important;"
      >

        <el-table-column align="center" label="ID" prop="id" sortable="custom" width="90">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentId }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="专利名称" min-width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentProperties.TI }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="申请日" sortable="custom" width="220px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentProperties.AD }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="发明人" width="110px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentProperties.PINN }}.el</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="主分类号" width="100px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.PNM }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="简介" width="300">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.patentProperties.CL }}</span>
          </template>
        </el-table-column>

      </el-table>

      <el-table
        v-if="ifupload"
        :key="tableKey"
        :data="files"
        :loading="listLoading"
        border
        fit
        style="width: 100%;  border-radius: 10px!important;"
      >
        <el-table-column align="center" label="报告ID" prop="id" sortable="custom" width="90">
          <template slot-scope="{row}">
            <span>{{ row.reportId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="文件名称" min-width="150px">
          <template slot-scope="{row}">
            <span>{{ row.files[0].FileName }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="上传日期" sortable="custom" width="220px">
          <template slot-scope="{row}">
            <span>{{ row.UpdatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="缩略图" width="300">
          <template slot-scope="{row}">
            <div v-if="ifupload" class="cards">

              <el-card
                v-for="(file, index) in row.files"
                :key="`file-${index}`"
                :style="{ padding: '0px', margin: '10px 10px 10px 0',border:'none' }"
                class="my-card"
                shadow="hover"
              >
                <div class="card-actions">
                  <div style="display: flex;flex-direction: row;align-items: center;justify-content: center">
                    <a :href="`http://${file.FilePath}`" target="_blank">
                      <el-button size="mini" type="primary">下载</el-button>
                    </a>
                  </div>
                </div>
                <div class="imageField">
                  <img v-if="isImage(file.FilePath)" :src="`http://${file.FilePath}`" alt="" class="image">
                  <svg v-else aria-hidden="true" class="image">
                    <use :xlink:href="`#icon-files2`" />
                  </svg>
                </div>
                <div style="font-size: 0.8rem;text-align: center">
                  {{ file.FileName }}
                </div>

              </el-card>

            </div>
          </template>
        </el-table-column>

      </el-table>

    </div>

  </div>

</template>
<script>
import waves from '@/directive/waves' // waves directive
import { userGetReportById, userGetReportListByType, userGetReportPatents, userReportList } from '@/api/report'

export default {
  name: 'ComplexTable',
  directives: { waves },

  data() {
    return {
      files: null,
      tableKey: 0,
      patents: null,
      ifquery: false,
      ifupload: false,
      ifpatent: false,
      ifreport: false,
      ifshow1: false,
      patentitems: null,
      patentlist: [],
      patentid: 0,
      centerDialogVisible: false,
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      queryid: '',
      options: [{
        value: 'infringement',
        label: '侵权报告'
      }, {
        value: 'valuation',
        label: '估值报告'
      }, {
        value: 'none',
        label: '查新报告'
      }],
      value: ''

    }
  },
  created() {
    this.getList()
  },
  methods: {
    isImage(filePath) {
      filePath = filePath || ''
      return filePath.endsWith('.jpg') || filePath.endsWith('.png') || filePath.endsWith('.jpeg')
    },
    preview(file) {
      this.centerDialogVisible = true
    },

    getReport(a) {
      this.listLoading = true
      if (a !== '') {
        userGetReportById(a).then(response => {
          if (response.data.data === null) {
            alert('您搜索的报告不存在或无权查看')
            this.ifquery = true
            this.listLoading = false
            return
          }
          this.ifquery = true
          this.list = [{}]
          this.list[0] = response.data.data
          this.listLoading = false
        })
      } else {
        console.log(this.value)
        userGetReportListByType(this.value).then(response => {
          if (response.data.data === null) {
            alert('您搜索的报告不存在或无权查看')
            this.ifquery = true
            this.listLoading = false
            return
          }
          this.ifquery = true
          this.list = [{}]

          for (var i = 0; i < response.data.data.length; i++) {
            this.list[i] = response.data.data[i]
            if (this.list[i].UpdatedAt === '') this.list[i].UpdatedAt = '无'
            if (this.list[i].Type === 'infringement') {
              this.list[i].Type = '侵权报告'
            }
            if (this.list[i].Type === 'valuation') {
              this.list[i].Type = '估值报告'
            }
          }

          this.listLoading = false
        })
      }
    },
    getFiles(id) {
      this.listLoading = true
      this.ifupload = !this.ifupload
      userGetReportById(id).then(response => {
        // 将数据转化为js对象
        this.files = response.data.data
        this.files.files = JSON.parse(response.data.data.files)
        this.files = [{}]
        this.files[0] = response.data.data
        this.listLoading = false
      })
    },

    getList() {
      this.listLoading = true
      userReportList().then(response => {
        this.list = response.data.data
        for (var i = 0; i < this.list.length; i++) {
          if (this.list[i].UpdatedAt === '') this.list[i].UpdatedAt = '无'
          if (this.list[i].Type === 'infringement') {
            this.list[i].Type = '侵权报告'
          }
          if (this.list[i].Type === 'valuation') {
            this.list[i].Type = '估值报告'
          }
        }
        this.listLoading = false
      })
    },
    getPatent(a) {
      this.listLoading = true
      userGetReportPatents(a).then(response => {
        this.patents = response.data.data
        this.patentlist[0] = this.patents
        this.ifpatent = true
        this.patents.patentProperties = JSON.parse(this.patents.patentProperties)
        this.listLoading = false
      })
    },

    back() {
      this.ifpatent = false
      this.ifupload = false
      this.ifquery = false
      this.queryid = ''
      this.getList()
    },
    flash() {
      this.getList()
      this.queryid = ''
    },

    download(path) {
      console.log(path)
      this.url = 'http://localhost:8000' + path
      console.log(this.url)
      window.open(this.url, '_self')
    }

  }
}
</script>
<style>
.detail {

  background-color: transparent;
  border: 1px solid #409EFF;
  color: #409EFF;

}

.container {
  padding: 10px 15px;
}

.cards {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 auto;
  flex-wrap: wrap;
}

.my-card {
  width: 180px;
  height: 200px;
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

.my-card {
  position: relative;
}

.my-card:hover .card-actions {
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
