import request from '@/utils/request'

/*
  * @param {String} scope : 'package' , 'all-claimed' , 'all-focused'
  * @param {String} type : 'rela' or 'tech'
  * @param {String} packageId
 */
export function getGraphData(params) {
  return request({
    url: '/user-agent/patent/graph',
    method: 'get',
    params
  })
}
