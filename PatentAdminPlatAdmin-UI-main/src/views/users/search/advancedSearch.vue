<template>
  <div class="container">
    <div class="my-card">
      <el-page-header content="高级检索" @back="$router.go(-1)" />
      <div class="row-center title">
        高级检索
      </div>
      <div class="form radius">
        <el-form class="demo-form-inline" label-position="right" label-width="140px" size="small">
          <el-row :gutter="20">
            <div v-for="(item,index) in form" :key="'item'+item.id" class="filterRow">
              <div style="width: 100px;margin-right: 10px">
                <el-select
                  slot="append"
                  v-model="item.operation"
                  :disabled="index===0"
                  placeholder="AND"
                  size="small"
                  style="width: 80px"
                >
                  <el-option label="AND" value="and" />
                  <el-option label="OR" value="or" />
                  <el-option label="NOT" value="not" />
                </el-select>
              </div>
              <el-input v-model="item.value" class="input-with-select" placeholder="请输入内容" size="small">
                <el-select slot="prepend" v-model="item.fieldName" placeholder="请选择" style="width: 150px">
                  <el-option
                    v-for="field in fields"
                    :key="field.cnName"
                    :label="field.cnName"
                    :value="field.fieldName"
                  />
                </el-select>
              </el-input>
              <div class="operation">
                <el-button
                  :disabled="index===0"
                  icon="el-icon-minus"
                  plain
                  size="mini"
                  type="info"
                  @click="removeCondition(index)"
                />
                <el-button icon="el-icon-plus" plain size="mini" type="info" @click="addCondition" />
              </div>
            </div>

          </el-row>
        </el-form>
      </div>
      <div class="row-center" style="margin-top: 30px">
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
      selected: '1',
      fields: [{ 'cnName': '公开(公告)号', 'fieldName': 'PNM' }, {
        'cnName': '公开(公告)日',
        'fieldName': 'PD'
      }, { 'cnName': '名称摘要权利要求书', 'fieldName': 'TAC' }, {
        'cnName': '名称摘要',
        'fieldName': 'TA'
      }, { 'cnName': '名称', 'fieldName': 'TI' }, { 'cnName': '摘要', 'fieldName': 'ABST' }, {
        'cnName': '权利要求书',
        'fieldName': 'blank'
      }, { 'cnName': '说明书', 'fieldName': 'DESCR' }, {
        'cnName': '申请人，当前专利权人',
        'fieldName': 'CAS'
      }, { 'cnName': '地址', 'fieldName': 'AR' }, {
        'cnName': '申请(专利权)人',
        'fieldName': 'PA'
      }, { 'cnName': '申请人集合', 'fieldName': 'PATMS' }, {
        'cnName': '当前专利权人',
        'fieldName': 'CAS'
      }, { 'cnName': '当前专利权人集合', 'fieldName': 'CASTMS' }, {
        'cnName': '发明(设计)人',
        'fieldName': 'PINN'
      }, { 'cnName': '发明人集合', 'fieldName': 'INNTMS' }, {
        'cnName': '存活期',
        'fieldName': 'PERIOD'
      }, { 'cnName': '国省代码', 'fieldName': 'CO' }, { 'cnName': '有效性', 'fieldName': 'LV' }, {
        'cnName': '当前法律状态',
        'fieldName': 'CLS'
      }, { 'cnName': 'IPC/LOC', 'fieldName': 'IPC' }, { 'cnName': '主IPC/LOC', 'fieldName': 'PICS' }, {
        'cnName': 'CPC',
        'fieldName': 'CPC'
      }, { 'cnName': 'UC', 'fieldName': 'UC' }, { 'cnName': '优先权', 'fieldName': 'PR' }, {
        'cnName': '国民经济分类',
        'fieldName': 'blank'
      }, { 'cnName': '分案原申请号', 'fieldName': 'DAN' }, {
        'cnName': '专利代理机构',
        'fieldName': 'AGC'
      }, { 'cnName': '代理人', 'fieldName': 'AGT' }, { 'cnName': '国际公布', 'fieldName': 'IPN' }, {
        'cnName': '国际申请',
        'fieldName': 'IAN'
      }, { 'cnName': '进入国家日期', 'fieldName': 'DEN' }, {
        'cnName': '解密专利',
        'fieldName': 'blank'
      }, { 'cnName': '同族号', 'fieldName': 'TZH' }, {
        'cnName': '期限调整PTA',
        'fieldName': 'PTAD'
      }, { 'cnName': '药品名称', 'fieldName': 'PTEN' }, { 'cnName': '期限延长PTE', 'fieldName': 'PTED' }],
      form: [
        {
          id: 0,
          fieldName: 'PNM',
          value: '',
          operation: ''
        }
      ]
    }
  },
  methods: {
    onSubmit() {
      const query = []
      for (const item of this.form) {
        query.push(`${item.operation} ${item.fieldName}='${item.value}'`)
      }
      this.$router.push({ path: '/search/results', query: { q: query.join(' ') }})
    },
    addCondition() {
      const id = this.form[this.form.length - 1].id + 1
      this.form.push({
        id,
        fieldName: 'PNM',
        value: '',
        operation: 'and'
      })
    },
    removeCondition(index) {
      this.form.splice(index, 1)
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
  width: 1000px;
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

.filterRow {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-bottom: 15px;
}

.operation {
  display: flex;
  width: 130px;
  flex-direction: row;
  margin-left: 10px;
  justify-content: space-between;
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
