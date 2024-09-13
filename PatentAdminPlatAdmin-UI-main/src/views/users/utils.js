import Docxtemplater from 'docxtemplater'
import PizZip from 'pizzip'
import { saveAs } from 'file-saver'
import PizZipUtils from 'pizzip/utils'

function loadFile(url, callback) {
  PizZipUtils.getBinaryContent(url, callback)
}

export function getTagColor(tag) {
  const tags = ['授权', '部分无效', '权利恢复', '公开', '实审', '全部无效', '届满', '终止', '撤回', '驳回', '未缴年费', '主动放弃', '视为放弃', '避重放弃']
  const colors = ['#67C23A', '#E6A23C', '#F56C6C', '#409EFF', '#67C23A', '#F56C6C', '#E6A23C', '#909399', '#909399', '#F56C6C', '#F56C6C', '#909399', '#909399', '#909399']
  return colors[tags.indexOf(tag)]
}

export function saveFile(data, name) {
  const urlObject = window.URL || window.webkitURL || window
  const export_blob = new Blob([data])
  const save_link = document.createElementNS('http://www.w3.org/1999/xhtml', 'a')
  save_link.href = urlObject.createObjectURL(export_blob)
  save_link.download = name
  save_link.click()
}

export function downloadFile(url, name) {
  const save_link = document.createElementNS('http://www.w3.org/1999/xhtml', 'a')
  save_link.href = url
  save_link.download = name
  save_link.click()
}

function renderDoc(data, fileName) {
  const self = this
  loadFile('http://10.112.138.178:8000/static/docs/novelty_report_template.docx', function(
    error,
    content
  ) {
    if (error) {
      throw error
    }
    const zip = new PizZip(content)
    const doc = new Docxtemplater().loadZip(zip)
    doc.setData(data)
    try {
      // render the document (replace all occurences of {first_name} by John, {last_name} by Doe, ...)
      doc.render()
    } catch (error) {
      // The error thrown here contains additional information when logged with JSON.stringify (it contains a properties object containing all suberrors).
      this.errorHandler(error)
    }
    const out = doc.getZip().generate({
      type: 'blob',
      mimeType:
        'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
    })
    saveAs(out, fileName + '.docx')
    self.loading = false
  })
}

export function generateNoveltyReport(fileName, data) {
  data = JSON.parse(data)
  data.hasKitty = fileName.length > 15
  data.QUERY_WORD = data.QUERY_WORD.split('\n')
  data.QUERY_EXPRESSION = data.QUERY_EXPRESSION.split('\n')
  data.SEARCH_RESULT = data.SEARCH_RESULT.split('\n')
  data.CONCLUSION = data.CONCLUSION.split('\n')
  renderDoc(data, fileName)
}
