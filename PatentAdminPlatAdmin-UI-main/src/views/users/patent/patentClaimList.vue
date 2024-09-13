<template>
  <div class="container" style="display: flex;flex-direction: row">
    <div class="my-card" style="width: 100% ;overflow-y: auto">
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
        <el-button icon="el-icon-refresh" size="mini" @click="getList" />

      </div>

      <el-table
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        fit
        highlight-current-row
        style="width: 100%; "
      >
        <el-table-column align="center" label="PNM" prop="id" width="140">
          <template slot-scope="{row}">
            <span>{{ row.patentProperties.PNM }}</span>
          </template>
        </el-table-column>
        <el-table-column label="专利名称" min-width="120">
          <template slot-scope="{row}">
            <router-link :to="`/search/detail/${ row.patentProperties.PNM}`">
              <span class="link-type">{{ row.patentProperties.TI }}</span>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="备注" width="400">
          <template slot-scope="{row}">
            <div style="display: flex;flex-direction: row;justify-content: space-between">
              <span style="width: calc(100% - 40px)">
                <span v-if="row.desc">{{ row.desc }}</span>
                <span v-else style="color:#909399;font-size: 13px">点击右侧添加/修改备注</span>
              </span>
              <div style="width: 40px;">
                <el-button
                  circle
                  icon="el-icon-edit"
                  size="mini"
                  style="margin-left: 10px"
                  type="primary"
                  @click="showDescDialog(row)"
                />
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="200">
          <template slot-scope="row">
            <el-button size="mini" type="danger" @click="unClaimClick(row)">
              取消认领
            </el-button>
            <!--             <el-button class="detail" size="mini" type="light" @click="showDialog(row)">-->
            <!--              生成-->
            <!--            </el-button>-->
            <addToPackage :patent="row.row.patentProperties">
              <el-button size="mini" type="primary">
                加入工艺包
              </el-button>
            </addToPackage>
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
      <el-dialog :visible.sync="reportDialogFormVisible" center title="申请生成报告" width="30%">
        <el-form :model="form">
          <el-form-item :label-width="formLabelWidth" label="报告类型">
            <el-select v-model="form.type" placeholder="请选择报告类型">
              <el-option label="侵权报告" value="infringement" />
              <el-option label="估值报告" value="valuation" />
            </el-select>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="reportDialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="InsertReport(form)">确 定</el-button>
        </div>
      </el-dialog>

    </div>

  </div>
</template>

<script>
import addToPackage from '@/views/users/components/AddToPackage'
import { getAllPatents, unClaimPatent, updateClaimPatentDescription } from '@/api/patent'

export default {
  name: 'ComplexTable',
  components: { addToPackage },
  data() {
    return {
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
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: ''
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
  created() {
    this.getList()
  },
  methods: {
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
    getList() {
      this.listLoading = true
      getAllPatents(this.listQuery).then(response => {
        const results = response.data.data
        results.map(item => {
          item.patentProperties = JSON.parse(item.patentProperties)
        })
        this.results = response.data.data
        this.list = results
        for (let i = 0; i < this.list.length; i++) {
          this.claim[i] = this.list[i].patentId
        }
        this.listLoading = false
      })
    },
    unClaimClick(row) {
      unClaimPatent(row.row.patentProperties.PNM).then(response => {
        this.$message({
          message: '取消认领成功',
          type: 'success',
          duration: 1000
        })
        this.getList()
      })
    },
    handleUpdateDesc() {
      updateClaimPatentDescription(this.currentPatent.patentProperties.PNM, this.description).then(res => {
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
      this.description = row.desc
    }
  }
}
</script>
<style scoped>

</style>
