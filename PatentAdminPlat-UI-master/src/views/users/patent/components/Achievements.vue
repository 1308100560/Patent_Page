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
        <el-button
          class="filter-item"
          icon="el-icon-plus"
          size="small"
          style="margin-left: 10px"
          @click="showAchievements()"
        >
          科技成果
        </el-button>

      </div>

    </div>
    <el-table
      :data="list"
      style="width: 100%"
    >
      <el-table-column
        fixed
        label="ID"
        prop="id"
        width="80"
      />
      <el-table-column
        label="名称"
        prop="name"
        width="120"
      >
        <template slot-scope="{row}">
          {{ row.properties.name }}
        </template>
      </el-table-column>
      <el-table-column
        label="详细内容"
        min-width="120"
        prop="province"
      >
        <template slot-scope="{row}">
          {{ row.properties.content }}
        </template>
      </el-table-column>

      <el-table-column
        fixed="right"
        label="操作"
        width="130"
      >
        <template slot-scope="{row}">
          <div class="actions" style="display: flex;flex-direction: row;justify-content: space-between">

            <el-tooltip class="item" content="编辑" effect="dark" placement="top">
              <el-button
                icon="el-icon-edit"
                size="mini"
                type="primary"
                @click="showAchievementEditDescForm(row.id,row.properties)"
              />
            </el-tooltip>
            <el-tooltip class="item" content="申请报告" effect="dark" placement="top">
              <router-link to="/report/generate/novelty">
                <el-button icon="el-icon-sunset" />
              </router-link>
            </el-tooltip>
            <el-tooltip class="item" content="删除" effect="dark" placement="top">
              <el-button icon="el-icon-delete" type="danger" @click="handleDeleteAchievement(row)" />
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
  </div>
</template>
<script>
import { deleteAchievement, getAchievementList } from '@/api/achievements'

export default {
  props: {
    showAchievements: {
      type: Function,
      default: () => {
      }
    },
    showAchievementEditDescForm: {
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
      listQuery: {
        pageIndex: 1,
        pageSize: 10,
        query: '',
        needEvalResult: true
      },
      list: [],
      results: {}
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
    handleDeleteAchievement(row) {
      this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteAchievement(row.id).then(res => {
          this.$message({
            type: 'success',
            message: '删除成功!'
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
    getList() {
      getAchievementList(this.listQuery).then(res => {
        this.results = res.data.data
        this.list = res.data.data.list
        this.list.map(item => {
          item.properties = item.properties !== 'null' ? JSON.parse(item.properties) : {}
        })
      })
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
