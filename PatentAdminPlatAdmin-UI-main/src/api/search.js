import request from '@/utils/request'
import { getToken } from '@/utils/auth'

export function searchSimple(searchForm) {
  const token = getToken()
  let path = '/user-agent/search'
  if (token) {
    path = '/user-agent/auth-search'
  }
  return request({
    url: path, method: 'post', data: searchForm
  })
}

// 记录用户搜索表达式
export function addQuery(query) {
  return request({
    url: '/user-agent/auth-search/queries', method: 'post', data: query
  })
}

// 删除用户搜索表达式
export function deleteQuery(queryId) {
  return request({
    url: `/user-agent/auth-search/queries/${queryId}`, method: 'delete'
  })
}

// 获取用户表达式列表
export function getQueryList(data) {
  return request({
    url: '/user-agent/auth-search/queries/search', method: 'get', params: data
  })
}

