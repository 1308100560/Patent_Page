import request from '@/utils/request'

export function uploadFile(data) {
  return request({
    url: '/public/uploadFile?type=1', // type=1单文件，type=2多文件
    method: 'post', data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
