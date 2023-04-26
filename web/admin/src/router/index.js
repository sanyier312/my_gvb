import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'

// 页面路由组件
import Index from '../components/admin/Index'
import AddArt from '../components/article/AddArt'
import ArtList from '../components/article/ArtList'
import CateList from '../components/category/CateList'
import UserList from '../components/user/UserList'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children: [
      {
        name: 'index',
        path: 'index',
        component: Index,
      },
      {
        name: 'addart',
        path: 'addart',
        component: AddArt,
      },
      {
        name: 'artlist',
        path: 'artlist',
        component: ArtList,
      },
      {
        name: 'catelist',
        path: 'catelist',
        component: CateList,
      },
      {
        name: 'userlist',
        path: 'userlist',
        component: UserList,
      },
    ],
  },
]

const router = new VueRouter({
  routes,
})

// router.beforeEach((to, from, next) => {
//   const token = window.sessionStorage.getItem('token')
//   if (to.path === '/login') return next()
//   if (!token && to.path === '/admin') {
//     next('/login')
//   } else {
//     next()
//   }
// })

export default router
