import request from '@/utils/request'

export function getChartOption(chartId, data) {
  return request({
    url: `/user-agent/search/charts/${chartId}`,
    method: 'post',
    data
  })
}
