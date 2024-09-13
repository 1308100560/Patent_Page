import request from '@/utils/request'

export function getDashboardData() {
  return request({
    url: `/user-agent/dashboard`,
    method: 'get'
  })
}
