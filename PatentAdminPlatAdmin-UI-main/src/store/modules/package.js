import { createPackage, getPackageList } from '@/api/package'

const state = {
  packageList: []
}

const mutations = {
  SET_PACKAGE_LIST: (state, packageList) => {
    state.packageList = packageList
  }
}

const actions = {
  // get package list
  getPackageList({ commit, state }) {
    return new Promise((resolve, reject) => {
      getPackageList().then(response => {
        const { data } = response
        commit('SET_PACKAGE_LIST', data.data)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  // create new package
  create({ commit }, packageInfo) {
    return new Promise((resolve, reject) => {
      createPackage(packageInfo).then(response => {
        const { data } = response
        console.log(state.packageList)
        if (data.code === 200) {
          commit('SET_PACKAGE_LIST', [])
          resolve()
        } else {
          reject(data.msg)
        }
      }).catch(error => {
        reject(error)
      })
    })
  }

}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
