<template>
  <div class="container">
    <div class="my-card">
      <el-page-header content="表格检索" @back="$router.go(-1)" />
      <div class="row-center title">
        表格检索
      </div>
      <div class="form radius">
        <el-form class="demo-form-inline" label-position="right" label-width="140px" size="small" style="width: 1000px">
          <el-row :gutter="10">
            <el-col v-for="field in fields" :key="field.cnName" :span="12">
              <el-form-item size="small">
                <div slot="label" style="font-size: 10px!important;">{{ field.cnName }}</div>
                <el-input v-model="field.value" :name="field.fieldName" />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <div class="row-center" style="margin-top: 20px">
        <el-button size="small">清空</el-button>
        <el-button size="small" type="primary" @click="onSubmit">检索</el-button>
      </div>
    </div>

  </div>
</template>
<script>
export default {
  name: 'AdvancedSearch',
  data() {
    return {
      fields: [
        { 'cnName': '公开(公告)号', 'fieldName': 'PNM', 'value': '' }, {
          'cnName': '公开(公告)日',
          'fieldName': 'PD',
          'value': ''
        }, { 'cnName': '名称摘要权利要求书', 'fieldName': 'TAC', 'value': '' }, {
          'cnName': '名称摘要',
          'fieldName': 'TA',
          'value': ''
        }, { 'cnName': '名称', 'fieldName': 'TI', 'value': '' }, {
          'cnName': '摘要',
          'fieldName': 'ABST',
          'value': ''
        }, { 'cnName': '权利要求书', 'fieldName': 'blank', 'value': '' }, {
          'cnName': '说明书',
          'fieldName': 'DESCR',
          'value': ''
        }, { 'cnName': '申请人，当前专利权人', 'fieldName': 'CAS', 'value': '' }, {
          'cnName': '地址',
          'fieldName': 'AR',
          'value': ''
        }, { 'cnName': '申请(专利权)人', 'fieldName': 'PA', 'value': '' }, {
          'cnName': '申请人集合',
          'fieldName': 'PATMS',
          'value': ''
        }, { 'cnName': '当前专利权人', 'fieldName': 'CAS', 'value': '' }, {
          'cnName': '当前专利权人集合',
          'fieldName': 'CASTMS',
          'value': ''
        }, { 'cnName': '发明(设计)人', 'fieldName': 'PINN', 'value': '' }, {
          'cnName': '发明人集合',
          'fieldName': 'INNTMS',
          'value': ''
        }, { 'cnName': '存活期', 'fieldName': 'PERIOD', 'value': '' }, {
          'cnName': '国省代码',
          'fieldName': 'CO',
          'value': ''
        }, { 'cnName': '有效性', 'fieldName': 'LV', 'value': '' }, {
          'cnName': '当前法律状态',
          'fieldName': 'CLS',
          'value': ''
        }, { 'cnName': 'IPC/LOC', 'fieldName': 'IPC', 'value': '' }, {
          'cnName': '主IPC/LOC',
          'fieldName': 'PICS',
          'value': ''
        }, { 'cnName': 'CPC', 'fieldName': 'CPC', 'value': '' }, {
          'cnName': 'UC',
          'fieldName': 'UC',
          'value': ''
        }, { 'cnName': '优先权', 'fieldName': 'PR', 'value': '' }, {
          'cnName': '国民经济分类',
          'fieldName': 'blank',
          'value': ''
        }, { 'cnName': '分案原申请号', 'fieldName': 'DAN', 'value': '' }, {
          'cnName': '专利代理机构',
          'fieldName': 'AGC',
          'value': ''
        }, { 'cnName': '代理人', 'fieldName': 'AGT', 'value': '' }, {
          'cnName': '国际公布',
          'fieldName': 'IPN',
          'value': ''
        }, { 'cnName': '国际申请', 'fieldName': 'IAN', 'value': '' }, {
          'cnName': '进入国家日期',
          'fieldName': 'DEN',
          'value': ''
        }, { 'cnName': '解密专利', 'fieldName': 'blank', 'value': '' }, {
          'cnName': '同族号',
          'fieldName': 'TZH',
          'value': ''
        }, { 'cnName': '期限调整PTA', 'fieldName': 'PTAD', 'value': '' }, {
          'cnName': '药品名称',
          'fieldName': 'PTEN',
          'value': ''
        }, { 'cnName': '期限延长PTE', 'fieldName': 'PTED', 'value': '' }]
    }
  },
  methods: {
    onSubmit() {
      const query = []
      for (const field of this.fields) {
        if (field.value) {
          query.push(`${field.fieldName}='${field.value}'`)
        }
      }
      this.$router.push({ path: '/search/results', query: { q: query.join(' and ') }})
    }
  }
}
</script>
<style scoped>
.container {
  padding: 15px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.form {
  padding: 20px;
  background-clip: border-box;
  border: 1px solid #f6f6f6;
  color: #666;
}

.row-center {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.title {
  color: #17233d;
  font-weight: bold;
  margin: 10px 0;
  font-size: 20px;
  cursor: pointer;
}

/deep/ .el-form-item {
  margin-bottom: 5px !important;
}
</style>
