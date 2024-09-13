import request from '@/utils/request'

// 更新关注的专利描述
export function updateClaimPatentDescription(PNM, desc) {
  return request({
    url: `/user-agent/patent/claim/${PNM}/desc`, method: 'put', data: { desc }
  })
}

// 认领专利
export function claimPatent(patent) {
  return request({
    url: '/user-agent/patent/claim', method: 'post', data: patent
  })
}

// 取消认领专利
export function unClaimPatent(PNM) {
  return request({
    url: `/user-agent/patent/claim/${PNM}`, method: 'delete'
  })
}

// 获取认领专利列表
export function getClaimedPatents(data) {
  return request({
    url: '/user-agent/patent/claim/search', method: 'get', params: data
  })
}

// 获取所有专利
export function getAllPatents(data) {
  return request({
    url: '/user-agent/patent', method: 'get', params: data
  })
}

// 更新关注的专利描述
export function updateFocusPatentDescription(PNM, desc) {
  return request({
    url: `/user-agent/patent/focus/${PNM}/desc`, method: 'put', data: { desc }
  })
}

// 关注专利
export function focusPatent(patent) {
  return request({
    url: '/user-agent/patent/focus', method: 'post', data: patent
  })
}

// 取消关注专利
export function unFocusPatent(PNM) {
  return request({
    url: `/user-agent/patent/focus/${PNM}`, method: 'delete'
  })
}

// 获取关注专利列表
export function getFocusedPatents(data) {
  return request({
    url: '/user-agent/patent/focus/search', method: 'get', params: data
  })
}

// 通过userid获得用户关注专利的关系图谱1
export function getGraphByUserId() {
  return request({
    url: `/user-agent/patent/relationship`,
    method: 'get'
  })
}

// 专利详情
export function getPatentDetail(PNM) {
  return request({
    url: `/user-agent/auth-search/full`,
    method: 'post',
    data: { query: `PNM=${PNM}` }
  })
}

// 获取关注专利的发明人关系图谱
export function getFocusRelationGraph() {
  return request({
    url: `/user-agent/patent/focus/graph/relation`,
    method: 'get'
  })
}

// 获取关注专利的技术图谱
export function getFocusTechGraph() {
  return request({
    url: `/user-agent/patent/focus/graph/tech`,
    method: 'get'
  })
}

