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

// modify password
export function updatePassword(data) {
  return request({
    url: '/user/pwd/set',
    method: 'put',
    data
  })
}

// reset password
export function resetPassword(data) {
  return request({
    url: '/user/pwd/reset',
    method: 'put',
    data
  })
}

export function createUser(data) {
  return request({
    url: '/sys-user',
    method: 'post',
    data
  })
}

// get all users
export function getAllUsers(params) {
  return request({
    url: '/sys-user',
    method: 'get',
    params
  })
}

export function getUserInfoByUserId(userId) {
  return request({
    url: `/sys-user/${userId}`,
    method: 'get'
  })
}

export function updateUserInfo(data) {
  return request({
    url: `/sys-user/`,
    method: 'put',
    data
  })
}

export function deleteAccount(userId) {
  return request({
    url: `/sys-user/${userId}`,
    method: 'delete'
  })
}
