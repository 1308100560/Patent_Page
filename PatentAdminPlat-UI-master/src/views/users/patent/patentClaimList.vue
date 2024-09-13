<template>
  <div>
    <div class="filter-container">
      <div style="display: flex;flex-direction: row;align-items: center;justify-content: center">
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
        <router-link
          style="margin-left: 10px"
          to="/search/index?h=1"
        >
          <el-button
            class="filter-item"
            icon="el-icon-plus"
            size="small"
          >专利认领
          </el-button>
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

    <el-table
      id="out-table"
      v-loading="listLoading"
      :data="list"
      :sort-method="sortMethod"
      :stripe="true"
      highlight-current-row
      style="width: 100%;"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        width="50"
      />
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
      <el-table-column label="PNM" prop="PNM" sortable width="200">
        <template slot-scope="{row}">
          <span>{{ row.patentProperties.PNM }}</span>
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
      <el-table-column label="法律状态" width="100">
        <template slot-scope="{row}">
          <el-tag
            :style="{backgroundColor:getTagColor(row.patentProperties.CLS),border:'none'}"
            effect="dark"
            size="mini"
            style="margin: 0 5px"
          >{{ row.patentProperties.CLS }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="专利发表年份" width="150">
        <template slot-scope="{row}">
          {{ row.patentProperties.PD }}
        </template>
      </el-table-column>
      <el-table-column label="申请人" width="200">
        <template slot-scope="{row}">
          <div style="font-size: 12px">
            {{ row.patentProperties.PA }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="发明人" width="200">
        <template slot-scope="{row}">
          <div style="font-size: 12px">
            {{ row.patentProperties.PINN }}
          </div>
        </template>
      </el-table-column>
      <el-table-column :sort-by="'price'" label="专利估值" min-width="250" prop="price" sortable>
        <template slot-scope="{row}">
          <div v-if="row.eval_result" class="space-between">
            <div>
              ¥{{
                (row.eval_result && row.eval_result.evalPrice).toFixed(2)
              }}
              <el-tag size="mini" style="margin-left: 3px">报告估值</el-tag>

            </div>
            <div>
              <router-link :to="`/report/list?ids=${[row.eval_result.evalReport.reportId]}`" class="link-type">来源
              </router-link>
            </div>
          </div>
          <div v-else-if="row.price" class="space-between">
            <div>
              ¥{{ (row.price).toFixed(2) }}
              <el-tag size="mini" style="margin-left: 3px" type="warning">算法估值</el-tag>
            </div>
            <router-link :to="`/report/generate/valuation?pId=${row.patentId}`" class="link-type">去估值</router-link>
          </div>
          <div v-else class="space-between">
            <div style="color: gray;display: flex;flex-direction: row;align-items: center">
              暂无估值
            </div>
            <router-link :to="`/report/generate/valuation?pId=${row.patentId}`" class="link-type">去估值</router-link>

          </div>

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
      <el-table-column label="归属专利组合" width="200">
        <template slot-scope="{row}">
          <div v-if="row.packageIDs">
            {{ row.packageIDs.length }}个
            <el-divider direction="vertical" />
            <router-link
              :to="{path:'/patent/package',query: {patent:row.patentProperties.TI,ids:row.packageIDs.join(',')}}"
            ><span
              class="link-type"
            >查看</span></router-link>
          </div>
          <div v-else>空</div>
        </template>
      </el-table-column>
      <el-table-column label="交易信息" width="160">
        <template>
          无
        </template>
      </el-table-column>
      <el-table-column label="关联视频" width="160">
        <template>
          尚未开放
        </template>
      </el-table-column>
      <el-table-column :width="fullActionWidth" class-name="small-padding fixed-width" fixed="right" label="操作">
        <template slot-scope="{row}">
          <div class="actions" style="display: flex;flex-direction: row;justify-content: space-between">

            <el-tooltip class="item" content="修改备注" effect="dark" placement="top">
              <el-button
                icon="el-icon-edit"
                size="mini"
                type="primary"
                @click="showDescDialog(row)"
              />
            </el-tooltip>
            <el-tooltip class="item" content="取消认领" effect="dark" placement="top">
              <el-button icon="el-icon-delete" type="danger" @click="unClaimClick(row)" />
            </el-tooltip>
            <el-tooltip class="item" content="加入专利组合" effect="dark" placement="top">
              <addToPackage :patent="row.patentProperties">
                <template #content="{showPopover}">
                  <el-button
                    icon="el-icon-files"
                    type="light"
                    @click.native="showPopover"
                  />
                </template>
              </addToPackage>
            </el-tooltip>

            <el-tooltip v-if="showFullAction" class="item" content="申请报告" effect="dark" placement="top">
              <router-link to="/report/list">
                <el-button icon="el-icon-sunset" />
              </router-link>
            </el-tooltip>
            <el-tooltip v-if="showFullAction" class="item" content="发起交易" effect="dark" placement="top">
              <router-link to="/trade/home">
                <el-button icon="el-icon-money" />
              </router-link>
            </el-tooltip>
            <el-tooltip v-if="showFullAction" class="item" content="上传视频" effect="dark" placement="top">
              <el-button icon="el-icon-upload" @click="$message.warning('尚未开放')" />
            </el-tooltip>
            <el-tooltip :content="showFullAction?'收起':'展开'" class="item" effect="dark" placement="top">
              <el-button :icon="showFullAction?'el-icon-s-unfold':'el-icon-s-fold'" @click="handleShowFullAction" />
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
    <el-dialog :visible.sync="editDescFromVisible" center title="修改备注" width="50%">
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
    <el-dialog :visible.sync="achievementsVisible" center title="添加科技成果" width="50%">
      <el-form>
        <el-form-item label="名称">
          <el-input v-model="achievementsForm.name" placeholder="修改名称" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="achievementsForm.content" placeholder="修改内容" type="textarea" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="achievementsVisible = false">取消</el-button>
        <el-button type="primary" @click="handleUpdateAchievements(false)">确定</el-button>
      </div>
    </el-dialog>
    <el-dialog :visible.sync="achievementsDetailVisible" center title="科技成果详情" width="50%">
      <div v-if="currentAchievementsEditable">
        <el-form>
          <div v-for="ac in currentAchievements" :key="ac.name">
            <el-form-item label="名称">
              <el-input v-model="ac.name" placeholder="修改名称" />
            </el-form-item>
            <el-form-item label="内容">
              <el-input v-model="ac.content" placeholder="修改内容" type="textarea" />
            </el-form-item>
            <div>
              <el-button
                size="mini"
                type="danger"
                @click="currentAchievements=currentAchievements.filter(item=>item.name!==ac.name)"
              >删除
              </el-button>
            </div>
          </div>
        </el-form>
      </div>
      <div v-else>
        <div v-for="ac in currentAchievements" :key="ac.name">
          <div style="display: flex;flex-direction: row;align-items: center">
            <h4>{{ ac.name }}</h4>
          </div>
          <div>{{ ac.content }}</div>
        </div>
      </div>
      <div style="margin-top: 40px;display: flex;flex-direction: row;justify-content: center">
        <el-button size="mini" @click="currentAchievementsEditable=!currentAchievementsEditable">编辑</el-button>
        <el-button size="mini" type="primary" @click="handleUpdateAchievements(true)">保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import addToPackage from '@/views/users/components/AddToPackage'
import { getClaimedPatents, unClaimPatent, updateClaimPatentProperty } from '@/api/patent'
import { getTagColor } from '@/views/users/utils'
import ReportStatus from '@/views/users/components/ReportStatus.vue'
import XLSX from 'xlsx'

export default {
  name: 'PatentList',
  components: { ReportStatus, addToPackage },
  props: {
    showAchievements: {
      type: Function,
      default: () => {
      }
    },
    switchTab: {
      type: Function,
      default: () => {
      }
    }
  },
  data() {
    return {
      achievementsVisible: false,
      achievementsDetailVisible: false,
      currentAchievements: [],
      currentAchievementsEditable: false,
      achievementsForm: {
        name: '',
        content: ''
      },
      multipleSelection: [],
      showFullAction: false,
      fullActionWidth: 165,
      patents: null,
      reportList: null,
      patentId: 0,
      editDescFromVisible: false,
      currentPatent: null,
      description: '',
      reportDialogFormVisible: false,
      list: null,
      claim: [],
      listLoading: true,
      currentRow: null,
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
      results: {}

    }
  },
  created() {
    this.getList()
  },
  methods: {
    handleSelectionChange(val) {
      const exportData = []
      val.forEach(item => {
        exportData.push({
          'ID': item.id,
          '专利名称': item.patentProperties.TI,
          'PNM': item.patentProperties.PNM,
          '备注': item.userProperties.desc,
          '法律状态': item.patentProperties.CLS,
          '专利发表年份': item.patentProperties.PD,
          '申请人': item.patentProperties.PA,
          '发明人': item.patentProperties.PINN,
          '专利估值': item.eval_result ? item.eval_result.evalPrice : item.price,
          '报告': item.relaTypedReports ? item.relaTypedReports.length : 0,
          '归属专利组合': item.packageIDs ? item.packageIDs.length : 0
        })
      })
      this.multipleSelection = exportData
    },
    exportExcel() {
      if (this.multipleSelection.length === 0) {
        this.$message({
          message: '请先选择要导出的专利',
          type: 'warning'
        })
        return
      }
      const workbook = XLSX.utils.book_new()
      const worksheet = XLSX.utils.json_to_sheet(this.multipleSelection)
      XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1')
      XLSX.writeFile(workbook, '专利管理.xlsx')
    },
    getAndShowReportByIds(ids) {
      this.$router.push({ path: '/report/list', query: { ids: ids.join(',') }})
    },
    getTagColor,

    sortMethod(a, b, prop) {
      console.log(a, b, prop)
    },
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
        this.fullActionWidth = 165
      } else {
        this.showFullAction = true
        this.fullActionWidth = 290
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
        unClaimPatent(row.patentProperties.PNM).then(response => {
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
    }
  }
}
</script>
<style scoped>

.space-between {
  display: flex;
  width: 100%;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.actions >>> .el-button {
  padding: 0 !important;
  min-width: 34px !important;
  height: 30px !important;
  margin: 0 !important;
}
</style>
