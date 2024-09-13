<template>
  <div v-if="dataList">
    <div class="filter-container">
      <div>
        <el-input
          v-model="listQuery.query"
          class="filter-item"
          placeholder="关键词"
          size="small"
          style="width: 200px;margin-right: 10px"
        />

        <el-button class="filter-item" icon="el-icon-search" size="small" type="primary" @click="getList">
          搜索
        </el-button>
      </div>

    </div>
    <el-table
      v-loading="listLoading"
      :data="results.list"
      :stripe="true"
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column align="center" label="ID" prop="id" width="100">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="专利名称" width="400">
        <template slot-scope="{row}">
          <router-link :to="`/search/detail/${ row.patentProperties.PNM}`">
            <span class="link-type">{{ row.patentProperties.TI }}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column label="PNM" prop="PNM" width="200">
        <template slot-scope="{row}">
          <span>{{ row.patentProperties.PNM }}</span>
        </template>
      </el-table-column>
      <el-table-column label="报告" width="220">
        <template slot-scope="{row}">
          <report-status
            :get-and-show-report-by-ids="getAndShowReportByIds"
            :patent-id="row.patentId"
            :report-list="row.relaTypedReports"
          />
        </template>
      </el-table-column>
      <el-table-column label="备注" width="400">
        <template slot-scope="{row}">
          <div style="display: flex;flex-direction: row;justify-content: space-between">
            <span style="width: calc(100% - 40px)">
              <span v-if="row.userProperties.desc">{{ row.userProperties.desc }}</span>
              <span v-else style="color:#909399;font-size: 13px">点击右侧添加/修改备注</span>
            </span>
          </div>
        </template>
      </el-table-column>

      <el-table-column :width="fullActionWidth" class-name="small-padding fixed-width" fixed="right" label="操作">
        <template slot-scope="row">
          <div class="actions" style="display: flex;flex-direction: row;justify-content: space-between">

            <el-tooltip class="item" content="修改备注" effect="dark" placement="top">
              <el-button
                icon="el-icon-edit"
                size="mini"
                type="primary"
                @click="showDescDialog(row.row)"
              />
            </el-tooltip>
            <el-tooltip class="item" content="取消认领" effect="dark" placement="top">
              <el-button icon="el-icon-delete" type="danger" @click="unClaimClick(row)" />

            </el-tooltip>
            <el-tooltip class="item" content="加入专利组合" effect="dark" placement="top">
              <addToPackage :patent="row.row.patentProperties">
                <template #content="{showPopover}">
                  <el-button
                    icon="el-icon-files"
                    type="light"
                    @click.native="showPopover"
                  />
                </template>
              </addToPackage>
            </el-tooltip>

          </div>
        </template>
      </el-table-column>
    </el-table>
    <div style="display:flex;flex-direction: row;justify-content: center;margin-top: 15px;">
      <el-pagination
        :current-page="listQuery.pageIndex"
        :hide-on-single-page="false"
        :page-size="100"
        :page-sizes="[10,20,40]"
        :total="results.count"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <el-dialog :visible.sync="editDescFromVisible" center title="修改备注" width="30%">
      <el-form>
        <el-form-item label="备注">
          <el-input v-model="description" placeholder="修改备注" type="textarea" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editDescFromVisible = false">取消</el-button>
        <el-button type="primary" @click="handleUpdateDesc">确定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import addToPackage from '@/views/users/components/AddToPackage'
import ReportStatus from '@/views/users/components/ReportStatus.vue'
import { getClaimedPatents, unClaimPatent, updateClaimPatentProperty } from '@/api/patent'
import { ApplyReport, userGetReportListByPaId } from '@/api/report'
import { getTagColor } from '@/views/users/utils'

export default {
  name: 'ComplexTable',
  components: { addToPackage, ReportStatus },
  props: {
    getAndShowReportByIds: {
      type: Function,
      default: () => {
      }
    },
    dataList: {
      type: Object,
      default: () => {
      }
    }
  },
  data() {
    return {
      showFullAction: false,
      fullActionWidth: 125,
      patents: null,
      reportList: null,
      patentId: 0,
      editDescFromVisible: false,
      currentPatent: null,
      description: '',
      reportDialogFormVisible: false,
      list: null,
      claim: [],
      listLoading: false,
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: '',
        needEvalResult: true
      },
      flag: 0,
      form: {
        patentId: '',
        type: ''
      },
      results: {},
      formLabelWidth: '120px'

    }
  },
  watch: {
    dataList: {
      handler(val) {
        console.log(val)
        this.results = val
      },
      deep: true
    }
  },
  created() {
    this.results = JSON.parse(JSON.stringify(this.dataList))
  },
  methods: {
    getTagColor,
    handleSearch() {
      this.listQuery.pageIndex = 1
      this.getList()
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.pageIndex = val
      this.getList()
    },
    handleShowFullAction() {
      if (this.showFullAction) {
        this.showFullAction = false
        this.fullActionWidth = 160
      } else {
        this.showFullAction = true
        this.fullActionWidth = 300
      }
    },
    getList() {
      this.listLoading = true
      getClaimedPatents(this.listQuery).then(response => {
        const results = response.data.data.list
        results.map(item => {
          item.patentProperties = JSON.parse(item.patentProperties)
        })
        this.results = response.data.data
        this.list = results
        this.list.forEach((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        for (let i = 0; i < this.list.length; i++) {
          this.claim[i] = this.list[i].patentId
        }
        this.listLoading = false
      })
    },
    unClaimClick(row) {
      this.$confirm('此操作将取消认领该专利, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        unClaimPatent(row.row.patentProperties.PNM).then(response => {
          this.$message({
            message: '取消认领成功',
            type: 'success',
            duration: 1000
          })
          this.getList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleUpdateDesc() {
      const properties = JSON.parse(JSON.stringify(this.currentPatent.userProperties))
      properties.desc = this.description
      updateClaimPatentProperty(this.currentPatent.patentProperties.PNM, properties).then(res => {
        this.$message({
          message: '修改成功',
          type: 'success',
          duration: 1000
        })
        this.editDescFromVisible = false
        this.getList()
      })
    },
    showDialog(row) {
      this.reportDialogFormVisible = true
      this.patentId = row.patentId
    },
    showDescDialog(row) {
      this.editDescFromVisible = true
      this.currentPatent = row
      this.description = row.userProperties.desc
    },

    InsertReport(form) {
      this.flag = 0
      this.reportDialogFormVisible = false
      this.form.patentId = this.patentId
      console.log(this.form)
      userGetReportListByPaId(this.form.patentId).then(response => {
        this.reportList = response.data.data
        this.listLoading = false
        if (this.reportList !== null) {
          for (let i = 0; i < this.reportList.length && this.flag === 0; i++) {
            if (this.reportList[i].Type === this.form.type) {
              this.$message({
                message: '您已申请该类型报告，点击详情查看',
                type: 'error',
                duration: 1000
              })
              this.flag = 1
              break
            }
          }
        }
        if (this.flag === 0) {
          ApplyReport(form).then(response => {
            if (response.data.code === 200) {
              this.$message({
                message: '申请成功',
                type: 'success',
                duration: 1000
              })
            }
          })
        }
      })
    }
  }
}
</script>
<style scoped>
.actions >>> .el-button {
  padding: 0 !important;
  min-width: 34px !important;
  height: 30px !important;
  margin: 0 !important;
}

.space-between {
  display: flex;
  width: 100%;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
