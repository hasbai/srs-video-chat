import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/:room?',
    component: () => import('@/views/Home.vue'),
    name: 'Home',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
