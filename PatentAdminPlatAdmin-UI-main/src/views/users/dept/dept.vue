<template>

  <div class="app-container" style="display: flex;flex-direction: row">
    <div style="width: 100% ;">
      <div class="filter-container">
        <el-input
          v-model="queryid"
          class="filter-item"
          placeholder="团队名称"
          size="small"
          style="width: 200px;margin-right: 10px"
        />

        <el-button
          v-waves
          class="filter-item"
          disabled
          icon="el-icon-search"
          size="small"
          type="primary"
          @click="getReport(queryid)"
        >
          搜索
        </el-button>

        <el-button
          v-waves
          class="filter-item"
          icon="el-icon-refresh"
          size="small"
          style="position:absolute;
            right:20px;"
          type="primary"
          @click="flash()"
        >
          刷新
        </el-button>
        <el-button
          v-if="ifquery || ifpatent"
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

      <el-table
        v-if="ifpatent === false"
        :key="tableKey"
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        border
        fit
        highlight-current-row
        style="width: 100%;
          border-radius: 10px!important;  "
      >

        <el-table-column align="center" label="团队ID" prop="id" sortable="custom" width="100">
          <template slot-scope="{row}">
            <span>{{ row.deptId }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="用户身份" min-width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.MemType }}</span>
          </template>

        </el-table-column>

        <el-table-column align="center" label="申请内容" width="110px">
          <template slot-scope="{row}">
            <span>{{ row.memStatus }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.createdAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="审核状态" min-width="90px">
          <template slot-scope="{row}">
            <span class="link-type"> {{ row.ExamineStatus }} </span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="最后更新时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.updatedAt }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="280">
          <template slot-scope="{row}">
            <el-button
              v-if="row.ExamineStatus === '已撤销' || row.ExamineStatus === '已驳回' && (row.ExamineStatus === '已完成' || row.MemType === '非组员')"

              size="mini"
              type="primary"
              @click="ReJoinDept(row)"
            >
              再次申请
            </el-button>
            <el-button
              v-if="row.ExamineStatus === '已驳回' || row.ExamineStatus === '未审核' || row.ExamineStatus === '处理中'"
              :disabled="row.memStatus ==='申请退出团队'"
              size="mini"
              type="danger"
              @click="CancelJoinDept(row)"
            >
              撤销
            </el-button>

            <el-button
              v-if="(row.ExamineStatus === '已完成' || row.MemType === '组员') && row.memStatus !== '申请退出团队' "
              size="mini"
              type="warning"
              @click="ExitDept(row)"
            >
              退出团队
            </el-button>

          </template>
        </el-table-column>

      </el-table>

    </div>

  </div>

</template>
<script>

import waves from '@/directive/waves' // waves directive
import { ExitDept, getDeptRELAUser, reJoinDept, UserCancelJoinDept } from '@/api/dept'

export default {
  name: 'ComplexTable',
  directives: { waves },

  data() {
    return {
      files: null,
      tableKey: 0,
      patents: null,
      ifquery: false,

      ifpatent: false,
      ifreport: false,
      ifshow1: false,
      patentitems: null,
      patentlist: [],
      patentid: 0,
      time1: '0:0:0',
      centerDialogVisible: false,
      a: 0,
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
      queryid: ''

    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getDeptRELAUser().then(response => {
        this.list = response.data.data
        for (var i = 0; i < this.list.length; i++) {
          if (this.list[i].UpdatedAt === '') this.list[i].UpdatedAt = '无'
        }
        this.listLoading = false
      })
    },
    CancelJoinDept(row) {
      UserCancelJoinDept(row.deptId).then(response => {
        this.$message({
          message: '撤销申请成功',
          type: 'success',
          duration: 1000
        })
        this.getList()
      })
    },
    ReJoinDept(row) {
      reJoinDept(row.deptId).then(response => {
        this.$message({
          message: '重新申请成功',
          type: 'success',
          duration: 1000
        })
        this.getList()
      })
    },
    ExitDept(row) {
      ExitDept(row.deptId).then(response => {
        this.$message({
          message: '申请退出成功',
          type: 'success',
          duration: 1000
        })
      })
    },

    back() {
      this.ifpatent = false
      this.ifquery = false
      this.queryid = ''
      this.getList()
    },
    flash() {
      this.getList()
      this.queryid = ''
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
