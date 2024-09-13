import request from '@/utils/request'

// 用户通过reportId查看对应专利详情（x）
export function userGetReportPatents(reportId) {
  return request({
    url: `/user-agent/report/patent/${reportId} `,
    method: 'get'
  })
}

// 用户申请报告
export function ApplyReport(form) {
  return request({
    url: '/user-agent/report', method: 'post', data: form
  })
}

// 用户撤销申请
export function cancelReport(reportId) {
  return request({
    url: `/user-agent/report/cancel/${reportId}`,
    method: 'put'
  })
}

// 用户重新申请
export function reAppReport(reportId) {
  return request({
    url: `/user-agent/report/reApp/${reportId}`,
    method: 'put'
  })
}

// 用户报告
export function userReportList() {
  return request({
    url: `/user-agent/reports`,
    method: 'get'
  })
}

// 用户通过patentId获取报告列表
export function userGetReportListByPaId(patentId) {
  return request({
    url: `/user-agent/report/reportList/${patentId}`,
    method: 'get'
  })
}

// 用户query报告
export function userGetReportById(reportId) {
  return request({
    url: `/user-agent/report/query/${reportId} `, method: 'get'
  })
}

// 用户按照类型查看报告
export function userGetReportListByType(type) {
  return request({
    url: `/user-agent/report/${type} `,
    method: 'get'
  })
}

// 获取当前时间
export function getNowTime() {
  var time = new Date()
  return time.toLocaleDateString() + ' ' + time.toLocaleTimeString()
}

export function getNoveltyReport(data) {
  return request({
    url: `/user-agent/report/novelty `,
    method: 'post',
    data: data
  })
}

// 报告工单提交
export function reportTicketSubmit(data) {
  return request({
    url: `/user-agent/tickets?type=report`,
    method: 'post',
    data: data
  })
}
