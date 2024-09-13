import router from './router'
import store from './store'

import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import getPageTitle from '@/utils/get-page-title'
import { getToken } from '@/utils/auth'

const whiteList = ['/login', '/auth-redirect', '/search/index', '/register']
NProgress.configure({ showSpinner: false }) // NProgress Configuration

router.beforeEach(async(to, from, next) => {
  // start progress bar
  NProgress.start()
  document.title = getPageTitle(to.meta.title)
  const role = 'admin'
  const accessRoutes = await store.dispatch('permission/generateRoutes', role)
  next()
  router.addRoutes(accessRoutes)
  if (getToken()) {
    store.dispatch('user/getInfo')
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      next('/login')
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
