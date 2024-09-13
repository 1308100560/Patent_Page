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

                <el-select slot="prepend" v-model="item.fieldName" filterable placeholder="请选择" style="width: 150px">
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
      fields: [{ 'cnName': '申请号', 'fieldName': 'AN' }, { 'cnName': '申请日', 'fieldName': 'AD' }, {
        'cnName': '申请年',
        'fieldName': 'ADY'
      }, { 'cnName': '公开(公告)号', 'fieldName': 'PNM' }, {
        'cnName': '公开(公告)日',
        'fieldName': 'PD'
      }, { 'cnName': '公开年', 'fieldName': 'PDY' }, { 'cnName': '公开类型', 'fieldName': 'KC' }, {
        'cnName': '优先权',
        'fieldName': 'PR'
      }, { 'cnName': '最早优先权日', 'fieldName': 'EPRD' }, {
        'cnName': '分案原申请号',
        'fieldName': 'DAN'
      }, { 'cnName': '名称', 'fieldName': 'TI' }, { 'cnName': '名称中', 'fieldName': 'TICN' }, {
        'cnName': '名称英',
        'fieldName': 'TIEN'
      }, { 'cnName': '名称日', 'fieldName': 'TIJP' }, { 'cnName': '名称OL', 'fieldName': 'TIOL' }, {
        'cnName': '摘要',
        'fieldName': 'ABST'
      }, { 'cnName': '摘要中', 'fieldName': 'ABSTCN' }, { 'cnName': '摘要英', 'fieldName': 'ABSTEN' }, {
        'cnName': '摘要日',
        'fieldName': 'ABSTJP'
      }, { 'cnName': '摘要OL', 'fieldName': 'ABSTOL' }, { 'cnName': '权利要求', 'fieldName': 'CLM' }, {
        'cnName': '主权项',
        'fieldName': 'CL'
      }, { 'cnName': '主权项字数', 'fieldName': 'CLZS' }, { 'cnName': '权项数', 'fieldName': 'CLMN' }, {
        'cnName': '独权数',
        'fieldName': 'ICLMN'
      }, { 'cnName': '说明书', 'fieldName': 'DESCR' }, {
        'cnName': '说明书页数',
        'fieldName': 'DPN'
      }, { 'cnName': '名称摘要', 'fieldName': 'TA' }, {
        'cnName': '名称摘要权利要求书',
        'fieldName': 'TAC'
      }, { 'cnName': '名称摘要权利要求书说明书', 'fieldName': 'TACD' }, {
        'cnName': '摘要权利要求书',
        'fieldName': 'AC'
      }, { 'cnName': '摘要权利要求书说明书', 'fieldName': 'ACD' }, {
        'cnName': '申请(专利权)人',
        'fieldName': 'PA'
      }, { 'cnName': '标准申请人', 'fieldName': 'SPATMS' }, {
        'cnName': '主申请人',
        'fieldName': 'PPA'
      }, { 'cnName': '申请人集合', 'fieldName': 'PATMS' }, {
        'cnName': '当前专利权人',
        'fieldName': 'CAS'
      }, { 'cnName': '标准当前专利权人', 'fieldName': 'SCASTMS' }, {
        'cnName': '当前主专利权人',
        'fieldName': 'PCAS'
      }, { 'cnName': '当前专利权人集合', 'fieldName': 'CASTMS' }, {
        'cnName': '预测专利权人',
        'fieldName': 'PAS'
      }, { 'cnName': '申请人地址', 'fieldName': 'AR' }, {
        'cnName': '当前专利权人地址',
        'fieldName': 'CAR'
      }, { 'cnName': '申请人（省市）', 'fieldName': 'ADDRP' }, {
        'cnName': '申请人（地市）',
        'fieldName': 'ADDRC'
      }, { 'cnName': '申请人（区县）', 'fieldName': 'ADDRDC' }, {
        'cnName': '当前专利权人（省市）',
        'fieldName': 'CASP'
      }, { 'cnName': '当前专利权人（地市）', 'fieldName': 'CASC' }, {
        'cnName': '当前专利权人（区县）',
        'fieldName': 'CASD'
      }, { 'cnName': '申请人数', 'fieldName': 'PAN' }, {
        'cnName': '申请人国家/地域',
        'fieldName': 'PACC'
      }, { 'cnName': '主申请人国家/地域', 'fieldName': 'PPACC' }, {
        'cnName': '发明（设计）人',
        'fieldName': 'INN'
      }, { 'cnName': '主发明人', 'fieldName': 'PINN' }, {
        'cnName': '发明人集合',
        'fieldName': 'INNTMS'
      }, { 'cnName': '发明人数', 'fieldName': 'INNN' }, { 'cnName': '分类号', 'fieldName': 'SIC' }, {
        'cnName': '主分类号',
        'fieldName': 'PIC'
      }, { 'cnName': '分类数', 'fieldName': 'SICN' }, {
        'cnName': 'IPC分类号',
        'fieldName': 'IPC'
      }, { 'cnName': 'IPC主分类号', 'fieldName': 'PIPC' }, {
        'cnName': 'IPC分类部',
        'fieldName': 'IPCS'
      }, { 'cnName': 'IPC大类', 'fieldName': 'IPCC' }, {
        'cnName': 'IPC小类',
        'fieldName': 'IPCSC'
      }, { 'cnName': 'IPC小类数', 'fieldName': 'IPCSCN' }, {
        'cnName': 'IPC大组',
        'fieldName': 'ICG'
      }, { 'cnName': 'IPC小组', 'fieldName': 'ICSG' }, { 'cnName': '主IPC部', 'fieldName': 'PICS' }, {
        'cnName': '主IPC大类',
        'fieldName': 'PICC'
      }, { 'cnName': '主IPC小类', 'fieldName': 'PICSC' }, {
        'cnName': '主IPC大组',
        'fieldName': 'PICG'
      }, { 'cnName': '主IPC小组', 'fieldName': 'PICSG' }, {
        'cnName': 'IPC分类数',
        'fieldName': 'IPCN'
      }, { 'cnName': 'IPC部数', 'fieldName': 'IPCBN' }, {
        'cnName': 'LOC分类号',
        'fieldName': 'LOC'
      }, { 'cnName': 'LOC主分类号', 'fieldName': 'PLOC' }, {
        'cnName': 'LOC分类数',
        'fieldName': 'LOCN'
      }, { 'cnName': '欧洲分类号', 'fieldName': 'SEC' }, {
        'cnName': '欧洲主分类号',
        'fieldName': 'PEC'
      }, { 'cnName': '欧洲分类数', 'fieldName': 'ECN' }, {
        'cnName': 'CPC分类号',
        'fieldName': 'CPC'
      }, { 'cnName': 'CPC主分类号', 'fieldName': 'PCPC' }, {
        'cnName': 'CPC分类数',
        'fieldName': 'CPCN'
      }, { 'cnName': 'CPC部', 'fieldName': 'CPCS' }, { 'cnName': 'CPC大类', 'fieldName': 'CPCC' }, {
        'cnName': 'CPC小类',
        'fieldName': 'CPCSC'
      }, { 'cnName': 'CPC大组', 'fieldName': 'CPCG' }, { 'cnName': 'CPC小组', 'fieldName': 'CPCSG' }, {
        'cnName': '主CPC部',
        'fieldName': 'PCPCS'
      }, { 'cnName': '主CPC大类', 'fieldName': 'PCPCC' }, {
        'cnName': '主CPC小类',
        'fieldName': 'PCPCSC'
      }, { 'cnName': '主CPC大组', 'fieldName': 'PCPCG' }, {
        'cnName': '主CPC小组',
        'fieldName': 'PCPCSG'
      }, { 'cnName': 'C-sets', 'fieldName': 'CPCSETS' }, {
        'cnName': 'GBC分类号',
        'fieldName': 'GBC'
      }, { 'cnName': 'GBC分类数', 'fieldName': 'GBCN' }, { 'cnName': 'GBC门类', 'fieldName': 'GBC1' }, {
        'cnName': 'GBC大类',
        'fieldName': 'GBC2'
      }, { 'cnName': 'GBC中类', 'fieldName': 'GBC3' }, {
        'cnName': 'GBC小类',
        'fieldName': 'GBC4'
      }, { 'cnName': 'GBC主分类号', 'fieldName': 'PGBC' }, {
        'cnName': '主GBC门类',
        'fieldName': 'PGBC1'
      }, { 'cnName': '主GBC大类', 'fieldName': 'PGBC2' }, {
        'cnName': '主GBC中类',
        'fieldName': 'PGBC3'
      }, { 'cnName': '主GBC小类', 'fieldName': 'PGBC4' }, {
        'cnName': 'EINDC分类号',
        'fieldName': 'EINDC'
      }, { 'cnName': 'EINDC分类数', 'fieldName': 'EINDCN' }, {
        'cnName': 'EINDC门类',
        'fieldName': 'EINDC1'
      }, { 'cnName': 'EINDC大类', 'fieldName': 'EINDC2' }, {
        'cnName': 'UC分类号',
        'fieldName': 'UC'
      }, { 'cnName': 'UC主分类号', 'fieldName': 'PUC' }, {
        'cnName': 'UC分类数',
        'fieldName': 'UCN'
      }, { 'cnName': 'FI分类号', 'fieldName': 'FIC' }, {
        'cnName': 'FI主分类号',
        'fieldName': 'PFIC'
      }, { 'cnName': 'FI分类数', 'fieldName': 'FICN' }, {
        'cnName': 'FTERM分类号',
        'fieldName': 'FTERM'
      }, { 'cnName': 'FTERM主分类号', 'fieldName': 'PFTERM' }, {
        'cnName': 'FTERM分类数',
        'fieldName': 'FTERMN'
      }, { 'cnName': 'DPI', 'fieldName': 'IDX' }, {
        'cnName': '技术价值(DPIT)',
        'fieldName': 'DPIT'
      }, { 'cnName': '法律价值(DPIL)', 'fieldName': 'DPIL' }, {
        'cnName': '经济价值(DPIE)',
        'fieldName': 'DPIE'
      }, { 'cnName': '市场价值(DPIM)', 'fieldName': 'DPIM' }, {
        'cnName': '战略价值(DPIS)',
        'fieldName': 'DPIS'
      }, { 'cnName': '国省代码', 'fieldName': 'CO' }, {
        'cnName': '实审公告日',
        'fieldName': 'ED'
      }, { 'cnName': '授权公告日', 'fieldName': 'GD' }, {
        'cnName': '权利失效日',
        'fieldName': 'QLFQR'
      }, { 'cnName': '专利代理机构', 'fieldName': 'AGC' }, {
        'cnName': '代理人',
        'fieldName': 'AGT'
      }, { 'cnName': '引证专利', 'fieldName': 'REFP' }, {
        'cnName': '引证专利数',
        'fieldName': 'REFPN'
      }, { 'cnName': '引用专利国别/地域数', 'fieldName': 'REFPCN' }, {
        'cnName': '引用非专利文献数',
        'fieldName': 'REFNPN'
      }, { 'cnName': '被引证数', 'fieldName': 'REFBYN' }, {
        'cnName': '被审查员引证数',
        'fieldName': 'REFBSCYN'
      }, { 'cnName': '国际申请', 'fieldName': 'IAN' }, {
        'cnName': '国际申请日',
        'fieldName': 'IAD'
      }, { 'cnName': '国际公布', 'fieldName': 'IPN' }, {
        'cnName': '国际公开日',
        'fieldName': 'IPD'
      }, { 'cnName': '进入国家日期', 'fieldName': 'DEN' }, {
        'cnName': '检索领域',
        'fieldName': 'SF'
      }, { 'cnName': '法律状态履历', 'fieldName': 'LS' }, {
        'cnName': '当前法律状态',
        'fieldName': 'CLS'
      }, { 'cnName': '当前法律状态公告日', 'fieldName': 'CLSD' }, {
        'cnName': '海关备案专利',
        'fieldName': 'CFP'
      }, { 'cnName': '受理局', 'fieldName': 'RO' }, { 'cnName': '转让人', 'fieldName': 'PRCP' }, {
        'cnName': '受让人',
        'fieldName': 'POCP'
      }, { 'cnName': '转让类型', 'fieldName': 'TSFT' }, {
        'cnName': '转让生效日',
        'fieldName': 'TEFD'
      }, { 'cnName': '转让法律状态公告日', 'fieldName': 'TLSD' }, {
        'cnName': '许可人',
        'fieldName': 'TRSF'
      }, { 'cnName': '被许可人', 'fieldName': 'ASSI' }, {
        'cnName': '许可类型',
        'fieldName': 'LICT'
      }, { 'cnName': '许可生效日', 'fieldName': 'LEFD' }, {
        'cnName': '许可变更日',
        'fieldName': 'LCHD'
      }, { 'cnName': '许可解除日', 'fieldName': 'LDSD' }, {
        'cnName': '许可法律状态公告日',
        'fieldName': 'LSBD'
      }, { 'cnName': '许可合同状态', 'fieldName': 'LSOC' }, {
        'cnName': '许可合同备案号',
        'fieldName': 'LCRN'
      }, { 'cnName': '出质人', 'fieldName': 'PLDO' }, {
        'cnName': '质权人',
        'fieldName': 'PLDE'
      }, { 'cnName': '质押保全类型', 'fieldName': 'PLDT' }, {
        'cnName': '质押生效日',
        'fieldName': 'PEFD'
      }, { 'cnName': '质押变更日', 'fieldName': 'PCHD' }, {
        'cnName': '质押解除日',
        'fieldName': 'PDSD'
      }, { 'cnName': '质押法律状态公告日', 'fieldName': 'PSBD' }, {
        'cnName': '质押合同状态',
        'fieldName': 'PSOC'
      }, { 'cnName': '质押合同登记号', 'fieldName': 'PCRN' }, {
        'cnName': '有效性',
        'fieldName': 'LV'
      }, { 'cnName': '同族/同族号', 'fieldName': 'TZH' }, {
        'cnName': '布局国家/地域',
        'fieldName': 'SFMLC'
      }, { 'cnName': 'INNOJOY同族', 'fieldName': 'ITZH' }, {
        'cnName': 'INNOJOY同族数',
        'fieldName': 'IFMLN'
      }, { 'cnName': 'INNOJOY同族布局国家/地域', 'fieldName': 'IFMLC' }, {
        'cnName': 'INNOJOY布局国家/地域数',
        'fieldName': 'IFMLCN'
      }, { 'cnName': 'PCT国际申请', 'fieldName': 'SFMLC' }, {
        'cnName': '三方专利（美日欧）',
        'fieldName': 'IFMLC'
      }, { 'cnName': '标准专利', 'fieldName': 'BZ' }, { 'cnName': '专利奖', 'fieldName': 'WINPS' }, {
        'cnName': '获奖年份',
        'fieldName': 'WINPDY'
      }, { 'cnName': '复审决定', 'fieldName': 'EAR' }, { 'cnName': '无效决定', 'fieldName': 'CRAI' }, {
        'cnName': 'ETSI标准',
        'fieldName': 'ETSIBZ'
      }, { 'cnName': '标准号', 'fieldName': 'ETSIN' }, {
        'cnName': '标准申报公司',
        'fieldName': 'BZC'
      }, { 'cnName': '标准申报日期', 'fieldName': 'BZD' }, {
        'cnName': '行业领域',
        'fieldName': 'HYLY'
      }, { 'cnName': '申请人类型', 'fieldName': 'SQRLX' }, {
        'cnName': '转让次数',
        'fieldName': 'TRA'
      }, { 'cnName': '许可次数', 'fieldName': 'LIC' }, { 'cnName': '质押次数', 'fieldName': 'PLE' }, {
        'cnName': '无效次数',
        'fieldName': 'INV'
      }, { 'cnName': '存活期', 'fieldName': 'PERIOD' }, {
        'cnName': '剩余有效期',
        'fieldName': 'RET'
      }, { 'cnName': '期限调整PTA', 'fieldName': 'PTAD' }, {
        'cnName': '期限延长PTE',
        'fieldName': 'PTED'
      }, { 'cnName': '药品名称', 'fieldName': 'PTEN' }, { 'cnName': '药品品牌', 'fieldName': 'PB' }, {
        'cnName': '药品成分',
        'fieldName': 'PI'
      }, { 'cnName': '药品剂型', 'fieldName': 'PF' }, { 'cnName': '药品用法', 'fieldName': 'PRT' }, {
        'cnName': '药品用量',
        'fieldName': 'PST'
      }],
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
