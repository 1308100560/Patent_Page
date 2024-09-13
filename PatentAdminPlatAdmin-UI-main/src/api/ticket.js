import request from '@/utils/request'

// 获取工单列表
export function getTicketList(query) {
  return request({
    url: `/admin-agent/tickets`,
    method: 'get',
    params: query
  })
}

// 撤销工单
export function closeTicket(ticketId) {
  return request({
    url: `/admin-agent/tickets/${ticketId}/close`,
    method: 'put'
  })
}

export function updateTicket(ticketId, data) {
  return request({
    url: `/admin-agent/tickets/${ticketId}?type=report`,
    method: 'put',
    data
  })
}

export function finishTicket(ticketId, data) {
  return request({
    url: `/admin-agent/tickets/${ticketId}/finish`,
    method: 'put',
    data
  })
}

export function getTicket(ticketId) {
  return request({
    url: `/admin-agent/tickets/${ticketId}`,
    method: 'get'
  })
}
