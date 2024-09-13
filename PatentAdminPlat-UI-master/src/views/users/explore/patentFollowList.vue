<template>
  <div class="container">
    <div class="my-card" style="width: 100% ;">
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
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        fit
        highlight-current-row
        style="width: 100%;"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="50"
        />
        <el-table-column align="center" label="ID" prop="id" width="80">
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column label="专利名称" min-width="400">
          <template slot-scope="{row}">
            <router-link :to="`/search/detail/${ row.patentProperties.PNM}`">
              <span class="link-type">{{ row.patentProperties.TI }}</span>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="PNM"
          prop="id"
          sortable
          width="140"
        >
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
        <el-table-column label="交易状态" width="100">
          <template>
            公开
          </template>
        </el-table-column>
        <el-table-column label="归属专利组合" width="200">
          <template slot-scope="{row}">
            <div v-if="row.packageIDs">
              {{ row.packageIDs.length }}个
              <el-divider direction="vertical" />
              <router-link
                :to="{path:'/explore/package',query: {patent:row.patentProperties.TI,ids:row.packageIDs.join(',')}}"
              ><span
                class="link-type"
              >查看</span></router-link>
            </div>
            <div v-else>空</div>
          </template>
        </el-table-column>
        <el-table-column class-name="small-padding fixed-width" fixed="right" label="操作" width="120">
          <template slot-scope="row">

            <div class="actions">

              <el-button
                icon="el-icon-edit"
                type="primary"
                @click="showDescDialog(row.row)"
              />
              <el-button icon="el-icon-delete" type="danger" @click="unFocusPatentClick(row)" />
              <el-tooltip class="item" content="加入关注组合" effect="dark" placement="top">
                <addToPackage :patent="row.row.patentProperties" type="focus">
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
    </div>
  </div>
</template>

<script>
import { getFocusedPatents, unFocusPatent, updateFocusPatentProperty } from '@/api/patent'
import { getTagColor } from '@/views/users/utils'
import AddToPackage from '@/views/users/components/AddToPackage.vue'
import XLSX from 'xlsx'

export default {
  components: { AddToPackage },
  data() {
    return {
      editDescFromVisible: false,
      currentPatent: null,
      description: '',
      list: null,
      listLoading: true,
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: ''
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
          '交易状态': item.patentProperties.CLS,
          '归属专利组合': item.packageIDs ? item.packageIDs.length : 0
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
      XLSX.writeFile(workbook, '关注的专利.xlsx')
    },
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
    showDescDialog(row) {
      this.editDescFromVisible = true
      this.currentPatent = row
      this.description = row.userProperties.desc
    },
    getList() {
      this.listLoading = true
      getFocusedPatents(this.listQuery).then(response => {
        const results = response.data.data
        results.list.map(item => {
          item.patentProperties = JSON.parse(item.patentProperties)
        })
        // add id by page
        results.list.map((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        this.list = results.list
        this.results = results
        this.listLoading = false
      })
    },
    unFocusPatentClick(row) {
      this.$confirm('确定取消关注该专利吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        unFocusPatent(row.row.patentProperties.PNM).then(response => {
          this.$message({
            message: '取消关注成功',
            type: 'success',
            duration: 5 * 1000
          })
          this.getList()
        })
      })
    },
    handleUpdateDesc() {
      const properties = JSON.parse(JSON.stringify(this.currentPatent.userProperties))
      properties.desc = this.description
      updateFocusPatentProperty(this.currentPatent.patentProperties.PNM, properties).then(res => {
        this.$message({
          message: '修改成功',
          type: 'success',
          duration: 1000
        })
        this.editDescFromVisible = false
        this.getList()
      })
    }
  }
}
</script>
<style scoped>
.actions {
  display: flex;
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
