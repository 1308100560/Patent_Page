import request from '@/utils/request'

export function getAchievementList(params) {
  return request({
    url: '/user-agent/tech-achieves',
    method: 'get',
    params
  })
}

export function createAchievement(data) {
  return request({
    url: '/user-agent/tech-achieves',
    method: 'post',
    data
  })
}

export function updateAchievement(id, data) {
  return request({
    url: '/user-agent/tech-achieves/' + id,
    method: 'put',
    data
  })
}

export function deleteAchievement(id) {
  return request({
    url: '/user-agent/tech-achieves/' + id,
    method: 'delete'
  })
}
