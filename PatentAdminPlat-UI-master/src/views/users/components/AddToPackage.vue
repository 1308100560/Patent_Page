<template>
  <div>
    <el-dialog
      :modal-append-to-body="false"
      :title="`加入${type==='focus'?'关注':'专利'}组合`"
      :visible.sync="dialogTableVisible"
    >
      <el-form
        :model="packageAddForm"
        label-position="left"
        label-width="100px"
        size="small"
        style="margin: 40px 10px"
      >
        <el-form-item label="专利" size="small">
          <el-input
            v-model="packageAddForm.patentName"

            disabled

            size="small"
          />
        </el-form-item>
        <el-form-item label="备注" size="small">
          <el-input
            v-model="packageAddForm.desc"
            :autosize="{ minRows: 2, maxRows: 4}"
            :maxlength="200"
            size="small"
            type="textarea"
          />
        </el-form-item>
        <el-form-item :label="`${type==='focus'?'关注':'专利'}组合`" size="small">
          <el-select
            v-model="packageAddForm.packageId"
            :loading="loadingPackageList"
            :placeholder="`请选择${type==='focus'?'关注':'专利'}组合`"
            :remote-method="getPackageListRemote"
            filterable
            remote
            size="small"
            style="width: 100%"
          >
            <el-option
              v-for="p in packageList"
              :key="'p'+p.packageName"
              :label="p.packageName"
              :value="p.packageId"
            />
          </el-select>
        </el-form-item>
        <div style="display: flex;flex-direction: row;align-items: center;justify-content: flex-end">
          <div style="text-align: right; margin: 0">

            <el-button
              :disabled="patentPackageExist"
              :loading="loadingRelation"
              size="mini"
              style="padding: 10px 25px!important;height: 35px!important;"
              type="primary"
              @click="handleAddPatentToPackage()"
            >
              {{ patentPackageExist ? '已添加' : '添加' }}
            </el-button>
          </div>
        </div>
      </el-form>
    </el-dialog>
    <slot :showPopover="showPopover" name="content" />

  </div>
</template>
<script>
import { addPatentToPackage, checkPatentToPackage, getPackageList } from '@/api/package'

export default {
  props: {
    type: {
      type: String,
      default: 'claim'
    },
    patent: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      dialogTableVisible: false,
      packageAddForm: { packageId: '', patentName: '', patentId: '', desc: '' },
      patentPackageExist: false,
      loadingRelation: false,
      loadingPackageList: false,
      packageList: []
    }
  },
  watch: {
    'packageAddForm.packageId': function(val) {
      if (val) {
        this.loadingRelation = true
        const { patentId } = this.packageAddForm
        checkPatentToPackage(val, patentId).then(res => {
          this.loadingRelation = false
          this.patentPackageExist = res.data.data.existed
        })
      } else {
        this.patentPackageExist = false
      }
    }
  },
  methods: {
    showPopover() {
      this.dialogTableVisible = true
      this.loadPackageList()
      this.packageAddForm.patentId = this.patent.PNM
      this.packageAddForm.patentName = this.patent.TI
      this.packageAddForm.packageId = ''
      this.packageAddForm.desc = ''
    },
    loadPackageList() {
      getPackageList({ type: this.type }).then(res => {
        this.packageList = res.data.data.list
        this.loadingPackageList = false
      })
    },
    getPackageListRemote(query) {
      if (query !== '') {
        this.loadingPackageList = true
        getPackageList(query).then(res => {
          this.packageList = res.data.data.list
          this.loadingPackageList = false
        })
      } else {
        this.packageList = []
      }
    },

    handleAddPatentToPackage() {
      const { packageId, patentId } = this.packageAddForm
      addPatentToPackage(packageId, patentId, this.patent).then(res => {
        if (res.data.code === 200) {
          this.$message.success('添加成功')
          this.dialogVisible = false
          this.patentPackageExist = true
        } else {
          this.$message.error(res.data.msg)
        }
      })
    }
  }
}
</script>
