<template>
  <div class="container">
    <div class="my-card" style="width: 100%">
      <create-package ref="createPack" :get-list="getList" type="focus" />
      <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
        <div style="display: flex;flex-direction: row;align-items: flex-end">
          <el-input
            v-model="listQuery.query"
            class="filter-item"
            placeholder="关注组合名称名称"
            size="small"
            style="width: 180px;margin-right: 10px"
          />

          <el-button class="filter-item" icon="el-icon-search" size="small" type="primary" @click="getList">
            搜索
          </el-button>
          <el-button icon="el-icon-folder-add" size="small" type="primary" @click="showCreatePack">
            创建关注组合
          </el-button>
          <div
            v-if="patentName"
            style="font-size: 14px;margin-left: 10px;margin-bottom: 3px;display: flex;flex-direction: row;align-items: center"
          >
            包含 <span style="margin: 0 5px;font-style:italic;font-weight: bold">{{ patentName }}</span>
            的关注组合
            <i class="el-icon-error" style="color: #909399;cursor: pointer;margin-left: 5px" @click="getList" />
          </div>
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
        style="width: 100%;margin-top: 20px"
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
        <el-table-column label="组合名称" min-width="220">
          <template slot-scope="{row}">
            <router-link :to="`/patent/package/${ row.packageId}`">
              <span class="link-type">{{ row.packageName }}</span>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="关注组合用途" width="200">
          <template slot-scope="{row}">
            <div style="display: flex;flex-direction: row;justify-content: space-between">
              <span style="width: calc(100% - 40px)">
                <span v-if="row.properties.usage">{{ row.properties.usage }}</span>
                <span v-else style="color:#909399;font-size: 13px">暂无</span>
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="备注" width="300">
          <template slot-scope="{row}">
            <div style="display: flex;flex-direction: row;justify-content: space-between">
              <span style="width: calc(100% - 40px)">
                <span v-if="row.properties.desc">{{ row.properties.desc }}</span>
                <span v-else style="color:#909399;font-size: 13px">暂无</span>
              </span>
            </div>
          </template>
        </el-table-column>

        <el-table-column align="center" label="专利数量" min-width="100">
          <template slot-scope="{row}">
            {{ row.patentsNum }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="附件数量" min-width="100">
          <template slot-scope="{row}">
            {{ row.files ? JSON.parse(row.files).length : 0 }}
          </template>
        </el-table-column>
        <el-table-column label="估值" min-width="200">
          <template slot-scope="{row}">
            ¥ {{ (row.totalPrice).toFixed(2) }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="150">
          <template slot-scope="{row}">
            <span>{{ row.CreatedAt|localTime }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="small-padding fixed-width" fixed="right" label="操作" min-width="120">
          <template slot-scope="{row}">
            <div class="actions" style="display: flex;flex-direction: row;justify-content: space-between">
              <el-button
                icon="el-icon-edit"
                type="primary"
                @click="packageDetailDialogVisible=true;packageDetail=row"
              />
              <el-button
                icon="el-icon-delete"
                style="margin-left: 10px"
                type="danger"
                @click="handleDeletePackage(row.packageId)"
              />
              <el-button
                icon="el-icon-plus"
                @click="$router.push(`/explore/follow`)"
              />
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

    </div>
    <el-dialog :visible.sync="packageDetailDialogVisible" title="编辑关注组合">
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

  </div>
</template>
<script>
import createPackage from '@/views/users/components/CreatePackage.vue'
import { deletePackage, getPackageList, getPackageListByIds, updatePackage } from '@/api/package'
import XLSX from 'xlsx'

export default {
  name: 'TechPack',
  components: {
    createPackage
  },
  data() {
    return {
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
      packageDetail: { packageName: '', properties: { usage: '', desc: '' }},
      packageDetailDialogVisible: false,
      listLoading: true,
      patentName: '',
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: '',
        type: 'focus'
      },
      list: [],
      results: { count: 0 },
      packageList: []
    }
  },
  mounted() {
    const { query } = this.$route
    if (query.ids) {
      const ids = JSON.stringify(query.ids.split(',').map(item => Number(item)))
      this.patentName = query.patent
      this.getListBydIds(ids)
    } else {
      this.getList()
    }
  },
  methods: {
    handleSelectionChange(val) {
      console.log(val)
      const exportData = []
      val.forEach(item => {
        exportData.push({
          'ID': item.id,
          '组合名称': item.packageName,
          '关注组合用途': item.properties.usage,
          '备注': item.properties.desc,
          '专利数量': item.patentsNum,
          '附件数量': item.files ? JSON.parse(item.files).length : 0,
          '估值': item.totalPrice,
          '创建时间': new Date(item.CreatedAt).toLocaleString().split(' ')[0]
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
      XLSX.writeFile(workbook, '关注组合.xlsx')
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
        this.getList()
      })
    },
    handleDeletePackage(packageId) {
      this.$confirm('此操作将永久删除该关注组合, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deletePackage(packageId).then(res => {
            this.$message.success('删除成功')
            this.getList()
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
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
    getListBydIds(ids) {
      this.listLoading = true
      getPackageListByIds({ ids }).then(res => {
        res.data.data.map(item => {
          item.properties = item.properties ? JSON.parse(item.properties) : {}
          return item
        })
        this.results = res.data.data
        this.list = res.data.data
        // add id to list
        this.list.forEach((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        this.listLoading = false
      })
    },
    getList() {
      this.patentName = ''
      this.listLoading = true
      getPackageList(this.listQuery).then(res => {
        res.data.data.list.map(item => {
          item.properties = item.properties ? JSON.parse(item.properties) : {}
          return item
        })
        this.results = res.data.data
        this.list = res.data.data.list
        // add id to list
        this.list.forEach((item, index) => {
          item.id = (this.listQuery.pageIndex - 1) * this.listQuery.pageSize + index + 1
        })
        this.listLoading = false
      })
    },
    showCreatePack() {
      this.$refs.createPack.show()
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
</style>
