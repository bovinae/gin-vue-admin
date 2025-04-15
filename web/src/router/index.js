import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  // {
  //   path: '/superAdmin/business/historicalVideo',
  //   name: 'HistoricalVideo',
  //   component: () => import('@/view/superAdmin/business/historicalVideo.vue'),
  //   meta: {
  //     title: 'Historical Video'
  //   }
  // },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
