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
          v-if="ifapply || ifmember || ifupload"
          v-waves
          class="filter-item"
          icon="el-icon-d-arrow-left"
          size="small"
          type="primary"
          @click="back"
        >
          返回
        </el-button>
        <el-button
          v-if="ifbtn"
          v-waves
          class="filter-item"
          icon="el-icon-thumb"
          size="small"
          type="danger"
          @click="UserJoinDept()"
        >
          申请加入团队
        </el-button>

      </div>
      <input v-show="false" ref="uploadInput" type="file">

      <el-table
        v-if="ifmember === false && ifapply === false"
        :key="tableKey"
        v-loading="listLoading"
        :data="list"
        :stripe="true"
        border
        fit
        highlight-current-row
        style="width: 100%;
          border-radius: 10px!important;"
      >

        <el-table-column align="center" label="ID" prop="团队id" sortable="custom" width="100">
          <template slot-scope="{row}">
            <span>{{ row.deptId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="团队名称" width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.deptName }}</span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="申请时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.CreatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="团队状态" width="90px">
          <template slot-scope="{row}">
            <span class="link-type"> {{ row.deptStatus }} </span>
          </template>

        </el-table-column>

        <el-table-column align="center" label="最后更新时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.UpdatedAt }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="团队简介" min-width="150px">

          <template slot-scope="{row}">
            <span>{{ row.deptProperties }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="研究方向" width="110px">

          <template slot-scope="{row}">
            <span>{{ row.researchInterest }}</span>
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="成员详情" width="150">
          <template slot-scope="{row}">
            <el-button
              class="detail"
              icon="el-icon-view"
              size="mini"
              type="warning"
              @click="getMembers(row.deptId)"
            >
              查看
            </el-button>
          </template>
        </el-table-column>

        <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="180">
          <template slot-scope="{row}">
            <el-button size="mini" type="danger" @click="ShowDeptApply(row)">
              详情
            </el-button>

          </template>
        </el-table-column>

      </el-table>

      <el-table
        v-if="ifmember && !ifapply"
        :key="tableKey"
        v-loading="listLoading"
        :data="memberlist"
        border
        fit
        style="width: 100%;  border-radius: 10px!important;"
      >

        <el-table-column align="center" label="团队ID" prop="id" sortable="custom" width="100">
          <template slot-scope="{row}">
            <span>{{ row.deptId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="用户ID" width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.userId }}</span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="用户身份" width="110px">
          <template slot-scope="{row}">
            <span>{{ row.MemType }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请内容" min-width="110px">
          <template slot-scope="{row}">
            <span>{{ row.memStatus }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.createdAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="审核状态" width="90px">
          <template slot-scope="{row}">
            <span class="link-type"> {{ row.ExamineStatus }} </span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="最后更新时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.updatedAt }}</span>
          </template>
        </el-table-column>
      </el-table>

      <el-table
        v-if="ifapply"
        :key="tableKey"
        v-loading="listLoading"
        :data="judgelist"
        border
        fit
        style="width: 100%;  border-radius: 10px!important;"
      >

        <el-table-column align="center" label="团队ID" prop="id" sortable="custom" width="100">
          <template slot-scope="{row}">
            <span>{{ row.deptId }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="用户ID" width="150px">
          <template slot-scope="{row}">
            <span class="link-type">{{ row.userId }}</span>
          </template>

        </el-table-column>
        <el-table-column align="center" label="用户身份" width="110px">
          <template slot-scope="{row}">
            <span>{{ row.MemType }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请内容" min-width="110px">
          <template slot-scope="{row}">
            <span>{{ row.memStatus }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="申请时间" sortable="custom" width="180px">
          <template slot-scope="{row}">
            <span>{{ row.createdAt }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="审核状态" width="90px">
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
              :disabled="row.memStatus ==='申请退出团队'"
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
              v-if="(row.ExamineStatus === '已完成' || row.MemType === '组员') && row.memStatus !=='申请退出团队'"
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
import {
  ExitDept,
  getDeptList,
  getDeptRELAListbyId,
  reJoinDept,
  RelaListOfUserId,
  UserCancelJoinDept,
  UserJoinDept
} from '@/api/dept'

export default {
  name: 'ComplexTable',
  directives: { waves },

  data() {
    return {
      files: null,
      tableKey: 0,
      memberlist: null,
      ifquery: false,
      ifupload: false,
      ifmember: false,
      ifreport: false,
      ifbtn: false,
      applyid: 0,
      ifapply: false,
      time1: '0:0:0',
      judgelist: null,
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
      queryid: '',
      form: {
        deptName: '',
        deptProperties: '',
        researchInterest: ''
        // deptStatus: false
        // CreatedAt: this.getNowTime()
      },
      formLabelWidth: '120px'

    }
  },
  created() {
    this.getList()
  },
  methods: {

    getList() {
      this.listLoading = true
      getDeptList().then(response => {
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
        RelaListOfUserId(row.deptId).then(response => {
          this.ifbtn = false
          this.judgelist = response.data.data
        })
      })
    },
    ReJoinDept(row) {
      reJoinDept(row.deptId).then(response => {
        this.$message({
          message: '重新申请成功',
          type: 'success',
          duration: 1000
        })
        RelaListOfUserId(row.deptId).then(response => {
          this.ifbtn = false
          this.judgelist = response.data.data
        })
      })
    },
    ExitDept(row) {
      ExitDept(row.deptId).then(response => {
        this.$message({
          message: '申请退出成功',
          type: 'success',
          duration: 1000
        })
        RelaListOfUserId(row.deptId).then(response => {
          this.ifbtn = false
          this.judgelist = response.data.data
        })
      })
    },
    getMembers(id) {
      this.ifmember = true
      getDeptRELAListbyId(id).then(response => {
        this.memberlist = response.data.data
        for (var i = 0; i < this.memberlist.length; i++) {
          if (this.memberlist[i].UpdatedAt === '') this.list[i].UpdatedAt = '无'
        }
        this.listLoading = false
      })
    },
    UserJoinDept() {
      UserJoinDept(this.applyid).then(response => {
        RelaListOfUserId(this.applyid).then(response => {
          this.ifbtn = false
          this.judgelist = response.data.data
          this.$message({
            message: '成功申请加入',
            type: 'success',
            duration: 1000
          })
        })
      })
    },
    ShowDeptApply(row) { // 点击详情
      this.applyid = row.deptId
      this.ifapply = true
      RelaListOfUserId(row.deptId).then(response => {
        this.judgelist = response.data.data
        console.log(this.judgelist)
        if (this.judgelist.length === 0) {
          this.ifbtn = true
        } else {
          this.ifbtn = false
        }
      })
    },
    back() {
      this.ifmember = false
      this.ifapply = false
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
