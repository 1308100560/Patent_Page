import request from '@/utils/request'

// 用户查看所有团队
export function getDeptList() {
  return request({
    url: '/dept',
    method: 'get'
  })
}

//  用户列表自己加入/管理的部门信息 (我的申请)
export function getDeptRELAUser() {
  return request({
    url: '/user-agent/dept/relaList',
    method: 'get'
  })
}

//  用户在该部门的状态
export function RelaListOfUserId(deptId) {
  return request({
    url: `/user-agent/dept/relaListOfUserId/${deptId} `,
    method: 'get'
  })
}

//  用户加入团队申请
export function UserJoinDept(deptId) {
  return request({
    url: `/user-agent/dept/joinApply/${deptId} `,
    method: 'post'

  })
}

//  用户撤销团队申请
export function UserCancelJoinDept(deptId) {
  return request({
    url: `/user-agent/dept/joinCancel/${deptId} `,
    method: 'put'

  })
}

// 用户重新申请
export function reJoinDept(deptId) {
  return request({
    url: `/user-agent/dept/reJoin/${deptId} `,
    method: 'put'

  })
}

//  用户申请退出团队
export function ExitDept(deptId) {
  return request({
    url: `/user-agent/dept/exit/${deptId} `,
    method: 'put'

  })
}
