import request from '@/utils/request'

export function getTraceList(params) {
  return request({
    url: `/user-agent/tracing/logs`,
    method: 'get',
    params
  })
}
