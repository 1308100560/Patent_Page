<template>
  <div class="container">
    <div class="my-card">
      <el-page-header content="竞争对手/行业人才" @back="$router.go(-1)" />
      <div class="row-center title">
        竞争对手/行业人才
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
      <div style="margin-top: 30px">
        <Chart v-if="query" ref="chartX" :chart="chart" :query="query" style="width: 100%" title="发明人排行榜" />
      </div>
    </div>

  </div>
</template>
<script>
import Chart from '@/views/users/search/components/Chart.vue'

export default {
  name: 'CompetitorSearch',
  components: { Chart },
  data() {
    return {
      selected: '1',
      query: '',
      fields: [{ 'cnName': 'ICG', 'fieldName': 'ICG' }],
      chart: {
        'id': '501',
        'name': '发明人排行榜',
        'active': false
      },
      form: [
        {
          id: 0,
          fieldName: 'ICG',
          value: '',
          operation: ''
        }
      ]
    }
  },
  mounted() {
    this.form = this.$route.query.q
    const query = []
    const first = this.form[0]
    query.push(`${first.fieldName}='${first.value}'`)
    for (const item of this.form.slice(1)) {
      query.push(`${item.operation} ${item.fieldName}='${item.value}'`)
    }
    this.query = query.join(' ')
  },
  methods: {
    onSubmit() {
      const query = []
      const first = this.form[0]
      query.push(`${first.fieldName}='${first.value}'`)
      for (const item of this.form.slice(1)) {
        query.push(`${item.operation} ${item.fieldName}='${item.value}'`)
      }
      this.query = query.join(' ')
      this.$refs.chartX.refresh()
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
