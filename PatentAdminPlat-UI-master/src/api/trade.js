import request from '@/utils/request'

export function getTradeList(params) {
  return request({
    url: `/user-agent/trading-records`,
    params
  })
}

export function createTrade(data) {
  return request({
    url: `/user-agent/trading-records`,
    method: 'post',
    data
  })
}

export function updateTrade(id, data) {
  return request({
    url: `/user-agent/trading-records/${id}`,
    method: 'put',
    data
  })
}

export function deleteTrade(id) {
  return request({
    url: `/user-agent/trading-records/${id}`,
    method: 'delete'
  })
}
