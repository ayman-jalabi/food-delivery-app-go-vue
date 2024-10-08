import { createRouter, createWebHashHistory } from 'vue-router'

import MainView from '@/views/MainView.vue'
import LoginView from '@/views/LoginView.vue'
import RegistrationView from '@/views/RegistrationView.vue'
import CartView from '@/views/CartView.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: MainView,
      meta: { breadcrumb: 'Home' }
    },
    {
      path: '/login',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: LoginView,
      meta: { breadcrumb: 'Login' }
    },
    {
      path: '/register',
      name: 'register',
      component: RegistrationView,
      meta: { breadcrumb: 'Register' }
    },
    {
      path: '/cart',
      name: 'cart',
      component: CartView,
      meta: { breadcrumb: 'Cart' }
    }
  ]
})

export default router