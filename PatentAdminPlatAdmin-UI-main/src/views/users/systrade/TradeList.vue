<template>
  <div class="container">
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

          <el-button
            class="filter-item"
            icon="el-icon-plus"
            size="small"
            type="primary"
            @click="addDialogVisible=true
            "
          >
            添加交易记录
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
        <el-table-column align="center" label="ID" prop="userId" width="80">
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column label="交易方" min-width="140" prop="phone">
          <template slot-scope="{row}">
            <span>{{ row.properties.buyer }}</span>
          </template>
        </el-table-column>
        <el-table-column label="交易联系方式" min-width="140" prop="phone">
          <template slot-scope="{row}">
            <span>{{ row.properties.buyerContact }}</span>
          </template>
        </el-table-column>
        <el-table-column label="交易内容" prop="phone" width="200">
          <template slot-scope="{row}">
            <span>{{ row.properties.content }}</span>
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="sex" width="200">
          <template slot-scope="{row}">
            <span v-if="row.properties.status==='completed'">
              <el-tag type="success">已完成</el-tag>
            </span>
            <span v-else-if="row.properties.status==='uncompleted'">
              <el-tag type="warning">未完成</el-tag>
            </span>
            <span v-else-if="row.properties.status==='closed'">
              <el-tag type="danger">已关闭</el-tag>
            </span>
          </template>
        </el-table-column>

        <el-table-column class-name="small-padding fixed-width" label="操作" width="280">
          <template slot-scope="{row}">
            <el-button size="mini" @click="currentRow=row;openEditDialog(row.properties)">
              完善资料
            </el-button>
            <el-button size="mini" @click="closeTrade(row)">
              关闭交易
            </el-button>

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
    <el-dialog :visible.sync="addDialogVisible" center title="添加交易记录" width="50%">
      <el-form :model="tradForm" :rules="tradFormRules">
        <el-form-item label="交易方" prop="buyer">
          <el-input v-model="tradForm.buyer" placeholder="请输入交易方" />
        </el-form-item>
        <el-form-item label="交易方联系方式" prop="buyerContact">
          <el-input v-model="tradForm.buyerContact" placeholder="请输入交易方" />
        </el-form-item>
        <el-form-item label="交易内容" prop="content">
          <el-input v-model="tradForm.content" placeholder="请输入交易内容" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="tradForm.status" placeholder="请选择状态">
            <el-option label="已完成" value="completed" />
            <el-option label="未完成" value="uncompleted" />
          </el-select>
        </el-form-item>
        <el-form-item label="时间" prop="time">
          <el-date-picker
            v-model="tradForm.time"
            placeholder="选择日期时间"
            type="datetime"
            value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="详细信息">
          <el-input v-model="tradForm.detail" placeholder="请输入详细信息" type="textarea" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible=false;currentRow=null">关闭</el-button>
        <el-button type="primary" @click="addTradeInfo">保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { createTrade, getTradeList, updateTrade } from '@/api/trade'

export default {
  data() {
    return {
      list: null,
      listLoading: true,
      addDialogVisible: false,
      currentRow: null,
      tradForm: {
        // 交易方
        buyer: '',
        // 交易联系方式
        buyerContact: '',
        // 交易内容
        content: '',
        // 状态
        status: '',
        // 时间
        time: '',
        // 详细信息
        detail: ''
      },
      tradFormRules: {
        buyer: [{ required: true, message: '请输入交易方', trigger: 'blur' }],
        content: [{ required: true, message: '请输入交易内容', trigger: 'blur' }],
        status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
        time: [{ required: true, message: '请选择时间', trigger: 'blur' }],
        buyerContact: [{ required: true, message: '请输入交易方联系方式', trigger: 'blur' }],
        detail: [{ required: true, message: '请输入详细信息', trigger: 'blur' }]
      },
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: ''
      },
      results: { count: 0 }
    }
  },
  mounted() {
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
    openEditDialog(properties) {
      this.addDialogVisible = true
      this.tradForm = JSON.parse(JSON.stringify(properties))
      console.log(JSON.parse(JSON.stringify(properties)))
    },
    addTradeInfo() {
      if (this.currentRow) {
        updateTrade(this.currentRow.id, { properties: this.tradForm }).then(response => {
          this.$message({
            message: '保存成功',
            type: 'success'
          })
          this.addDialogVisible = false
          this.getList()
        })
      } else {
        createTrade({ properties: this.tradForm }).then(response => {
          this.$message({
            message: '保存成功',
            type: 'success'
          })
          this.addDialogVisible = false
          this.getList()
        })
      }
    },
    closeTrade(row) {
      const data = JSON.parse(JSON.stringify(row.properties))
      data.status = 'closed'
      updateTrade(row.id, { properties: data }).then(response => {
        this.$message({
          message: '关闭成功',
          type: 'success'
        })
        this.getList()
      })
    },
    getList() {
      this.listLoading = true
      getTradeList(this.listQuery).then(response => {
        this.results = response.data.data
        this.list = response.data.data.list
        this.list.map(item => {
          item.properties = JSON.parse(item.properties)
          return item
        })
        console.log(this.results)
        this.listLoading = false
      })
    }
  }
}
</script>
