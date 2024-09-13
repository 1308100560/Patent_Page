import request from '@/utils/request'

export function getDashboardData() {
  return request({
    url: `/user-agent/dashboard`,
    method: 'get'
  })
}

export function getDashboardDataByPackage(packageId) {
  return request({
    url: `/user-agent/dashboard/package/${packageId}`,
    method: 'get'
  })
}
