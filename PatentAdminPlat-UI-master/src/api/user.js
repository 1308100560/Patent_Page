import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/user/getinfo',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}

export function getCaptcha() {
  return request({
    url: '/captcha',
    method: 'get'
  })
}

export function register(data) {
  return request({
    url: '/register',
    method: 'post',
    data
  })
}

export function updateInfo(data) {
  return request({
    url: '/user/updateinfo',
    method: 'put',
    data
  })
}

export function modifyPassword(data) {
  return request({
    url: '/user/pwd/set',
    method: 'put',
    data
  })
}
