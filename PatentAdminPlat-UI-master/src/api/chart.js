import request from '@/utils/request'

export function getChartOption(chartId, data) {
  return request({
    url: `/user-agent/auth-search/charts/${chartId}`,
    method: 'post',
    data
  })
}
