import request from '@/utils/request'

export function getTradeList(params) {
  return request({
    url: `/admin-agent/trading-records`,
    params
  })
}

export function createTrade(data) {
  return request({
    url: `/admin-agent/trading-records`,
    method: 'post',
    data
  })
}

export function updateTrade(id, data) {
  return request({
    url: `/admin-agent/trading-records/${id}`,
    method: 'put',
    data
  })
}

export function deleteTrade(id) {
  return request({
    url: `/admin-agent/trading-records/${id}`,
    method: 'delete'
  })
}
