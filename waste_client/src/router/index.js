import Vue from 'vue'
import VueRouter from 'vue-router'
import homeRouter from './home'
import orderRouter from './order'
import aboutRouter from './about'

Vue.use(VueRouter)

const routes = [
  homeRouter,
  orderRouter,
  aboutRouter,
  {
    path: '/*',
    redirect:'/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
