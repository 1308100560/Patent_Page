<template>
  <div class="app-container">
    <div v-if="userinfo">
      <el-row :gutter="20">

        <el-col :span="6" :xs="24">
          <user-card :user="userinfo" />
        </el-col>

        <el-col :span="18" :xs="24">
          <el-card>
            <el-tabs v-model="activeTab">
              <el-tab-pane label="基础资料" name="account">
                <account :user="userinfo" />
              </el-tab-pane>
              <el-tab-pane label="账户安全" name="password">
                <account-safety :user="userinfo" />
              </el-tab-pane>
            </el-tabs>
          </el-card>
        </el-col>

      </el-row>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import UserCard from './components/UserCard'
import Account from './components/Account'
import AccountSafety from './components/AccountSafety'

export default {
  name: 'Profile',
  components: { UserCard, Account, AccountSafety },
  data() {
    return {
      user: {},
      activeTab: 'account'
    }
  },
  computed: {
    ...mapGetters([
      'userinfo'
    ])
  },
  created() {
    this.user = JSON.parse(localStorage.getItem('user'))
  },
  methods: {
    getUser() {
      const response = {
        'requestId': 'a12ccf34-09fc-43ae-8251-9a2efdd05b20',
        'code': 200,
        'msg': '查询成功',
        'data': {
          'userId': 2,
          'username': '测试用户',
          'nickName': 'user',
          'phone': '13818888888',
          'roleId': 2,
          'avatar': 'https://img0.baidu.com/it/u=236085137,1979895699&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1665853200&t=ce6b32f431170809ecfba71babd1b103',
          'sex': '1',
          'email': '1@qq.com',
          'remark': '',
          'interest': '机器学习 /JavaScript /Css /ESLint',
          'bio': 'hello world',
          'departure': '',
          'status': '2',
          'roleIds': [
            2
          ],
          'createBy': 0,
          'updateBy': 0,
          'createdAt': '2021-05-13T19:56:37.914+08:00',
          'updatedAt': '2021-05-13T19:56:40.205+08:00'
        }
      }
      this.user = response.data
    }
  }
}
</script>
