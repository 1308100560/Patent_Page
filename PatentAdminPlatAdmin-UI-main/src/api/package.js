import request from '@/utils/request'

export function getPackageList(params) {
  return request({
    url: '/user-agent/package',
    method: 'get',
    params
  })
}

export function createPackage(data) {
  return request({
    url: `/user-agent/package`,
    method: 'post',
    data
  })
}

export function deletePackage(id) {
  return request({
    url: `/user-agent/package/${id}`,
    method: 'delete'
  })
}

export function updatePackage(id, data) {
  return request({
    url: `/user-agent/package/${id}`,
    method: 'put',
    data
  })
}

export function getPackage(id) {
  return request({
    url: `/user-agent/package/${id}`,
    method: 'get'
  })
}

export function addPatentToPackage(packageId, patentId, patent) {
  return request({
    url: `/user-agent/package/${packageId}/patent`,
    method: 'post',
    data: patent
  })
}

export function removePatentFromPackage(packageId, pnm) {
  return request({
    url: `/user-agent/package/${packageId}/patent/${pnm}`,
    method: 'delete'
  })
}

export function checkPatentToPackage(packageId, patentId, patent) {
  return request({
    url: `/user-agent/package/${packageId}/patent/${patentId}/isExist`,
    method: 'get',
    data: patent
  })
}

export function getPatentListByPackageId(packageId) {
  return request({
    url: `/user-agent/package/${packageId}/patent`,
    method: 'get'
  })
}

// 通过专利包id获得专利包中专利的用户关系图谱
export function getGraphByPackageId(packageId) {
  return request({
    url: `/user-agent/package/${packageId}/relationship`,
    method: 'get'
  })
}

// 通过专利包id获得专利包中专利的发明人关系图谱
export function getRelationGraphByPackageId(packageId) {
  return request({
    url: `/user-agent/package/${packageId}/graph/relation`,
    method: 'get'
  })
}

// 通过专利包id获得专利包中专利的技术图谱
export function getTechGraphByPackageId(packageId) {
  return request({
    url: `/user-agent/package/${packageId}/graph/tech`,
    method: 'get'
  })
}
